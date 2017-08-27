package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
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
	messages = make([]db.Message, numOfRecipient)
	var messageContent = "Sample SMS"

	//generate individual messages for each recipient
	for i, rep := range event.TriggeredEvent.Recipient {
		dateCreated := time.Now()
		message := db.Message{}
		message.AlertId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
		message.Content = messageContent + " " + rep //temp
		message.Recipient = rep
		message.DateCreated = dateCreated
		message.ReferenceId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
		messages[i] = message
	}
	return messages, nil
}

func (event EventForSMS) SendMessage() db.MessageResponse {
	return db.MessageResponse{}
}

