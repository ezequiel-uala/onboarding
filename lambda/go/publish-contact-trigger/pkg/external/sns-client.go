package external

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/ezequiel-uala/contacts/lambda/go/publish-contact-trigger/pkg/models"
)

type SnsClient struct {
	Sns      snsiface.SNSAPI
	TopicARN string
}

func InitSnsClient(topicName string) SnsClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return SnsClient{
		Sns:      sns.New(sess),
		TopicARN: topicName,
	}
}

func (s *SnsClient) PublishMessage(message models.Contact) (*sns.PublishOutput, error) {
	var publishResult *sns.PublishOutput

	byteMsg, err := json.Marshal(message)
	if err != nil {
		return publishResult, err
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(byteMsg)),
		TopicArn: aws.String(s.TopicARN),
	}

	publishResult, err = s.Sns.Publish(input)
	if err != nil {
		fmt.Printf("error on publisher: %s", err.Error())
		return publishResult, err
	}

	return publishResult, nil
}
