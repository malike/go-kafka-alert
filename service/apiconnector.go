package service

import "st.malike.go.kafka.alert/db"

type EventForAPI struct {
	TriggeredEvent db.Event
}

func (event EventForAPI) ParseTemplate() (db.Message, error) {
	message := db.Message{}
	message.Content = "Sample API Webhook"
	return message, nil
}

func (event EventForAPI) SendMessage() db.MessageResponse {
	return nil
}
