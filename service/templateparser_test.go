package service

import (
	"testing"
	"go-kafka-alert/db"
)


var fakeTempEvent = db.Event{
	EventId:"SUBSCRIPTION",
	Description:"Notification based on subscribing to a service",
	UnmappedData:map[string]string{
		"Name":"Malike",
		"ItemName":"Monthly Delivery of Awesomeness",
	},
	Recipient: []string{fakeRecipient},
	Channel: map[string]bool{
		"SMS": true,
	},
}

func TestParseTemplateForMessage404Template(t *testing.T) {
	fakeTempEvent.EventId="DUMMY_SERVICE"
 parsed,err:= ParseTemplateForMessage(fakeTempEvent,"SMS")
	if err != nil{
		t.Error("Error parsing template "+err.Error())
		t.FailNow()
	}
	if parsed != fakeTempEvent.Description{
		t.Error("Parser didn't pick description of event as default template as fallback ")
		t.FailNow()
	}
	t.Log("Parsed Template == '"+parsed)
}

func TestParseTemplateForMessage(t *testing.T) {
 parsed,err := ParseTemplateForMessage(fakeTempEvent,"SMS")
	if err != nil{
		t.Error("Error parsing template "+err.Error())
		t.FailNow()
	}
	if len(parsed) <= 0{
		t.Error("Invalid template parsed")
		t.FailNow()
	}
	t.Log("Parsed Template == '"+parsed)
}

func BenchmarkParseTemplateForMessage(b *testing.B) {

}