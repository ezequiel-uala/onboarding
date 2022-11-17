package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/handler"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/processor"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/storage"
)

func main() {
	p := processor.New(&storage.DynamoClient{})
	lambda.Start(handler.New(p).HandleRequest)
}
