package processor

import (
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/dto"
)

type Processor struct {
	DynamoClient dto.DynamoClient
}

func New(client dto.DynamoClient) *Processor {
	return &Processor{
		DynamoClient: client,
	}
}

func (p *Processor) Process(id string) (dto.ContactResponse, error) {
	contact := dto.ContactResponse{}

	err := p.DynamoClient.GetContactById(id, &contact)
	if err != nil {
		return contact, err
	}

	return contact, nil
}
