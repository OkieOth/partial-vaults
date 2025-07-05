package yaml_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/pvault/pkg/types"
	"github.com/okieoth/pvault/pkg/yaml"
)

func printValue(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, error) {
	fmt.Println("key: ", keyPath, "value: ", string(input))
	if vt == types.STRING {
		return string(input) + "_changed", types.STRING, nil
	} else {
		return input, vt, nil
	}
}

func TestProcessJsonFile(t *testing.T) {
	inputFile := "../../resources/tests/example.yaml"
	outputFile := "../../temp/example_output.yaml"
	yaml.ProcessYamlFile(inputFile, outputFile, printValue)
}
