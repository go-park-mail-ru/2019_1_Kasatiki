package main

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)
import (
	"models"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

var Users []models.User

type App struct {
	Router *gin.Engine
}

func (instance *App) initializeRoutes() {

	// GET ( get exist data )
	instance.Router.GET("/leaderboard", instance.getLeaderboard)
	instance.Router.GET("/isauth", instance.isAuth)
	instance.Router.GET("/me", instance.getMe)
	instance.Router.GET("/logout", instance.logout)

	// POST ( create new data )
	instance.Router.POST("/signup", instance.createUser)
	instance.Router.POST("/upload", instance.upload)
	instance.Router.POST("/login", instance.login)

	// PUT ( update data )
	instance.Router.PUT("/users/{Nickname}", instance.editUser)
	instance.Router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Static path
	instance.Router.Use(static.Serve("/", static.LocalFile("./static", true)))
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
