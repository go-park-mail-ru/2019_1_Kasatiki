package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
)

// ToDo бахнуть обработку ошибок

type User struct {
	ID       string
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Points   int
}

type Order struct {
	Sequence string `json:"order"`
}

var users []User

func (u *User) setUniqueId() {
	// DB incremental or smth
	out, _ := exec.Command("uuidgen").Output()
	u.Points = 0
	u.ID = string(out[:len(out)-1])
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	errorCreateUser := map[string]string{
		"Error": "Nickname/mail already exists",
	}
	_ = json.NewDecoder(r.Body).Decode(&newUser) // ToDo: Log error
	for _, existUser := range users {
		if newUser.Nickname == existUser.Nickname || newUser.Email == existUser.Email {
			json.NewEncoder(w).Encode(errorCreateUser)
			return
		}
	}
	newUser.setUniqueId()
	users = append(users, newUser) // Check succesfull append? ( in db clearly )
	json.NewEncoder(w).Encode(newUser)

}

func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	errorOrder := map[string]string{
		"Error": "Unknown order",
	}

	_ = json.NewDecoder(r.Body).Decode(&order)
	if order.Sequence == "ASC" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Points < users[j].Points
		})
	} else if order.Sequence == "DESC" {
		sort.Slice(users, func(i, j int) bool {
			return users[i].Points > users[j].Points
		})
	} else {
		json.NewEncoder(w).Encode(errorOrder)
		return
	}
	json.NewEncoder(w).Encode(users)
}

//func editUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for i, item := range users {
//		id, _ := strconv.Atoi(params["ID"])
//		if item.ID == id {
//			var user User
//			_ = json.NewDecoder(r.Body).Decode(&user)
//			user.ID = id
//			u := &users[i]
//			*u = user
//			return
//		}
//	}
//}

//func getUser(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for _, item := range users {
//		id, _ := strconv.Atoi(params["ID"])
//		if item.ID == id {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&User{})
//}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./static/img/"+"test1.jpg", os.O_WRONLY|os.O_CREATE, 0666) // ToDo: Change way to handle img
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	var mockedUser = User{"1", "Ag1", "123", "123213", 100}
	var mockedUser1 = User{"1", "Ag1", "123", "123213", -100}
	users = append(users, mockedUser)
	users = append(users, mockedUser1)
	reciever := mux.NewRouter()
	reciever.HandleFunc("/signup", createUser).Methods("POST")
	//reciever.HandleFunc("/users/{ID}", getUser).Methods("GET")
	//reciever.HandleFunc("/settings/{ID}", editUser).Methods("POST")
	reciever.HandleFunc("/upload", upload).Methods("POST")
	reciever.HandleFunc("/leaderboard", getLeaderboard).Methods("POST")
	reciever.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
