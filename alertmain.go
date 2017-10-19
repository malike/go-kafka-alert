package main

import (
	"go-kafka-alert/util"
	"flag"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {

	logLevel := flag.String("loglevel", "error", "Possible options warn,trace,error,info")
	flag.Parse()
	util.SetLogLevel(util.LOG_LEVEL(logLevel))
	util.Trace.Println("Starting up Service with Log level '" + *logLevel + "'")
	util.NewConfiguration()
	util.Trace.Println("Configuration file loaded successfully with '" +
		strconv.Itoa(len(util.AppConfiguration.Templates)) + "' templates and " +
		strconv.Itoa(util.AppConfiguration.Workers) +" workers processing events")
}