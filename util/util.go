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
}

