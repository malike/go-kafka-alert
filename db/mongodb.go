package db

import (
	"gopkg.in/mgo.v2"
	"go-kafka-alert/util"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const (
	MESSAGE_ID = "messageId"
	MESSAGE_REFERENCE = "reference"
	MESSAGE_RESPONSE = "messageResponse"
)

var db, _ = dialDB()

func (message *Message) IndexMessage() error {
	var er error
	if er = db.C(util.AppConfiguration.DbConfig.Collection).Insert(message); er != nil {
		util.Error.Println("Error indexing message " + er.Error())
	}
	return er
}

func (message Message) FindMessage(Id string) (Message, error) {
	var msg Message
	var err error
	if err := db.C(util.AppConfiguration.DbConfig.Collection).Find(bson.M{MESSAGE_ID:Id}).One(&msg); err != nil {
		util.Error.Println("Error finding message by Id : " + Id + err.Error())
	}
	return msg, err
}

func (message *Message) RemoveMessage(Id string) bool {
	if err := db.C(util.AppConfiguration.DbConfig.Collection).Remove(bson.M{MESSAGE_ID:Id}); err != nil {
		return false
	}
	return true
}
func (message *Message) UpdateResponse(Id string, response MessageResponse) (Message, error) {
	var msg Message
	err := db.C(util.AppConfiguration.DbConfig.Collection).Update(bson.M{MESSAGE_ID:Id},
		bson.M{"$set":bson.M{MESSAGE_RESPONSE: response}})
	if err != nil {
		util.Error.Println("Error updating message " + err.Error())
		return msg, err
	}
	msg.MessageResponse = response
	return msg, err
}

func FindAllMessagesByReference(reference string) ([]Message, error) {
	var msgs []Message //add limit and sort
	var err error
	if err = db.C(util.AppConfiguration.DbConfig.Collection).Find(bson.M{MESSAGE_REFERENCE:reference}).All(&msgs);
		err != nil {
		util.Error.Println("Error finding message by reference " + err.Error())
	}
	return msgs, err
}

func CountAllMessagesByReference(reference string) int {
	size, _ := db.C(util.AppConfiguration.DbConfig.Collection).Find(bson.M{MESSAGE_REFERENCE:reference}).Count()
	return size
}

func RemoveAllMessagesByReference(reference string) {
	db.C(util.AppConfiguration.DbConfig.Collection).RemoveAll(bson.M{MESSAGE_REFERENCE:reference})
}

func dialDB() (*mgo.Database, error) {
	util.NewConfiguration()
	var db *mgo.Database
	_, err := mgo.Dial(util.AppConfiguration.DbConfig.MongoHost)
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{util.AppConfiguration.DbConfig.MongoHost},
		Timeout:  60 * time.Second,
		Database: util.AppConfiguration.DbConfig.MongoDB,
		Username: util.AppConfiguration.DbConfig.MongoDBUsername,
		Password: util.AppConfiguration.DbConfig.MongoDBPassword,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		util.Error.Println("Error connecting to database " + err.Error())
		return db, err
	}
	index := mgo.Index{
		Key:        []string{MESSAGE_ID},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db = session.DB(util.AppConfiguration.DbConfig.MongoDB)
	db.C(util.AppConfiguration.DbConfig.Collection).EnsureIndex(index)
	return db, err
}
