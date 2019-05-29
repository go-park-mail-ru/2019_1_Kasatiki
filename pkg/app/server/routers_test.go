package server

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var testInstance App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	testInstance.Router.ServeHTTP(recorder, req)
	return recorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	testInstance = App{}
	configBytes, _ := ioutil.ReadFile("../../../cmd/server/config.json")
	config := &models.Config{}
	_ = json.NewDecoder(bytes.NewReader(configBytes)).Decode(&config)
	testInstance.Initialize(config)
	code := m.Run()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	failedUser, _ := http.NewRequest("POST", "/api/login",
		strings.NewReader(`{"nickname":"tested","password":"qqq"}`))
	failedUser.Header.Set("Content-Type", "application/json")
	response_404 := executeRequest(failedUser)

	checkResponseCode(t, http.StatusNotFound, response_404.Code)

	wrongFormat, _ := http.NewRequest("POST", "/api/login",
		strings.NewReader(`sdfsdfsdfsdf`))
	wrongFormat.Header.Set("Content-Type", "application/json")

	response_400 := executeRequest(wrongFormat)
	checkResponseCode(t, http.StatusBadRequest, response_400.Code)

	succesUser, _ := http.NewRequest("POST", "/api/login",
		strings.NewReader(`{"nickname":"XVlBzgbaiC","password":"MRAjWwhTHc"}`))
	succesUser.Header.Set("Content-Type", "application/json")
	response_201 := executeRequest(succesUser)
	checkResponseCode(t, http.StatusCreated, response_201.Code)
}

func TestCreateUser(t *testing.T) {
	decodeFailed, _ := http.NewRequest("POST", "/api/signup",
		strings.NewReader(`{"nickdfsname":"tested","password":"qqq"}`))
	decodeFailed.Header.Set("Content-Type", "application/json")
	response_400 := executeRequest(decodeFailed)
	checkResponseCode(t, http.StatusBadRequest, response_400.Code)

	duplicateUser, _ := http.NewRequest("POST", "/api/signup",
		strings.NewReader(`{"nickname":"XVlBzgbaiC","password":"MRAjWwhTHc", "email":"123"}`))
	duplicateUser.Header.Set("Content-Type", "application/json")
	response_409 := executeRequest(duplicateUser)
	checkResponseCode(t, http.StatusConflict, response_409.Code)

	createdUser, _ := http.NewRequest("POST", "/api/signup",
		strings.NewReader(`{"nickname":"testUser","password":"testUser", "email":"123"}`))
	duplicateUser.Header.Set("Content-Type", "application/json")
	response_201 := executeRequest(createdUser)
	checkResponseCode(t, http.StatusCreated, response_201.Code)
}

func TestLeaderboard(t *testing.T) {
	wrongOffset, _ := http.NewRequest("GET", "/api/leaderboard?offset=qwqwr", nil)
	response_400 := executeRequest(wrongOffset)
	checkResponseCode(t, http.StatusBadRequest, response_400.Code)

	limit, _ := http.NewRequest("GET", "/api/leaderboard?offset=32523532", nil)
	response_404 := executeRequest(limit)
	checkResponseCode(t, http.StatusNotFound, response_404.Code)

	valid, _ := http.NewRequest("GET", "/api/leaderboard?offset=1", nil)
	response_200 := executeRequest(valid)
	checkResponseCode(t, http.StatusOK, response_200.Code)
}

func TestEditUser(t *testing.T) {
	decodeFailed, _ := http.NewRequest("PUT", "/api/edit",
		strings.NewReader(`{"nickdfsname":"tested","password":"qqq"}`))
	decodeFailed.Header.Set("Content-Type", "application/json")
	decodeFailed.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_409 := executeRequest(decodeFailed)
	checkResponseCode(t, http.StatusConflict, response_409.Code)

	nickname, _ := http.NewRequest("PUT", "/api/edit",
		strings.NewReader(`{"nickname":"mBTvKSJfjz","email":"qqq"}`))
	nickname.Header.Set("Content-Type", "application/json")
	nickname.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_300 := executeRequest(nickname)
	checkResponseCode(t, 300, response_300.Code)

	email, _ := http.NewRequest("PUT", "/api/edit",
		strings.NewReader(`{"nickname":"qwrad","email":"LDnJObCsNV@ya.ru"}`))
	email.Header.Set("Content-Type", "application/json")
	email.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_301 := executeRequest(email)
	checkResponseCode(t, 301, response_301.Code)

	sc, _ := http.NewRequest("PUT", "/api/edit",
		strings.NewReader(`{"nickname":"XVlBzgbaiC","password":"qqq"}`))
	sc.Header.Set("Content-Type", "application/json")
	sc.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_200 := executeRequest(sc)
	checkResponseCode(t, http.StatusOK, response_200.Code)
}

func TestIsAuh(t *testing.T) {
	isauth, _ := http.NewRequest("GET", "/api/isauth", nil)
	isauth.Header.Set("Content-Type", "application/json")
	isauth.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_200 := executeRequest(isauth)
	checkResponseCode(t, http.StatusOK, response_200.Code)
}

func TestLogout(t *testing.T) {
	logout, _ := http.NewRequest("DELETE", "/api/logout", nil)
	logout.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_200 := executeRequest(logout)
	checkResponseCode(t, http.StatusOK, response_200.Code)
}

func TestPayout(t *testing.T) {
	decodeFailed, _ := http.NewRequest("POST", "/api/payments",
		strings.NewReader(`{"money": "all"}`))
	decodeFailed.Header.Set("Content-Type", "application/json")
	//decodeFailed.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.XF3YG1m_vtJDN6tfO5iKWgZdpIcFgXnpG_fuDVBn0Uc; path=/; domain=0.0.0.0; HttpOnly; Expires=Wed, 17 Apr 2222 03:38:00 GMT;")
	response_400 := executeRequest(decodeFailed)
	checkResponseCode(t, http.StatusBadRequest, response_400.Code)
}
