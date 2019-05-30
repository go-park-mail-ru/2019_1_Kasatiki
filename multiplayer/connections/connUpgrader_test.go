package connections

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
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

func TestConnUpgrader_StartGame(t *testing.T) {
	c := NewConnUpgrader()
	resp1 := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(resp1)
	req, _ := http.NewRequest("GET", "/test", nil)

	ctx.Request = req
	c.StartGame(ctx)
	checkResponseCode(t, 404, resp1.Code)

	resp2 := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	w := http.ResponseWriter()
	ctx2, _ := gin.CreateTestContext()
	req2, _ := http.NewRequest("GET", "/test", nil)
	//req1.Header.Set("Content-type", "application/json")

	req2.Header.Set("Host", "0.0.0.0:8080")
	req2.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
	req2.Header.Set("Accept", "*/*")
	req2.Header.Set("Origin", "http://0.0.0.0:8080")
	req2.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req2.Header.Set("Accept-Encoding", "gzip, deflate")
	req2.Header.Set("Sec-WebSocket-Version", "13")

	req2.Header.Set("Sec-WebSocket-Extensions", "permessage-deflate")
	req2.Header.Set("Sec-WebSocket-Key", "gsxP/sLa0cNlHejBy7SyVg==")
	req2.Header.Set("Connection", "keep-alive, Upgrade")
	req2.Header.Set("Cookie", "session_id=4290630630")
	req2.Header.Set("Pragma", "no-cache")
	req2.Header.Set("Cache-Control", "no-cache")
	req2.Header.Set("Upgrade", "websocket")

	ctx2.Request = req2

	c.StartGame(ctx2)
	checkResponseCode(t, 400, resp2.Code)

}
