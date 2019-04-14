package main

import (
	"github.com/gin-contrib/static"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gopkg.in/olahol/melody.v1"
)
import (
	"github.com/gin-gonic/gin"
	"log"
	"models"
)

var Users []models.User

type App struct {
	Router *gin.Engine
}

func (instance *App) initializeRoutes() {

	m := melody.New()
	// GET ( get exist data )
	instance.Router.GET("/api/leaderboard", instance.getLeaderboard)
	instance.Router.GET("/api/isauth", instance.isAuth)
	instance.Router.GET("/api/me", instance.getMe)
	instance.Router.GET("/api/logout", instance.logout)

	// POST ( create new data )
	instance.Router.POST("/api/signup", instance.createUser)
	instance.Router.POST("/api/upload", instance.upload)
	instance.Router.POST("/api/login", instance.login)

	// PUT ( update data )
	instance.Router.PUT("/api/users/{Nickname}", instance.editUser)
	instance.Router.GET("/api/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	instance.Router.GET("/api/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	//Static path
	instance.Router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Echo websocket for test

}

func (instance *App) Run(port string) {
	log.Fatal(instance.Router.Run()) // ToDO change logFatal?
}

func (instance *App) Initialize() {
	var mockedUser = models.User{"1", "evv", "onetaker@gmail.com",
		"evv", -100, 23, "test",
		"Voronezh", "В левой руке салам"}
	var mockedUser1 = models.User{"2", "tony", "trendpusher@hydra.com",
		"qwerty", 100, 22, "test",
		"Moscow", "В правой алейкум"}
	// Mocked users

	Users = append(Users, mockedUser)
	Users = append(Users, mockedUser1)
	instance.Router = gin.New()
	instance.initializeRoutes()
}

type Order struct {
	Sequence string `json:"order"`
}

//ToDo: Move to another package
var errorLogin = map[string]string{
	"Error": "User dont exist",
}

var errorCreateUser = map[string]string{
	"Error": "Nickname/mail already exists",
}
