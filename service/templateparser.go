package service

import (
	"go-kafka-alert/db"
	"go-kafka-alert/util"
	"text/template"
	"bytes"
)

func ParseTemplateForMessage(event db.Event, channel string) (string,error) {
	var parse string
	temp := util.AppConfiguration.Templates[event.EventId+"_"+channel]
	if len(temp) == 0{
		return event.Description,nil
	}
	t := template.New("Template")
	t, err := t.Parse(temp)
	if err != nil {
		return parse,err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, event)
	if err != nil {
		return parse,err
	}
	return tpl.String(),err
}

