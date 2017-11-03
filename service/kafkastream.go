package service

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go-kafka-alert/db"
	"go-kafka-alert/util"
	"strconv"
)

var KafkaConsumer *kafka.Consumer

func GetEventFromKafkaStream() ([]db.Event, error) {
	events := []db.Event{}
	var err error
	ev := <-KafkaConsumer.Events()
	switch e := ev.(type) {
	case *kafka.Message:
		err = json.Unmarshal([]byte(string(e.Value)), events)
		if err != nil {
			event := db.Event{}
			json.Unmarshal([]byte(string(e.Value)), event)
			if len(event.EventType) > 0 {
				events = append(events, event)
			}
		}
	case kafka.Error:
		util.Error.Println("Error : " + e.Error())
		return events, errors.New("Error : " + e.Error())
	}
	return events, err
}

func NewKafkaConsumer() {
	var err error
	KafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               util.AppConfiguration.KafkaConfig.BootstrapServers,
		"group.id":                        util.AppConfiguration.KafkaConfig.KafkaGroupId,
		"session.timeout.ms":              util.AppConfiguration.KafkaConfig.KafkaTimeout,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": util.AppConfiguration.KafkaConfig.KafkaTopicConfig}})
	if err != nil {
		util.Error.Println("Error creating consumer : " + err.Error())
		return
	}
	kafkaTopic := []string{util.AppConfiguration.KafkaConfig.KafkaTopic}
	util.Trace.Println("Kafka Consumer created successfully. Listening on " + util.AppConfiguration.KafkaConfig.KafkaTopic)
	err = KafkaConsumer.SubscribeTopics(kafkaTopic, nil)

}

func EventProcessorForChannel(events []db.Event) {
	if len(events) > 0 {
		util.Info.Print("Processing " + strconv.Itoa(len(events)))
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
}

func ProcessEvent(eventForMessage EventForMessage) {
	messages, err := eventForMessage.ParseTemplate()
	if err != nil {
		util.Info.Print("Error parsing template Error :" + err.Error() + "")
	} else {
		for _, msg := range messages {
			//index message
			msg.IndexMessage()

			response := eventForMessage.SendMessage(msg)
			//index response
			msg.UpdateResponse(msg.MessageId, response)

		}
	}
}
