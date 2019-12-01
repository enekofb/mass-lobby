package aws

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestICanCreateDefaultDynamoDB(t *testing.T) {
	dynamoDB, e := NewDefaultDynamoDB()

	require.Nil(t, e)
	require.NotNil(t, dynamoDB)

}

func TestIShouldDescribeTableForExistingTable(t *testing.T) {

	existingTableName, ok := os.LookupEnv("ExistingTableName")

	if !ok {
		t.Errorf("not defined 'ExistingTableName' environment variable")
	}

	dynamoDB, e := NewDefaultDynamoDB()

	require.Nil(t, e)
	require.NotNil(t, dynamoDB)

	dynamoDBTable, e := dynamoDB.DescribeTable(existingTableName)

	require.Nil(t, e)
	require.NotNil(t, dynamoDBTable)

}

//
//func TestIShouldReturnErrorWhenDontExistSecret(t *testing.T) {
//
//	existingSecretName := "secret-idontexist"
//
//	_, e := GetSecretValue(existingSecretName)
//
//	require.NotNil(t, e)
//	require.NotEmpty(t, e.Error())
//}
