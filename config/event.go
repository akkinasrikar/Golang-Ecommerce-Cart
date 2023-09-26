package config

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaEvent "github.com/akkinasrikar/ecommerce-cart/kafka"

)

func StartKafkaProducer(ctx context.Context, cfg KafkaConfig) kafkaEvent.Producer {
	conf := make(kafka.ConfigMap)
	if err := conf.SetKey("bootstrap.servers", cfg.ProducerBootstrapServers); err != nil {
		log.Fatal(err)
	}
	if err := conf.SetKey("security.protocol", cfg.ProducerSecurityProtocol); err != nil {
		log.Fatal(err)
	}
	if err := conf.SetKey("sasl.mechanisms", cfg.ProducerSASLMechanisms); err != nil {
		log.Fatal(err)
	}
	if err := conf.SetKey("sasl.username", cfg.ProducerSASLUsername); err != nil {
		log.Fatal(err)
	}
	if err := conf.SetKey("sasl.password", cfg.ProducerSASLPassword); err != nil {
		log.Fatal(err)
	}
	producerConfig := &kafkaEvent.ProducerConfig{
		ConfigMap:       &conf,
		ProducerTopic:   cfg.ProducerTopic,
		NumberOfWorkers: cfg.ProducerNumberOfWorkers,
		BufferSize:      cfg.ProducerBufferSize,
		FlushTimeout:    cfg.ProducerFlushTimeout,
	}

	producer := kafkaEvent.NewProducer(producerConfig)
	producer.Start(ctx)
	return producer
}
