package service

import (
	"errors"
	"github.com/malike/go-kafka-alert/db"
)

type EventForAPI struct {
	TriggeredEvent db.Event
}

func (event EventForAPI) ParseTemplate() ([]db.Message, error) {
	var message []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "API")
	if !channelSupported {
		return message, errors.New("API channel not supported")
	}
	return message, nil
}

func (event EventForAPI) SendMessage(message db.Message) db.MessageResponse {
	return db.MessageResponse{}
}
