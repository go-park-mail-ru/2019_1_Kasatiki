package AdvCookie

import (
	"2019_1_Kasatiki/domestic/Models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)

func GetClaims(r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err
	}
	claims, err := checkAuth(cookie)
	return claims, nil
}

func CreateSessionId(user Models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	spiceSalt, err := ioutil.ReadFile("secret.conf")
	if err != nil {
		fmt.Println("Can't read secret.conf")
		fmt.Println(err)
		return "", err
	}
	secretStr, err := token.SignedString(spiceSalt)
	if err != nil {
		fmt.Println("Can't take secret string")
		fmt.Println(err)
		return "", err
	}
	return secretStr, nil
}

func checkAuth(cookie *http.Cookie) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		spiceSalt, _ := ioutil.ReadFile("secret.conf")
		return spiceSalt, nil
	})
	if err != nil {
		fmt.Println("JWT parse error")
		fmt.Println(err)
		return nil, err
	}
	claims, err := token.Claims.(jwt.MapClaims)
	if err != nil {
		fmt.Println("Getting MapClaims from token error")
		fmt.Println(err)
		return nil, err
	}
	return claims, nil
}
