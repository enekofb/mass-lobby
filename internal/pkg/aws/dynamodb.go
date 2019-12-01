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

func NewDefaultDynamoDB() (DynamoDB, error) {

	defaultConfig, err := NewDefaultConfig()

	if err != nil {
		return DynamoDB{}, errors.Wrap(err, "could not create default config")
	}

	return DynamoDB{svc: dynamodb.New(defaultConfig.awsConfig)}, nil
}

//func GetSecretValue(name string) (string, error) {
//
//	sess := session.Must(session.NewSession(aws.NewConfig()))
//
//	svc := secretsmanager.New(sess)
//	input := &secretsmanager.GetSecretValueInput{
//		SecretId: aws.String(name),
//	}
//
//	result, err := svc.GetSecretValue(input)
//	if err != nil {
//		if aerr, ok := err.(awserr.Error); ok {
//			return "", errors.Wrap(err, fmt.Sprintf("%s:%s", aerr.Code(), aerr.Message()))
//		}
//		return "", err
//	}
//
//	return aws.StringValue(result.SecretString), nil
//}
