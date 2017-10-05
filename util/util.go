package util

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const (
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

type SMTPConfig struct {
	Host string "json:'smtpServerHost'"
        Port string "json:'smtpServerHost'"
        Username string "json:'smtpServerHost'"
        Password string "json:'smtpServerHost'"
        EmailSender string "json:'smtpServerHost'"
}

type SMSConfig struct {
	UserName string "json:'smtpServerHost'"
        Password   string "json:'smtpServerHost'"
        SenderName string "json:'smtpServerHost'"
        URL        string "json:'smtpServerHost'"
}


type Configuration struct {
	SmsConfig  SMSConfig
	SmtpConfig SMTPConfig
}

func LoadConfiguration() Configuration {
	jsonConfig, err := os.Open("configuration.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonConfig.Close()
	byteValue, _ := ioutil.ReadAll(jsonConfig)
	conf := new(Configuration)
        json.Unmarshal(byteValue, &conf)
	return conf
}
