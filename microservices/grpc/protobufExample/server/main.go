package main

import (
	fmt "fmt"
	"io/ioutil"
	"net/http"

	"github.com/williamnoble/goExercises/microservices/grpc/protobufExample/transport"

	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myClient := transport.Client{}
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		if err := proto.Unmarshal(data, &myClient); err != nil {
			fmt.Println(err)
		}
		println(myClient.Id, ":", myClient.Name, ":", myClient.Email, ":", myClient.Country)
		for _, mail := range myClient.Inbox {
			fmt.Println(mail.RemoteEmail, ":", mail.Body)
		}
	})
	http.ListenAndServe(":3030", nil)
}
