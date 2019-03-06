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
)

// ToDo бахнуть обработку ошибок

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

//ToDo: Use get with key order? (ASC/DESC )
func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order

	_ = json.NewDecoder(r.Body).Decode(&order)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})

	json.NewEncoder(w).Encode(users)
}

func createSessionId(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": user.Nickname,
		"email":    user.Email,
		"about":    user.About,
		"region":   user.Region,
		"img":      user.ImgUrl,
		"age":      user.Age,
	})
	// ToDo: Error handle
	spiceSalt, _ := ioutil.ReadFile("secret.conf")
	secretStr, _ := token.SignedString(spiceSalt)
	return secretStr
}

func checkAuth(cookie *http.Cookie) jwt.MapClaims {
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

func editUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.Write([]byte("=("))
		return
	}
	claims := checkAuth(cookie)
	w.Write([]byte(claims["email"].(string)))
}

func login(w http.ResponseWriter, r *http.Request) {
	var sessionId string
	var userExistFlag bool
	var existUser User
	userExistFlag = false
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
	sessionId = createSessionId(existUser)
	fmt.Println(sessionId)
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: sessionId,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(existUser)
}

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
	f, err := os.OpenFile("./static/img/"+"test1.jpg", os.O_WRONLY|os.O_CREATE, 0666) // ToDo: Change way to handle img
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	// Mocked part for leaderboard
	var mockedUser = User{"1", "evv", "onetaker@gmail.com",
		"evv", -100, 23, "test",
		"Voronezh", "В левой руке салам"}
	var mockedUser1 = User{"2", "tony", "trendpusher@hydra.com",
		"qwerty", 100, 22, "test",
		"Moscow", "В правой алейкум"}
	// Mocker part end
	users = append(users, mockedUser)
	users = append(users, mockedUser1)
	reciever := mux.NewRouter()
	// GET  ( get exists data )
	reciever.HandleFunc("/users/{Nickname}", getUser).Methods("GET")
	reciever.HandleFunc("/leaderboard", getLeaderboard).Methods("GET")
	reciever.HandleFunc("/edit", editUser).Methods("GET")

	// POST ( create new data )
	reciever.HandleFunc("/signup", createUser).Methods("POST")
	reciever.HandleFunc("/upload", upload).Methods("POST")
	reciever.HandleFunc("/login", login).Methods("POST")

	//reciever.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/"))) // Uncomment if want to run locally
	log.Fatal(http.ListenAndServe(":8080", reciever))
}
