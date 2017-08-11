package db

import "time"

type Message struct {
	Id          string
	ReferenceId string
	AlertId     string
	Content     string
	Recipient   string
	ApiResponse MessageResponse
	DateCreated time.Time
}

type MessageResponse struct {
	Response       string
	TimeOfResponse time.Time
}

type Template struct {
	ID             string
	Content        string
	TimeOfResponse time.Time
}
