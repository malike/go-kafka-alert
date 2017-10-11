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
	err := db.C(util.AppConfiguration.DbConfig.Collection).Find(bson.M{"_id":Id}).One(&msg)
	return msg, err
}

func (Message *Message) UpdateResponse(response MessageResponse) bool {
	return false
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
	db = session.DB(util.AppConfiguration.DbConfig.MongoDB)
	return db, err
}
