package connections

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type UserConnection struct {
	Login      string
	Token      string
	Connection *websocket.Conn
	TypeGame   string
	Id         int
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
	// Проверяет SessionId из cookie.
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		fmt.Println("Ошибка", err)
		c.JSON(404, "Bad cookie")
		return
	}

	gettinId, _ := c.Get("id")
	// Меняет протокол.
	WSConnection, err := up.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(409)
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
		Id:         int(gettinId.(float64)),
		Token:      sessionID,
		Connection: WSConnection,
		//TypeGame:   "Multiplayer",
		TypeGame: "sp",
	}

	up.Queue <- connection
	return
}
