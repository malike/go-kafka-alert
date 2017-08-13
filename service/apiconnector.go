package service

import (
	"st.malike.go.kafka.alert/db"
	"errors"
)

type EventForAPI struct {
	TriggeredEvent db.Event
}

func (event EventForAPI) ParseTemplate() (db.Message, error) {
	channelSupported := CheckChannel(event.TriggeredEvent,"API")
	if !channelSupported{
		return nil, errors.New("API channel not supported")
	}
	message := db.Message{}
	message.Content = "Sample API Webhook"
	return message, nil
}

func (event EventForAPI) SendMessage() db.MessageResponse {
	return nil
}
