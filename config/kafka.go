package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type KafkaConfig struct {
	ProducerTopic            string `json:"KAFKA_PRODUCERTOPIC"`
	ProducerBootstrapServers string `json:"KAFKA_PRODUCERBOOTSTRAPSERVERS"`
	ProducerBufferSize       int    `json:"KAFKA_PRODUCERBUFFERSIZE"`
	ProducerFlushTimeout     int    `json:"KAFKA_PRODUCERFLUSHTIMEOUT"`
	ProducerNumberOfWorkers  int    `json:"KAFKA_PRODUCERNUMBEROFWORKERS"`
	ProducerSASLMechanisms   string `json:"KAFKA_PRODUCERSASLMECHANISMS"`
	ProducerSASLPassword     string `json:"KAFKA_PRODUCERSASLPASSWORD"`
	ProducerSASLUsername     string `json:"KAFKA_PRODUCERSASLUSERNAME"`
	ProducerSecurityProtocol string `json:"KAFKA_PRODUCERSECURITYPROTOCOL"`
}

var Kafka *KafkaConfig

func loadKafkaConfig() {
	Kafka = &KafkaConfig{}
	err := envconfig.Process("kafka", Kafka)
	if err != nil {
		log.Fatal(err.Error())
	}
}
