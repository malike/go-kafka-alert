package main

import (
	"go-kafka-alert/util"
	"log"
)

func main() {
	if util.AppConfiguration == nil {
		var err error
		util.AppConfiguration, err = util.NewConfiguration()
		if err != nil {
			log.Fatal("Application can not start without configuration. Error " + err.Error())
		}
	}
	log.Println("Starting up Service")
}