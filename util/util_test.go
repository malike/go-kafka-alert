package util

import (
	"testing"
	"fmt"
)

func TestLoadConfiguration(t *testing.T) {
	conf, e := NewConfiguration()
	if e != nil {
		t.Errorf("Configuration can't be nil " + e.Error())
		t.FailNow()
	}
	if conf.SmtpConfig.EmailSender == "" {
		t.Errorf("Required configuration not loaded ")
		t.FailNow()
	}
	t.Log("Email Sender available as " + conf.SmtpConfig.EmailSender)
}

func TestLoadConfigurationWithTemplates(t *testing.T) {
	conf, e := NewConfiguration()
	if e != nil {
		t.Error("Configuration can't be nil " + e.Error())
		t.FailNow()
	}
	if len(conf.Templates) == 0 {
		t.Error("Required configuration not loaded. No Templats found ")
		t.FailNow()
	}
	t.Log(fmt.Print("Templates found '%d' ", len(conf.Templates)))
	//for _, temp := range conf.Templates {
	for k, v := range conf.Templates {
		t.Log("Template ID '" + k + "' ==> " + v)
	}
	//}
}

func TestConfiguration_GetTemplate(t *testing.T) {
	conf, e := NewConfiguration()
	if e != nil {
		t.Errorf("Configuration can't be nil " + e.Error())
		t.FailNow()
	}
	if len(conf.Templates) == 0 {
		t.Errorf("Required configuration not loaded. No Templats found ")
		t.FailNow()
	}
	var randomTemplateId string
	var randomTemplateContent string
	for k, v := range conf.Templates {
		randomTemplateId = k
		randomTemplateContent = v
		break
	}
	templateContent := conf.GetTemplate(randomTemplateId)
	if templateContent != randomTemplateContent {
		t.Errorf("Could not fetch template")
		t.FailNow()
	}
	t.Log("Temaplate ID '" + randomTemplateId + "' matched it's  content '" + randomTemplateContent + "'")
}

func BenchmarkConfiguration_GetTemplate(b *testing.B) {
	conf, _ := NewConfiguration()
	for i := 0; i < b.N; i++ {
		var randomTemplateId string
		for k := range conf.Templates {
			randomTemplateId = k
			break
		}
		conf.GetTemplate(randomTemplateId)
	}
}