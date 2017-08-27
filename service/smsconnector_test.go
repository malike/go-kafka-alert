package service

import (
	"testing"
	"time"
	"go-kafka-alert/db"
)

type FakeEventForSMS struct {
	Message db.Message
	Err     error
}

func (event FakeEventForSMS) ParseTemplate() (db.Message, error) {
	var message db.Message
	if event.Err != nil {
		return message, event.Err
	}
	return event.Message, nil
}

func (event FakeEventForSMS) SendMessage() db.MessageResponse {
	return db.MessageResponse{Response:"SUCCESS", TimeOfResponse:time.Now()}
}

func TestParseTemplate(t *testing.T) {
	f := FakeEventForSMS{
		Message: db.Message{
			Id          : "",
			ReferenceId : "",
			AlertId     : "",
			Content     : "",
			Recipient   : "",
			ApiResponse : db.MessageResponse{},
			DateCreated : time.Now()},
		Err: nil,
	}
	if f.Err != nil {
		 t.Errorf("Test failed. Result unexpected")
	}
}

func TestSendMessage(t *testing.T) {

}

