package aws

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestICanCreateDefaultDynamoDB(t *testing.T) {
	dynamoDB, e := NewDefaultDynamoDB()

	require.Nil(t, e)
	require.NotNil(t, dynamoDB)

}

func TestICouldDescribeTableForExistingTable(t *testing.T) {

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

func TestICouldPutItemForExistingTable(t *testing.T) {

	existingTableName, ok := os.LookupEnv("ExistingTableName")

	if !ok {
		t.Errorf("not defined 'ExistingTableName' environment variable")
	}

	itemKey, ok := os.LookupEnv("ExistingTablePrimaryKeyName")

	if !ok {
		t.Errorf("not defined 'ExistingTablePrimaryKeyName' environment variable")
	}

	dynamoDB, e := NewDefaultDynamoDB()

	require.Nil(t, e)
	require.NotNil(t, dynamoDB)

	itemValue := Randomize("myDynamoDBItem")

	dynamoDBItem, e := dynamoDB.PutItem(existingTableName, itemKey, itemValue)

	require.Nil(t, e)
	require.NotNil(t, dynamoDBItem)
	fmt.Println(dynamoDBItem)

}
