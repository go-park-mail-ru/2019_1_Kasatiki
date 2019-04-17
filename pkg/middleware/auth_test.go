package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func mocked(c *gin.Context) {
	fmt.Println("Chain")
}
