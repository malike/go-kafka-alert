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
	UnmappedData   map[string]string
	EventType   string
	Description string
	DateCreated time.Time
}

type Message struct {
	_Id             bson.ObjectId `bson:"_id,omitempty"`
	MessageId       string `bson:"messageId,omitempty"`
	Reference       string `bson:"reference,omitempty"`
	AlertId         string `bson:"alertId,omitempty"`
	Subject         string `bson:"subject,omitempty"`
	Content         string `bson:"content,omitempty"`
	Recipient       string `bson:"recipient,omitempty"`
	FileAttached    string `bson:"fileAttached,omitempty"`
	MessageResponse MessageResponse `bson:"messageResponse,omitempty"`
	DateCreated     time.Time `bson:"dateCreated,omitempty"`
}

type MessageResponse struct {
	Response       string `bson:"response,omitempty"`
	Status         string `bson:"status,omitempty"`
	APIStatus      string `bson:"apiStatus,omitempty"`
	TimeOfResponse time.Time `bson:"timeOfResponse,omitempty"`
}

