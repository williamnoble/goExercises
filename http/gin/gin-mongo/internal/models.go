package internal

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `json:"username,omitempty" bson:"username"`
	Password string        `json:"password,omitempty" bson:"password"`
	Fullname string        `json:"fullname,omitempty" bson:"fullname"`
}

type Update struct {
	Fullname string `json:"fullname" bson:"fullname"`
}
