package typedetect_test

import (
	"testing"

	"github.com/okieoth/pvault/internal/pkg/typedetect"
	"github.com/stretchr/testify/require"
)

func TestDetectFormat(t *testing.T) {
	testData := []struct {
		inputFile    string
		expectedType typedetect.InputType
	}{
		{
			inputFile:    "../../../resources/tests/example.json",
			expectedType: typedetect.INPUT_JSON,
		},
		{
			inputFile:    "../../../resources/tests/example.yaml",
			expectedType: typedetect.INPUT_YAML,
		},
		{
			inputFile:    "../../../resources/tests/example_encrypted.json",
			expectedType: typedetect.INPUT_UNKNOWN,
		},
		{
			inputFile:    "../../../resources/tests/partial.yaml",
			expectedType: typedetect.INPUT_YAML,
		},
		{
			inputFile:    "../../../README.md",
			expectedType: typedetect.INPUT_UNKNOWN,
		},
	}
	for _, td := range testData {
		detectedType, err := typedetect.DetectFormat(td.inputFile)
		require.Nil(t, err)
		require.Equal(t, td.expectedType, detectedType, "file: %s, expected: %s, got: %s", td.inputFile, td.expectedType, detectedType)
	}
}
