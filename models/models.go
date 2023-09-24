package models

import "github.com/confluentinc/confluent-kafka-go/kafka"

type OutBox struct {
	Id      int
	Payload string
	Status  int
	DT      string
}

type Message struct {
	OutBoxId     int
	KafkaMessage kafka.Message
	Result       bool
	Retry        bool
	Error        error
}

type Payload struct {
	Topic   string            `json:"Topic"`
	Key     string            `json:"Key"`
	Message string            `json:"Message"`
	Headers map[string]string `json:"Headers"`
}
