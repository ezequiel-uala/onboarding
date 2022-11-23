package mocks

import (
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(req dto.CreateContactRequest) (models.Contact, error) {
	args := m.Called(req)
	if err := args.Get(1); err != nil {
		return models.Contact{}, err.(error)
	}
	return args.Get(0).(models.Contact), nil
}
