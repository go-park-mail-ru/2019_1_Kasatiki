package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
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

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./public/img/"+"test1.jpg", os.O_WRONLY|os.O_CREATE, 0666) // ToDo: Change way to handle img
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	reciever := mux.NewRouter()
	reciever.HandleFunc("/signup", createUser).Methods("POST")
	reciever.HandleFunc("/users/{ID}", getUser).Methods("GET")
	reciever.HandleFunc("/settings/{ID}", editUser).Methods("POST")
	reciever.HandleFunc("/upload", upload).Methods("POST")
	reciever.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
