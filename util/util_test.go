package util

import "testing"

func TestLoadConfiguration(t *testing.T) {
	conf := LoadConfiguration()
	if conf == nil || conf.SmsConfig.SenderName == ""{
		t.Error("Configuration can't be nil")
	}
	t.Log("Email Sender available as "+conf.SmtpConfig.EmailSender)
}

//func BenchmarkLoadConfiguration(b *testing.B) {
//
//}
