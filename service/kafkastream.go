package service

import (
	"go-kafka-alert/db"
	"go-kafka-alert/util"
)

func GetEventFromKafkaStream() ([]db.Event) {
	events := []db.Event{}
	return events
}

func EventProcessorForChannel(events []db.Event) {
	for _, event := range events {
		if CheckChannel(event, "SMS") {
			util.Info.Print("Processing " + event.EventId + " for SMS")
			smsChannel := EventForSMS{event}
			ProcessEvent(smsChannel)
		}
		if CheckChannel(event, "EMAIL") {
			util.Info.Print("Processing " + event.EventId + " for EMAIL")
			emailChannel := EventForEmail{event}
			ProcessEvent(emailChannel)
		}
		if CheckChannel(event, "API") {
			util.Info.Print("Processing " + event.EventId + " for API")
			apiChannel := EventForAPI{event}
			ProcessEvent(apiChannel)
		}
	}
}

func ProcessEvent(eventForMessage EventForMessage) {
	messages, err := eventForMessage.ParseTemplate()
	if err != nil {
		for _, msg := range messages {
			//index message
			msg.IndexMessage()

			response := eventForMessage.SendMessage(msg)
			//index response
			msg.UpdateResponse(msg.MessageId, response)

		}
	}

}