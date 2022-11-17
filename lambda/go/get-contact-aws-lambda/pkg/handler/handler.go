package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/dto"
)

type Handler struct {
	processor dto.Processor
}

func New(in dto.Processor) *Handler {
	return &Handler{
		processor: in,
	}
}

func (h *Handler) HandleRequest(ctx context.Context, payload json.RawMessage) (dto.ContactResponse, error) {
	req := dto.ContactRequest{}
	fmt.Printf("Payload: %s\n", payload)
	if err := json.Unmarshal(payload, &req); err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return dto.ContactResponse{}, err
	}
	fmt.Printf("Req: %s\n", req)
	if len(req.Id) == 0 {
		fmt.Println("ERR: missing id path parameter /contacts/{id}")
		return dto.ContactResponse{}, ErrorResponse(400, "missing id path parameter /contacts/{id}")
	}

	res, err := h.processor.Process(req.Id)
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return dto.ContactResponse{}, ErrorResponse(500, err.Error())
	}

	return res, nil
}

func ErrorResponse(status int, msg string) error {
	errMsg := dto.ErrorMessage{
		Status:  status,
		Message: msg,
	}

	byteErrMsg, _ := json.Marshal(errMsg)
	return errors.New(string(byteErrMsg))
}
