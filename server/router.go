package server

import (
	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
