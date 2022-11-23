package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/pkg/handler"
	"github.com/ezequiel-uala/contacts/lambda/go/get-contact-lambda/test/mocks"
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
		mock      mocks.Processor
		init      func(in *mocks.Processor)
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
						"id": "1"
					}
					`,
				),
			},
			mock: mocks.Processor{},
			init: func(in *mocks.Processor) {
				in.On("Process", "1").Return(
					dto.ContactResponse{
						Id:        "1",
						FirstName: "first-name",
						LastName:  "last-name",
						Status:    "CREATED",
					}, nil)
			},
			wantValue: assert.NotNil,
			wantErr:   assert.NoError,
		},
		{
			name: "error path: invalid id",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": 1
					}
					`,
				),
			},
			mock:      mocks.Processor{},
			init:      func(in *mocks.Processor) {},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: missing id",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": ""
					}
					`,
				),
			},
			mock:      mocks.Processor{},
			init:      func(in *mocks.Processor) {},
			wantValue: assert.NotNil,
			wantErr:   assert.Error,
		},
		{
			name: "error path: internal server error: process failed",
			args: args{
				ctx: mocks.Context(),
				payload: json.RawMessage(
					`
					{
						"id": "1"
					}
					`,
				),
			},
			mock: mocks.Processor{},
			init: func(in *mocks.Processor) {
				in.On("Process", "1").
					Return(dto.ContactResponse{}, errors.New("internal server error"))
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
