package main

import (
	"fmt"
	"natsEventPattern/Transport"
	"time"

	"github.com/nats-io/nats.go"
	natsProto "github.com/nats-io/nats.go/encoders/protobuf"
)

const (
	natsURL string = "nats://localhost:4444"
)

var nc *nats.Conn

func main() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		fmt.Println(err)

	}
	ec, err := nats.NewEncodedConn(nc, natsProto.PROTOBUF_ENCODER)
	defer ec.Close()

	/*
		§§ PUBLISH PATTERN(S) §§
	*/

	// 1. Publish(Subject, Payload)
	for i := 0; i < 5; i++ {
		myMessage := Transport.TextMessage{Id: int32(i), Body: "Hello over Standard"}
		err := ec.Publish("Messaging.Text.Standard", &myMessage)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 2. Request(Subject, Payload, Response, Timeout)
	for i := 5; i < 10; i++ {
		myMessage := Transport.TextMessage{Id: int32(i), Body: "Hello, Please respond in a timely fashion"}
		response := Transport.TextMessage{}
		err := ec.Request("Messaging.Text.Respond", &myMessage, &response, 200*time.Millisecond)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response.Body, " with id ", response.Id)
	}

	// 3. Send via Channel(Subject, Channel)
	sendChannel := make(chan *Transport.TextMessage)
	_ = ec.BindSendChan("Messaging.Text.Channel", sendChannel)
	for i := 10; i < 15; i++ {
		myMessage := Transport.TextMessage{Id: int32(i), Body: "Hello over a simple Channel!"}
		sendChannel <- &myMessage
	}
}
