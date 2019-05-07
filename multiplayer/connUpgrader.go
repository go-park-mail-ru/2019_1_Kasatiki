package main

import (
	"github.com/gin-gonic/gin"
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

func (up *ConnUpgrader) StartGame(c *gin.Context) {
	log.Printf("New connection: %#v", c.Request)
	// Проверяет SessionId из cookie.
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(404, "error of cookie")
		return
	}

	//
	var login string
	if debug {
		// просто создаёт случайный логин
		login = "Anon" + time.Now().Format(time.RFC3339)
	}

	// Меняет протокол.
	WSConnection, err := up.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(409, "error of creating WS")
		return
	}

	// Todo убрать хардкод
	connection := &UserConnection{
		Login:      login,
		Token:      sessionID,
		Connection: WSConnection,
		TypeGame:   "Multiplayer",
	}

	up.Queue <- connection
	return
}
