package util

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"log"
	"io"
	"strings"
	"fmt"
)

const (
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

const (
	TRACE = "TRACE"
	ERROR = "ERROR"
	WARNING = "WARNING"
	INFO = "INFO"
)

var (
	AppConfiguration *Configuration
	Error *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Trace *log.Logger
)

type LOG_LEVEL *string

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
	Workers    int `json:"workers"`
	DbConfig   DBConfig `json:"dbConfig"`
	SmsConfig  SMSConfig `json:"smsConfig"`
	SmtpConfig SMTPConfig `json:"emailConfig"`
	Templates  map[string]string `json:"templates"`
}

func SetLogLevel(logLvl *string) {
	f, err := os.OpenFile("go_kafka_alert.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err.Error())
	}
	logLevel := *logLvl
	switch strings.ToUpper(logLevel) {
	case TRACE:
		initLog(f, f, f, f, true)
		return
	case INFO:
		initLog(ioutil.Discard, f, f, f, true)
		return
	case WARNING:
		initLog(ioutil.Discard, ioutil.Discard, f, f, true)
		return
	case ERROR:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, f, true)
		return
	default:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard,
			false)
		f.Close()
		return
	}
}

func NewConfiguration() {
	var jsonConfig *os.File
	dir, _ := filepath.Abs("../")
	jsonConfig, err := os.Open(dir + "/configuration.json")
	if err != nil {
		dir, _ := filepath.Abs("./")
		jsonConfig, err = os.Open(dir + "/configuration.json")
		if err != nil {
			fmt.Println("Error reading configuration file " + err.Error())
			return
		}
	}
	defer jsonConfig.Close()
	byteValue, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		fmt.Println("Error reading configuration file " + err.Error())
		return
	}
	er := json.Unmarshal(byteValue, &AppConfiguration)
	if er != nil {
		fmt.Println("Error parsing json configuration file " + err.Error())
		return
	}
	fmt.Println("Setting configuration")
	return
}

func (config *Configuration) GetTemplate(templateId string) string {
	return AppConfiguration.Templates[templateId]
}

func initLog(traceHandle io.Writer, infoHandle io.Writer,
warningHandle io.Writer, errorHandle io.Writer, isFlag bool) {
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile | log.LstdFlags
	}
	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)
}