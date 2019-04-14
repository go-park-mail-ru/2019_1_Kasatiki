package main

import (
	"github.com/jackc/pgx"
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
	Router     *gin.Engine
	Connection *pgx.Conn
}

func CORSMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "www.advhater.ru")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Next()
}

func AuthMiddleware(c *gin.Context) {

}

func (instance *App) initializeRoutes() {

	instance.Router.Use(gin.Logger())
	instance.Router.Use(gin.Recovery())
	instance.Router.Use(CORSMiddleware)

	// GET ( get exist data )
	api := instance.Router.Group("/api")
	{
		api.GET("/leaderboard", instance.getLeaderboard)
		api.GET("/isauth", instance.isAuth)
		api.GET("/me", instance.getMe)
		api.GET("/logout", instance.logout)

		// POST ( create new data )
		api.POST("/signup", instance.createUser)
		api.POST("/upload", instance.upload)
		api.POST("/login", instance.login)

		// PUT ( update data )
		api.PUT("/users/{Nickname}", instance.editUser)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//Static path
	instance.Router.Use(static.Serve("/", static.LocalFile("./static", true)))
}

func (instance *App) Run(port string) {
	log.Fatal(instance.Router.Run()) // ToDO change logFatal?
}

func (instance *App) GetDBConnection() error {
	conf := pgx.ConnConfig{
		User:      "sayonara",
		Password:  "boy",
		Host:      "localhost",
		Port:      5432,
		Database:  "Kasatiki",
		TLSConfig: nil,
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		return err
	}
	instance.Connection = conn
	return err
}

func (instance *App) Initialize() {
	_ = instance.GetDBConnection()
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
