package service

import (
	"errors"
	"go-kafka-alert/db"
)

type EventForEmail struct {
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() ([]db.Message, error) {
	var message []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "EMAIL")
	if !channelSupported {
		return message, errors.New("Email channel not supported")
	}
	return message, nil
}

func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	return db.MessageResponse{};
}
