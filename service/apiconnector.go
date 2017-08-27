package service

import (
	"errors"
	"go-kafka-alert/db"
)

type EventForAPI struct {
	TriggeredEvent db.Event
}

func (event EventForAPI) ParseTemplate() (db.Message, error) {
	var message db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "API")
	if !channelSupported {
		return message, errors.New("API channel not supported")
	}
	message = db.Message{}
	message.Content = "Sample API Webhook"
	return message, nil
}

func (event EventForAPI) SendMessage() db.MessageResponse {
	return db.MessageResponse{}
}
