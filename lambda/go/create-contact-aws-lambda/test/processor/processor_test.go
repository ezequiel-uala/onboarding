package processor

import (
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/dto"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/models"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/pkg/processor"
	"github.com/ezequiel-uala/contacts/lambda/go/create-contact-aws-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process(t *testing.T) {
	type args struct {
		req dto.CreateContactRequest
	}

	tests := []struct {
		name      string
		args      args
		mock      mocks.DynamoClient
		init      func(in *mocks.DynamoClient)
		wantValue assert.ValueAssertionFunc
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			args: args{
				req: dto.CreateContactRequest{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
				},
			},
			mock: mocks.DynamoClient{},
			init: func(in *mocks.DynamoClient) {
				in.On("CreateContact", models.Contact{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
					Status:    "CREATED",
				}).Return(models.Contact{
					Id:        "1",
					FirstName: "first-name",
					LastName:  "last-name",
					Status:    "CREATED",
				}, nil)
			},
			wantErr:   assert.NoError,
			wantValue: assert.NotNil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.init(&tt.mock)
				p := processor.New(&tt.mock)
				res, err := p.Process(tt.args.req)
				tt.wantErr(t, err)
				tt.wantValue(t, res)
			},
		)
	}
}
