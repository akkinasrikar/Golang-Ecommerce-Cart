package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	fmt.Println("Mapping Url's!...")
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
