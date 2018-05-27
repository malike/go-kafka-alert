package service

import (
	"bytes"
	"text/template"

	"github.com/malike/go-kafka-alert/config"
	"github.com/malike/go-kafka-alert/db"
)

//ParseTemplateForMessage : Parses Template
func ParseTemplateForMessage(event db.Event, channel string) (string, error) {
	var parse string
	temp := config.AppConfiguration.Templates[event.EventType+"_"+channel]
	if len(temp) == 0 {
		config.Trace.Println("Template not available. Sending description of event as content")
		return event.Description, nil
	}
	t := template.New("Template")
	t, err := t.Parse(temp)
	if err != nil {
		config.Error.Println("Error parsing template. Event dropped. Reason: " + err.Error())
		return parse, err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, event)
	if err != nil {
		config.Error.Println("Error parsing template. Event dropped. Reason: " + err.Error())
		return parse, err
	}
	return tpl.String(), err
}
