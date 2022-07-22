package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		nowTime := time.Now()
		c.Next()
		contTime := time.Since(nowTime)
		log.Println("time const %s", contTime)
	}
}
