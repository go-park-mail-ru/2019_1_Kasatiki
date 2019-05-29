package middleware

import "github.com/gin-gonic/gin"

func (instance *Middlewares) CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "www.advhater.ru")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Kasatiki-X-28")
	c.Next()
}
