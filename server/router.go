package server

import (
	"github.com/akkinasrikar/ecommerce-cart/controllers"
	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/middleware"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	servicesLogin "github.com/akkinasrikar/ecommerce-cart/services/login"
	validatorsLogin "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func setUpRoutes(router *gin.Engine, db database.DB, redisClient *redis.Client) {
	ecomStore := repositories.NewRepository(db)
	servicesLogin := servicesLogin.NewLoginService(ecomStore, redisClient)
	LoginHandler := controllers.NewLoginHandler(servicesLogin, validatorsLogin.NewValidator(), ecomStore)
	loginHandler(router, *LoginHandler)
}

func loginHandler(router *gin.Engine, LoginHandler controllers.LoginHandler) {
	router.POST("/signup", LoginHandler.SignUp)
	router.POST("/login", LoginHandler.Login)

	router.Use(middleware.ValidateJwtAuthToken())
	router.Use((middleware.TraceIDMiddleware()))
	router.GET("/homePage", LoginHandler.HomePage)
}
