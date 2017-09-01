package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
)

type EventForEmail struct {
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "EMAIL")
	if !channelSupported {
		return messages, errors.New("SMS channel not supported")
	}
	emailContent := ParseTemplateForMessage(event.TriggeredEvent,"EMAIL")
	for  _ ,em := range event.TriggeredEvent.Recipient{
		if validateEmail(em) {
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.DateCreated = dateCreated
			message.AlertId = event.TriggeredEvent.EventId + "_EMAIL_" + em
			message.Content = emailContent
			message.ReferenceId = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventId
			message.Id = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventId
			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	return db.MessageResponse{};
}

func validateEmail(email string) bool {
	return true
}
