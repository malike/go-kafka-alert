package service

import (
	"errors"
	"go-kafka-alert/db"
)

type EventForSMS struct {
	DefaultPhoneNumber string
	TriggeredEvent     db.Event
}

func (event EventForSMS) ParseTemplate() (db.Message, error) {
	var message db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "SMS")
	if !channelSupported {
		return message, errors.New("SMS channel not supported")
	}
	message = db.Message{}
	message.Content = "Sample SMS"
	return message, nil

}

func (event EventForSMS) SendMessage() db.MessageResponse {
	return db.MessageResponse{}
}

