package service

import "go-kafka-alert/db"

func sendmessage(event EventForMessage) {

	msg, err := event.ParseTemplate()
	if err == nil {
		//index message
		db.IndexMessage(msg)

		response := event.SendMessage()
		//index response
		db.UpdateResponse(msg, response)

	}

}