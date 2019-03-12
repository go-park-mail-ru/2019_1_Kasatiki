package main

import (
	"2019_1_Kasatiki/domestic/AdvCookie"
	"2019_1_Kasatiki/domestic/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

// ToDo: delete this struct
type Order struct {
	Sequence string `json:"order"`
}

//ToDo: set all errors in one map and set this map in AdvErrors.go
var errorLogin = map[string]string{
	"Error": "User dont exist",
}

var errorCreateUser = map[string]string{
	"Error": "Nickname/mail already exists",
}

var users []Models.User

// ToDo: set this func in Handlers.go or API.go or something else
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser Models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser) // ToDo: Log error
	for _, existUser := range users {
		if newUser.Nickname == existUser.Nickname || newUser.Email == existUser.Email {
			json.NewEncoder(w).Encode(errorCreateUser)
			return
		}
	}
	newUser.SetUniqueId()
	users = append(users, newUser) // Check succesfull append? ( in db clearly )
	//json.NewEncoder(w).Encode(newUser)

}

// ToDo: set this func in Handlers.go or API.go or something else
//ToDo: Use get with key order? (ASC/DESC )
//ToDo: Check and simplify conditions !!!
func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	var pageSize int
	// Initilize pagesize
	pageSize = 1
	_ = json.NewDecoder(r.Body).Decode(&order)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})
	offset, getOffset := r.URL.Query()["offset"]
	if getOffset {
		offsetInt, _ := strconv.ParseInt(offset[0], 10, 32) // ToDo Handle error
		if int(offsetInt) > len(users) {
			json.NewEncoder(w).Encode(users)
			return
		} else if int(offsetInt) == len(users) {
			json.NewEncoder(w).Encode(users)
			return
		}
		if int(offsetInt)+pageSize < len(users) {
			json.NewEncoder(w).Encode(users[offsetInt : int(offsetInt)+pageSize])
			return
		} else {
			json.NewEncoder(w).Encode(users[offsetInt:len(users)])
			return
		}
	} else {
		if pageSize < len(users) {
			json.NewEncoder(w).Encode(users[:pageSize])
			return
		} else {
			json.NewEncoder(w).Encode(users)
			return
		}
	}

}

// ToDo: set this func in Handlers.go or API.go or something else
func isAuth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}

	claims := AdvCookie.CheckAuth(cookie)
	for _, user := range users {
		if user.Nickname == claims["id"].(string) {
			json.NewEncoder(w).Encode(map[string]bool{"is_auth": true})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]bool{"is_auth": false})
}

// ToDo: set this func in Handlers.go or API.go or something else
func editUser(w http.ResponseWriter, r *http.Request) {
	//Checking cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	// Taking JSON of modified user from edit form
	var modUser Models.User
	_ = json.NewDecoder(r.Body).Decode(&modUser)
	// Getting claims from current cookie
	claims := AdvCookie.CheckAuth(cookie)

	// Finding user from claims in users and changing old data to modified data
	for i, user := range users {
		if user.ID == claims["id"].(string) {
			u := &users[i]
			if modUser.Nickname != "" {
				u.Nickname = modUser.Nickname
			}
			if modUser.Email != "" {
				u.Email = modUser.Email
			}
			if modUser.Password != "" {
				u.Password = modUser.Password
			}
			if modUser.Region != "" {
				u.Region = modUser.Region
			}
			if modUser.Age != 0 {
				u.Age = modUser.Age
			}
			if modUser.About != "" {
				u.About = modUser.About
			}
			if modUser.ImgUrl != "" {
				u.ImgUrl = modUser.ImgUrl
			}
			json.NewEncoder(w).Encode(*u)
			break
		}
	}
}

// ToDo: set this func in Handlers.go or API.go or something else
func login(w http.ResponseWriter, r *http.Request) {
	var userExistFlag bool
	var existUser Models.User
	_ = json.NewDecoder(r.Body).Decode(&existUser)
	for _, user := range users {
		if user.Nickname == existUser.Nickname && user.Password == existUser.Password {
			userExistFlag = true
			existUser = user
		}
	}
	if !userExistFlag {
		json.NewEncoder(w).Encode(errorLogin)
		return
	}
	sessionId := AdvCookie.CreateSessionId(existUser)
	// ToDo: set this move in another func
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(existUser)
}

// ToDo: set this func in Handlers.go or API.go or something else
func getMe(w http.ResponseWriter, r *http.Request) {
	// ToDo: set all process of getting cookies and claims in one func and return claims and error
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	claims := AdvCookie.CheckAuth(cookie)
	for _, user := range users {
		if user.ID == claims["id"].(string) {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

// ToDo: set this func in Handlers.go or API.go or something else
// ToDO: Add case sensitive ( high/low )
func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		//id, _ := strconv.Atoi(params["ID"])
		if item.Nickname == params["Nickname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Models.User{})
}

// ToDo: set this func in Handlers.go or API.go or something else
func upload(w http.ResponseWriter, r *http.Request) {
	// Tacking file from request
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		return
	}
	defer file.Close()
	// Tacking cookie of current user
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	claims := AdvCookie.CheckAuth(cookie)
	// Path to users avatar
	// ToDo: set process of saving pict in some func
	picpath := "./static/img/" + claims["id"].(string) + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	// Changing ImgURL field in current user
	for i, user := range users {
		if user.ID == claims["id"].(string) {
			u := &users[i]
			u.ImgUrl = "http://advhater.ru/img/" + claims["id"].(string) + ".jpeg"
		}
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	// ToDo: Set this mocking in some func and file
	// Mocked part for leaderboard
	var mockedUser = Models.User{"1", "evv", "onetaker@gmail.com",
		"evv", -100, 23, "test",
		"Voronezh", "В левой руке салам"}
	var mockedUser1 = Models.User{"2", "tony", "trendpusher@hydra.com",
		"qwerty", 100, 22, "test",
		"Moscow", "В правой алейкум"}

	// Mocker part end
	users = append(users, mockedUser)
	users = append(users, mockedUser1)

	// ToDo: set server in struct
	reciever := mux.NewRouter()

	// GET  ( get exists data )
	reciever.HandleFunc("/users/{Nickname}", getUser).Methods("GET")
	reciever.HandleFunc("/leaderboard", getLeaderboard).Methods("GET")
	reciever.HandleFunc("/isauth", isAuth).Methods("GET")
	reciever.HandleFunc("/me", getMe).Methods("Get")

	// POST ( create new data )
	reciever.HandleFunc("/signup", createUser).Methods("POST")
	reciever.HandleFunc("/upload", upload).Methods("POST")
	reciever.HandleFunc("/login", login).Methods("POST")

	// Todo: change method of request on PUT (CRUD)
	reciever.HandleFunc("/users/{Nickname}", editUser).Methods("POST")

	reciever.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/"))) // Uncomment if want to run locally
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
