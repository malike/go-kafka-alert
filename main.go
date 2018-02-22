package main

import (
	"flag"
	"github.com/malike/go-kafka-alert/db"
	"github.com/malike/go-kafka-alert/service"
	"github.com/malike/go-kafka-alert/util"
	"os"
	"strconv"
	"sync"
)

func main() {

	logLevel := flag.String("loglevel", "error", "Possible options warn,trace,error,info")
	flag.Parse()
	util.LogLevel = *logLevel
	util.NewConfiguration()
	util.Trace.Println("Starting up Service with Log level '" + *logLevel + "'")
	util.Trace.Println("Configuration file loaded successfully with '" +
		strconv.Itoa(len(util.AppConfiguration.Templates)) + "' templates and " +
		strconv.Itoa(util.AppConfiguration.Workers) + " workers processing events")

	service.NewKafkaConsumer()
	if service.KafkaConsumer == nil {
		util.Error.Println("Error starting Kafka Consumer ")
		os.Exit(1)
	}

	for {
		events, _ := service.GetEventFromKafkaStream()

		if len(events) > 0 {
			var wg sync.WaitGroup

			//if event is enough for one worker, let it handle it
			if len(events) <= util.AppConfiguration.Workers {
				util.Info.Println("Distributing " + strconv.Itoa(len(events)) + " worker of the month")
				wg.Add(1)
				go service.EventProcessorForChannel(events)
			} else {
				wg.Add(util.AppConfiguration.Workers)
				batchSize := len(events) / util.AppConfiguration.Workers
				util.Info.Println("Distributing '" + strconv.Itoa(len(events)) + "' events for '" +
					strconv.Itoa(util.AppConfiguration.Workers) +
					"' workers '" + strconv.Itoa(batchSize) + "' each.")

				//..else share
				currentPointer := 0
				eventBatch := []db.Event{}
				for i := 1; i <= util.AppConfiguration.Workers; i++ {
					//slice events ..using batchSize
					if i == util.AppConfiguration.Workers {
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
