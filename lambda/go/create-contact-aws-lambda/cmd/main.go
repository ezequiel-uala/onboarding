package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/handler"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/processor"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/storage"
)

func main() {
	p := processor.New(&storage.DynamoClient{})
	lambda.Start(handler.New(p).HandleRequest)
}
