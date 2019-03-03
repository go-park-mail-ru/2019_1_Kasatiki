package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID         int
	Nickname   string
	password   string
	CardNumber string
	CashPoints int
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func editUser(w http.ResponseWriter, r *http.Request) {

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		fmt.Println(users)
		if item.Nickname == params["Nickname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func main() {
	reciever := mux.NewRouter()
	reciever.HandleFunc("/createUser", createUser).Methods("POST")
	reciever.HandleFunc("/users/{Nickname}", getUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
