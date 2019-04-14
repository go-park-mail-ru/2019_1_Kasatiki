package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"models"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func (instance *App) createUser(c *gin.Context) {
	var newUser models.User
	_ = json.NewDecoder(c.Request.Body).Decode(&newUser) // ToDo: Log error
	for _, existUser := range Users {
		if newUser.Nickname == existUser.Nickname || newUser.Email == existUser.Email {
			//json.NewEncoder(w).Encode(errorCreateUser)
			c.JSON(407, errorCreateUser)
			return
		}
	}
	newUser.SetUniqueId()
	Users = append(Users, newUser) // Check succesfull append? ( in db clearly )
}

//ToDo: Use get with key order? (ASC/DESC )
//ToDo: Check and simplify conditions !!!
func (instance *App) getLeaderboard(c *gin.Context) {
	//var order Order
	var pageSize int
	// Initilize pagesize
	pageSize = 1
	//_ = json.NewDecoder(r.Body).Decode(&order)
	sort.Slice(Users, func(i, j int) bool {
		return Users[i].Points > Users[j].Points
	})
	offset, getOffset := c.Request.URL.Query()["offset"]
	if getOffset {
		offsetInt, _ := strconv.ParseInt(offset[0], 10, 32) // ToDo Handle error
		if int(offsetInt) > len(Users) {
			c.JSON(200, Users)
			//json.NewEncoder(w).Encode(Users)
			return
		} else if int(offsetInt) == len(Users) {
			c.JSON(200, Users)
			//json.NewEncoder(w).Encode(Users)
			return
		}
		if int(offsetInt)+pageSize < len(Users) {
			c.JSON(200, Users[offsetInt:int(offsetInt)+pageSize])
			//json.NewEncoder(w).Encode(Users[offsetInt : int(offsetInt)+pageSize])
			return
		} else {
			c.JSON(200, Users[offsetInt:len(Users)])
			//json.NewEncoder(w).Encode(Users[offsetInt:len(Users)])
			return
		}
	} else {
		if pageSize < len(Users) {
			c.JSON(200, Users[:pageSize])
			//json.NewEncoder(w).Encode(Users[:pageSize])
			return
		} else {
			c.JSON(200, Users)
			//json.NewEncoder(w).Encode(Users)
			return
		}
	}

}

func (instance *App) createSessionId(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}

func (instance *App) checkAuth(cookie *http.Cookie) jwt.MapClaims {
	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		spiceSalt, _ := ioutil.ReadFile("secret.conf")
		return spiceSalt, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	// ToDo: Handle else case
	return claims
}

func (instance *App) isAuth(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error of cookie")
		//c.Write([]byte("{}"))
		return
	}

	claims := instance.checkAuth(cookie)
	for _, user := range Users {
		if user.ID == claims["id"].(string) {
			c.JSON(200, map[string]bool{"is_auth": true})
			//json.NewEncoder(w).Encode(map[string]bool{"is_auth": true})
			return
		}
	}
	c.JSON(200, map[string]bool{"is_auth": false})
	//json.NewEncoder(w).Encode(map[string]bool{"is_auth": false})
}

func (instance *App) editUser(c *gin.Context) {
	fmt.Println(Users)
	//Checking cookie
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error of cookie")
		//w.Write([]byte("{}"))
		return
	}
	// Taking JSON of modified user from edit form
	var modUser models.User
	_ = json.NewDecoder(c.Request.Body).Decode(&modUser)
	file, _, err := c.Request.FormFile("avatar")
	fmt.Println(file)
	// Getting claims from current cookie
	claims := instance.checkAuth(cookie)

	// Finding user from claims in Users and changing old data to modified data
	for i, user := range Users {
		if user.ID == claims["id"].(string) {
			u := &Users[i]
			if modUser.Nickname != "" {
				u.Nickname = modUser.Nickname
			}
			if modUser.Email != "" {
				u.Email = modUser.Email
			}
			if modUser.Password != "" {
				u.Password = modUser.Password
			}
			if modUser.Region != "" {
				u.Region = modUser.Region
			}
			if modUser.Age != 0 {
				u.Age = modUser.Age
			}
			if modUser.About != "" {
				u.About = modUser.About
			}
			if modUser.ImgUrl != "" {
				u.ImgUrl = modUser.ImgUrl
			}
			c.JSON(200, *u)
			//json.NewEncoder(w).Encode(*u)
			break
		}
	}
}

func (instance *App) login(c *gin.Context) {
	var sessionId string
	var userExistFlag bool
	var existUser models.User
	_ = json.NewDecoder(c.Request.Body).Decode(&existUser)
	for _, user := range Users {
		if user.Nickname == existUser.Nickname && user.Password == existUser.Password {
			userExistFlag = true
			existUser = user
		}
	}
	if !userExistFlag {
		c.JSON(404, errorLogin)
		//json.NewEncoder(w).Encode(errorLogin)
		return
	}
	sessionId = instance.createSessionId(existUser)
	//cookie := &http.Cookie{
	//	Name:     "session_id",
	//	Value:    sessionId,
	//	HttpOnly: true,
	//}
	c.SetCookie("sesion_id", sessionId, 3600, "/", "", true, true) // Todo:check this params
	//http.SetCookie(w, cookie)
	c.JSON(200, existUser)
	//json.NewEncoder(w).Encode(existUser)
}

func (instance *App) getMe(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error of cookie")
		//w.Write([]byte("{}"))
		return
	}
	claims := instance.checkAuth(cookie)
	for _, user := range Users {
		if user.ID == claims["id"].(string) {
			c.JSON(200, user)
			//json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func (instance *App) getUser(c *gin.Context) {

	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(c.Request) // ToDo: illuminate mux
	for _, item := range Users {
		//id, _ := strconv.Atoi(params["ID"])
		if item.ID == params["id"] {
			c.JSON(200, item)
			//json.NewEncoder(w).Encode(item)
			return
		}
	}
	c.Header("Content-Type", "application/json")
	c.JSON(404, &models.User{}) // ToDo: replace &User{}
	//json.NewEncoder(w).Encode(&User{})
}

func (instance *App) upload(c *gin.Context) {
	// Tacking file from request
	fmt.Println("UPLOAD")
	_ = c.Request.ParseMultipartForm(32 << 20)
	fmt.Println(c.Request)
	file, _, err := c.Request.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("FILE IS HERE")

	// Tacking cookie of current user
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error")
		return
	}
	claims := instance.checkAuth(cookie)
	// Path to Users avatar
	picpath := "./static/img/" + claims["id"].(string) + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Changing ImgURL field in current user
	for i, user := range Users {
		if user.ID == claims["id"].(string) {
			u := &Users[i]
			u.ImgUrl = "https://advhater.ru/img/" + claims["id"].(string) + ".jpeg"
		}
	}
	defer f.Close()
	io.Copy(f, file)
}

// ToDo: Add cookie handle
func (instance *App) logout(c *gin.Context) {
	_, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error of cookie")
		//w.WriteHeader(http.StatusNotFound)
		return
	}

	//http.SetCookie(w, &http.Cookie{
	//	Name:     "session_id",
	//	Value:    "",
	//	Expires:  time.Now().AddDate(0, 0, -1),
	//	Path:     "/",
	//	HttpOnly: true,
	//})
	c.SetCookie("sesion_id", "", -1, "/", "", true, true) // Todo:check this params

	//w.WriteHeader(http.StatusOK)
}
