package kafka

import (
	"context"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer interface {
	Start(ctx context.Context) error
	Stop() error
	Publish(topic string, message interface{}) error
}

type kafkaProducer struct {
	config   *ProducerConfig
	jobs     chan interface{}
	wg       sync.WaitGroup
	producer *kafka.Producer
}

type ProducerConfig struct {
	BufferSize      int
	NumberOfWorkers int
	FlushTimeout    int
	*kafka.ConfigMap
}

func NewProducer(config *ProducerConfig) Producer {
	return &kafkaProducer{ config: config }
}

func (kf *kafkaProducer) Start(ctx context.Context) error {
	
}
