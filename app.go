package main

import (
	"./domestic/AdvCookie"
	"./domestic/Models"
	"2019_1_Kasatiki/domestic/Errors"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
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
	instance.Router.HandleFunc("/users/{Nickname}", instance.getUser).Methods("GET")
	instance.Router.HandleFunc("/leaderboard", instance.getLeaderboard).Methods("GET")
	instance.Router.HandleFunc("/isauth", instance.isAuth).Methods("GET")
	instance.Router.HandleFunc("/me", instance.getMe).Methods("GET")

	// POST ( create new data )
	instance.Router.HandleFunc("/signup", instance.createUser).Methods("POST")
	instance.Router.HandleFunc("/upload", instance.upload).Methods("POST")
	instance.Router.HandleFunc("/login", instance.login).Methods("POST")
	instance.Router.HandleFunc("/users/{Nickname}", instance.editUser).Methods("POST")

	// PUT ( update data )
}

func (instance *App) Run(port string) {
	instance.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Fatalln(http.ListenAndServe(port, instance.Router))
}

type Order struct {
	Sequence string `json:"order"`
}

var users []Models.User

func (instance *App) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser Models.User
	// Taking JSON with new user's data
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		// Sending error response
		err := json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	// Checking new user's data on existing
	for _, existUser := range users {
		// If nickname exist
		if newUser.Nickname == existUser.Nickname {
			err := json.NewEncoder(w).Encode(Errors.Error["Nickname already exist"])
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
			// If email exist
		} else if newUser.Email == existUser.Email {
			err := json.NewEncoder(w).Encode(Errors.Error["Email already exist"])
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
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
			log.Fatalln("Parsing ofset error")
			log.Fatalln(err)
		}
		if int(offsetInt) > len(users) {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		} else if int(offsetInt) == len(users) {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		}
		if int(offsetInt)+pageSize < len(users) {
			err := json.NewEncoder(w).Encode(users[offsetInt : int(offsetInt)+pageSize])
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		} else {
			err := json.NewEncoder(w).Encode(users[offsetInt:len(users)])
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		}
	} else {
		if pageSize < len(users) {
			err := json.NewEncoder(w).Encode(users[:pageSize])
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		} else {
			err := json.NewEncoder(w).Encode(users)
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
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
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	for _, user := range users {
		if user.Nickname == claims["id"].(string) {
			err := json.NewEncoder(w).Encode(map[string]bool{"is_auth": true})
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		}
	}
	err = json.NewEncoder(w).Encode(map[string]bool{"is_auth": false})
	if err != nil {
		log.Fatalln("Can't sand response")
		log.Fatalln(err)
		return
	}
}

func (instance *App) editUser(w http.ResponseWriter, r *http.Request) {
	//Checking cookie
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	// Taking JSON of modified user from edit form
	var modUser Models.User
	err = json.NewDecoder(r.Body).Decode(&modUser)
	if err != nil {
		log.Fatalln("Bad JSON")
		log.Fatalln(err)
		err = json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
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
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
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
		log.Fatalln("Bad JSON")
		log.Fatalln(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Invalid JSON"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
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
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
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
		log.Fatalln("Can't sand response")
		log.Fatalln(err)
		return
	}
}

func (instance *App) getMe(w http.ResponseWriter, r *http.Request) {
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	for _, user := range users {
		if user.ID == claims["id"].(string) {
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
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
				log.Fatalln("Can't sand response")
				log.Fatalln(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Models.User{})
	if err != nil {
		log.Fatalln("Can't sand response")
		log.Fatalln(err)
	}
}

func (instance *App) upload(w http.ResponseWriter, r *http.Request) {
	// Tacking file from request
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatalln("Memory error")
		log.Fatalln(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Bad File"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Fatalln("Bad file read")
		log.Fatalln(err)
		err := json.NewEncoder(w).Encode(Errors.Error["Bad File"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
		}
		return
	}
	defer file.Close()
	// Tacking cookie of current user
	claims, err := AdvCookie.GetClaims(r)
	if err != nil {
		err := json.NewEncoder(w).Encode(Errors.Error["Bad cookies"])
		if err != nil {
			log.Fatalln("Can't sand response")
			log.Fatalln(err)
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
