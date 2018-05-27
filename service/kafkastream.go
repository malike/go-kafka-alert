package service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/malike/go-kafka-alert/config"

	"github.com/malike/go-kafka-alert/db"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer instance
var KafkaConsumer = NewKafkaConsumer()

// GetEventFromKafkaStream : Reads events from Kafka
func GetEventFromKafkaStream() ([]db.Event, error) {
	var event db.Event
	var events []db.Event
	var err error
	ev := <-KafkaConsumer.Events()
	switch e := ev.(type) {
	case kafka.AssignedPartitions:
		config.Trace.Println(e.String())
		KafkaConsumer.Assign(e.Partitions)
	case kafka.RevokedPartitions:
		config.Trace.Println(e.String())
		KafkaConsumer.Unassign()
	case *kafka.Message:
		kafkaMessage := string(e.Value)
		config.Trace.Println("DEBUG : Recieved message " + kafkaMessage)
		err = json.Unmarshal([]byte(kafkaMessage), &event)
		if len(event.EventType) > 0 {
			events = append(events, event)
		}
		if err != nil {
			json.Unmarshal([]byte(kafkaMessage), &events)
		}
		if len(events) == 0 {
			config.Error.Println("Error Casting JSON " + kafkaMessage)
		}
	case kafka.Error:
		config.Error.Println("Error : os.Stderr " + e.Error())
		return events, errors.New("Error : " + e.Error())
	case kafka.PartitionEOF:
		config.Trace.Println("Reached " + e.String())
	}
	return events, err
}

// NewKafkaConsumer : New Kafka Consumer
func NewKafkaConsumer() *kafka.Consumer {
	// var err error
	var KafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               config.AppConfiguration.KafkaConfig.BootstrapServers,
		"group.id":                        config.AppConfiguration.KafkaConfig.KafkaGroupID,
		"session.timeout.ms":              config.AppConfiguration.KafkaConfig.KafkaTimeout,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": config.AppConfiguration.KafkaConfig.KafkaTopicConfig}})
	if err != nil {
		config.Error.Println("Error creating consumer : " + err.Error())
		return nil
	}
	kafkaTopic := []string{config.AppConfiguration.KafkaConfig.KafkaTopic}
	config.Trace.Println("Kafka Consumer created successfully. Listening on " + config.AppConfiguration.KafkaConfig.KafkaTopic)
	err = KafkaConsumer.SubscribeTopics(kafkaTopic, nil)
	return KafkaConsumer
}

// EventProcessorForChannel : Event Processor For Channel
func EventProcessorForChannel(events []db.Event) {
	if len(events) > 0 {
		config.Info.Print("Processing " + strconv.Itoa(len(events)))
		for _, event := range events {
			if CheckChannel(event, "SMS") {
				config.Info.Print("Processing " + event.EventID + " for SMS")
				smsChannel := EventForSMS{event}
				ProcessEvent(smsChannel)
			}
			if CheckChannel(event, "EMAIL") {
				config.Info.Print("Processing " + event.EventID + " for EMAIL")
				emailChannel := EventForEmail{event}
				ProcessEvent(emailChannel)
			}
			if CheckChannel(event, "API") {
				config.Info.Print("Processing " + event.EventID + " for API")
				apiChannel := EventForAPI{event}
				ProcessEvent(apiChannel)
			}
		}
	}
}

// ProcessEvent : Process Event
func ProcessEvent(eventForMessage EventForMessage) {
	messages, err := eventForMessage.ParseTemplate()
	if err != nil {
		config.Info.Print("Error parsing template Error :" + err.Error() + "")
	} else {
		for _, msg := range messages {
			//index message
			msg.IndexMessage()

			response := eventForMessage.SendMessage(msg)
			//index response
			msg.UpdateResponse(msg.MessageID, response)

		}
	}
}
