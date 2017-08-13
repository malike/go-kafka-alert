package service

import (
	"st.malike.go.kafka.alert/db"
	"errors"
)

type EventForSMS struct {
	DefaultPhoneNumber string
	TriggeredEvent     db.Event
}

func (event EventForSMS) ParseTemplate() (db.Message, error) {
	channelSupported := CheckChannel(event.TriggeredEvent,"SMS")
	if !channelSupported{
		return nil, errors.New("SMS channel not supported")
	}
	message := db.Message{}
	message.Content = "Sample SMS"
	return message, nil

}

func (event EventForSMS) SendMessage() db.MessageResponse {
	return nil
}

