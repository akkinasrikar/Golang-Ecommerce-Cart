package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type KafkaConfig struct {
	ProducerTopic            string `json:"KAFKA_PRODUCERTOPIC"`
	ProducerBootstrapServers string `json:"KAFKA_PRODUCERBOOTSTRAPSERVERS"`
	ProducerBufferSize       int    `json:"KAFKA_PRODUCERBUFFERSIZE"`
	ProducerNumberOfWorkers  int    `json:"KAFKA_PRODUCERNUMBEROFWORKERS"`
	ProducerSASLMechanisms   string `json:"KAFKA_PRODUCERSASLMECHANISMS"`
	ProducerSASLPassword     string `json:"KAFKA_PRODUCERSASLPASSWORD"`
	ProducerSASLUsername     string `json:"KAFKA_PRODUCERSASLUSERNAME"`
	ProducerSecurityProtocol string `json:"KAFKA_PRODUCERSECURITYPROTOCOL"`
	ProducerGroupId          string `json:"KAFKA_PRODUCERGROUPID"`
	ProducerOffsetReset      string `json:"KAFKA_PRODUCEROFFSETRESET"`
}

var Kafka *KafkaConfig

func loadKafkaConfig() {
	Kafka = &KafkaConfig{}
	err := envconfig.Process("kafka", Kafka)
	if err != nil {
		log.Fatal(err.Error())
	}
}
