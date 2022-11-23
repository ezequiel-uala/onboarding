package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/models"
)

type Handler struct {
	processor dto.Processor
}

func New(processor dto.Processor) *Handler {
	return &Handler{processor: processor}
}

func (h *Handler) HandleRequest(ctx context.Context, e events.DynamoDBEvent) error {

	fmt.Printf("Event %s", e)

	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		id := record.Change.NewImage["id"].String()
		firstName := record.Change.NewImage["firstName"].String()
		lastName := record.Change.NewImage["lastName"].String()
		status := record.Change.NewImage["status"].String()

		fmt.Printf("HandleRequest publish Id: %s Firstname: %s, Lastname: %s, Status: %s\n", id, firstName, lastName, status)

		err := h.processor.Process(models.Contact{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Status:    status,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
