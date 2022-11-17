package processor

import (
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"
)

type Processor struct {
	client dto.DynamoClient
}

func New(client dto.DynamoClient) *Processor {
	return &Processor{
		client: client,
	}
}

func (p *Processor) Process(req dto.CreateContactRequest) (models.Contact, error) {
	item := models.Contact{
		Id:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	contact, err := p.client.CreateContact(item)
	if err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}
