package server

import (
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
	setUpRoutes(router, db, redisClient)
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
