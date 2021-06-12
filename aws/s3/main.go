package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	fmt.Println("listing buckets")

	c := credentials.NewStaticCredentials("test", "test", "")
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(endpoints.EuWest2RegionID),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://100.88.225.67:4566"),
		Credentials:      c,
	},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	//createBucket(sess, "buckettwo")
	// listBuckets(sess)
	// uploadItems(sess)
	listBucketItems(sess)
	// downloadBucketItems(sess)
	// deleteBucketItem(sess)

}

func createBucket(s *session.Session, bucketName string) {
	fmt.Println("**Creating S3 Buckets**")
	svc := s3.New(s)

	input := s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	result, err := svc.CreateBucket(&input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				log.Fatal("Bucket already exists")

			case s3.ErrCodeBucketAlreadyOwnedByYou:
				log.Fatal("Bucket already owned by user")

			default:
				log.Fatal(aerr.Error())
			}
		} else {
			log.Fatal(err.Error())
		}
	}

	log.Println(result)
}

func listBuckets(s *session.Session) {
	fmt.Println("** Listing S3 Buckets")
	svc := s3.New(s)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("error listing buckets")
	}

	count := len(result.Buckets)
	fmt.Printf("Found %d buckets:\n", count)
	for bucketIndex, bucket := range result.Buckets {
		log.Printf("Bucket %d: %s\n", bucketIndex, aws.StringValue(bucket.Name))
	}

	fmt.Println("")
}

func listBucketItems(s *session.Session) {
	fmt.Println("** Listing Bucket Items")
	svc := s3.New(s)

	resp, err := svc.ListObjectsV2(
		&s3.ListObjectsV2Input{
			Bucket: aws.String("bucketone"),
		},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range resp.Contents {
		fmt.Printf("File: %s, Size %dbytes\n", *item.Key, *item.Size)
	}

	fmt.Println("")
}

func uploadItems(s *session.Session) {
	uploadResults := make(map[string][]string)
	var rawResults []s3manager.UploadOutput
	fmt.Println("**Preparing to Upload Items**")
	f, err := os.Open("upload.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}

	defer f.Close()

	uploader := s3manager.NewUploader(s)
	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("bucketone"),
		Key:    aws.String("upload.txt"),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	rawResults = append(rawResults, *result)
	uploadResults["bucketone"] = append(uploadResults["bucketone"], "upload.txt")
	result, err = uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String("bucketone"),
		Key:    aws.String("remove.txt"),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	rawResults = append(rawResults, *result)
	uploadResults["bucketone"] = append(uploadResults["bucketone"], "remove.txt")
	for i := range uploadResults["bucketone"] {
		fmt.Printf("Successfully Uploaded: %v\n", uploadResults["bucketone"][i])
		// log.Printf("Upload Result(s): %+v\n\n", result)
	}

	fmt.Println("Raw Results from AWS")
	for i := range rawResults {
		fmt.Printf("%+v\n", rawResults[i])
	}
	fmt.Println("\n")
}

func downloadBucketItems(s *session.Session) {
	fmt.Println("**Attempting to Download an item from S3**")
	file, err := os.Create("downloaded.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(s)
	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String("bucketone"),
			Key:    aws.String("upload.txt"),
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("successfully downloaded file\n")

}

func deleteBucketItem(s *session.Session) {
	fmt.Println("Attempting to delete a single s3 object**")
	svc := s3.New(s)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("bucketone"),
		Key:    aws.String("remove.txt"),
	}

	result, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			log.Fatal(err.Error())
		}
	}

	log.Printf("Result: %+v\n", result)
}
