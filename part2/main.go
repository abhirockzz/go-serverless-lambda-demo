package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamoDBClient *dynamodb.DynamoDB
var tableName string

func init() {
	tableName = os.Getenv("TABLE_NAME")
	if tableName == "" {
		log.Fatal("missing env variable TABLE_NAME")
	}

	dynamoDBClient = dynamodb.New(session.New())
}

func main() {
	lambda.Start(handler)
}

type User struct {
	Email string
	Name  string
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	payload := req.Body
	log.Println("got payload -", payload)

	var user User
	err := json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		log.Println("failed to unmarshal payload")
		return events.APIGatewayV2HTTPResponse{}, err
	}

	item := make(map[string]*dynamodb.AttributeValue)
	emailAV := dynamodb.AttributeValue{S: aws.String(user.Email)}
	nameAV := dynamodb.AttributeValue{S: aws.String(user.Name)}

	item["email"] = &emailAV
	item["name"] = &nameAV

	_, err = dynamoDBClient.PutItem(&dynamodb.PutItemInput{TableName: aws.String(tableName), Item: item})
	if err != nil {
		log.Println("failed to put item in dynamodb")
		return events.APIGatewayV2HTTPResponse{}, err
	}

	log.Println("created user with email", user.Email)

	return events.APIGatewayV2HTTPResponse{StatusCode: http.StatusOK}, nil
}
