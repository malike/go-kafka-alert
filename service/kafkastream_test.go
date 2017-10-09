package service

import (
	"testing"
	"go-kafka-alert/db"
)

var event = db.Event{}
var eventSMS = EventForSMS{event}
var eventEmail = EventForEmail{event}
var eventAPI = EventForAPI{event}

func TestProcessEvent(t *testing.T) {

}

func TestEventProcessorForChannel(t *testing.T) {

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
		EventProcessorForChannel(event)
	}
}