package dto

type Processor interface {
	Process(string) (ContactResponse, error)
}

type DynamoClient interface {
	GetContactById(string, *ContactResponse) error
}
