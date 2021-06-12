package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/williamnoble/goExercises/microservices/messagebrokers/nats/GoNatsTwoService/Transport"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
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
	m := mux.NewRouter()
	fmt.Println("listening on port 8080")
	m.HandleFunc("/{id}", handleUserWithTime)
	_ = http.ListenAndServe(":8080", m)

}

func handleUserWithTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	myUser := Transport.User{Id: vars["id"]}
	curTime := Transport.Time{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("go1-user")
		data, err := proto.Marshal(&myUser)
		fmt.Println("error", err)
		if err != nil || len(myUser.Id) == 0 {
			fmt.Println(err)
			w.WriteHeader(500)
			fmt.Println("Problem with parsing the User Id")
			return
		}
		fmt.Println("Requesting Username Info..")
		msg, err := nc.Request("UserNameById", data, 100*time.Millisecond)
		if err == nil && msg != nil {
			myUserWithName := Transport.User{}
			err := proto.Unmarshal(msg.Data, &myUserWithName)
			fmt.Println("main:  My User:", &myUserWithName)
			if err == nil {
				myUser = myUserWithName
			}
		}
		wg.Done()
	}()

	go func() {
		msg, err := nc.Request("TimeTeller", nil, 100*time.Millisecond)
		if err == nil && msg != nil {
			receivedTime := Transport.Time{}
			err := proto.Unmarshal(msg.Data, &receivedTime)
			if err == nil {
				curTime = receivedTime
			}
		}
		wg.Done()
	}()

	wg.Wait()
	//	fmt.Println("name", myUser.Name)
	_, _ = fmt.Fprintln(w, "Hello ", myUser.Name, " with id ", myUser.Id, ", the time is ", curTime.Time, ".")
}
