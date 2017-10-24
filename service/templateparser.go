package service

import (
	"go-kafka-alert/db"
	"go-kafka-alert/util"
	"text/template"
	"bytes"
)

func ParseTemplateForMessage(event db.Event, channel string) (string, error) {
	var parse string
	temp := util.AppConfiguration.Templates[event.EventType + "_" + channel]
	if len(temp) == 0 {
		util.Trace.Println("Template not available. Sending description of event as content")
		return event.Description, nil
	}
	t := template.New("Template")
	t, err := t.Parse(temp)
	if err != nil {
		util.Error.Println("Error parsing template. Event dropped. Reason: "+err.Error())
		return parse, err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, event)
	if err != nil {
		util.Error.Println("Error parsing template. Event dropped. Reason: "+err.Error())
		return parse, err
	}
	return tpl.String(), err
}