package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/handler"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandle_Request(t *testing.T) {

	type args struct {
		ctx     context.Context
		payload json.RawMessage
	}

	tests := []struct {
		name      string
		args      args
		mock      mocks.MockProcessor
		init      func(in *mocks.MockProcessor)
		wantValue assert.ValueAssertionFunc
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": "1",
						"firstName": "first-name",
						"lastName": "last-name"
					}
					`,
				),
			},
			mock: mocks.MockProcessor{},
			init: func(in *mocks.MockProcessor) {
				in.On("Process", dto.CreateContactRequest{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return(models.Contact{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
				}, nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.NoError,
		},
		{
			name: "error path: internal server error: process failed",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": "1",
						"firstName": "first-name",
						"lastName": "last-name"
					}
					`,
				),
			},
			mock: mocks.MockProcessor{},
			init: func(in *mocks.MockProcessor) {
				in.On("Process", dto.CreateContactRequest{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
				}).Return(models.Contact{}, errors.New("internal server error"))
			},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				h := handler.New(&tt.mock)
				res, err := h.HandleRequest(tt.args.ctx, tt.args.payload)
				tt.wantErr(t, err)
				tt.wantValue(t, res)
			},
		)
	}
}
