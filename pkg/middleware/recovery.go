package middleware

import (
	"github.com/gin-gonic/gin"
)

func (instance *Middlewares) Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if instance.Logger != nil {
				instance.Logger.Error(err)
				instance.Logger.Println("PANIC", err)
			}
		}
	}()
	c.Next()
}
