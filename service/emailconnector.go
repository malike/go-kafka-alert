package service

import "st.malike.go.kafka.alert/db"

type EventForEmail struct {
	DefaultEmail   string
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() (db.Message, error) {
	message := db.Message{}
	message.Content = "<html><body></body>Sample Email</html>"
	return message, nil
}

func (event EventForEmail) SendMessage() db.MessageResponse {
	return nil;
}
