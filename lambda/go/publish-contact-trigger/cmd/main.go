package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/internal/processor"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/external"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/handler"
)

func main() {
	sns := external.InitSnsClient("arn:aws:sns:us-east-1:620097380428:ContactsTopicArangue")
	p := processor.New(&sns)
	lambda.Start(handler.New(p).HandleRequest)
}
