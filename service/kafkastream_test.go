package service

import (
	"testing"
	"go-kafka-alert/db"
	"time"
)

var event = db.Event{}
var eventSMS = EventForSMS{event}
var eventEmail = EventForEmail{event}
var eventAPI = EventForAPI{event}
var fakeSteamEvent = db.Event{
	EventId:"eventid123456",
	DateCreated:time.Now(),
	Description:"Subscrption Desc",
	EventType:"SUBSCRIPTION",
	UnmappedData:map[string]string{
		"Name":"Malike",
		"ItemName":"Monthly Delivery of Awesomeness",
	},
	Recipient: []string{"0201234567","st.malike@gmail.com"},
	Channel: map[string]bool{
		"SMS": true,

	},
	Subject:"Test Subscription",
}


func MockGetEventFromKafkaStream() ([]db.Event, error) {
	return []db.Event{fakeEmailEvent},nil
}

func TestEventProcessorForChannel(t *testing.T) {
	db.RemoveAllMessagesByReference(fakeSteamEvent.EventId+"EMAIL")
	db.RemoveAllMessagesByReference(fakeSteamEvent.EventId+"SMS")
	fakeEvents,_ := MockGetEventFromKafkaStream()
	EventProcessorForChannel(fakeEvents)

	emailMsgs,_ := db.FindAllMessagesByReference(fakeSteamEvent.EventId+"EMAIL")

	if len(emailMsgs) <= 0{
          t.Error("Error: Email channel not indexed")
	}

	smsMsgs,_ := db.FindAllMessagesByReference(fakeSteamEvent.EventId+"EMAIL")

	if len(smsMsgs) <= 0{
         t.Error("Error: SMS channel not indexed")
	}
}


func TestProcessEvent(t *testing.T) {

}


func BenchmarkGetEventFromKafkaStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetEventFromKafkaStream()
	}
}

func BenchmarkProcessEventForSMS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventSMS)
	}
}
func BenchmarkProcessEventForEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventEmail)
	}
}
func BenchmarkProcessEventForAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventAPI)
	}
}

func BenchmarkEventProcessorForChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EventProcessorForChannel([]db.Event{event})
	}
}