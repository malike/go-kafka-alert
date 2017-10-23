package main

import (
	"go-kafka-alert/util"
	"flag"
	"strconv"
	"sync"
	"go-kafka-alert/service"
)

var wg sync.WaitGroup

func main() {

	logLevel := flag.String("loglevel", "error", "Possible options warn,trace,error,info")
	flag.Parse()
	util.LogLevel = *logLevel
	util.NewConfiguration()
	util.Trace.Println("Starting up Service with Log level '" + *logLevel + "'")
	util.Trace.Println("Configuration file loaded successfully with '" +
		strconv.Itoa(len(util.AppConfiguration.Templates)) + "' templates and " +
		strconv.Itoa(util.AppConfiguration.Workers) + " workers processing events")
	for {
		//one extractor
		events := service.GetEventFromKafkaStream()

		if len(events) > 0 {

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
				for i := 1; i < util.AppConfiguration.Workers; i++ {
					//slice events ..using batchSize
					eventBatch := events
					go service.EventProcessorForChannel(events)
				}
			}
			wg.Wait()
		}
	}
}