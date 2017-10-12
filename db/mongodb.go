package db

import (
	"gopkg.in/mgo.v2"
	"go-kafka-alert/util"
	"time"
	"gopkg.in/mgo.v2/bson"
)

var db, _ = dialDB()

func (message *Message) IndexMessage() error {
	er := db.C(util.AppConfiguration.DbConfig.Collection).Insert(message)
	return er
}

func (message Message) FindMessage(Id string) (Message, error) {
	var msg Message
	err := db.C(util.AppConfiguration.DbConfig.Collection).Find(bson.M{"messageid":Id}).One(&msg)
	return msg, err
}

func (message *Message) RemoveMessage(Id string) bool {
	if err := db.C(util.AppConfiguration.DbConfig.Collection).Remove(bson.M{"messageid":Id}); err != nil {
		return false
	}
	return true
}
func (message *Message) UpdateResponse(Id string, response MessageResponse) (Message, error) {
	var msg Message
	err := db.C(util.AppConfiguration.DbConfig.Collection).Update(bson.M{"messageid":Id},
		bson.M{"$set":bson.M{"messageresponse": response}})
	if err != nil {
		return msg, err
	}
	msg.MessageResponse = response
	return msg, err
}

func GetTemplate(templateId string) Template {
	return Template{}
}

func dialDB() (*mgo.Database, error) {
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
		return db, err
	}
	index := mgo.Index{
		Key:        []string{"messageid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db = session.DB(util.AppConfiguration.DbConfig.MongoDB)
	db.C(util.AppConfiguration.DbConfig.Collection).EnsureIndex(index)
	return db, err
}
