package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/williamnoble/goExercises/messagebrokers/nats/GoNatsTwoService/Transport"
	// "github.com/nats-io/go-nats"
)

var nc *nats.Conn
var err error
var users map[string]string

func main() {

	natsURL := "nats://localhost:4444"
	nc, err = nats.Connect(natsURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	users = make(map[string]string)
	users["1"] = "Bob"
	users["2"] = "John"
	users["3"] = "Dan"
	users["4"] = "Kate"

	nc.QueueSubscribe("UserNameById", "userNameByIdProviders", replyWithUserId)
	select {}
}

func replyWithUserId(m *nats.Msg) {
	myUser := Transport.User{}
	err := proto.Unmarshal(m.Data, &myUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	myUser.Name = users[myUser.Id]
	fmt.Println("user service:  User with Id", myUser.Id, "is", users[myUser.Id])
	data, err := proto.Marshal(&myUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user service:  Replying to ", m.Reply)
	fmt.Println("user service:  Replying with User Data: ", data)
	nc.Publish(m.Reply, data)
}
