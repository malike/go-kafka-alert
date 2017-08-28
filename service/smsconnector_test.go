package service

import (
	"testing"
	"go-kafka-alert/db"
	"fmt"
)

var fakeRecipient = "233201234567"
var fakeEvent = db.Event{
	Recipient: []string{fakeRecipient},
	Channel: map[string]bool{
		"SMS": true,
	},
}

func TestParseTemplate(t *testing.T) {
	result, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Errorf("Test failed. Result unexpected")
	}
	if result[0].Content != ("Sample SMS " + fakeRecipient) {
		t.Errorf("Test failed. Result unexpected")
	}
}

func TestParseTemplateInvalidChannel(t *testing.T) {
	fakeEvent.Channel = map[string]bool{
		"EMAIL" : true,
	}
	_, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Log("Success. Channel Not supported")
	}
}

func TestParseTemplateInvalidRecipient(t *testing.T) {
	fakeEvent.Recipient = []string{}
	_, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Log("Success. Recipient Unknown")
	}
}

func TestParseTemplateForAllMessages(t *testing.T) {
	fakeEvent.Recipient = []string{
		fakeRecipient,
		"233241234567",
		"233271234567",
	}
	msg, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated")
	}
	if len(msg) != len(fakeEvent.Recipient) {
		t.Error(fmt.Printf("Messages not generated for all recipients. Expected %d ," +
			" Got  %d",len(fakeEvent.Recipient),len(msg)))
	}

}
func TestParseTemplateAllMessagesExceptInvalidRecipients(t *testing.T) {
	fakeEvent.Recipient = []string{
		fakeRecipient,
		"st.malike@gmail.com",
		"233271234567",
	}
	msg, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated")
	}
	if len(msg) != (len(fakeEvent.Recipient) - 1) {
		t.Error(fmt.Printf("Messages not generated for all recipients, Expected %d Got %d",
			(len(fakeEvent.Recipient) - 1), len(msg)))
	}

}

func TestSendMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("Testing is running in short mode")
	}

}

func BenchmarkSendMessage(b *testing.B) {

}

