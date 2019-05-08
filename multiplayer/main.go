package main

import (
	"2019_1_Kasatiki/multiplayer/connections"
	"2019_1_Kasatiki/multiplayer/lobby"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	gameService := gin.New()
	gameService.Use(gin.Logger())
	gameService.Use(gin.Recovery())
	upgrader := connections.NewConnUpgrader()
	// Создаем Лобби
	Lobby := lobby.NewLobby()
	// Запускаем горутину в котрой заполянем комнаты
	go Lobby.Run(upgrader.Queue)

	// Преобразование HTTP запроса в ws
	gameService.Use(static.Serve("/", static.LocalFile("../static/", true)))
	gameService.GET("/game/start", upgrader.StartGame)
	gameService.Run(":8080")
}
