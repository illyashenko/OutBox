package appconfig

import (
	"encoding/json"
	"log"
	"os"
)

type Postgres struct {
	ConnectionString string `json:"ConnectionString"`
}

type Kafka struct {
	KafkaConfig        []map[string]string `json:"KafkaConfig"`
	AdditionalSettings string              `json:"AdditionalSettings"`
}

type AppConfig struct {
	DbConfig Postgres `json:"Postgres"`
	MqConfig Kafka    `json:"Kafka"`
}

func (config *AppConfig) configure() {

	data, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatal("Config file reading error", err)
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal("Config json deserialization error", err)
	}

}

func NewConfig() AppConfig {

	config := AppConfig{}
	config.configure()

	return config
}
