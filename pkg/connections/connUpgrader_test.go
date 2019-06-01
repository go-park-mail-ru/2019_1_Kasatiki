package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestNewConnUpgrader(t *testing.T) {
	c := NewConnUpgrader()
	if c == nil {
		t.Errorf("Failed connUpgrader creating")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestConnUpgrader_StartGameBadCookie(t *testing.T) {
	c := NewConnUpgrader()
	resp1 := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(resp1)
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Cookie", "sessin_id=4290630630")
	ctx.Request = req
	c.StartGame(ctx.Writer, ctx.Request)
	checkResponseCode(t, 404, resp1.Code)
}

// Успешная работа фунции StartGame
func TestConnUpgrader_StartGame(t *testing.T) {
	c := NewConnUpgrader()
	s := httptest.NewServer(http.HandlerFunc(c.StartGame))
	defer s.Close()
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	h := http.Header{}
	h.Set("Host", "0.0.0.0:8080")
	h.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
	h.Set("Accept", "*/*")
	h.Set("Origin", "http://0.0.0.0:8080")
	h.Set("Accept-Language", "en-US,en;q=0.5")
	h.Set("Accept-Encoding", "gzip, deflate")
	h.Set("Cookie", "session_id=429945o4f")
	h.Set("Pragma", "no-cache")
	h.Set("Cache-Control", "no-cache")
	ws, response, err := websocket.DefaultDialer.Dial(u, h)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()
	status, _ := strconv.Atoi(response.Status[0:3])
	checkResponseCode(t, 101, status)
}
