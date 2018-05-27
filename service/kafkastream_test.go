package service

import (
	"testing"
	"time"

	"github.com/malike/go-kafka-alert/db"
)

var fakeStreamEvent = db.Event{
	EventID:     "KafkaStream123456",
	DateCreated: time.Now(),
	Description: "Subscrption Desc",
	EventType:   "SUBSCRIPTION",
	UnmappedData: map[string]string{
		"Name":     "Malike St",
		"ItemName": "Monthly Delivery of Awesomeness",
	},
	Recipient: []string{"0201234567", "0241234567", "st.malike@gmail.com"},
	Channel: map[string]bool{
		"SMS":   true,
		"EMAIL": true,
	},
	Subject: "Test Subscription from Kafa Stream",
}
var fakeStreamEventService = db.Event{
	EventID:     "KafkaStream123456",
	DateCreated: time.Now(),
	Description: "Metrics on Service A",
	EventType:   "SERVICEHEALTH",
	UnmappedData: map[string]string{
		"ServiceName":     "Service A",
		"FailureCount":    "4",
		"FailureDuration": "15",
	},
	Recipient: []string{"0201234567", "0241234567", "st.malike@gmail.com"},
	Channel: map[string]bool{
		"SMS":   true,
		"EMAIL": true,
	},
	Subject: "Service Health Alert [Kafa Stream]",
}
var eventSMS = EventForSMS{fakeStreamEvent}
var eventEmail = EventForEmail{fakeStreamEvent}
var eventAPI = EventForAPI{fakeStreamEvent}

func MockGetEventFromKafkaStream() ([]db.Event, error) {
	return []db.Event{fakeStreamEvent, fakeStreamEventService}, nil
}

func TestEventProcessorForChannel(t *testing.T) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")

	fakeKafkaEvents, _ := MockGetEventFromKafkaStream()
	EventProcessorForChannel(fakeKafkaEvents)

	emailMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")

	if CheckChannel(fakeKafkaEvents[0], "EMAIL") && len(emailMsgs) <= 0 {
		t.Error("Error: Email channel not indexed")
		t.FailNow()
	}

	smsMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventID + "SMS")

	if CheckChannel(fakeKafkaEvents[0], "SMS") && len(smsMsgs) <= 0 {
		t.Error("Error: SMS channel not indexed")
		t.FailNow()
	}
}

func TestProcessEventSMS(t *testing.T) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")

	ProcessEvent(eventSMS)

	smsMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventID + "SMS")

	if CheckChannel(eventSMS.TriggeredEvent, "SMS") && len(smsMsgs) <= 0 {
		t.Error("Error: SMS channel not indexed")
	}
}

func TestProcessEventEmail(t *testing.T) {

	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")

	ProcessEvent(eventEmail)

	smsMsgs, _ := db.FindAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")

	if CheckChannel(eventSMS.TriggeredEvent, "EMAIL") && len(smsMsgs) <= 0 {
		t.Error("Error: EMAIL channel not indexed")
	}

}

func BenchmarkProcessEventForSMS(b *testing.B) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventSMS)
	}
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")
}

func BenchmarkProcessEventForEmail(b *testing.B) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventEmail)
	}
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")
}

func BenchmarkProcessEventForAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessEvent(eventAPI)
	}
}

func BenchmarkEventProcessorForChannel(b *testing.B) {
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")

	for i := 0; i < b.N; i++ {

		fakeKafkaEvents, _ := MockGetEventFromKafkaStream()
		EventProcessorForChannel(fakeKafkaEvents)

	}
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "EMAIL")
	db.RemoveAllMessagesByReference(fakeStreamEvent.EventID + "SMS")
}
