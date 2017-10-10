package db

import (
	"gopkg.in/mgo.v2"
	"go-kafka-alert/util"
	"time"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Session

func (message *Message) IndexMessage() error {
	collection, err := dialdb()
	if err != nil {
		return err
	}
	er := collection.Insert(message)
	defer closedb()
	return er
}

func (message Message) FindMessage(Id string) (Message,error){
	var msg Message
	var err error
	collection, err := dialdb()
	if err != nil {
		return msg,err
	}
	err = collection.FindId(bson.ObjectIdHex(Id)).One(&msg)
	return msg,err
}

func (Message *Message) UpdateResponse(response MessageResponse) bool {
	return false
}

func GetTemplate(templateId string) Template {
	return Template{}
}

func dialdb() (*mgo.Collection, error) {
	var collection *mgo.Collection
	var err error
	db, err = mgo.Dial(util.AppConfiguration.DbConfig.MongoHost)
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{util.AppConfiguration.DbConfig.MongoHost},
		Timeout:  60 * time.Second,
		Database: util.AppConfiguration.DbConfig.MongoDB,
		Username: util.AppConfiguration.DbConfig.MongoDBUsername,
		Password: util.AppConfiguration.DbConfig.MongoDBPassword,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		return collection, err
	}
	collection = session.DB(util.AppConfiguration.DbConfig.MongoDB).
		C(util.AppConfiguration.DbConfig.Collection)
	return collection, err
}

func closedb() {
	db.Close()
}