package room

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/connections"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/game_logic"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoom_GameEngine(t *testing.T) {
	c := connections.NewConnUpgrader()

	// Приконнектился один игрок
	s1 := httptest.NewServer(http.HandlerFunc(c.StartGame))
	defer s1.Close()
	u := "ws" + strings.TrimPrefix(s1.URL, "http")
	h := http.Header{}
	h.Set("Host", "0.0.0.0:8080")
	h.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
	h.Set("Accept", "*/*")
	h.Set("Origin", "http://0.0.0.0:8080")
	h.Set("Accept-Language", "en-US,en;q=0.5")
	h.Set("Accept-Encoding", "gzip, deflate")
	h.Set("Cookie", "session_id=429r06iu30630")
	h.Set("Pragma", "no-cache")
	h.Set("Cache-Control", "no-cache")
	ws, _, err := websocket.DefaultDialer.Dial(u, h)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()
	// Приконнектился второй игрок
	s2 := httptest.NewServer(http.HandlerFunc(c.StartGame))
	defer s2.Close()
	u2 := "ws" + strings.TrimPrefix(s2.URL, "http")
	h2 := http.Header{}
	h2.Set("Host", "0.0.0.0:8080")
	h2.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
	h2.Set("Accept", "*/*")
	h2.Set("Origin", "http://0.0.0.0:8080")
	h2.Set("Accept-Language", "en-US,en;q=0.5")
	h2.Set("Accept-Encoding", "gzip, deflate")
	h2.Set("Cookie", "session_id=42963rhhh0630")
	h2.Set("Pragma", "no-cache")
	h2.Set("Cache-Control", "no-cache")
	ws2, _, err := websocket.DefaultDialer.Dial(u2, h2)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws2.Close()

	testPlayers := make(map[string]*connections.UserConnection, 2)

	select {
	case connection, _ := <-c.Queue:
		testPlayers[connection.Login] = connection
	}
	g, _ := game_logic.GameIni(testPlayers)
	if g == nil {
		t.Fatalf("%v", "Game initialization failed")
	}
	mes := &game_logic.InputMessage{}
	mes.Up = true
	mes.Left = true
	var logins []string
	for k, _ := range testPlayers {
		logins = append(logins, k)
	}
	g.EventListener(*mes, logins[0])

}
