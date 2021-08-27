package internal

import (
	"gopkg.in/mgo.v2"
	"log"
)

func InitialiseDatabase() *mgo.Collection {
	mongoURL := "mongodb://localhost:27017"
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatal("failed to connect to mongo", err)
	}
	collection := session.DB("gin-mongo").C("users")
	return collection
}
