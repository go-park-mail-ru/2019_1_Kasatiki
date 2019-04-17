package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mocked(c *gin.Context) {
	c.Status(200)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGoodAuth(t *testing.T) {
	mid := Middlewares{}
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAxfQ.XHkfXuuU5H5i-vW6j2GVgQUjB2FELzTQff0OcKd7gbY; path=/; domain=192.168.100.32; Secure; HttpOnly; Expires=Sun, 14 Apr 2019 23:28:13 GMT;")
	c.Request = req
	r.GET("/test", mid.AuthMiddleware(mocked))
	r.ServeHTTP(resp, c.Request)
	checkResponseCode(t, 200, resp.Code)
}

func TestBadCookieName(t *testing.T) {
	mid := Middlewares{}
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Cookie", "session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAxfQ.XHkfXuuU5H5i-vW6j2GVgQUjB2FELzTQff0OcKd7gbY; path=/; domain=192.168.100.32; Secure; HttpOnly; Expires=Sun, 14 Apr 2019 23:28:13 GMT;")
	c.Request = req
	r.GET("/test", mid.AuthMiddleware(mocked))
	r.ServeHTTP(resp, c.Request)
	checkResponseCode(t, 404, resp.Code)
}

func TestBadCookie(t *testing.T) {
	mid := Middlewares{}
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Cookie", "session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAxfQ.XHkfXuuU5H5i-vW6j2GVgQUjB2FELzTQff0OcKd0gbY; path=/; domain=192.168.100.32; Secure; HttpOnly; Expires=Sun, 14 Apr 2019 23:28:13 GMT;")
	c.Request = req
	r.GET("/test", mid.AuthMiddleware(mocked))
	r.ServeHTTP(resp, c.Request)
	checkResponseCode(t, 404, resp.Code)
}
