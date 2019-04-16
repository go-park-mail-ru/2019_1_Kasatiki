package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/dbhandler"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/middleware"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
	"io/ioutil"
	"net/http"
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

func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "www.advhater.ru")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.Next()
}

func checkAuth(cookie *http.Cookie) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		spiceSalt, _ := ioutil.ReadFile("secret.conf")
		return spiceSalt, nil
	})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	// ToDo: Handle else case
	return claims, nil
}

func AuthMiddleware(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil {

			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		claims, err := checkAuth(cookie)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		id := claims["id"].(float64)
		c.Set("id", id)
		fmt.Println(c.Get("id"))
		handlerFunc(c)
	}
}

func (instance *App) initializeRoutes() {

	m := melody.New()
	instance.Router.Use(instance.Middleware.LoggerMiddleware)
	instance.Router.Use(gin.Recovery())
	instance.Router.Use(CORSMiddleware)

	api := instance.Router.Group("/api")
	{
		api.DELETE("/logout", AuthMiddleware(instance.logout))

		api.GET("/leaderboard", instance.getLeaderboard)
		api.GET("/isauth", AuthMiddleware(instance.isAuth))

		// POST ( create new data )
		api.POST("/signup", instance.createUser)
		api.POST("/login", instance.login)
		api.POST("/upload", AuthMiddleware(instance.upload))

		// PUT ( update data )
		api.PUT("/edit", AuthMiddleware(instance.editUser))

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
