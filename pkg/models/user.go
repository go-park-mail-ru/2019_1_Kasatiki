package models

import (
	"errors"
)

//easyjson:json
type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
	Points   int    `json:"points"`
}

//easyjson:json
type EditUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
}

//easyjson:json
type SignupUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//easyjson:json
type PublicUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
}

//easyjson:json
type LoginInfo struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

//easyjson:json
type LeaderboardUsers struct {
	Imgurl string   `json:"imgurl"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
}

func (u *User) Validation() error {
	if u.Nickname == "" || u.Email == "" || u.Password == "" {
		return errors.New("Bad request")
	}
	return nil
}
