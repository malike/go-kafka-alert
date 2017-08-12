package service

import "st.malike.go.kafka.alert/db"

type SendMessage interface {

	SendMessage(message db.Message) db.MessageResponse
}
