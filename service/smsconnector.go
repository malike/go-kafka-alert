package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
	"regexp"
)

type EventForSMS struct {
	TriggeredEvent db.Event
}

func (event EventForSMS) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "SMS")
	if !channelSupported {
		return messages, errors.New("SMS channel not supported")
	}
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		return messages, errors.New("No recipients found")
	}
	var messageContent = ParseTemplateForMessage(event.TriggeredEvent, "SMS")

	//generate individual messages for each recipient
	for _, rep := range event.TriggeredEvent.Recipient {
		if validatePhone(rep) {
			dateCreated := time.Now()
			message := db.Message{}
			message.AlertId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
			message.Content = messageContent + " " + rep //temp
			message.Recipient = rep
			message.DateCreated = dateCreated
			message.ReferenceId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
			messages = append(messages, message)
		}
	}
	return messages, nil
}

func (event EventForSMS) SendMessage() db.MessageResponse {
	return db.MessageResponse{}
}

func validatePhone(phone string) bool {
	re := regexp.MustCompile("[0-9]+")
	return re.MatchString(phone)
}

