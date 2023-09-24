package mq

import (
	"OutBox/appconfig"
	"OutBox/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type KafkaMq struct {
	Producer *kafka.Producer
}

func NewKafkaProducer(config appconfig.Kafka) (KafkaMq, error) {

	producer, err := kafka.NewProducer(configure(config))

	if err != nil {
		log.Fatal("Fail create Kafka producer ")
	}

	return KafkaMq{Producer: producer}, nil
}

func configure(config appconfig.Kafka) *kafka.ConfigMap {

	kafkaConfig := kafka.ConfigMap{}

	for _, element := range config.KafkaConfig {
		for key, value := range element {
			_ = kafkaConfig.SetKey(key, value)
		}
	}
	return &kafkaConfig
}

func (kafkaMq KafkaMq) SendMassages(messages []models.Message) {

	for _, message := range messages {
		delChan := make(chan kafka.Event)

		err := kafkaMq.Producer.Produce(&message.KafkaMessage, delChan)
		if err != nil {
			message.Error = err
			continue
		}
		answer := <-delChan
		msg := answer.(*kafka.Message)

		if msg.TopicPartition.Error != nil {
			message.Error = msg.TopicPartition.Error
			message.Retry = true
		} else {
			message.Result = true
		}
		kafkaMq.Producer.Flush(10000)
		close(delChan)
	}
}
