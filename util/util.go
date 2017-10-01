package util

const (
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

type SMTPConfig struct{
	Host string
	Port string
}

type Configuration struct {
	TwilioAccountId string
	TwilioAuthToken string
	SMSSenderName   string
	SmtpConfig  SMTPConfig

}

func NewConfiguration() *Configuration {
	conf := new(Configuration)
	return conf
}

