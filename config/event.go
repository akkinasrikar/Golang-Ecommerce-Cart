package config

import (
	"context"
	"log"

	kafkaEvent "github.com/akkinasrikar/ecommerce-cart/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hibiken/asynq"
)

func StartKafkaProducer(ctx context.Context, cfg KafkaConfig, asynq *asynq.Client) kafkaEvent.Producer {
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
	if err := conf.SetKey("group.id", cfg.ProducerGroupId); err != nil {
		log.Fatal(err)
	}
	if err := conf.SetKey("auto.offset.reset", cfg.ProducerOffsetReset); err != nil {
		log.Fatal(err)
	}
	producerConfig := &kafkaEvent.ProducerConfig{
		ConfigMap:       &conf,
		ProducerTopic:   cfg.ProducerTopic,
		NumberOfWorkers: cfg.ProducerNumberOfWorkers,
		BufferSize:      cfg.ProducerBufferSize,
	}

	producer := kafkaEvent.NewProducer(producerConfig, asynq)
	producer.Start(ctx)
	return producer
}
