package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func CheckAuth(cookie *http.Cookie) (jwt.MapClaims, error) {
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

func (instance *Middlewares) AuthMiddleware(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		claims, err := CheckAuth(cookie)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		c.Set("id", claims["id"])
		fmt.Println(c.Get("id"))
		handlerFunc(c)
	}
}
