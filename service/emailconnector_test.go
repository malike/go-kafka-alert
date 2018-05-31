package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/malike/go-kafka-alert/config"
	"github.com/malike/go-kafka-alert/db"
)

var fakeEmailRecipient = "st.malike@gmail.com"
var fakeEmailEvent = db.Event{
	EventID:   "eventid123456",
	EventType: "SUBSCRIPTION",
	UnmappedData: map[string]string{
		"Name":     "Malike",
		"ItemName": "Monthly Delivery of Awesomeness",
	},
	Recipient: []string{fakeEmailRecipient},
	Channel: map[string]bool{
		"EMAIL": true,
	},
}

func TestParseTemplateInvalidChannelEmail(t *testing.T) {
	config.LoadConfiguration()
	fakeEmailEvent.Channel = map[string]bool{
		"SMS": true,
	}
	_, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Log("Success. Channel Not supported")
	}
}

func TestParseTemplateForAllMessagesEmail(t *testing.T) {
	config.LoadConfiguration()
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
		t.Error(fmt.Printf("Messages not generated for all recipients. Expected %d ,"+
			" Got  %d", len(fakeEmailEvent.Recipient), len(msg)))
	}
}

func TestParseTemplateAllMessagesExceptInvalidRecipientsEmail(t *testing.T) {
	config.LoadConfiguration()
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
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
	}
	result, err := EventForEmail{fakeEmailEvent}.ParseTemplate()
	if err != nil {
		t.Errorf("Test failed. Result unexpected " + err.Error())
	}
	if result == nil || result[0].Content == "" {
		t.Errorf("Test failed. Result unexpected")
	}
}

func TestSendMessageWithNilEmail(t *testing.T) {
	msg := db.Message{}
	emailEvent := EventForEmail{fakeEmailEvent}
	response := emailEvent.SendMessage(msg)
	if response.Status != config.FAILED {
		t.Error("Empty message was sent.")
	}
}

func TestSendMessageWithContentEmptyEmail(t *testing.T) {
	msg := db.Message{AlertID: "1234", Content: "", DateCreated: time.Now(), Recipient: "+233201234567"}
	emailEvent := EventForEmail{fakeEmailEvent}
	emResponse := emailEvent.SendMessage(msg)
	if emResponse.Status != config.FAILED {
		t.Error("Empty message should fail.")
	}

}

func TestSendMessageEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Test doesn't run short mode")
	}
	fakeEmailEvent.Recipient = []string{
		fakeEmailRecipient,
	}
	fakeEmailEvent.Channel = map[string]bool{
		"EMAIL": true,
	}
	emailEvent := EventForEmail{fakeEmailEvent}
	msg, err := emailEvent.ParseTemplate()
	if err != nil {
		t.Error("Messages not generated")
	}
	if msg == nil {
		t.Error("Messages not generated")
		t.FailNow()
	}
	emailResponse := emailEvent.SendMessage(msg[0])
	if emailResponse.Status != config.SUCCESS {
		t.Error(fmt.Printf("Message not sent , Expected 'SUCCESS'. Got '%s'. Error : %s",
			emailResponse.Status, emailResponse.Response))
	}
}

func BenchmarkParseTemplateForMessageEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EventForEmail{fakeEmailEvent}.ParseTemplate()
	}
}

func BenchmarkSendMessageEmail(b *testing.B) {
	if testing.Short() {
		b.Skip("Test is running in short mode")
	}
	for i := 0; i < b.N; i++ {
		fakeEmailEvent.Recipient = []string{
			fakeEmailRecipient,
		}
		fakeEmailEvent.Channel = map[string]bool{
			"EMAIL": true,
		}
		emailEvent := EventForEmail{fakeEmailEvent}
		msg, _ := emailEvent.ParseTemplate()
		emailEvent.SendMessage(msg[0])
	}
}
