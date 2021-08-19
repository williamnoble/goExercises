package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("mongodb://localhost:27100/test")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"William", "10293 304930"},
		&Person{"Cla", "29383 837272"},
	)
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "William"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone", result.Phone)
}

//localhost:32768
