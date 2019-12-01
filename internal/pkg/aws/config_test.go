package aws

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	defaultConfig, err := NewDefaultConfig()
	require.NoError(t, err)
	require.NotNil(t, defaultConfig)
}
