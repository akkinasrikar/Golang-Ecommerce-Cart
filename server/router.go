package server

import (
	"github.com/akkinasrikar/ecommerce-cart/api"
	"github.com/akkinasrikar/ecommerce-cart/controllers"
	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/middleware"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	servicesLogin "github.com/akkinasrikar/ecommerce-cart/services/login"
	services "github.com/akkinasrikar/ecommerce-cart/services/products"
	validator "github.com/akkinasrikar/ecommerce-cart/validators"
	validatorsLogin "github.com/akkinasrikar/ecommerce-cart/validators/login"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func setUpRoutes(router *gin.Engine, db database.DB, redisClient *redis.Client) {
	ecomStore := repositories.NewRepository(db)
	validatorsLogin := validatorsLogin.NewValidator()

	servicesLogin := servicesLogin.NewLoginService(ecomStore, redisClient)
	LoginHandler := controllers.NewLoginHandler(servicesLogin, validatorsLogin, ecomStore)
	loginHandler(router, *LoginHandler)

	apiServices := api.NewService()
	validatorServices := validator.NewValidator(ecomStore)
	ecomServices := services.NewService(apiServices, ecomStore)
	ecomHandler := controllers.NewProductHandler(validatorServices, ecomServices)
	productHandler(router, *ecomHandler)
}

func loginHandler(router *gin.Engine, LoginHandler controllers.LoginHandler) {
	router.POST("/signup", LoginHandler.SignUp)
	router.POST("/login", LoginHandler.Login)

	router.Use(middleware.ValidateJwtAuthToken())
	router.Use((middleware.TraceIDMiddleware()))
	router.GET("/homePage", LoginHandler.HomePage)
}

func productHandler(router *gin.Engine, ecomHandler controllers.ProductHandler) {
	router.Use(middleware.ValidateJwtAuthToken())
	router.Use((middleware.TraceIDMiddleware()))
	router.GET("/user", ecomHandler.GetUserDetails)
	router.GET("/products", ecomHandler.GetProducts)
	router.GET("/products-by-id", ecomHandler.GetProductById)
	router.POST("/card-details", ecomHandler.CardDetails)
	router.GET("/card-details", ecomHandler.GetCardDetails)
	router.POST("/add-address", ecomHandler.AddAddress)
	router.GET("/get-address", ecomHandler.GetAddress)
	router.POST("/cart", ecomHandler.AddOrDeleteToCart)
	router.GET("/cart", ecomHandler.GetProductsFromCart)
	router.POST("/order", ecomHandler.OrderProducts)
}
