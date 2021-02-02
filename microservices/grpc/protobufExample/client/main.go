package main

import (
	"bytes"
	fmt "fmt"
	"net/http"

	"github.com/williamnoble/projects/grpc/protobufExample/transport"

	"github.com/golang/protobuf/proto"
)

func main() {
	myClient := transport.Client{Id: 211, Name: "William Noble", Email: "william@mail.com", Country: "Wales"}
	clientInbox := make([]*transport.Client_Mail, 0, 20)
	clientInbox = append(clientInbox,
		&transport.Client_Mail{RemoteEmail: "William@AnotherMail.Com", Body: "First First First"},
		&transport.Client_Mail{RemoteEmail: "Wills@NotAnotherMail.Co", Body: "Second Mail Box checking in!"})
	myClient.Inbox = clientInbox

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3030", "", bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}
