package data

import (
	"OutBox/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type OutBoxRepository struct {
	Context *sql.DB
}

func (repository OutBoxRepository) GetOutBoxes() []models.OutBox {

	query := "SELECT * FROM outbox WHERE Status = $1 AND DT <= $2"
	rows, err := repository.Context.Query(query, 0, time.Now())
	if err != nil {
		log.Fatal("Error reading records", err)
	}

	defer rows.Close()

	var outBoxes []models.OutBox

	for rows.Next() {
		outBox := models.OutBox{}
		if err := rows.Scan(&outBox.Id, &outBox.Payload, &outBox.Status, &outBox.DT); err != nil {
			log.Fatal(err)
		}
		outBoxes = append(outBoxes, outBox)
	}

	return outBoxes
}

func (repository OutBoxRepository) UpdateOutBoxes(messages []models.Message) {

	query := "UPDATE outbox SET %v WHERE Id = %v"
	for _, message := range messages {
		var field string

		if !message.Result && !message.Retry && message.Error != nil {
			field = "status=2"
		} else if !message.Result && message.Retry {
			field = fmt.Sprintf("DT=%s", time.Now().Add(time.Minute*15))
		} else if message.Result {
			field = "status=1"
		}
		_, err := repository.Context.Query(fmt.Sprintf(query, field, message.OutBoxId))
		if err != nil {
			log.Println(err)
		}
	}
}
