package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// ToDo бахнуть обработку ошибок

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range users {
		id, _ := strconv.Atoi(params["ID"])
		if item.ID == id {
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = id
			u := &users[i]
			*u = user
			return
		}
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		id, _ := strconv.Atoi(params["ID"])
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func main() {
	reciever := mux.NewRouter()
	reciever.HandleFunc("/createUser", createUser).Methods("POST")
	reciever.HandleFunc("/users/{ID}", getUser).Methods("GET")
	reciever.HandleFunc("/settings/{ID}", editUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
