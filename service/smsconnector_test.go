package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/malike/go-kafka-alert/db"

	"github.com/malike/go-kafka-alert/config"
)

var fakeRecipient = "233201234567"
var fakeEvent = db.Event{
	EventID:   "eventid123456",
	EventType: "SUBSCRIPTION",
	UnmappedData: map[string]string{
		"Name":     "Malike",
		"ItemName": "Monthly Delivery of Awesomeness",
	},
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
	if result[0].Content == "" {
		t.Errorf("Test failed. Result unexpected")
	}
}

func TestParseTemplateInvalidChannel(t *testing.T) {
	fakeEvent.Channel = map[string]bool{
		"EMAIL": true,
	}
	_, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err == nil {
		t.Error("Error. Channel Not supported")
	}
}

func TestParseTemplateInvalidRecipient(t *testing.T) {
	fakeEvent.Recipient = []string{}
	_, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err == nil {
		t.Error("Error. Recipient Unknown")
	}
}

func TestParseTemplateForAllMessages(t *testing.T) {
	fakeEvent.Recipient = []string{
		fakeRecipient,
		"233241234567",
		"233271234567",
	}
	fakeEvent.Channel = map[string]bool{
		"SMS": true,
	}
	msg, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated ", err)
	}
	if len(msg) != len(fakeEvent.Recipient) {
		t.Error(fmt.Printf("Messages not generated for all recipients. Expected %d ,"+
			" Got  %d", len(fakeEvent.Recipient), len(msg)))
	}

}

func TestParseTemplateAllMessagesExceptInvalidRecipients(t *testing.T) {
	fakeEvent.Recipient = []string{
		fakeRecipient,
		"st.malike@gmail.com",
		"233271234567",
	}
	fakeEvent.Channel = map[string]bool{
		"SMS": true,
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
		t.Skip("Test doesn't run short mode")
	}
	fakeEvent.Recipient = []string{
		"+233208358615",
	}
	fakeEvent.Channel = map[string]bool{
		"SMS": true,
	}
	smsEvent := EventForSMS{fakeEvent}
	msg, err := smsEvent.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated")
	}
	if msg == nil {
		t.Error("Messages not generated")
	}
	smsResponse := smsEvent.SendMessage(msg[0])
	if smsResponse.Status != config.SUCCESS {
		t.Error(fmt.Printf("Message not sent , Expected 'SUCCESS'. Got '%s'. Error '%s'",
			smsResponse.Status, smsResponse.Response))
	}

}

func TestSendMessageWithNil(t *testing.T) {
	msg := db.Message{}
	smsEvent := EventForSMS{fakeEvent}
	smsResponse := smsEvent.SendMessage(msg)
	if smsResponse.Status != config.FAILED {
		t.Error("Empty message was sent.")
	}

}

func TestSendMessageWithContentEmpty(t *testing.T) {
	msg := db.Message{AlertID: "1234", Content: "", DateCreated: time.Now(), Recipient: "+233201234567"}
	smsEvent := EventForSMS{fakeEvent}
	smsResponse := smsEvent.SendMessage(msg)
	if smsResponse.Status == config.SUCCESS {
		t.Error("Empty message should fail.")
	}

}

func BenchmarkParseTemplateForMessageSMS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EventForSMS{fakeEvent}.ParseTemplate()
	}
}

func BenchmarkSendMessage(b *testing.B) {
	if testing.Short() {
		b.Skip("Test doesn't run short mode")
	}
	for i := 0; i < b.N; i++ {
		fakeEvent.Recipient = []string{
			"+233201234567",
		}
		fakeEvent.Channel = map[string]bool{
			"SMS": true,
		}
		smsEvent := EventForSMS{fakeEvent}
		msg, _ := smsEvent.ParseTemplate()
		smsEvent.SendMessage(msg[0])
	}
}
