package service

import "go-kafka-alert/db"

func ParseTemplateForMessage(event db.Event, channel string) string {
	if channel == "SMS" {
		return "Sample SMS"
	} else if channel == "EMAIL" {
		return "<html><head></head><body> Sample HTML </body></html>"
	} else if channel == "API" {
		return "Sample API"
	} else {
		return event.Description
	}
}