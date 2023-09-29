package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/database"
	"github.com/akkinasrikar/ecommerce-cart/models"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/repositories"
	services "github.com/akkinasrikar/ecommerce-cart/services/products"
	"github.com/akkinasrikar/ecommerce-cart/worker/cron"
	"github.com/go-co-op/gocron"
	redislock "github.com/go-co-op/gocron-redis-lock"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"default": 3,
			},
		},
	)

	db := database.ConnectDataBase()
	dbStore := database.NewDb(db)
	ecomStore := repositories.NewRepository(dbStore)
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	productService := services.NewAsynqService(ecomStore, asynqClient)

	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{redisAddr},
	})

	locker, err := redislock.NewRedisLocker(redisClient, redislock.WithTries(1))
	if err != nil {
		log.Fatal(err)
	}
	s := gocron.NewScheduler(time.Local)
	ctx := context.Background()
	s.WithDistributedLocker(locker)
	go cron.Start(ctx, s, productService)
	defer s.Stop()

	mux := asynq.NewServeMux()
	mux.Use(func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			err := h.ProcessTask(ctx, t)
			if err != nil {
				log.Printf("error processing task: %v", err)
			}
			return err
		})
	})
	mux.HandleFunc(constants.ProcessTasks.IMAGERESIZE, func(ctx context.Context, t *asynq.Task) error {
		task := models.ImageResize{}
		err := json.Unmarshal(t.Payload(), &task)
		if err != nil {
			return err
		}
		return productService.ImageResize(ctx, task)
	})
	mux.HandleFunc(constants.ProcessTasks.CONSUMEDATA, func(ctx context.Context, t *asynq.Task) error {
		task := entities.Consume{}
		err := json.Unmarshal(t.Payload(), &task)
		if err != nil {
			return err
		}
		return productService.SubscribeDataFromKafka(ctx, task)
	})
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
