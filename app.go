package main

import (
	"./domestic/AdvCookie"
	"./domestic/Models"
	"2019_1_Kasatiki/domestic/Errors"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB // ToDo: __future__
}

func (instance *App) Initialize() {
	instance.Router = mux.NewRouter()
	instance.initializeRoutes()
}

func (instance *App) initializeRoutes() {
	// GET ( get exist data )
	instance.Router.HandleFunc("/leaderboard", instance.getLeaderboard).Methods("GET")
	instance.Router.HandleFunc("/isauth", instance.isAuth).Methods("GET")
	instance.Router.HandleFunc("/me", instance.getMe).Methods("GET")
	instance.Router.HandleFunc("/logout", instance.logout).Methods("GET") // ToDO: Cors added ( maybe post?)

	// POST ( create new data )
	instance.Router.HandleFunc("/signup", instance.createUser).Methods("POST")
	instance.Router.HandleFunc("/upload", instance.upload).Methods("POST")
	instance.Router.HandleFunc("/login", instance.login).Methods("POST")

	// PUT ( update data )
	instance.Router.HandleFunc("/users/{Nickname}", instance.editUser).Methods("PUT")

	//Static path
	instance.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
}

func (instance *App) Run(port string) {
	instance.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatalln(http.ListenAndServe(port, instance.Router))
}

type Order struct {
	Sequence string `json:"order"`
}

var users []Models.User

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

func (instance *App) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser Models.User
	// Taking JSON with new user's data
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// Sending error response
		err := json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	// Checking new user's data on existing
	for _, existUser := range users {
		// If nickname exist
		if newUser.Nickname == existUser.Nickname {
			err := json.NewEncoder(w).Encode(Errors.Error["Nickname already exist"])
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
			// If email exist
		} else if newUser.Email == existUser.Email {
			err := json.NewEncoder(w).Encode(Errors.Error["Email already exist"])
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	}
	newUser.SetUniqueId()
	users = append(users, newUser)
}

func (instance *App) getLeaderboard(w http.ResponseWriter, r *http.Request) {
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
		offsetInt, err := strconv.ParseInt(offset[0], 10, 32) // ToDo Handle error
		if err != nil {
			fmt.Println("Parsing ofset error")
			fmt.Println(err)
		}
		if int(offsetInt) > len(users) {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		} else if int(offsetInt) == len(users) {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
		if int(offsetInt)+pageSize < len(users) {
			err := json.NewEncoder(w).Encode(users[offsetInt : int(offsetInt)+pageSize])
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		} else {
			err := json.NewEncoder(w).Encode(users[offsetInt:len(users)])
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	} else {
		if pageSize < len(users) {
			err := json.NewEncoder(w).Encode(users[:pageSize])
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		} else {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	}

}

func (instance *App) isAuth(w http.ResponseWriter, r *http.Request) {
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	for _, user := range users {
		if user.Nickname == claims["id"].(string) {
			err := json.NewEncoder(w).Encode(map[string]bool{"is_auth": true})
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	}
	err = json.NewEncoder(w).Encode(map[string]bool{"is_auth": false})
	if err != nil {
		fmt.Println("Can't sand response")
		fmt.Println(err)
		return
	}
}

func (instance *App) editUser(w http.ResponseWriter, r *http.Request) {
	//Checking cookie
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	// Taking JSON of modified user from edit form
	var modUser Models.User
	err = json.NewDecoder(r.Body).Decode(&modUser)
	if err != nil {
		fmt.Println("Bad JSON")
		fmt.Println(err)
		err = json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
	}

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
			err := json.NewEncoder(w).Encode(*u)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
				return
			}
			break
		}
	}
}

func (instance *App) login(w http.ResponseWriter, r *http.Request) {
	var userExistFlag bool
	var existUser Models.User
	err := json.NewDecoder(r.Body).Decode(&existUser)
	if err != nil {
		fmt.Println("Bad JSON")
		fmt.Println(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
	}
	for _, user := range users {
		if user.Nickname == existUser.Nickname && user.Password == existUser.Password {
			userExistFlag = true
			existUser = user
		}
	}
	if !userExistFlag {
		err := json.NewEncoder(w).Encode(Errors.Error["User does not exist"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	sessionId, err := AdvCookie.CreateSessionId(existUser)
	if err != nil {

	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	err = json.NewEncoder(w).Encode(existUser)
	if err != nil {
		fmt.Println("Can't sand response")
		fmt.Println(err)
		return
	}
}

func (instance *App) getMe(w http.ResponseWriter, r *http.Request) {
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	for _, user := range users {
		if user.ID == claims["id"].(string) {
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	}
}

// ToDO: Add case sensitive ( high/low )
func (instance *App) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				fmt.Println("Can't sand response")
				fmt.Println(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Models.User{})
	if err != nil {
		fmt.Println("Can't sand response")
		fmt.Println(err)
	}
}

func (instance *App) upload(w http.ResponseWriter, r *http.Request) {
	// Tacking file from request
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("Memory error")
		fmt.Println(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Bad File"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println("Bad file read")
		fmt.Println(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Bad File"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	defer file.Close()
	// Tacking cookie of current user
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			fmt.Println("Can't sand response")
			fmt.Println(err)
		}
		return
	}
	// Path to users avatar
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
