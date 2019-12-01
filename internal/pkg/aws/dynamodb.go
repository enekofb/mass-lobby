package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

type DynamoDB struct {
	svc *dynamodb.Client
}

func NewDefaultDynamoDB() (DynamoDB, error) {

	defaultConfig, err := NewDefaultConfig()

	if err != nil {
		return DynamoDB{}, errors.Wrap(err, "could not create default config")
	}

	return DynamoDB{svc: dynamodb.New(defaultConfig.awsConfig)}, nil
}

func (db DynamoDB) DescribeTable(tableName string) (string, error) {

	inputRequest := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	describeTableRequest := db.svc.DescribeTableRequest(inputRequest)

	describeTableResponse, err := describeTableRequest.Send(context.Background())

	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("could not describe table: '%s'", tableName))
	}

	return describeTableResponse.String(), nil

}

func (db DynamoDB) PutItem(tableName string, itemKey string, itemValue string) (string, error) {

	inputRequest := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]dynamodb.AttributeValue{
			itemKey: {
				S: aws.String(itemValue),
			},
		},
	}
	putItemRequest := db.svc.PutItemRequest(inputRequest)

	putItemResponse, err := putItemRequest.Send(context.Background())

	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("could not put item '%s' into table: '%s'", itemKey, tableName))
	}

	return aws.StringValue(putItemResponse.Attributes[itemKey].S), nil

}
