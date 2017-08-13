package service

import (
	"st.malike.go.kafka.alert/db"
)

type EventForSMS struct {
	DefaultPhoneNumber string
	TriggeredEvent     db.Event
}

func (event EventForSMS) ParseTemplate() (db.Message, error) {
	message := db.Message{}
	message.Content = "Sample SMS"
	return message, nil
}

func (event EventForSMS) SendMessage() db.MessageResponse {
	return nil
}

