package main

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"io"
	"io/ioutil"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/models"
	"net/http"
	"os"
	"strconv"
)

func (instance *App) createUser(c *gin.Context) {
	var newUser models.User
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)

	if err != nil || newUser.Validation() != nil {
		c.Status(400)
		return
	}
	_, err = instance.InsertUser(newUser)
	if err != nil {
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
	_, id, err := instance.LoginCheck(data)
	if err != nil {
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
		c.Status(400)
		return
	}
	from := coef * pageSize
	users, err := instance.GetUsers("DESC", from, pageSize)
	if getOffset {
		if len(users) == 0 || err != nil {
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

func (instance *App) checkAuth(cookie *http.Cookie) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		spiceSalt, _ := ioutil.ReadFile("secret.conf")
		return spiceSalt, nil
	})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	// ToDo: Handle else case
	return claims, nil
}

func (instance *App) isAuth(c *gin.Context) {
	//cookie, err := c.Request.Cookie("session_id")
	//if err != nil {
	//	c.Status(404)
	//	return
	//}
	//claims, err := instance.checkAuth(cookie)
	//if err != nil {
	//
	//	c.Status(404)
	//	return
	//}
	id, _ := c.Get("id")

	user, err := instance.GetUser(id.(int))
	if err != nil {
		fmt.Println(err)
		c.Status(404)
		return
	}
	c.JSON(200, user)
}

func (instance *App) editUser(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.Status(404)
		return
	}
	claims, err := instance.checkAuth(cookie)
	if err != nil {

		c.Status(404)
		return
	}
	id := claims["id"].(float64)
	_, _, err = c.Request.FormFile("avatar")

	var edUser models.EditUser
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&edUser)
	fmt.Println(err)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}
	err = instance.UpdateUser(id, edUser)
	if err != nil {
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
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		c.Status(400)
		return
	}
	_, id, err := instance.LoginCheck(data)
	if err != nil {
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
		instance.Logger.Warn(err)
		fmt.Println(err)
		c.Status(409)
		return
	}
	file, _, err := c.Request.FormFile("avatar")
	if err != nil {
		instance.Logger.Warn(err)
		fmt.Println(err)
		c.Status(409)
		return
	}
	defer file.Close()

	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.Status(404)
		return
	}
	claims, err := instance.checkAuth(cookie)
	id := int(claims["id"].(float64))
	// Может падать из-за отсутствия этой папки
	picpath := "./static/img" + strconv.Itoa(id) + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		c.Status(404)
		return
	}

	ImgUrl := "https://advhater.ru/img/" + strconv.Itoa(id) + ".jpeg"
	err = instance.ImgUpdate(id, ImgUrl)
	if err != nil {
		c.Status(404)
		return
	}
	c.Status(200)
	defer f.Close()
	io.Copy(f, file)
}

func (instance *App) logout(c *gin.Context) {
	//_, err := c.Request.Cookie("session_id")
	//if err != nil {
	//	c.Status(404)
	//	return
	//}
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.Status(200)
}
