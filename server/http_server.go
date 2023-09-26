package server

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/akkinasrikar/ecommerce-cart/config"
	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/kafka"
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
	producer := InitKafkaProducer()

	// create postgres db connection database.DB
	dbStore := database.NewDb(db)
	// check if data is seeded
	setUpRoutes(router, dbStore, redisClient, producer)
	return server, nil
}

func (s *HttpServer) Start() error {
	s.Router.Run(":8081")
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

func InitKafkaProducer() (kafka.Producer) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	producer := config.StartKafkaProducer(ctx, *config.Kafka)
	return producer
}
