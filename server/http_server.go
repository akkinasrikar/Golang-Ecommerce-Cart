package server

import (
	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type HttpServer struct {
	Router *gin.Engine
}

func Init(db *gorm.DB) (*HttpServer, error) {
	server := &HttpServer{}
	router := gin.Default()
	server.Router = router
	redisClient := InitRedisCache()
	// create postgres db connection database.DB
	dbStore := database.NewDb(db)
	// check if data is seeded
	if !database.IsDataSeeded(dbStore) {
		database.SeedData(dbStore)
	}
	setUpRoutes(router, dbStore, redisClient)
	return server, nil
}

func (s *HttpServer) Start() error {
	s.Router.Run(":8080")
	return nil
}

func InitRedisCache() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
