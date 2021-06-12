package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleRequest(ctx context.Context, event Event) (string, error) {
	fmt.Println("Here")
	log.Println("LogHere")
	log.Printf("First Name: %s\n", event.FirstName)
	log.Printf("Last Name: %s\n", event.LastName)
	log.Printf("Age: %d\n", event.Age)
	r := Response{
		Name: fmt.Sprintf("%s %s", event.FirstName, event.LastName),
		Age:  event.Age,
	}
	data, _ := json.Marshal(r)
	rToString := string(data)
	log.Println(rToString)
	return rToString, nil
}

func main() {
	lambda.Start(handleRequest)
}
