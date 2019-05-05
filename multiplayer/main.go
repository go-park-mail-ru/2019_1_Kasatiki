package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gameService := gin.New()
	gameService.Use(gin.Logger())

	upgrader := NewConnUpgrader()
	// Создаем Лобби
	Lobby := NewLobby()
	// Запускаем горутину в котрой заполянем комнаты
	go Lobby.Run(upgrader.Queue)

	// Преобразование HTTP запроса в ws
	gameService.GET("/game/start", upgrader.StartGame)
	gameService.GET("/", WebSocketTestPage)
	gameService.Run(":8080")
	//log.Print(http.ListenAndServe("0.0.0.0:8082", nil))
}
