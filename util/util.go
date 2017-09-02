package util

const (
	SERVER_ERROR = "SERVER_ERROR"
	SUCCESS = "SUCCESS"
	FAILED = "FAILED"
)

type Configuration struct {
	TwilioAccountId string
	TwilioAuthToken string
	SMSSenderName   string
	EmailHost string
	EmailSender string
	AuthName string
	Password string
}

