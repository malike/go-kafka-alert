package util

import "testing"

func TestLoadConfiguration(t *testing.T) {
	conf,e := LoadConfiguration()
	if e !=  nil {
		t.Errorf("Configuration can't be nil"+e.Error())
		t.FailNow()
	}
	if conf.SmtpConfig.EmailSender == ""{
		t.Errorf("Required configuration not loaded ")
		t.FailNow()
	}
	t.Log("Email Sender available as "+conf.SmtpConfig.EmailSender)
}

//func BenchmarkLoadConfiguration(b *testing.B) {
//
//}
