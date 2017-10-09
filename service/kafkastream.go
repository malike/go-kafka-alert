package service

import (
	"go-kafka-alert/db"
)

func EventProcessorForChannel(event db.Event) {

	if CheckChannel(event,"SMS"){
		smsChannel  := EventForSMS{event}
		ProcessEvent(smsChannel)
	}
	if CheckChannel(event,"EMAIL"){
		emailChannel  := EventForEmail{event}
		ProcessEvent(emailChannel)

	}
	if CheckChannel(event,"API"){
		apiChannel  := EventForAPI{event}
		ProcessEvent(apiChannel	)
	}
}


func ProcessEvent(eventForMessage EventForMessage){

	messages, err := eventForMessage.ParseTemplate()
	if err != nil {
		for _, msg := range messages {
			//index message
			db.IndexMessage(msg)

			response := eventForMessage.SendMessage(msg)
			//index response
			db.UpdateResponse(msg, response)

		}
	}


}