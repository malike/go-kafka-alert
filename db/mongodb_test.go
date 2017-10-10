package db

import (
	"testing"
	"time"
)

var msg = Message{
	Id :"1234",
	ReferenceId:"1234",
	Content:"Sample Message",
	AlertId:"Test1234",
	Subject:"Sample Subject",
	Recipient:"st.malike@gmail.com",
	DateCreated: time.Now(),
}

func TestIndexMessage(t *testing.T) {
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}
	message, _ := msg.FindMessage(msg.Id)
	if message.Subject != msg.Subject {
		t.Fatal("Error saving and finding message ")
		t.FailNow()
	}
}

func TestUpdateResponse(t *testing.T) {

}

func TestGetTemplate(t *testing.T) {

}
