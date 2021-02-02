package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/golang/protobuf/proto"

	"github.com/williamnoble/Projects/messagebrokers/nats/GoNatsTwoService/Transport"
)

var nc *nats.Conn
var err error

func main() {

	natsURL := "nats://localhost:4444"
	nc, err = nats.Connect(natsURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	nc.QueueSubscribe("TimeTeller", "TimeTellers", replyWithTime)
	select {}
}

func replyWithTime(m *nats.Msg) {
	t := string(time.Now().Format(time.RFC3339))
	curTime := Transport.Time{Time: t}
	data, err := proto.Marshal(&curTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("time service:  Replying to ", m.Reply)
	nc.Publish(m.Reply, data)

}
