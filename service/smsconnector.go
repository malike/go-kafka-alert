package service

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/malike/go-kafka-alert/db"

	"github.com/malike/go-kafka-alert/config"

	"github.com/sfreiberg/gotwilio"
)

//EventForSMS : SMS implementation for SMS
type EventForSMS struct {
	TriggeredEvent db.Event
}

//ParseTemplate : Parsing Template for SMS
func (event EventForSMS) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "SMS")
	if !channelSupported {
		config.Trace.Println("Dropping event ['" + event.TriggeredEvent.EventID + "']. SMS channel not supported.")
		return messages, errors.New("SMS channel not supported")
	}
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		config.Trace.Println("Dropping event ['" + event.TriggeredEvent.EventID + "']. No recipient found.")
		return messages, errors.New("No recipients found")
	}
	var messageContent, _ = ParseTemplateForMessage(event.TriggeredEvent, "SMS")

	//generate indivIDual messages for each recipient
	for _, rep := range event.TriggeredEvent.Recipient {
		if validatePhone(rep) {
			dateCreated := time.Now()
			message := db.Message{}
			message.AlertID = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventID
			message.Subject = event.TriggeredEvent.Subject
			message.Reference = event.TriggeredEvent.EventID + "SMS"
			message.Content = messageContent + " " + rep //temp
			message.Recipient = rep
			message.DateCreated = dateCreated
			message.UnmappedData = event.TriggeredEvent.UnmappedData
			message.MessageID = strconv.Itoa(dateCreated.Nanosecond()) + rep + event.TriggeredEvent.EventID
			messages = append(messages, message)
		} else {
			config.Error.Println("Phone number not valid ['" + rep + "']")
		}
	}
	return messages, nil
}

//SendMessage : Message Sending  for SMS
func (event EventForSMS) SendMessage(message db.Message) db.MessageResponse {
	var response = db.MessageResponse{}
	if message.Content == "" {
		config.Error.Println("Sending  Failed. Message body empty")
		return db.MessageResponse{Status: config.FAILED, Response: "MESSAGE HAS NO CONTENT", TimeOfResponse: time.Now()}
	}
	if config.AppConfiguration.SmsConfig.UserName == "" || config.AppConfiguration.SmsConfig.Password == "" ||
		config.AppConfiguration.SmsConfig.SenderName == "" {
		config.Error.Println("Sending  Failed. SMS Config not available")
		return db.MessageResponse{Status: config.FAILED, Response: "SMS Config not available", TimeOfResponse: time.Now()}
	}
	twilio := gotwilio.NewTwilioClient(config.AppConfiguration.SmsConfig.UserName, config.AppConfiguration.SmsConfig.Password)
	twilioSmsResponse, smsEx, _ := twilio.SendSMS(config.AppConfiguration.SmsConfig.SenderName, message.Recipient, message.Content, "", "")
	if smsEx != nil {
		response.Response = smsEx.Message
		response.APIStatus = strconv.Itoa(smsEx.Status)
		response.Status = config.SUCCESS
		response.TimeOfResponse = time.Now()
		config.Info.Println("SMS sent to  ['" + message.Recipient + "']")
		return response
	}
	timeSent, err := twilioSmsResponse.DateSentAsTime()
	if err != nil {
		timeSent = time.Now()
		config.Error.Println("Sending  Failed. " + err.Error())
	}
	response.Response = twilioSmsResponse.Body
	response.APIStatus = strconv.Itoa(smsEx.Status)
	response.Status = config.FAILED
	response.TimeOfResponse = timeSent
	return response
}

func validatePhone(phone string) bool {
	re := regexp.MustCompile("[0-9]+") //temporal regex
	return re.MatchString(phone)
}
