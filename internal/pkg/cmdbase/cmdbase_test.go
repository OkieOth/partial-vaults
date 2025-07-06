package cmdbase_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/internal/pkg/yamlreader"
	"github.com/stretchr/testify/require"
)

func getJsonData(t *testing.T, inputFile string) *jsonreader.OrderedValue {
	inputBytes, err := os.ReadFile(inputFile)
	require.Nil(t, err)
	var jsonData jsonreader.OrderedValue
	err = json.Unmarshal(inputBytes, &jsonData)
	require.Nil(t, err)
	return &jsonData
}

func getYamlData(t *testing.T, inputFile string) *yaml.Node {
	node, err := yamlreader.ReadYAML(inputFile)
	require.Nil(t, err)
	return node
}

func TestNumber(t *testing.T) {
	inputJsonFile := "../../../resources/tests/number.json"
	inputYamlFile := "../../../resources/tests/number.yaml"
	jsonData := getJsonData(t, inputJsonFile)
	yamlData := getYamlData(t, inputYamlFile)
	fmt.Println(jsonData, yamlData) // TODO
}

func TestInteger(t *testing.T) {
	inputJsonFile := "../../../resources/tests/integer.json"
	inputYamlFile := "../../../resources/tests/integer.yaml"
	jsonData := getJsonData(t, inputJsonFile)
	yamlData := getYamlData(t, inputYamlFile)
	fmt.Println(jsonData, yamlData) // TODO
}

func TestBool(t *testing.T) {
	inputJsonFile1 := "../../../resources/tests/bool_01.json"
	inputJsonFile2 := "../../../resources/tests/bool_02.json"
	inputYamlFile1 := "../../../resources/tests/bool_01.yaml"
	inputYamlFile2 := "../../../resources/tests/bool_02.yaml"
	jsonData1 := getJsonData(t, inputJsonFile1)
	jsonData2 := getJsonData(t, inputJsonFile2)
	yamlData1 := getYamlData(t, inputYamlFile1)
	yamlData2 := getYamlData(t, inputYamlFile2)
	fmt.Println(jsonData1, jsonData2, yamlData1, yamlData2) // TODO

}

func TestString(t *testing.T) {
	inputJsonFile := "../../../resources/tests/string.json"
	inputYamlFile := "../../../resources/tests/string.yaml"
	jsonData := getJsonData(t, inputJsonFile)
	yamlData := getJsonData(t, inputYamlFile)
	fmt.Println(jsonData, yamlData) // TODO
}
