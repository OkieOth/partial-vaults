package types_test

import (
	"testing"

	"github.com/okieoth/pvault/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestNewKeyPath(t *testing.T) {
	tests := []struct {
		keyPath  string
		key      string
		expected string
	}{
		{
			keyPath:  "",
			key:      "test",
			expected: "test",
		},
		{
			keyPath:  "x",
			key:      "test",
			expected: "x.test",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, types.NewKeyPath(test.keyPath, test.key))
	}
}
