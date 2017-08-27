package service

import (
	"errors"
	"go-kafka-alert/db"
)

type EventForEmail struct {
	DefaultEmail   string
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() (db.Message, error) {
	var message db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "EMAIL")
	if !channelSupported {
		return message, errors.New("Email channel not supported")
	}
	message = db.Message{}
	message.Content = "<html><body></body>Sample Email</html>"
	return message, nil
}

func (event EventForEmail) SendMessage() db.MessageResponse {
	return db.MessageResponse{};
}
