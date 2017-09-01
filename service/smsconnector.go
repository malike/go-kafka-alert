package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
	"regexp"
	"github.com/sfreiberg/gotwilio"
	"go-kafka-alert/util"
)

type EventForSMS struct {
	TriggeredEvent db.Event
}

func (event EventForSMS) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "SMS")
	if !channelSupported {
		return messages, errors.New("SMS channel not supported")
	}
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		return messages, errors.New("No recipients found")
	}
	var messageContent = ParseTemplateForMessage(event.TriggeredEvent, "SMS")

	//generate individual messages for each recipient
	for _, rep := range event.TriggeredEvent.Recipient {
		if validatePhone(rep) {
			dateCreated := time.Now()
			message := db.Message{}
			message.AlertId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
			message.Content = messageContent + " " + rep //temp
			message.Recipient = rep
			message.DateCreated = dateCreated
			message.ReferenceId = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventId
			messages = append(messages, message)
		}
	}
	return messages, nil
}

func (event EventForSMS) SendMessage(message db.Message) db.MessageResponse {
	var response = db.MessageResponse{}
	if (db.Message{}) == message {
		return db.MessageResponse{Status:util.FAILED,Response:"MESSAGE EMPTY", TimeOfResponse: time.Now()}
	}
	if message.Content == "" {
		return db.MessageResponse{Status:util.FAILED,Response:"MESSAGE HAS NO CONTENT", TimeOfResponse: time.Now()}
	}
	twilio := gotwilio.NewTwilioClient("", "")
	smsResponse, smsEx, _ := twilio.SendSMS("+15005550006", message.Recipient, message.Content, "", "")
	if smsEx != nil {
		response.Response = smsEx.Message
		response.Status = strconv.Itoa(smsEx.Status)
		response.TimeOfResponse = time.Now()
		return response
	}
	timeSent, err := smsResponse.DateSentAsTime()
	if err != nil {
		timeSent = time.Now()
	}
	response.Response = smsResponse.Body
	response.Status = smsResponse.Status
	response.TimeOfResponse = timeSent
	return response
}

func validatePhone(phone string) bool {
	re := regexp.MustCompile("[0-9]+")
	return re.MatchString(phone)
}

