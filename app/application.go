package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Started Ecommerce BackEnd Server!")
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	mapUrls()
	router.Run(":8080")
}