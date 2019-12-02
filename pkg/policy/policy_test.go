package policy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPolicy(t *testing.T) {

	client, err := NewClient()

	require.Nil(t, err)
	require.NotNil(t, client)
	require.NotEmpty(t, client.config.Url)

}

func TestICouldSearchAllPolicies(t *testing.T) {

	policyClient, err := NewClient()

	require.Nil(t, err)
	require.NotNil(t, policyClient)

	policyClient.searchPolicies()

}
