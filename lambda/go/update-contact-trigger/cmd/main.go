package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ezequiel-uala/contacts/lambda/go/update-contact-trigger/pkg/handler"
	"github.com/ezequiel-uala/contacts/lambda/go/update-contact-trigger/pkg/processor"
	"github.com/ezequiel-uala/contacts/lambda/go/update-contact-trigger/pkg/storage"
)

func main() {
	p := processor.New(&storage.DynamoClient{})
	lambda.Start(handler.New(p).HandleRequest)
}
