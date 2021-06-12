//https://github.com/shobhitsharma/lamba-local/blob/master/func/main.go

package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrNameNotProvided = errors.New("no name supplied in request body")
)

// Handler is a lambda fn handler. It uses Amazon API Gateway
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout + stderr => cloudwatch logs
	log.Printf("Processing your lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	response := events.APIGatewayProxyResponse{
		Body:       "Hello" + request.Body,
		StatusCode: 200,
	}
	return response, nil

}

func main() {
	lambda.Start(Handler)
}
