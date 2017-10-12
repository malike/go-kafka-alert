package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

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
	_Id             bson.ObjectId `bson:"_id,omitempty"`
	MessageId       string
	AlertId         string
	Subject         string
	Content         string
	Recipient       string
	FileAttached    string
	MessageResponse MessageResponse
	DateCreated     time.Time
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

