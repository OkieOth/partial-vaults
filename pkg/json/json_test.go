package json_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/pvault/pkg/json"
	"github.com/okieoth/pvault/pkg/types"
)

func printValue(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	fmt.Printf("key: %s, value: %v", keyPath, input)
	return input, vt, types.HANDLING_PROCESS, nil
}

func TestProcessJsonFile(t *testing.T) {
	inputFile := "../../resources/tests/example.json"
	outputFile := "../../temp/example_output2.json"
	json.ProcessJsonFile(inputFile, outputFile, printValue, []string{})
}
