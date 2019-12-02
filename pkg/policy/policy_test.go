package policy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPolicy(t *testing.T) {

	policyClient, err := NewPolicyClient()

	require.Nil(t, err)
	require.NotNil(t, policyClient)

}
