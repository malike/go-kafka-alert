package service

import (
	"go-kafka-alert/db"
	"strings"
)

//EventForMessage : Interface to be Implemented For Delivery Channels
type EventForMessage interface {
	ParseTemplate() ([]db.Message, error)

	SendMessage(message db.Message) db.MessageResponse
}

//CheckChannel : Validates Delivery Channel Support By Event
func CheckChannel(event db.Event, channel string) bool {
	return event.Channel[strings.ToUpper(channel)]
}
