package main

import "github.com/aws/aws-lambda-go/lambda"

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func show() (*book, error) {
	bk := &book{
		ISBN:   "394-2839480293",
		Title:  "Charting the unchartered chartrists",
		Author: "Bilbo Braggins",
	}
	return bk, nil
}

func main() {
	lambda.Start(show)
}
