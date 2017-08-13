package service

import (
	"st.malike.go.kafka.alert/db"
	"errors"
)

type EventForEmail struct {
	DefaultEmail   string
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() (db.Message, error) {
	channelSupported := CheckChannel(event.TriggeredEvent,"EMAIL")
	if !channelSupported{
		return nil, errors.New("Email channel not supported")
	}
	message := db.Message{}
	message.Content = "<html><body></body>Sample Email</html>"
	return message, nil
}

func (event EventForEmail) SendMessage() db.MessageResponse {
	return nil;
}
