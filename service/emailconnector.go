package service

import (
	"errors"
	"github.com/malike/go-kafka-alert/db"
	"github.com/malike/go-kafka-alert/util"
	"github.com/smancke/mailck"
	"gopkg.in/gomail.v2"
	"strconv"
	"time"
)

var smtpDialer = gomail.NewPlainDialer(util.AppConfiguration.SmtpConfig.Host,
	util.AppConfiguration.SmtpConfig.Port,
	util.AppConfiguration.SmtpConfig.Username,
	util.AppConfiguration.SmtpConfig.Password)

//EventForEmail : Email implementation for SMS
type EventForEmail struct {
	TriggeredEvent db.Event
}

//ParseTemplate : Template Parser Implementation for Email
func (event EventForEmail) ParseTemplate() ([]db.Message, error) {
	var messages []db.Message
	channelSupported := CheckChannel(event.TriggeredEvent, "EMAIL")
	if !channelSupported {
		util.Trace.Println("Dropping event ['" + event.TriggeredEvent.EventID + "']. EMAIL channel not supported.")
		return messages, errors.New("EMAIL channel not supported")
	}
	numOfRecipient := len(event.TriggeredEvent.Recipient)
	if numOfRecipient <= 0 {
		util.Trace.Println("Dropping event ['" + event.TriggeredEvent.EventID + "']. No recipient found.")
		return messages, errors.New("no recipients found")
	}
	emailContent, _ := ParseTemplateForMessage(event.TriggeredEvent, "EMAIL")
	//parse each mail separately because it may vary by recipient
	for _, em := range event.TriggeredEvent.Recipient {
		if validateEmail(em) {
			dateCreated := time.Now()
			message := db.Message{}
			message.Recipient = em
			message.Subject = event.TriggeredEvent.Subject
			message.Reference = event.TriggeredEvent.EventID + "EMAIL"
			message.DateCreated = dateCreated
			message.AlertID = event.TriggeredEvent.EventID + "_EMAIL_" + em
			message.Content = emailContent
			message.UnmappedData = event.TriggeredEvent.UnmappedData
			message.MessageID = strconv.Itoa(dateCreated.Nanosecond()) + em + event.TriggeredEvent.EventID
			messages = append(messages, message)
		} else {
			util.Error.Println("Email address not valid ['" + em + "']")
		}
	}
	return messages, nil
}

//SendMessage : Messaging Sending for Email
func (event EventForEmail) SendMessage(message db.Message) db.MessageResponse {
	if message.Content == "" {
		util.Error.Println("Sending  Failed. Message body empty")
		return db.MessageResponse{Status: util.FAILED, Response: "MESSAGE EMPTY", TimeOfResponse: time.Now()}
	}

	emailResponse := db.MessageResponse{}
	m := gomail.NewMessage()

	s, err := smtpDialer.Dial()
	if err != nil {
		util.Error.Println("Error sending email " + err.Error())
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

	er := gomail.Send(s, m)
	if er != nil {
		emailResponse.Response = er.Error()
		emailResponse.Status = util.FAILED
		emailResponse.TimeOfResponse = time.Now()
		util.Error.Println("Error sending email " + err.Error())
	} else {
		emailResponse.Response = "SENT"
		emailResponse.Status = util.SUCCESS
		emailResponse.TimeOfResponse = time.Now()
		util.Info.Println("Email sent to  ['" + message.Recipient + "']")
	}
	return emailResponse
}

func validateEmail(email string) bool {
	return mailck.CheckSyntax(email)
}
