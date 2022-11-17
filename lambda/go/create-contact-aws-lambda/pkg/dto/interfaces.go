package dto

import "github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"

type Processor interface {
	Process(request CreateContactRequest) (models.Contact, error)
}

type DynamoClient interface {
	CreateContact(contact models.Contact) (models.Contact, error)
}
