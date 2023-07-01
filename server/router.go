package server

import (
	"github.com/akkinasrikar/ecommerce-cart/controllers"
	servicesLogin "github.com/akkinasrikar/ecommerce-cart/services/login"
	validatorsLogin "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.Engine) {
	LoginHandler := controllers.NewLoginHandler(servicesLogin.NewLoginService(), validatorsLogin.NewValidator())
	loginHandler(router, *LoginHandler)
}

func loginHandler(router *gin.Engine, LoginHandler controllers.LoginHandler) {
	router.POST("/signup", LoginHandler.SignUp)
}
