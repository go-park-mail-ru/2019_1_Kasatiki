package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"time"
)

type App struct {
	Router *mux.Router // ToDo: __future__
}

func (instance *App) Initialize() {
	var mockedUser = User{"1", "evv", "onetaker@gmail.com",
		"evv", -100, 23, "test",
		"Voronezh", "В левой руке салам"}
	var mockedUser1 = User{"2", "tony", "trendpusher@hydra.com",
		"qwerty", 100, 22, "test",
		"Moscow", "В правой алейкум"}
	// Mocked users

	Users = append(Users, mockedUser)
	Users = append(Users, mockedUser1)
	instance.Router = mux.NewRouter()
	instance.initializeRoutes()
}

func (instance *App) initializeRoutes() {
	// GET ( get exist data )
	instance.Router.HandleFunc("/api/leaderboard", instance.getLeaderboard).Methods("GET")
	instance.Router.HandleFunc("/api/isauth", instance.isAuth).Methods("GET")
	instance.Router.HandleFunc("/api/me", instance.getMe).Methods("GET")
	instance.Router.HandleFunc("/api/logout", instance.logout).Methods("GET") // ToDO: Cors added ( maybe post?)

	// POST ( create new data )
	instance.Router.HandleFunc("/api/signup", instance.createUser).Methods("POST")
	instance.Router.HandleFunc("/api/upload", instance.upload).Methods("POST")
	instance.Router.HandleFunc("/api/login", instance.login).Methods("POST")

	// PUT ( update data )
	instance.Router.HandleFunc("/api/users/{Nickname}", instance.editUser).Methods("PUT")

	//Static path
	//instance.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
}

func (instance *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, instance.Router))
}

type User struct {
	ID       string
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Points   int
	Age      int
	ImgUrl   string
	Region   string
	About    string
}

type Order struct {
	Sequence string `json:"order"`
}

//ToDo: Move to another package
var errorLogin = map[string]string{
	"Error": "User dont exist",
}

var errorCreateUser = map[string]string{
	"Error": "Nickname/mail already exists",
}

var Users []User

func (u *User) setUniqueId() {
	// DB incremental or smth
	out, _ := exec.Command("uuidgen").Output()
	u.Points = 0
	u.ID = string(out[:len(out)-1])
}

func (instance *App) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Users)
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser) // ToDo: Log error
	for _, existUser := range Users {
		if newUser.Nickname == existUser.Nickname || newUser.Email == existUser.Email {
			json.NewEncoder(w).Encode(errorCreateUser)
			return
		}
	}
	newUser.setUniqueId()
	Users = append(Users, newUser) // Check succesfull append? ( in db clearly )
	//json.NewEncoder(w).Encode(newUser)

}

//ToDo: Use get with key order? (ASC/DESC )
//ToDo: Check and simplify conditions !!!
func (instance *App) getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var order Order
	var pageSize int
	// Initilize pagesize
	pageSize = 1
	//_ = json.NewDecoder(r.Body).Decode(&order)
	sort.Slice(Users, func(i, j int) bool {
		return Users[i].Points > Users[j].Points
	})
	offset, getOffset := r.URL.Query()["offset"]
	if getOffset {
		offsetInt, _ := strconv.ParseInt(offset[0], 10, 32) // ToDo Handle error
		if int(offsetInt) > len(Users) {
			json.NewEncoder(w).Encode(Users)
			return
		} else if int(offsetInt) == len(Users) {
			json.NewEncoder(w).Encode(Users)
			return
		}
		if int(offsetInt)+pageSize < len(Users) {
			json.NewEncoder(w).Encode(Users[offsetInt : int(offsetInt)+pageSize])
			return
		} else {
			json.NewEncoder(w).Encode(Users[offsetInt:len(Users)])
			return
		}
	} else {
		if pageSize < len(Users) {
			json.NewEncoder(w).Encode(Users[:pageSize])
			return
		} else {
			json.NewEncoder(w).Encode(Users)
			return
		}
	}

}

func (instance *App) createSessionId(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}

func (instance *App) checkAuth(cookie *http.Cookie) jwt.MapClaims {
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

func (instance *App) isAuth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}

	claims := instance.checkAuth(cookie)
	for _, user := range Users {
		if user.ID == claims["id"].(string) {
			json.NewEncoder(w).Encode(map[string]bool{"is_auth": true})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]bool{"is_auth": false})
}

func (instance *App) editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Users)
	//Checking cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	// Taking JSON of modified user from edit form
	var modUser User
	_ = json.NewDecoder(r.Body).Decode(&modUser)
	file, _, err := r.FormFile("avatar")
	fmt.Println(file)
	// Getting claims from current cookie
	claims := instance.checkAuth(cookie)

	// Finding user from claims in Users and changing old data to modified data
	for i, user := range Users {
		if user.ID == claims["id"].(string) {
			u := &Users[i]
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

func (instance *App) login(w http.ResponseWriter, r *http.Request) {
	var sessionId string
	var userExistFlag bool
	var existUser User
	_ = json.NewDecoder(r.Body).Decode(&existUser)
	for _, user := range Users {
		if user.Nickname == existUser.Nickname && user.Password == existUser.Password {
			userExistFlag = true
			existUser = user
		}
	}
	if !userExistFlag {
		json.NewEncoder(w).Encode(errorLogin)
		return
	}
	sessionId = instance.createSessionId(existUser)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(existUser)
}

func (instance *App) getMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	claims := instance.checkAuth(cookie)
	for _, user := range Users {
		if user.ID == claims["id"].(string) {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

// ToDO: Add case sensitive ( high/low )
func (instance *App) getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Users {
		//id, _ := strconv.Atoi(params["ID"])
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func (instance *App) upload(w http.ResponseWriter, r *http.Request) {
	// Tacking file from request
	fmt.Println("UPLOAD")
	r.ParseMultipartForm(32 << 20)
	fmt.Println(r)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("FILE IS HERE")

	// Tacking cookie of current user
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("{}"))
		return
	}
	claims := instance.checkAuth(cookie)
	// Path to Users avatar
	picpath := "./static/img/" + claims["id"].(string) + ".jpeg"
	f, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Changing ImgURL field in current user
	for i, user := range Users {
		if user.ID == claims["id"].(string) {
			u := &Users[i]
			u.ImgUrl = "https://advhater.ru/img/" + claims["id"].(string) + ".jpeg"
		}
	}
	defer f.Close()
	io.Copy(f, file)
}

// ToDo: Add cookie handle
func (instance *App) logout(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_id")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		HttpOnly: true,
	})
	w.WriteHeader(http.StatusOK)
}
