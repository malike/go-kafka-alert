package service

import (
	"go-kafka-alert/db"
	"strings"
)

type EventForMessage interface {
	ParseTemplate() ([]db.Message, error)

	SendMessage(message db.Message) db.MessageResponse
}

func CheckChannel(event db.Event, channel string) bool {
	return event.Channel[strings.ToUpper(channel)]
}
