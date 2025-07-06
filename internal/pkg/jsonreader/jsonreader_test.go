package jsonreader_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
)

func TestJsonKeepOrder(t *testing.T) {
	inputFile := "../../../resources/tests/example.json"
	outputFile := "../../../temp/example_output.json"

	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	var root jsonreader.OrderedValue
	if err := json.Unmarshal(inputBytes, &root); err != nil {
		panic(err)
	}

	// Optional: print keys for inspection
	fmt.Println("Parsed JSON structure:")
	jsonreader.PrintOrdered(&root, 0)

	// Write back to output.json
	outputBytes, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputFile, outputBytes, 0644); err != nil {
		panic(err)
	}
}
