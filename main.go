package main

import (
	"OutBox/appconfig"
	"OutBox/data"
	mq "OutBox/mq"
	service "OutBox/service"
)

func main() {
	config := appconfig.NewConfig()

	context := data.PostgreSqlContext(config.DbConfig)
	defer context.Close()

	mqKafka, _ := mq.NewKafkaProducer(config.MqConfig)
	defer mqKafka.Producer.Close()

	repos := data.OutBoxRepository{Context: context}

	serv := service.NewService(repos, mqKafka)
	serv.Run()
}
