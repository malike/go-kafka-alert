package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// SUCCESS Log Level
	SUCCESS = "SUCCESS"
	// FAILED Log Level
	FAILED = "FAILED"
	// TRACE Log Level
	TRACE = "TRACE"
	// ERROR Log Level
	ERROR = "ERROR"
	// WARNING Log Level
	WARNING = "WARNING"
	// INFO Log Level
	INFO = "INFO"
)

var (
	// Error Logger
	Error *log.Logger
	// Info Logger
	Info *log.Logger
	// Warning Logger
	Warning *log.Logger
	// Trace Logger
	Trace *log.Logger
	// LogLevel default value "ERROR"
	LogLevel = "ERROR"
	// ConfigProfile config profile
	ConfigProfile string
	//ApplicationName application name
	ApplicationName string
	//ConfigServer base url
	ConfigServer string
	// AppConfiguration configuration
	AppConfiguration, _ = LoadConfiguration()
)

// SMTPConfig represents SMTPConfig Properties
type SMTPConfig struct {
	Host        string `json:"smtpServerHost"`
	Port        int    `json:"smtpServerPort"`
	Username    string `json:"emailAuthUserName"`
	Password    string `json:"emailAuthPassword"`
	EmailFrom   string `json:"emailFrom"`
	EmailSender string `json:"emailSender"`
	TLS         bool   `json:"tls"`
}

// SMSConfig represents Twilio SMS Config Properties
type SMSConfig struct {
	UserName   string `json:"twilioAccountId"`
	Password   string `json:"twilioAuthToken"`
	SenderName string `json:"smsSender"`
}

// WebhookConfig : Webhook Config Properties
type WebhookConfig struct {
	AppURL string `json:"appURL"`
	AppKey string `json:"appKey"`
}

// KafkaConfig represents Apache Kafka Config Properties
type KafkaConfig struct {
	BootstrapServers string `json:"bootstrapServers"`
	KafkaTopic       string `json:"kafkaTopic"`
	KafkaTopicConfig string `json:"kafkaTopicConfig"`
	KafkaGroupID     string `json:"kafkaGroupId"`
	KafkaTimeout     int    `json:"kafkaTimeout"`
}

// DBConfig represents MongoDB Config Properties
type DBConfig struct {
	MongoHost       string `json:"mongoHost"`
	MongoPort       int    `json:"mongoPort"`
	MongoDBUsername string `json:"mongoDBUsername"`
	MongoDBPassword string `json:"mongoDBPassword"`
	MongoDB         string `json:"mongoDB"`
	Collection      string `json:"collection"`
}

// Configuration represents configuration File
type Configuration struct {
	Workers         int               `json:"workers"`
	LogFileLocation string            `json:"logFileLocation"`
	Log             bool              `json:"log"`
	KafkaConfig     KafkaConfig       `json:"kafkaConfig"`
	DbConfig        DBConfig          `json:"dbConfig"`
	SmsConfig       SMSConfig         `json:"smsConfig"`
	SMTPConfig      SMTPConfig        `json:"emailConfig"`
	WebhookConfig   WebhookConfig     `json:"webhookConfig"`
	Templates       map[string]string `json:"templates"`
}

// LoadConfiguration loads App Config from File
func LoadConfiguration() (Configuration, error) {
	var err error
	var AppConfiguration Configuration
	if len(ConfigProfile) > 0 {
		fmt.Println("Loading configuration ['" + ConfigProfile + "'] from server.")
		AppConfiguration, err = loadConfigFromServer(AppConfiguration)
	} else {
		fmt.Println("Loading configuration file.")
		AppConfiguration, err = loadConfigFromFile(AppConfiguration)
	}
	if err != nil {
		fmt.Println("Error loading configuration. " + err.Error())
		os.Exit(1)
	}
	setLogLevel(LogLevel, &AppConfiguration)
	return AppConfiguration, nil
}

func loadConfigFromServer(configuration Configuration) (Configuration, error) {
	configFile := ConfigServer + ApplicationName + "-" + ConfigProfile + ".json"
	configFileResponse, err := http.Get(configFile)
	if err != nil {
		fmt.Println("Error downloading file [\"" + ConfigProfile + "\" for \"" + ApplicationName + "\"] " + err.Error())
		return configuration, err
	}
	defer configFileResponse.Body.Close()
	byteValue, err := ioutil.ReadAll(configFileResponse.Body)
	if err != nil {
		fmt.Println("Error reading configuration from Config Server " + err.Error())
		return configuration, err
	}
	er := json.Unmarshal(byteValue, &configuration)
	if er != nil {
		fmt.Println("Error parsing json configuration file ")
		return configuration, err
	}
	return configuration, nil
}

func loadConfigFromFile(configuration Configuration) (Configuration, error) {
	var jsonConfig *os.File
	dir, _ := filepath.Abs("../")
	jsonConfig, err := os.Open(dir + "/configuration.json")
	if err != nil {
		dir, _ := filepath.Abs("./")
		jsonConfig, err = os.Open(dir + "/configuration.json")
		if err != nil {
			fmt.Println("Error reading configuration file " + err.Error())
			return configuration, err
		}
	}
	defer jsonConfig.Close()
	byteValue, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		fmt.Println("Error reading configuration file " + err.Error())
		return configuration, err
	}
	er := json.Unmarshal(byteValue, &configuration)
	if er != nil {
		fmt.Println("Error parsing json configuration file ")
		return configuration, err
	}
	return configuration, nil
}

func setLogLevel(logLevel string, Config *Configuration) {
	if Config.Log {
		f, err := os.OpenFile(Config.LogFileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			log.Fatalf("Error opening log file: %s", err.Error())
		}
		if !Config.Log {
			initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard,
				false)
			return
		}
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
			f.Close() // nolint
			return
		}
	}
	initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard,
		false)
}

// GetTemplate gets Template From Config File
func (config *Configuration) GetTemplate(templateID string) string {
	return AppConfiguration.Templates[templateID]
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
