package server

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"github.com/jackc/pgx"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func (instance *App) createUser(c *gin.Context) {
	var newUser models.User
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)
	if err != nil || newUser.Validation() != nil {
		instance.Middleware.Logger.Warnln("Create user error: ", err)
		fmt.Println(err)
		c.Status(400)
		return
	}
	_, err = instance.DB.InsertUser(newUser)
	if err != nil {
		instance.Middleware.Logger.Warnln("Create user error: ", err)
		fmt.Println(err)
		if err.(pgx.PgError).Code == "23505" {
			c.Status(409)
			return
		}
	}
	nickname := newUser.Nickname
	password := newUser.Password

	// Login after signup tmp
	var data models.LoginInfo
	data.Nickname = nickname
	data.Password = password
	_, id, err := instance.DB.LoginCheck(data)
	if err != nil {
		instance.Middleware.Logger.Warnln("Create user error: ", err)
		fmt.Println(err)
		c.Status(404)
		return
	}
	sessionId := instance.createSessionId(id)
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, true)
	c.Status(201)
}

func (instance *App) getLeaderboard(c *gin.Context) {
	var pageSize int64
	pageSize = 6
	offset, getOffset := c.Request.URL.Query()["offset"]
	coef, err := strconv.ParseInt(offset[0], 10, 32)
	if err != nil {
		instance.Middleware.Logger.Warnln("Get Leaderboard error: ", err)
		fmt.Println(err)
		c.Status(400)
		return
	}
	from := coef * pageSize
	users, err := instance.DB.GetUsers("DESC", from, pageSize)
	if getOffset {
		if len(users) == 0 || err != nil {
			instance.Middleware.Logger.Warnln("Get Leaderboard error: ", err)
			fmt.Println(err)
			c.Status(404)
			return
		}
		c.JSON(200, users)
	}
}

func (instance *App) createSessionId(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}

func (instance *App) isAuth(c *gin.Context) {
	id, _ := c.Get("id")
	user, err := instance.DB.GetUser(int(id.(float64)))
	if err != nil {
		instance.Middleware.Logger.Warnln("IsAuth error: ", err)
		fmt.Println(err)
		c.Status(404)
		return
	}
	c.JSON(200, user)
}

func (instance *App) editUser(c *gin.Context) {
	_, _, err := c.Request.FormFile("avatar")
	var edUser models.EditUser
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&edUser)
	if err != nil {
		instance.Middleware.Logger.Warnln("Edit User error: ", err)
		fmt.Println(err)
		c.Status(409)
		return
	}
	id, _ := c.Get("id")
	err = instance.DB.UpdateUser(int(id.(float64)), edUser)
	if err != nil {
		instance.Middleware.Logger.Warnln("Edit User error: ", err)
		fmt.Println(err)
		constrain := err.(pgx.PgError).ConstraintName
		if constrain == "users_nickname_key" {
			c.Status(300)
			return
		}
		if constrain == "users_email_key" {
			c.Status(301)
			return
		}
		c.Status(303)
		return
	}
	c.Status(200)
}

func (instance *App) login(c *gin.Context) {
	var data models.LoginInfo
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	fmt.Println("!#!##!#")
	if err != nil {
		instance.Middleware.Logger.Warnln("Login error: ", err)
		fmt.Println(err)
		c.Status(400)
		return
	}
	_, id, err := instance.DB.LoginCheck(data)
	if err != nil {
		instance.Middleware.Logger.Warnln("Login error: ", err)
		fmt.Println(err)
		c.Status(404)
		return
	}
	sessionId := instance.createSessionId(id)
	c.SetCookie("session_id", sessionId, 3600, "/", "", false, true)
	c.Status(201)
}

func (instance *App) upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		instance.Middleware.Logger.Warnln("Upload error: ", err)
		fmt.Println(err)
		c.Status(409)
		return
	}
	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		instance.Middleware.Logger.Warnln("Upload error: ", err)
		fmt.Println(err)
		c.Status(409)
		return
	}
	defer file.Close()
	id, _ := c.Get("id")
	picpath := "./static/img" + strconv.Itoa(int(id.(float64))) + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		instance.Middleware.Logger.Warnln("Upload error: ", err)
		fmt.Println(err)
		c.Status(404)
		return
	}
	ImgUrl := "https://advhater.ru/img/" + strconv.Itoa(int(id.(float64))) + ".jpeg"
	err = instance.DB.ImgUpdate(int(id.(float64)), ImgUrl)
	if err != nil {
		instance.Middleware.Logger.Warnln("Upload error: ", err)
		fmt.Println(err)
		c.Status(404)
		return
	}
	c.Status(200)
	defer f.Close()
	io.Copy(f, file)
}

func (instance *App) logout(c *gin.Context) {
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.Status(200)
}
