package main

import (
	"flag"
	"os"
	"strconv"
	"sync"

	"go-kafka-alert/config"
	"go-kafka-alert/db"
	"go-kafka-alert/service"
)

func main() {

	logLevel := flag.String("loglevel", "error", "Possible options warn,trace,error,info")
	profile := flag.String("profile", "default", "Configuration profile")
	flag.Parse()
	config.LogLevel = *logLevel
	_, configErr := config.LoadConfiguration(*profile)
	if configErr != nil {
		config.Error.Println("Error loading config. Shutting down ")
		os.Exit(1)
	}
	db.DialDB()
	config.Trace.Println("Starting up Service with Log level '" + *logLevel + "'")
	config.Trace.Println("Configuration file loaded successfully with '" +
		strconv.Itoa(len(config.AppConfiguration.Templates)) + "' templates and " +
		strconv.Itoa(config.AppConfiguration.Workers) + " workers processing events")

	service.NewKafkaConsumer()
	if service.KafkaConsumer == nil {
		config.Error.Println("Error starting Kafka Consumer ")
		os.Exit(1)
	}

	for {
		events, _ := service.GetEventFromKafkaStream()

		if len(events) > 0 {
			var wg sync.WaitGroup

			//if event is enough for one worker, let it handle it
			if len(events) <= config.AppConfiguration.Workers {
				config.Info.Println("Distributing " + strconv.Itoa(len(events)) + " worker of the month")
				wg.Add(1)
				go service.EventProcessorForChannel(events)
			} else {
				wg.Add(config.AppConfiguration.Workers)
				batchSize := len(events) / config.AppConfiguration.Workers
				config.Info.Println("Distributing '" + strconv.Itoa(len(events)) + "' events for '" +
					strconv.Itoa(config.AppConfiguration.Workers) +
					"' workers '" + strconv.Itoa(batchSize) + "' each.")

				//..else share
				currentPointer := 0
				var eventBatch []db.Event
				for i := 1; i <= config.AppConfiguration.Workers; i++ {
					//slice events ..using batchSize
					if i == config.AppConfiguration.Workers {
						eventBatch = events[currentPointer:]
					} else {
						eventBatch = events[currentPointer:batchSize]
					}
					go service.EventProcessorForChannel(eventBatch)
					currentPointer = currentPointer + batchSize + 1
				}
			}
			wg.Wait()
			wg.Done()
		}
	}
}
