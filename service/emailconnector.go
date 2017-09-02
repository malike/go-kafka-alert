package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
	"github.com/smancke/mailck"
	"net/smtp"
	"go-kafka-alert/util"
)

type EventForEmail struct {
	TriggeredEvent db.Event
}

func (event EventForEmail) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "EMAIL")
	if !channelSupported {
		return messages, errors.New("SMS channel not supported")
	}
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		return messages, errors.New("No recipients found")
	}
	emailContent := ParseTemplateForMessage(event.TriggeredEvent,"EMAIL")
	//parse each mail separately because it may vary by recipient
	for  _ ,em := range event.TriggeredEvent.Recipient{
		if validateEmail(em) {
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.DateCreated = dateCreated
			message.AlertId = event.TriggeredEvent.EventId + "_EMAIL_" + em
			message.Content = emailContent
			message.ReferenceId = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventId
			message.Id = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventId
			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	auth := smtp.PlainAuth(util.Configuration{}.EmailSender, util.Configuration{}.AuthName,
		util.Configuration{}.Password, util.Configuration{}.EmailHost)

	err := smtp.SendMail(util.Configuration{}.EmailHost, auth,
		util.Configuration{}.EmailSender, []string{message.Recipient}, messageToByte(message))
	if err ==nil{
		emailResponse := db.MessageResponse{}
		emailResponse.Response = "SENT"
		emailResponse.Status = util.SUCCESS
		emailResponse.TimeOfResponse = time.Now()
	}
	return db.MessageResponse{};
}

func attachFile() db.Message{
	return db.Message{}
}

func messageToByte(message db.Message) []byte{
 return []byte{}
}

func validateEmail(email string) bool {
	return mailck.CheckSyntax(email)
}



