package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gopkg.in/olahol/melody.v1"
	"io/ioutil"
	"net/http"
	"os"
)
import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"models"
)

var Users []models.User

type App struct {
	Router     *gin.Engine
	Connection *pgx.Conn
	Logger     *logrus.Logger
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		fmt.Println("!!!")
		claims, err := checkAuth(cookie)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			return
		}
		id := claims["id"].(float64)
		c.Set("id", id)
		c.Next()
	}
}

func (instance *App) initializeRoutes() {

	loggerFilename := "logfile.log"
	loggerFile, err := os.OpenFile(loggerFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	log := logrus.New()
	log.SetOutput(loggerFile)
	instance.Logger = log

	m := melody.New()
	instance.Router.Use(instance.LoggerMiddleware)
	instance.Router.Use(gin.Recovery())
	//instance.Router.Use(CORSMiddleware)

	api := instance.Router.Group("/api")
	{
		auth := api.Use(AuthMiddleware())
		{
			auth.GET("/isauth", instance.isAuth)
			auth.DELETE("/logout", instance.logout)

		}
		api.GET("/leaderboard", instance.getLeaderboard)
		//api.
		//api.GET("/me", instance.getMe)

		// POST ( create new data )
		api.POST("/signup", instance.createUser)
		api.POST("/upload", instance.upload)
		api.POST("/login", instance.login)

		// PUT ( update data )
		api.PUT("/edit", instance.editUser)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	instance.Connection = conn
	return err
}

func (instance *App) Initialize() {
	err := instance.GetDBConnection()
	err = instance.CreateTables()
	fmt.Println(err)
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
