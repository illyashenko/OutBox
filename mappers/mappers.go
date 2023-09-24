package mappers

import (
	"OutBox/models"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func OutBoxToMessage(outBox models.OutBox) models.Message {

	message := models.Message{}
	payload := models.Payload{}

	message.OutBoxId = outBox.Id

	err := json.Unmarshal([]byte(outBox.Payload), &payload)
	if err != nil {
		message.Error = err
		return message
	}
	message.KafkaMessage = PayloadToKafkaMessage(payload)

	return message
}

func PayloadToKafkaMessage(payload models.Payload) kafka.Message {
	return kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &payload.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(payload.Key),
		Value:          []byte(payload.Message),
		Headers:        toHeaders(payload.Headers),
	}
}

func toHeaders(headers map[string]string) []kafka.Header {
	var kafkaHeaders []kafka.Header
	for key, value := range headers {
		header := kafka.Header{
			Key:   key,
			Value: []byte(value),
		}
		kafkaHeaders = append(kafkaHeaders, header)
	}
	return kafkaHeaders
}
