package yaml_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/pvault/pkg/types"
	"github.com/okieoth/pvault/pkg/yaml"
)

func printValue(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	fmt.Printf("key: %s, value:= %v", keyPath, input)
	return input, vt, types.HANDLING_PROCESS, nil
}

func TestProcessYamlFile(t *testing.T) {
	inputFile := "../../resources/tests/example.yaml"
	outputFile := "../../temp/example_output.yaml"
	yaml.ProcessYamlFile(inputFile, outputFile, printValue)
}
