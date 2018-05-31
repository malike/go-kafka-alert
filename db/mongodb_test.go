package db

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/malike/go-kafka-alert/config"
)

var msg = Message{
	MessageID:   "1234",
	Content:     "Sample Message",
	AlertID:     "Test1234",
	Subject:     "Sample Subject",
	Recipient:   "st.malike@gmail.com",
	DateCreated: time.Now(),
}

var messageIds = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

var fakeMessageResponse = MessageResponse{
	APIStatus:      "DELIVERED:123456",
	Response:       "SENT",
	Status:         "SUCCESS",
	TimeOfResponse: time.Now(),
}

func TestIndexFindAndRemoveMessage(t *testing.T) {
	config.LoadConfiguration()
	DialDB()
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}
	message, _ := msg.FindMessage(msg.MessageID)
	if message.Subject != msg.Subject {
		t.Fatal("Error finding message ")
		t.FailNow()
	}
	if deleted := msg.RemoveMessage(msg.MessageID); !deleted {
		t.Fatal("Message not removed")
		t.FailNow()
	}
}

func TestUpdateResponseAndRemove(t *testing.T) {
	config.LoadConfiguration()
	DialDB()
	err := msg.IndexMessage()
	if err != nil {
		t.Fatal("Error saving message " + err.Error())
		t.FailNow()
	}

	message, _ := msg.UpdateResponse(msg.MessageID, fakeMessageResponse)
	if message.MessageResponse != fakeMessageResponse {
		t.Fatal(fmt.Printf("Error updating  message response. Expected '%s', Got '%s'",
			message.MessageResponse.Status, fakeMessageResponse.Status))
		t.FailNow()
	}
	if deleted := msg.RemoveMessage(msg.MessageID); !deleted {
		t.Fatal("Message not removed")
		t.FailNow()
	}
}

func TestFindAllMessagesByReferenceAndRemoveAll(t *testing.T) {
	config.LoadConfiguration()
	DialDB()
	RemoveAllMessagesByReference("SaveMultipleTest")
	for _, it := range messageIds {
		msg.MessageID = it
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
	config.LoadConfiguration()
	DialDB()
	RemoveAllMessagesByReference("CountByReferenceTest")
	for _, it := range messageIds {
		msg.MessageID = it
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
	config.LoadConfiguration()
	DialDB()
	RemoveAllMessagesByReference("BenchMarkTest")
	for i := 0; i < b.N; i++ {
		msg.MessageID = strconv.Itoa(i)
		msg.Reference = "BenchMarkTest"
		msg.IndexMessage()
	}
	RemoveAllMessagesByReference("BenchMarkTest")

}

func BenchmarkMessage_IndexAndUpdateMessage(b *testing.B) {
	config.LoadConfiguration()
	DialDB()
	RemoveAllMessagesByReference("BenchMarkTestIndexUpdate")
	for i := 0; i < b.N; i++ {
		msg.MessageID = strconv.Itoa(i)
		msg.Reference = "BenchMarkTestIndexUpdate"
		msg.IndexMessage()

		fakeMessageResponse.Response = strconv.Itoa(i) + "Response"
		msg.UpdateResponse(msg.MessageID, fakeMessageResponse)
	}
	RemoveAllMessagesByReference("BenchMarkTestIndexUpdate")
}
