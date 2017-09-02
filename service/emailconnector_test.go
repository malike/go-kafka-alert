package service

import (
	"testing"
	"go-kafka-alert/db"
	"fmt"
)

var fakeEmailRecipient = "st.malike@gmail.com"
var fakeEmailEvent = db.Event{
	Recipient: []string{fakeEmailRecipient},
	Channel: map[string]bool{
		"EMAIL": true,
	},
}


func TestParseTemplateInvalidChannelEmail(t *testing.T) {
	fakeEmailEvent.Channel = map[string]bool{
		"SMS" : true,
	}
	_, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Log("Success. Channel Not supported")
	}
}


func TestParseTemplateForAllMessagesEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
		"st.malike@yahoo.com",
		"st.malike@outlook.com",
	}
	fakeEmailEvent.Channel = map[string]bool{
		"EMAIL": true,
	}
	msg, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated ", err)
	}
	if len(msg) != len(fakeEmailEvent.Recipient) {
		t.Error(fmt.Printf("Messages not generated for all recipients. Expected %d ," +
			" Got  %d", len(fakeEmailEvent.Recipient), len(msg)))
	}
}

func TestParseTemplateAllMessagesExceptInvalidRecipientsEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
		"st.malike@malike.com",
		"233271234567",
	}
	fakeEmailEvent.Channel = map[string]bool{
		"EMAIL": true,
	}
	msg, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated")
	}
	if len(msg) != (len(fakeEmailEvent.Recipient) - 1) {
		t.Error(fmt.Printf("Messages not generated for all recipients, Expected %d Got %d",
			(len(fakeEmailEvent.Recipient) - 1), len(msg)))
	}

}

func TestParseTemplateInvalidRecipientEmail(t *testing.T) {
	fakeEmailEvent.Recipient = []string{}
	_, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err == nil {
		t.Error("Error. Recipient Unknown")
	}
}

func TestParseTemplateEmail(t *testing.T) {
	result, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Errorf("Test failed. Result unexpected")
	}
	if result[0].Content == "" {
		t.Errorf("Test failed. Result unexpected")
	}
}
