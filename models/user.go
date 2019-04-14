package models

import (
	"errors"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
	Points   int    `json"points"`
}

type EditUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
}

type SignupUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
	Age      int    `json:"age"`
	ImgUrl   string `json:"imgurl"`
	Region   string `json:"region"`
	About    string `json:"about"`
}

type LoginInfo struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LeaderboardUsers struct {
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
