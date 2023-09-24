package service

import (
	"OutBox/data"
	"OutBox/mappers"
	"OutBox/models"
	"OutBox/mq"
)

type Service struct {
	Repository data.OutBoxRepository
	KafkaMq    mq.KafkaMq
}

func NewService(repository data.OutBoxRepository, kafkaMq mq.KafkaMq) Service {
	return Service{repository, kafkaMq}
}

func (service Service) Run() {

	for {
		outBoxes := service.Repository.GetOutBoxes()
		if len(outBoxes) > 0 {
			var messagesToSend []models.Message
			var errorMessages []models.Message

			for _, outBox := range outBoxes {
				message := mappers.OutBoxToMessage(outBox)
				if message.Error != nil {
					errorMessages = append(errorMessages, message)
					continue
				}
				messagesToSend = append(messagesToSend, message)
			}
			if len(errorMessages) > 0 {
				service.Repository.UpdateOutBoxes(errorMessages)
			}

			if len(messagesToSend) > 0 {
				service.KafkaMq.SendMassages(messagesToSend)
				service.Repository.UpdateOutBoxes(messagesToSend)
			}

		}

	}
}
