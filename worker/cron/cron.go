package cron

import (
	"context"
	"log"
	"github.com/go-co-op/gocron"
	services "github.com/akkinasrikar/ecommerce-cart/services/products"
)

func Start(ctx context.Context, s *gocron.Scheduler, productService services.ProducAsynqService) {
	log.Println("Starting cron worker")
	
	_, err := s.Every(60).Seconds().Do(func() {
		err := productService.UpdateDeliveryStatus(ctx)
		if err != nil {
			log.Println("Error resizing image", err)
		}
	})
	log.Printf("Scheduled cron job every 10 seconds")
	if err != nil {
		log.Println("Error scheduling cron job", err)
	}

	s.StartAsync()
}
