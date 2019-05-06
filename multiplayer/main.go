package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gameService := gin.New()
	gameService.Use(gin.Logger())
	gameService.Use(gin.Recovery())
	upgrader := NewConnUpgrader()
	// Создаем Лобби
	Lobby := NewLobby()
	// Запускаем горутину в котрой заполянем комнаты
	go Lobby.Run(upgrader.Queue)

	// Преобразование HTTP запроса в ws
	gameService.Use(static.Serve("/", static.LocalFile("../static/", true)))
	gameService.GET("/game/start", upgrader.StartGame)
	gameService.Run(":8080")
}
