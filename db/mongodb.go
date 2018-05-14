package db

import (
	"go-kafka-alert/config"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	messageID       = "messageId"
	messageRef      = "reference"
	messageResponse = "messageResponse"
)

var db, _ = dialDB()
var configuration = config.AppConfiguration
var errorLogger = config.Error

// IndexMessage : Index Message
func (message *Message) IndexMessage() error {
	var er error
	if er = db.C(configuration.DbConfig.Collection).Insert(message); er != nil {
		errorLogger.Println("Error indexing message " + er.Error())
	}
	return er
}

//FindMessage : Find Message by ID
func (message Message) FindMessage(ID string) (Message, error) {
	var msg Message
	var err error
	if err := db.C(configuration.DbConfig.Collection).Find(bson.M{messageID: ID}).One(&msg); err != nil {
		errorLogger.Println("Error finding message by Id : " + ID + err.Error())
	}
	return msg, err
}

// RemoveMessage : Remove Message by ID
func (message *Message) RemoveMessage(ID string) bool {
	if err := db.C(configuration.DbConfig.Collection).Remove(bson.M{messageID: ID}); err != nil {
		return false
	}
	return true
}

// UpdateResponse : Update Message with Response
func (message *Message) UpdateResponse(ID string, response MessageResponse) (Message, error) {
	var msg Message
	err := db.C(configuration.DbConfig.Collection).Update(bson.M{messageID: ID},
		bson.M{"$set": bson.M{messageResponse: response}})
	if err != nil {
		errorLogger.Println("Error updating message " + err.Error())
		return msg, err
	}
	msg.MessageResponse = response
	return msg, err
}

// FindAllMessagesByReference : Find messages by Reference
func FindAllMessagesByReference(reference string) ([]Message, error) {
	var msgs []Message //add limit and sort
	var err error
	if err = db.C(configuration.DbConfig.Collection).Find(bson.M{messageRef: reference}).All(&msgs); err != nil {
		errorLogger.Println("Error finding message by reference " + err.Error())
	}
	return msgs, err
}

// CountAllMessagesByReference : Count by Reference
func CountAllMessagesByReference(reference string) int {
	size, _ := db.C(configuration.DbConfig.Collection).Find(bson.M{messageRef: reference}).Count()
	return size
}

// RemoveAllMessagesByReference : Remove Messages by Reference
func RemoveAllMessagesByReference(reference string) {
	db.C(configuration.DbConfig.Collection).RemoveAll(bson.M{messageRef: reference})
}

func dialDB() (*mgo.Database, error) {
	var db *mgo.Database
	_, err := mgo.Dial(configuration.DbConfig.MongoHost)
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{configuration.DbConfig.MongoHost},
		Timeout:  60 * time.Second,
		Database: configuration.DbConfig.MongoDB,
		Username: configuration.DbConfig.MongoDBUsername,
		Password: configuration.DbConfig.MongoDBPassword,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		errorLogger.Println("Error connecting to database " + err.Error())
		return db, err
	}
	index := mgo.Index{
		Key:        []string{messageID},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db = session.DB(configuration.DbConfig.MongoDB)
	db.C(configuration.DbConfig.Collection).EnsureIndex(index)
	return db, err
}
