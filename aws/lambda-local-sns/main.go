//https://github.com/Bouwdie/localstack-example-go/blob/master/main.go
//
// go run main.go --queue "input-queue" --topic-arn "arn:aws:sns:us-east-1:000000000000:result-topic" --bucket "result-bucket"

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	LocalAWSEndpoint = "http://100.88.225.67:4566"
)

type ResultNotification struct {
	DocumentURL string    `json:"document_url"`
	FinishedAt  time.Time `json:"finished_at"`
}

var (
	queueName   = kingpin.Flag("queue", "Queue to read messages from").Required().String()                        //input-queue
	topicARN    = kingpin.Flag("topic-arn", "AWS topic ARN to notify about finished results").Required().String() //arn:aws:sns:eu-west-1:123456789012:result-topic
	bucket      = kingpin.Flag("bucket", "S3 bucket name to publish results to").Required().String()              //result-bucket
	sqsEndpoint = kingpin.Flag("sqs-endpoint", "AWS endpoint for SQS service").Default("").String()
	snsEndpoint = kingpin.Flag("sns-endpoint", "AWS endpoint for SNS service").Default("").String()
	s3Endpoint  = kingpin.Flag("s3-endpoint", "AWS endpoint for s3 service").Default("").String()
)

func main() {

	kingpin.Parse()
	c := credentials.NewStaticCredentials("test", "test", "")
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(endpoints.EuWest2RegionID),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://100.88.225.67:4566"),
		Credentials:      c,
	},
	)
	if err != nil {
		fmt.Println(err)
	}

	sqsClient := sqs.New(sess)
	snsClient := sns.New(sess)

	quo, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: queueName})
	if err != nil {
		log.Fatal(err)
	}

	sendMessageInput := &sqs.SendMessageInput{
		QueueUrl:    quo.QueueUrl,
		MessageBody: aws.String("Hey There!"),
	}

	_, err = sqsClient.SendMessage(sendMessageInput)

	if err != nil {
		log.Fatal(err)
	}

	receiveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl: quo.QueueUrl,
		AttributeNames: []*string{
			aws.String("All"),
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String("All"),
		},
		VisibilityTimeout: aws.Int64(3600), // 1 hour to process a request
	}

	receivedMessageOutput, err := sqsClient.ReceiveMessage(receiveMessageInput)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(receivedMessageOutput)

	//Let's do some S3 stuff!
	s3Client := s3.New(sess)
	uploadInput := &s3manager.UploadInput{
		Bucket: bucket,
		Body:   strings.NewReader("asdasd"),
		Key:    aws.String("dasd"),
	}
	uploader := s3manager.NewUploaderWithClient(s3Client)
	uploadOutput, err := uploader.Upload(uploadInput)

	if err != nil {
		log.Fatal(err)
	}

	//Boy we did it, lets publish the location of the file to S3
	resultNotification := &ResultNotification{
		DocumentURL: uploadOutput.Location,
		FinishedAt:  time.Now(),
	}

	resultJSON, _ := json.Marshal(resultNotification)
	publishOutput, err := snsClient.Publish(&sns.PublishInput{Message: aws.String(string(resultJSON)), TopicArn: topicARN})

	if err != nil {
		log.Fatal(err)
	}

	log.Print(publishOutput)

	//Now lets check if we have something in the result-queue -> through the sns subscription
	quo, err = sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String("result-queue")})

	if err != nil {
		log.Fatal(err)
	}

	receiveMessageInput = &sqs.ReceiveMessageInput{
		QueueUrl: quo.QueueUrl,
		AttributeNames: []*string{
			aws.String("All"),
		},
		MaxNumberOfMessages: aws.Int64(10),
		MessageAttributeNames: []*string{
			aws.String("All"),
		},
		VisibilityTimeout: aws.Int64(3600), // 1 hour to process a request
	}

	receivedMessageOutput, err = sqsClient.ReceiveMessage(receiveMessageInput)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(receivedMessageOutput)
}
