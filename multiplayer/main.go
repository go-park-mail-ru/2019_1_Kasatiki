package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
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
	gameService.Use(static.Serve("/", static.LocalFile("/home/gel0/Desktop/multiplayer/static/", true)))
	gameService.GET("/game/start", upgrader.StartGame)
	gameService.Run(":8080")
	//log.Print(http.ListenAndServe("0.0.0.0:8082", nil))
}
