package server

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/dbhandler"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/middleware"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
	"os"
)
import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"log"
)

var Users []models.User

type App struct {
	Router     *gin.Engine
	DB         *dbhandler.DBHandler
	Middleware *middleware.Middlewares
}

func (instance *App) initializeRoutes() {

	m := melody.New()
	instance.Router.Use(instance.Middleware.LoggerMiddleware)
	instance.Router.Use(gin.Recovery())
	instance.Router.Use(instance.Middleware.CORSMiddleware)

	api := instance.Router.Group("/api")
	{
		api.DELETE("/logout", instance.Middleware.AuthMiddleware(instance.logout))

		api.GET("/leaderboard", instance.getLeaderboard)
		api.GET("/isauth", instance.Middleware.AuthMiddleware(instance.isAuth))

		// POST ( create new data )
		api.POST("/signup", instance.createUser)
		api.POST("/login", instance.login)
		api.POST("/upload", instance.Middleware.AuthMiddleware(instance.upload))

		// PUT ( update data )
		api.PUT("/edit", instance.Middleware.AuthMiddleware(instance.editUser))

		api.GET("/ws", func(c *gin.Context) {
			m.HandleRequest(c.Writer, c.Request)
		})
		m.HandleMessage(func(s *melody.Session, msg []byte) {
			m.Broadcast(msg)
		})
	}

	//Static path
	instance.Router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Echo websocket for test

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
		Database:  "kasatiki",
		TLSConfig: nil,
	}
	conn, err := pgx.Connect(conf)
	fmt.Print(err)
	if err != nil {
		return err
	}
	instance.DB = &dbhandler.DBHandler{conn}

	loggerFilename := "logs/logfile.log"
	loggerFile, err := os.OpenFile(loggerFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	log := logrus.New()
	log.SetOutput(loggerFile)
	instance.Middleware = &middleware.Middlewares{log}
	fmt.Print(instance.DB)
	return err
}

func (instance *App) Initialize() {
	err := instance.GetDBConnection()
	err = instance.DB.CreateTables()
	fmt.Println(err)
	instance.Router = gin.New()
	instance.initializeRoutes()
}
