package db

import "time"

type Event struct {
	EventId     string
	Subject     string
	Channel     map[string]bool
	Recipient   []string
	EventType   string
	Description string
	DateCreated time.Time
}

type Message struct {
	Id           string
	ReferenceId  string
	AlertId      string
	Subject      string
	Content      string
	Recipient    string
	FileAttached string
	ApiResponse  MessageResponse
	DateCreated  time.Time
}

type MessageResponse struct {
	Response       string
	Status         string
	APIStatus      string
	TimeOfResponse time.Time
}

type Template struct {
	Id             string
	Content        string
	TimeOfResponse time.Time
}

