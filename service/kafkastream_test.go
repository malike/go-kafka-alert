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
var fakeStreamEvent = db.Event{
	EventId:"KafkaStream123456",
	DateCreated:time.Now(),
	Description:"Subscrption Desc",
	EventType:"SUBSCRIPTION",
	UnmappedData:map[string]string{
		"Name":"Malike St",
		"ItemName":"Monthly Delivery of Awesomeness",
	},
	Recipient: []string{"0201234567", "st.malike@gmail.com"},
	Channel: map[string]bool{
		"SMS": true,
		"EMAIL": true,
	},
	Subject:"Test Subscription from Kafa Stream",
}

func MockGetEventFromKafkaStream() ([]db.Event, error) {
	return []db.Event{fakeStreamEvent}, nil
}

func TestEventProcessorForChannel(t *testing.T) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventId + "EMAIL")
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventId + "SMS")

	fakeKafkaEvents, _ := MockGetEventFromKafkaStream()
	EventProcessorForChannel(fakeKafkaEvents)

	emailMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventId + "EMAIL")

	if CheckChannel(fakeKafkaEvents[0], "EMAIL") && len(emailMsgs) <= 0 {
		t.Error("Error: Email channel not indexed")
	}

	smsMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventId + "SMS")

	if CheckChannel(fakeKafkaEvents[0], "SMS") && len(smsMsgs) <= 0 {
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