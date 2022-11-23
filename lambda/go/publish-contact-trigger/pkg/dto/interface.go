package dto

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/models"
)

type Processor interface {
	Process(models.Contact) error
}

type SnsClient interface {
	PublishMessage(models.Contact) (*sns.PublishOutput, error)
}
