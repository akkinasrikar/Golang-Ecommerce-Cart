package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/akkinasrikar/ecommerce-cart/constants"
	"github.com/akkinasrikar/ecommerce-cart/models/entities"
	"github.com/akkinasrikar/ecommerce-cart/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hibiken/asynq"
)

type Producer interface {
	Start(ctx context.Context)
	Stop()
	Publish(message interface{})
	Consumer(ctx context.Context)
}

type kafkaProducer struct {
	config   *ProducerConfig
	jobs     chan interface{}
	wg       sync.WaitGroup
	producer *kafka.Producer
	asynq    *asynq.Client
}

type ProducerConfig struct {
	ProducerTopic   string
	BufferSize      int
	NumberOfWorkers int
	FlushTimeout    int
	*kafka.ConfigMap
}

func NewProducer(config *ProducerConfig, asynqClient *asynq.Client) Producer {
	return &kafkaProducer{
		config: config,
		asynq:  asynqClient,
	}
}

func (kf *kafkaProducer) initPools() {
	// create channels fo the jobs
	kf.jobs = make(chan interface{}, kf.config.BufferSize)
	for i := 0; i < kf.config.NumberOfWorkers; i++ {
		kf.wg.Add(1)
		go func(id int) {
			defer kf.wg.Done()
			kf.process(id)
		}(i)
	}
}

func (kf *kafkaProducer) Start(ctx context.Context) {
	kf.initPools()
	kfProducer, err := kafka.NewProducer(kf.config.ConfigMap)
	if err != nil {
		log.Println("failed to create kafka producer")
	}
	kf.producer = kfProducer

	go func() {
		for e := range kf.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("context cancelled")
				return
			case <-sigChan:
				log.Println("terminating: via signal")
				return
			}
		}
	}()
}

func (kf *kafkaProducer) Stop() {
	close(kf.jobs)
	kf.wg.Wait()
	kf.producer.Close()
}

func (pf *kafkaProducer) Publish(message interface{}) {
	pf.jobs <- message
}

func (kf *kafkaProducer) process(id int) {
	for message := range kf.jobs {
		log.Printf("worker %d: received message: %v", id, message)
		kf.processMessage(message)
	}
}

func (kf *kafkaProducer) processMessage(message interface{}) {
	msg, err := json.Marshal(message)
	if err != nil {
		log.Printf("failed to marshal message: %v", err)
		return
	}

	err = kf.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kf.config.ProducerTopic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, nil)

	if err != nil {
		var kafkaError kafka.Error
		if errors.As(err, &kafkaError) && kafkaError.Code() == kafka.ErrQueueFull {
			time.Sleep(time.Second)
			if err := kf.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &kf.config.ProducerTopic,
					Partition: kafka.PartitionAny,
				},
				Value: msg,
			}, nil); err != nil {
				log.Printf("failed to produce message: %v", err)
			}
			return
		}
		log.Printf("failed to produce message: %v", err)
	}
	return
}

func (kf *kafkaProducer) Consumer(ctx context.Context) {
	consumer, err := kafka.NewConsumer(kf.config.ConfigMap)
	if err != nil {
		log.Println("failed to create kafka consumer")
	}
	consumer.SubscribeTopics([]string{kf.config.ProducerTopic}, nil)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("context cancelled")
				return
			default:
				msg, err := consumer.ReadMessage(-1)
				if err != nil {
					log.Printf("failed to read message: %v", err)
					continue
				}
				data := entities.Consume{
					ProcessId:   string(msg.Key),
					ProcessData: string(msg.Value),
					ProcessName: constants.ProcessTasks.CONSUMEDATA,
				}
				// marshal data
				jsonData, err := json.Marshal(data)
				if err != nil {
					log.Printf("failed to marshal message: %v", err)
					continue
				}
				task := asynq.NewTask(constants.ProcessTasks.CONSUMEDATA, jsonData, asynq.TaskID(utils.GenerateTaskID()))
				_, err = kf.asynq.Enqueue(task)
				if err != nil {
					log.Printf("failed to enqueue task: %v", err)
					continue
				}
				log.Printf("******************* Consumer Task Enqueued *******************")
				log.Printf("Consumer worker received message: %s\n", string(jsonData))
			}
		}
	}()
}
