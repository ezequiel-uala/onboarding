package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/ezequiel-uala/contacts/lambda/go/update-contact-trigger/pkg/dto"
)

type Handler struct {
	processor dto.Processor
}

func New(in dto.Processor) *Handler {
	return &Handler{
		processor: in,
	}
}

func (h *Handler) HandleRequest(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		var contact dto.Contact
		snsRecord := record.SNS

		if err := json.Unmarshal([]byte(snsRecord.Message), &contact); err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
			return err
		}

		err := h.processor.Process(contact.Id)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
			return err
		}

		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
	}
	return nil
}
