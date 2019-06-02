package server

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/connections"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/dbhandler"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/lobby"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"os"
)

var Users []models.User

type App struct {
	Router     *gin.Engine
	DB         *dbhandler.DBHandler
	Middleware *middleware.Middlewares
	Upgrader   *connections.ConnUpgrader
	Lobby      *lobby.Lobby
}

func (instance *App) initializeRoutes() {
	instance.Router.Use(instance.Middleware.Recovery)
	instance.Router.Use(instance.Middleware.LoggerMiddleware)
	instance.Router.Use(instance.Middleware.CORSMiddleware)

	//instance.Router.Use(gin.Recovery())
	//instance.Router.Use(gin.Logger())

	//m := melody.New()
	api := instance.Router.Group("/api")
	{
		api.DELETE("/logout", instance.Middleware.AuthMiddleware(instance.logout))

		api.GET("/leaderboard", instance.getLeaderboard)
		api.GET("/isauth", instance.Middleware.AuthMiddleware(instance.isAuth))
		api.GET("/balance", instance.Middleware.AuthMiddleware(instance.balance))

		// POST ( create new data )
		api.POST("/signup", instance.createUser)
		api.POST("/login", instance.login)
		api.POST("/upload", instance.Middleware.AuthMiddleware(instance.upload))
		api.POST("/payments", instance.Middleware.AuthMiddleware(instance.payout))
		api.POST("/checkPointHandler", instance.Middleware.AuthMiddleware(instance.checkPoints))
		// PUT ( update data )
		api.PUT("/edit", instance.Middleware.AuthMiddleware(instance.editUser))

		api.GET("/game/start", gin.WrapF(instance.Upgrader.StartGame))

		//api.GET("/ws", func(c *gin.Context) {
		//	m.HandleRequest(c.Writer, c.Request)
		//})

		//m.HandleMessage(func(s *melody.Session, msg []byte) {
		//	m.Broadcast(msg)
		//})

	}

	//Static path

	instance.Router.Use(static.Serve("/", static.LocalFile("../../static", true)))
	// Echo websocket for test

}

func (instance *App) Run(port string) (err error) {
	err = instance.Router.Run()
	return
}

// Todo Обернуть в конфиг
func (instance *App) GetDBConnection(config *models.Config) error {

	conf := pgx.ConnConfig{
		User:      config.DBUser,
		Password:  config.DBPassword,
		Host:      config.DBHost,
		Port:      config.DBPort,
		Database:  config.DBSpace,
		TLSConfig: nil,
	}
	confPool := pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: 16,
	}
	conn, err := pgx.NewConnPool(confPool)
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

func (instance *App) Initialize(conf *models.Config) {
	err := instance.GetDBConnection(conf)
	err = instance.DB.CreateTables()
	fmt.Println(err)
	err = instance.DB.CreateAdvTable()
	fmt.Println(err)
	instance.DB.AdvsIserting()
	instance.Router = gin.New()
	instance.Upgrader = connections.NewConnUpgrader()
	instance.Lobby = lobby.NewLobby()
	go instance.Lobby.Run(instance.Upgrader.Queue)
	instance.initializeRoutes()
}
