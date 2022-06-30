package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	payload := req.Body
	log.Println("got payload -", payload)

	response := "hello " + payload + "!"
	return events.APIGatewayV2HTTPResponse{Body: response}, nil
}
