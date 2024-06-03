package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type OnTweetPostedEvent struct {
	Tweet Tweet `json:"detail"`
}

type Tweet struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, event *OnTweetPostedEvent) (*OnTweetPostedEvent, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	fmt.Printf("%+v\n", event)

	return event, nil
}

func main() {
	lambda.Start(HandleRequest)
}
