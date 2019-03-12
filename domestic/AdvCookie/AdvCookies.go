package AdvCookie

import (
	"2019_1_Kasatiki/domestic/Models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)

func GetClaims(r *http.Request) jwt.MapClaims {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		// ToDo: Error handle
		return nil
	}
	claims := checkAuth(cookie)
	return claims
}

func CreateSessionId(user Models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}

func checkAuth(cookie *http.Cookie) jwt.MapClaims {
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
