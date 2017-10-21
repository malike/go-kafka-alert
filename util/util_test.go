package util

import (
	"testing"
	"fmt"
)

func TestLoadConfiguration(t *testing.T) {
	NewConfiguration()
	LogLevel = "INFO"
	if AppConfiguration.SmtpConfig.EmailSender == "" {
		t.Errorf("Required configuration not loaded ")
		t.FailNow()
	}
	t.Log("Email Sender available as " + AppConfiguration.SmtpConfig.EmailSender)
}

func TestLoadConfigurationWithTemplates(t *testing.T) {
	NewConfiguration()
	if len(AppConfiguration.Templates) == 0 {
		t.Error("Required configuration not loaded. No Templats found ")
		t.FailNow()
	}
	t.Log(fmt.Print("Templates found '%d' ", len(AppConfiguration.Templates)))
	for k, v := range AppConfiguration.Templates {
		t.Log("Template ID '" + k + "' ==> " + v)
	}
}

func TestConfiguration_GetTemplate(t *testing.T) {
	NewConfiguration()
	if len(AppConfiguration.Templates) == 0 {
		t.Errorf("Required configuration not loaded. No Templats found ")
		t.FailNow()
	}
	var randomTemplateId string
	var randomTemplateContent string
	for k, v := range AppConfiguration.Templates {
		randomTemplateId = k
		randomTemplateContent = v
		break
	}
	templateContent := AppConfiguration.GetTemplate(randomTemplateId)
	if templateContent != randomTemplateContent {
		t.Errorf("Could not fetch template")
		t.FailNow()
	}
	t.Log("Temaplate ID '" + randomTemplateId + "' matched it's  content '" + randomTemplateContent + "'")
}

func BenchmarkConfiguration_GetTemplate(b *testing.B) {
	NewConfiguration()
	for i := 0; i < b.N; i++ {
		var randomTemplateId string
		for k := range AppConfiguration.Templates {
			randomTemplateId = k
			break
		}
		AppConfiguration.GetTemplate(randomTemplateId)
	}
}