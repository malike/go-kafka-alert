package service

import "go-kafka-alert/db"

func ProcessEvent(event EventForMessage) {

	messages, err := event.ParseTemplate()
	if err != nil {
		for _, msg := range messages {
			//index message
			db.IndexMessage(msg)

			response := event.SendMessage()
			//index response
			db.UpdateResponse(msg, response)

		}
	}

}