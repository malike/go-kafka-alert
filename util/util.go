package util

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

var AppConfiguration, _ = NewConfiguration()

const (
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

type SMTPConfig struct {
	Host        string `json:"smtpServerHost"`
	Port        int `json:"smtpServerPort"`
	Username    string `json:"emailAuthUserName"`
	Password    string `json:"emailAuthPassword"`
	EmailFrom   string `json:"emailFrom"`
	EmailSender string `json:"emailSender"`
	TLS         bool `json:"tls"`
}

type SMSConfig struct {
	UserName   string `json:"twilioAccountId"`
	Password   string `json:"twilioAuthToken"`
	SenderName string `json:"smsSender"`
}

type DBConfig struct {
	MongoHost       string `json:"mongoHost"`
	MongoPort       int `json:"mongoPort"`
	MongoDBUsername string `json:"mongoDBUsername"`
	MongoDBPassword string `json:"mongoDBPassword"`
	MongoDB         string `json:"mongoDB"`
	Collection      string `json:"collection"`
}

type Configuration struct {
	DbConfig   DBConfig `json:"dbConfig"`
	SmsConfig  SMSConfig `json:"smsConfig"`
	SmtpConfig SMTPConfig `json:"emailConfig"`
}

func NewConfiguration() (*Configuration, error) {
	var conf *Configuration = nil
	var jsonConfig *os.File
	var err error
	dir, _ := filepath.Abs("../")
	jsonConfig, err = os.Open(dir + "/configuration.json")
	if err != nil {
		dir, _ := filepath.Abs("./")
		jsonConfig, err = os.Open(dir + "/configuration.json")
		if err == nil {
			return conf, err
		}
	}
	defer jsonConfig.Close()
	byteValue, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		return conf, err
	}
	er := json.Unmarshal(byteValue, &conf)
	if er != nil {
		return conf, er
	}
	return conf, nil
}
