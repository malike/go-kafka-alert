package service

import (
	"errors"
	"go-kafka-alert/db"
	"time"
	"strconv"
	"github.com/smancke/mailck"
	"go-kafka-alert/util"
	"gopkg.in/gomail.v2"
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
	emailContent := ParseTemplateForMessage(event.TriggeredEvent, "EMAIL")
	//parse each mail separately because it may vary by recipient
	for _, em := range event.TriggeredEvent.Recipient {
		if validateEmail(em) {
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.Subject = event.TriggeredEvent.Subject
			message.DateCreated = dateCreated
			message.AlertId = event.TriggeredEvent.EventId + "_EMAIL_" + em
			message.Content = emailContent
			message.MessageId = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventId
			messages = append(messages, message)
		}
	}

	return messages, nil
}

func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	if message.Content == "" {
		return db.MessageResponse{Status:util.FAILED, Response:"MESSAGE EMPTY", TimeOfResponse: time.Now()}
	}

	emailResponse := db.MessageResponse{}
	m := gomail.NewMessage()

	d := gomail.NewDialer(util.AppConfiguration.SmtpConfig.Host,
		util.AppConfiguration.SmtpConfig.Port,
		util.AppConfiguration.SmtpConfig.Username,
		util.AppConfiguration.SmtpConfig.Password)

	s, err := d.Dial()
	if err != nil {
		emailResponse.Response = err.Error()
		emailResponse.Status = util.FAILED
		emailResponse.TimeOfResponse = time.Now()
		return emailResponse
	}

	m.SetHeader("From", util.AppConfiguration.SmtpConfig.EmailFrom)
	m.SetAddressHeader("To", message.Recipient, message.Recipient)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Content)
	if message.FileAttached != "" {
		m.Attach(message.FileAttached)
	}

	if er := gomail.Send(s, m); er != nil {
		emailResponse.Response = er.Error()
		emailResponse.Status = util.FAILED
		emailResponse.TimeOfResponse = time.Now()
	} else {
		emailResponse.Response = "SENT"
		emailResponse.Status = util.SUCCESS
		emailResponse.TimeOfResponse = time.Now()
	}
	return emailResponse
}

func validateEmail(email string) bool {
	return mailck.CheckSyntax(email)
}



