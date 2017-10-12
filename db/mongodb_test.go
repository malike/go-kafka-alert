package db

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)

var msg = Message{
	MessageId :"1234",
	Content:"Sample Message",
	AlertId:"Test1234",
	Subject:"Sample Subject",
	Recipient:"st.malike@gmail.com",
	DateCreated: time.Now(),
}

var messageIds = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

var messageResponse = MessageResponse{
	APIStatus:"DELIVERED:123456",
	Response:"SENT",
	Status:"SUCCESS",
	TimeOfResponse: time.Now(),
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

func TestUpdateResponseAndRemove(t *testing.T) {
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}

	message, _ := msg.UpdateResponse(msg.MessageId, messageResponse)
	if message.MessageResponse != messageResponse {
		t.Fatal(fmt.Printf("Error updating  message response. Expected '%s', Got '%s'",
			message.MessageResponse.Status, messageResponse.Status))
		t.FailNow()
	}
	if deleted := msg.RemoveMessage(msg.MessageId); !deleted {
		t.Fatal("Message not removed")
		t.FailNow()
	}
}

func TestFindAllMessagesByReferenceAndRemoveAll(t *testing.T) {
	RemoveAllMessagesByReference("SaveMultipleTest")
	for _, it := range messageIds {
		msg.MessageId = it
		msg.Reference = "SaveMultipleTest"
		msg.IndexMessage()
	}
	messages, _ := FindAllMessagesByReference("SaveMultipleTest")
	if len(messages) != len(messageIds) {
		t.Fatal(fmt.Printf("Not all messages were retrieved. Expected '%d'. Got '%d' ",
			len(messageIds), len(messages)))
	}
	RemoveAllMessagesByReference("SaveMultipleTest")
}

func TestCountAllMessagesByReference(t *testing.T) {
	RemoveAllMessagesByReference("CountByReferenceTest")
	for _, it := range messageIds {
		msg.MessageId = it
		msg.Reference = "CountByReferenceTest"
		msg.IndexMessage()
	}
	messageSize := CountAllMessagesByReference("CountByReferenceTest")
	if messageSize != len(messageIds) {
		t.Fatal(fmt.Printf("Not all messages were retrieved. Expected '%d'. Got '%d' ",
			len(messageIds), messageSize))
	}
	RemoveAllMessagesByReference("CountByReferenceTest")
}

func BenchmarkMessage_IndexMessage(b *testing.B) {
	RemoveAllMessagesByReference("BenchMarkTest")
	for i := 0; i < b.N; i++ {
		msg.MessageId = strconv.Itoa(i)
		msg.Reference = "BenchMarkTest"
		msg.IndexMessage()
	}
	RemoveAllMessagesByReference("BenchMarkTest")

}

func BenchmarkMessage_IndexAndUpdateMessage(b *testing.B) {
	RemoveAllMessagesByReference("BenchMarkTestIndexUpdate")
	for i := 0; i < b.N; i++ {
		msg.MessageId = strconv.Itoa(i)
		msg.Reference = "BenchMarkTestIndexUpdate"
		msg.IndexMessage()

		messageResponse.Response = strconv.Itoa(i) + "Response"
		msg.UpdateResponse(msg.MessageId, messageResponse)
	}
	RemoveAllMessagesByReference("BenchMarkTestIndexUpdate")
}
