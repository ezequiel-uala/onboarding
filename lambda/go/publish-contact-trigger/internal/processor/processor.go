package processor

import (
	"fmt"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/models"
)

type Processor struct {
	client dto.SnsClient
}

func New(client dto.SnsClient) *Processor {
	return &Processor{client: client}
}

func (p *Processor) Process(contact models.Contact) error {
	res, err := p.client.PublishMessage(contact)
	if err != nil {
		fmt.Printf("ERR: error publishing the message: %s\n", err)
	}

	fmt.Printf("Publish message result: %v\n", res)

	return nil
}
