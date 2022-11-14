package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Contact struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
}

func HandleRequest(ctx context.Context, id string) (Contact, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	if id == "" {
		return Contact{}, errors.New("empty id")
	}

	svc := dynamodb.New(sess)

	contactKey := Contact{
		Id:        "10",
		FirstName: "Ezequiel",
		LastName:  "Arangue",
	}

	key, err := dynamodbattribute.MarshalMap(contactKey)
	if err != nil {
		fmt.Println(err.Error())
		return Contact{}, nil
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("Contacts-Arangue"),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return Contact{}, nil
	}

	contact := Contact{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &contact)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return contact, nil
}

func main() {
	lambda.Start(HandleRequest)
}
