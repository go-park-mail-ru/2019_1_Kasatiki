package models

import (
	"testing"
)

func TestUser_Validation(t *testing.T) {
	var newUser User
	err := newUser.Validation()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	newUser.Nickname = "1"
	newUser.Password = "1"
	newUser.Email = "1"
	valid := newUser.Validation()
	if valid != nil {
		t.Errorf("Expected nil but got err: %s", valid.Error())
	}
}
