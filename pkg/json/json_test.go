package json_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/pvault/pkg/json"
	"github.com/okieoth/pvault/pkg/types"
)

func printValue(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, error) {
	fmt.Println("key: ", keyPath, "value: ", string(input))
	return string(input) + "_changed", types.STRING, nil
}

func TestProcessJsonFile(t *testing.T) {
	inputFile := "../../resources/tests/example.json"
	outputFile := "../../temp/example_output2.json"
	json.ProcessJsonFile(inputFile, outputFile, printValue)
}
