package db

import (
	"testing"
	"time"
)

var msg = Message{
	MessageId :"1234",
	Content:"Sample Message",
	AlertId:"Test1234",
	Subject:"Sample Subject",
	Recipient:"st.malike@gmail.com",
	DateCreated: time.Now(),
}

func TestIndexFindAndRemoveMessage(t *testing.T) {
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}
	message, _ := msg.FindMessage(msg.MessageId)
	if message.Subject != msg.Subject {
		t.Fatal("Error finding message ")
		t.FailNow()
	}
	if deleted := msg.RemoveMessage(msg.MessageId); !deleted {
		t.Fatal("Message not removed")
		t.FailNow()
	}
}

func TestUpdateResponse(t *testing.T) {
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}
	messageResponse := MessageResponse{
		APIStatus:"DELIVERED:123456",
		Response:"SENT",
		Status:"SUCCESS",
		TimeOfResponse: time.Now(),
	}
	msg.UpdateResponse(msg.MessageId, messageResponse)
	if msg.MessageResponse.Status != messageResponse.Status {
		t.Fatal("Error updating  message response")
		t.FailNow()
	}
}

func TestGetTemplate(t *testing.T) {

}
