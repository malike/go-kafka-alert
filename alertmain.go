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
		events := service.GetEventFromKafkaStream()
		if len(events) > 0 {
			util.Info.Println("Processing " + strconv.Itoa(len(events)) + " events")
			for _, event := range events {
				service.EventProcessorForChannel(event)
			}
		}
	}
}