package main_test

import (
	"."
	"encoding/json"
	//"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var testInstance main.App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	testInstance.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	testInstance = main.App{}
	testInstance.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestLeaderboard(t *testing.T) {
	req, _ := http.NewRequest("GET", "/leaderboard?offset=2", nil)
	response := executeRequest(req)
	decoder := json.NewDecoder(response.Body)
	var sortedUsers []main.User
	err := decoder.Decode(&sortedUsers)
	if err != nil {
		t.Errorf("Trouble with decoding: %s", err)
	}

	if len(sortedUsers) != 2 {
		t.Errorf("Wrong len. Expected 2, got %d", len(sortedUsers))
	}
	if sortedUsers[0].Points < sortedUsers[1].Points {
		t.Error("Expexted sortes list of users ( 2 users )")
	}
	checkResponseCode(t, http.StatusOK, response.Code)
	main.Users = []main.User{}
}

func TestSignup(t *testing.T) {
	reqGet, _ := http.NewRequest("GET", "/signup", nil)
	reqPost, _ := http.NewRequest("POST", "/signup", strings.NewReader(`{"nickname":"tested","email":"tested@gmail.com","password":"qqq"}`))
	responsePost := executeRequest(reqPost)
	responseGet := executeRequest(reqGet)
	checkResponseCode(t, http.StatusOK, responsePost.Code)
	checkResponseCode(t, http.StatusNotFound, responseGet.Code)
	usersCreated := false
	for _, user := range main.Users {
		if user.Nickname == "tested" && user.Email == "tested@gmail.com" {
			usersCreated = true
		}
	}
	if !usersCreated {
		t.Error("New users not created")
	}
}

// ToDo: Set cookie before put request
func TestEditUser(t *testing.T) {
	reqGet, _ := http.NewRequest("GET", "/nickname/evv", nil)
	reqPost, _ := http.NewRequest("put", "/nickname/tested", strings.NewReader(`{"Age": 25}`))

	//existUser, _:= http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	//responseExist := executeRequest(existUser)
	//authGet.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	//authResp := executeRequest(authGet)

	//reqPut, _ := http.NewRequest("PUT", "/nickname/tested" , strings.NewReader(`{"Age": 25}`))
	responsePost := executeRequest(reqPost)
	responseGet := executeRequest(reqGet)
	//responsePut := executeRequest(reqPut)
	checkResponseCode(t, http.StatusNotFound, responsePost.Code)
	checkResponseCode(t, http.StatusNotFound, responseGet.Code)
	//fmt.Println(main.Users)

}

func TestLogin(t *testing.T) {
	existUser, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	notExistUser, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"NOTEXIST","password":"NOTEXIST"}`))
	responseNotExist := executeRequest(notExistUser)
	responseExist := executeRequest(existUser)

	compare := strings.Compare(strings.TrimRight(responseNotExist.Body.String(), "\n"), `{"Error":"User dont exist"}`)
	if compare != 0 {
		t.Errorf(`Expected answer: {"Error":"User dont exist"}, got %s`, responseNotExist.Body.String())
	}

	decoder := json.NewDecoder(responseExist.Body)
	var userInfo main.User
	_ = decoder.Decode(&userInfo)
	if userInfo.Nickname != "tested" {
		t.Errorf("Expected nickname for tested user, got %s", responseExist.Body.String())
	}

}

//ToDo: Remove cookie from test )
func TestIsAuth(t *testing.T) {
	nonAuthGet, _ := http.NewRequest("GET", "/isauth", nil)
	responseExist := executeRequest(nonAuthGet)
	//fmt.Println(responseExist.Body)
	if strings.Compare(strings.TrimRight(responseExist.Body.String(), "\n"), `{}`) != 0 {
		t.Errorf("Expected {}, got %s", responseExist.Body.String())
	}

	authGet, _ := http.NewRequest("GET", "/isauth", nil)
	// Login
	existUser, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	responseExist = executeRequest(existUser)
	authGet.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	authResp := executeRequest(authGet)
	if strings.Compare(strings.TrimRight(authResp.Body.String(), "\n\t"), `{"is_auth":true}`) != 0 {
		t.Errorf(`Expected auth is {"is_auth":true}, got !%s!`, `{"is_auth":true}`)
	}

	//nonAuthGet, _ := http.NewRequest("GET", "/isauth", nil)
	//authGet, _ := http.NewRequest("GET", "/isauth", nil)
	//existUser, _:= http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	//responseExist := executeRequest(existUser)
	//authGet.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	//nonAuthResp := executeRequest(nonAuthGet)
	//authResp := executeRequest(authGet)
	//fmt.Println(nonAuthResp.Body)
	//fmt.Println(authResp.Body)

}

func TestLogout(t *testing.T) {
	authGet, _ := http.NewRequest("GET", "/isauth", nil)
	// Login
	existUser, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	responseExist := executeRequest(existUser)
	authGet.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	authResp := executeRequest(authGet)
	if strings.Compare(strings.TrimRight(authResp.Body.String(), "\n\t"), `{"is_auth":true}`) != 0 {
		t.Errorf(`Expected auth is {"is_auth":true}, got !%s!`, authResp.Body.String())
	}

	logout, _ := http.NewRequest("GET", "/logout", nil)
	logout.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	afLogResp := executeRequest(logout)
	if !strings.Contains(afLogResp.Header()["Set-Cookie"][0], "session_id=;") {
		t.Errorf("Session id not deleted %s", afLogResp.Header())
	}
}

func TestGetMe(t *testing.T) {
	authGet, _ := http.NewRequest("GET", "/isauth", nil)
	// Login
	existUser, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	responseExist := executeRequest(existUser)
	authGet.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	//r := executeRequest(authGet)
	//fmt.Println(r.Body)
	getMe, _ := http.NewRequest("GET", "/me", nil)
	getMe.Header.Set("Cookie", responseExist.Header()["Set-Cookie"][0])
	responseExist = executeRequest(getMe)

	decoder := json.NewDecoder(responseExist.Body)
	var userInfo main.User
	_ = decoder.Decode(&userInfo)

	if userInfo.Nickname != "tested" {
		t.Errorf("Expected nickname for tested user, got %s", responseExist.Body.String())
	}

}

func TestUpload(t *testing.T) {
	uploadGet, _ := http.NewRequest("GET", "/upload", nil)
	uploadPost, _ := http.NewRequest("POST", "/upload", strings.NewReader(`invalid`))
	responseGet := executeRequest(uploadGet)
	responsePost := executeRequest(uploadPost)

	if strings.Compare(responsePost.Body.String(), "") != 0 {
		t.Errorf(`Expected auth is empty sting, got %s`, responsePost.Body.String())
	}

	checkResponseCode(t, http.StatusNotFound, responseGet.Code)
	checkResponseCode(t, http.StatusOK, responsePost.Code)
}
