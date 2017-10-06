package util

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

const (
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

type SMTPConfig struct {
	Host string `json:"smtpServerHost"`
        Port string `json:"smtpServerPort"`
        Username string `json:"emailAuthUserName"`
        Password string `json:"emailAuthPassword"`
        EmailSender string `json:"emailSender"`
}

type SMSConfig struct {
	UserName string `json:"twilioAccountId"`
        Password   string `json:"twilioAuthToken"`
        SenderName string `json:"smsSender"`
}


type Configuration struct {
	SmsConfig  SMSConfig `json:"smsConfig"`
	SmtpConfig SMTPConfig `json:"emailConfig"`
}

func LoadConfiguration() (Configuration,error) {
	jsonConfig, err := os.Open("../configuration.json")
	if err != nil {
		return Configuration{},err
	}
	defer jsonConfig.Close()
	byteValue, err := ioutil.ReadAll(jsonConfig)
	if err != nil{
		return Configuration{},err
	}
	var conf Configuration
        er := json.Unmarshal(byteValue, &conf)
	if er != nil{
		return Configuration{},er
	}
	return conf,nil
}
