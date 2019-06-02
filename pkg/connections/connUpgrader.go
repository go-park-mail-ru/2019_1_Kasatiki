package connections

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type UserConnection struct {
	Login      string
	Token      string
	Connection *websocket.Conn
	TypeGame   string
}

type ConnUpgrader struct {
	// Настройки WebSocket.
	upgrader websocket.Upgrader
	// Канал-очередь, для добавления игроков в Lobby
	Queue chan *UserConnection
}

// Создание нового апгрейдера
func NewConnUpgrader() (cu *ConnUpgrader) {
	cu = &ConnUpgrader{
		upgrader: websocket.Upgrader{

			HandshakeTimeout: time.Duration(1 * time.Second),
			CheckOrigin: func(r *http.Request) bool { // Токен не проверяется.
				return true
			},
			//EnableCompression: true,
		},

		Queue: make(chan *UserConnection, 50),
	}
	return
}

var debug = true

func (up *ConnUpgrader) StartGame(res http.ResponseWriter, req *http.Request) {
	log.Printf("New connection: %#v", req)
	// Проверяет SessionId из cookie.
	fmt.Println(req.Cookies())
	sessionID, err := req.Cookie("session_id")
	if err != nil {
		fmt.Println(err)
		http.Error(res, "wrong cookie", 404)
		return
	}
	var login string
	if debug {
		// просто создаёт случайный логин
		login = "Anon" + time.Now().Format(time.RFC3339)
	}

	// Меняет протокол.
	WSConnection, err := up.upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(409)
		//c.JSON(409, "error of creating WS")
		return
	}
	//typeGame := req.Header.Get("Mode")
	// Todo убрать хардкод
	//if typeGame == "" {
	//	res.WriteHeader(400)
	//	return
	//}
	connection := &UserConnection{
		Login:      login,
		Token:      sessionID.Value,
		Connection: WSConnection,
		//TypeGame:   "Multiplayer",
		TypeGame: "sp",
	}

	up.Queue <- connection
	return
}
