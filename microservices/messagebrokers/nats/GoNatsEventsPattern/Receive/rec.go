package main

import (
	"fmt"
	natsProto "github.com/nats-io/nats.go/encoders/protobuf"
	"natsEventPattern/Transport"

	"github.com/nats-io/nats.go"
)

const (
	natsURL string = "nats://localhost:4444"
)

func main() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		fmt.Println(err)
	}
	ec, err := nats.NewEncodedConn(nc, natsProto.PROTOBUF_ENCODER)
	defer ec.Close()
	// Subscribe(Subject, CallbackFn(Subject, Reply, Response))
	// then Publish(subject, interface)

	_, _ = ec.Subscribe("Messaging.Text.Standard", func(m *Transport.TextMessage) {
		fmt.Println("Got standard message: \"", m.Body, "\" with the Id ", m.Id, ".")
	})

	_, _ = ec.Subscribe("Messaging.Text.Respond", func(subject, reply string, m *Transport.TextMessage) {
		fmt.Println("Go ask for response message: \"", m.Body, "\" with the Id ", m.Id, ".")
		newMessage := Transport.TextMessage{Id: int32(m.Id), Body: "Responding!!!"}
		_ = ec.Publish(reply, &newMessage)
	})

	receiveChannel := make(chan *Transport.TextMessage)
	_, _ = ec.BindRecvChan("Messaging.Text.Channel", receiveChannel)
	for m := range receiveChannel {
		fmt.Println("Got channel message: \"", m.Body, "\" with the Id ", m.Id, ".")
	}
}
