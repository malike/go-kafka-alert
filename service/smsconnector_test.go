package service

import (
	"testing"
	"go-kafka-alert/db"
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
	_, err := EventForSMS{fakeEvent}.ParseTemplate()
	if err != nil {
		t.Log("Success. Channel Not supported")
	}
}

func TestSendMessage(t *testing.T) {
	if testing.Short(){
		t.Skip("Testing is running in short mode")
	}


}

func BenchmarkSendMessage(b *testing.B) {

}

