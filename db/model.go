package db

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Event struct
type Event struct {
	EventID      string            `json:"eventId,omitempty"`
	Subject      string            `json:"subject"`
	Channel      map[string]bool   `json:"channel"`
	Recipient    []string          `json:"recipient"`
	UnmappedData map[string]string `json:"unmappedData"`
	EventType    string            `json:"eventType"`
	Description  string            `json:"description"`
	DateCreated  time.Time         `json:"dateCreated,omitempty"`
}

// Message struct
type Message struct {
	ID              bson.ObjectId     `bson:"_id,omitempty"`
	MessageID       string            `bson:"messageId,omitempty"`
	Reference       string            `bson:"reference,omitempty"`
	AlertID         string            `bson:"alertId,omitempty"`
	UnmappedData    map[string]string `json:"unmappedData"`
	Subject         string            `bson:"subject,omitempty"`
	Content         string            `bson:"content,omitempty"`
	Recipient       string            `bson:"recipient,omitempty"`
	FileAttached    string            `bson:"fileAttached,omitempty"`
	MessageResponse MessageResponse   `bson:"messageResponse,omitempty"`
	DateCreated     time.Time         `bson:"dateCreated,omitempty"`
}

// MessageResponse struct
type MessageResponse struct {
	Response       string    `bson:"response,omitempty"`
	Status         string    `bson:"status,omitempty"`
	APIStatus      string    `bson:"apiStatus,omitempty"`
	TimeOfResponse time.Time `bson:"timeOfResponse,omitempty"`
}
