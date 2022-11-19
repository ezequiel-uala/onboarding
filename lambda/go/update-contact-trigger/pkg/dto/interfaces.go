package dto

type Processor interface {
	Process(string) error
}

type DynamoClient interface {
	UpdateStatus(string, string) error
}
