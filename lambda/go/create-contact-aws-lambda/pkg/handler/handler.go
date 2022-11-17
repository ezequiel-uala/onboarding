package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"
)

type Handler struct {
	processor dto.Processor
}

func New(processor dto.Processor) *Handler {
	return &Handler{
		processor: processor,
	}
}

func (h *Handler) HandleRequest(ctx context.Context, payload json.RawMessage) (models.Contact, error) {
	req := dto.CreateContactRequest{}

	if err := json.Unmarshal(payload, &req); err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return models.Contact{}, err
	}

	contact, err := h.processor.Process(req)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return models.Contact{}, err
	}

	return contact, nil
}
