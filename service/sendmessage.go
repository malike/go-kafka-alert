package service

import (
	"st.malike.go.kafka.alert/db"
)

type EventForMessage interface {
	ParseTemplate() (db.Message, error)

	SendMessage() db.MessageResponse
}
