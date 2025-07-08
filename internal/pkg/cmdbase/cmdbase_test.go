package cmdbase_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/internal/pkg/yamlreader"
	"github.com/okieoth/pvault/pkg/types"
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

	v1, err := jsonData.NumberValue()
	require.Nil(t, err)
	require.Equal(t, v1, float64(24.71))

	v2, err := yamlreader.NumberValue(yamlData)
	require.Nil(t, err)

	require.Equal(t, v1, v2)

	_, err = jsonData.StringValue()
	require.NotNil(t, err)
	_, err = jsonData.BoolValue()
	require.NotNil(t, err)
	_, err = jsonData.IntValue()
	require.NotNil(t, err) // for json specifics

	_, err = yamlreader.StringValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.BoolValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.IntValue(yamlData)
	require.NotNil(t, err)

	vv, vt, err := yamlreader.GetValue(yamlData)
	require.Nil(t, err)
	require.Equal(t, types.NUMBER, vt)
	require.Equal(t, float64(24.71), vv)

	vv, vt, err = jsonData.GetValue()
	require.Nil(t, err)
	require.Equal(t, types.NUMBER, vt)
	require.Equal(t, float64(24.71), vv)
}

func TestInteger(t *testing.T) {
	inputJsonFile := "../../../resources/tests/integer.json"
	inputYamlFile := "../../../resources/tests/integer.yaml"
	jsonData := getJsonData(t, inputJsonFile)
	yamlData := getYamlData(t, inputYamlFile)

	v1, err := jsonData.IntValue()
	require.Nil(t, err)
	require.Equal(t, v1, int64(13))

	v2, err := yamlreader.IntValue(yamlData)
	require.Nil(t, err)

	require.Equal(t, v1, v2)

	_, err = jsonData.StringValue()
	require.NotNil(t, err)
	_, err = jsonData.BoolValue()
	require.NotNil(t, err)
	_, err = jsonData.NumberValue()
	require.Nil(t, err) // for JSON specifics

	_, err = yamlreader.StringValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.BoolValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.NumberValue(yamlData)
	require.NotNil(t, err)

	vv, vt, err := yamlreader.GetValue(yamlData)
	require.Nil(t, err)
	require.Equal(t, types.INTEGER, vt)
	require.Equal(t, int64(13), vv)

	vv, vt, err = jsonData.GetValue()
	require.Nil(t, err)
	require.Equal(t, types.INTEGER, vt)
	require.Equal(t, int64(13), vv)
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

	v1, err := jsonData1.BoolValue()
	require.Nil(t, err)
	require.Equal(t, v1, true)
	v2, err := yamlreader.BoolValue(yamlData1)
	require.Nil(t, err)
	require.Equal(t, v1, v2)

	v1, err = jsonData2.BoolValue()
	require.Nil(t, err)
	require.Equal(t, v1, false)
	v2, err = yamlreader.BoolValue(yamlData2)
	require.Nil(t, err)
	require.Equal(t, v1, v2)

	_, err = jsonData1.StringValue()
	require.NotNil(t, err)
	_, err = jsonData1.IntValue()
	require.NotNil(t, err)
	_, err = jsonData1.NumberValue()
	require.NotNil(t, err)
	_, err = jsonData2.StringValue()
	require.NotNil(t, err)
	_, err = jsonData2.IntValue()
	require.NotNil(t, err)
	_, err = jsonData2.NumberValue()
	require.NotNil(t, err)

	_, err = yamlreader.StringValue(yamlData1)
	require.NotNil(t, err)
	_, err = yamlreader.IntValue(yamlData1)
	require.NotNil(t, err)
	_, err = yamlreader.NumberValue(yamlData1)
	require.NotNil(t, err)
	_, err = yamlreader.StringValue(yamlData2)
	require.NotNil(t, err)
	_, err = yamlreader.IntValue(yamlData2)
	require.NotNil(t, err)
	_, err = yamlreader.NumberValue(yamlData2)
	require.NotNil(t, err)

	vv, vt, err := yamlreader.GetValue(yamlData1)
	require.Nil(t, err)
	require.Equal(t, types.BOOL, vt)
	require.Equal(t, true, vv)

	vv, vt, err = jsonData1.GetValue()
	require.Nil(t, err)
	require.Equal(t, types.BOOL, vt)
	require.Equal(t, true, vv)

}

func TestString(t *testing.T) {
	inputJsonFile := "../../../resources/tests/string.json"
	inputYamlFile := "../../../resources/tests/string.yaml"
	jsonData := getJsonData(t, inputJsonFile)
	yamlData := getYamlData(t, inputYamlFile)
	fmt.Println(jsonData, yamlData) // TODO

	v1, err := jsonData.StringValue()
	require.Nil(t, err)
	require.Equal(t, v1, "I am a longer text")
	v2, err := yamlreader.StringValue(yamlData)
	require.Nil(t, err)
	require.Equal(t, v1, v2)

	_, err = jsonData.IntValue()
	require.NotNil(t, err)
	_, err = jsonData.BoolValue()
	require.NotNil(t, err)
	_, err = jsonData.NumberValue()
	require.NotNil(t, err)

	_, err = yamlreader.IntValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.BoolValue(yamlData)
	require.NotNil(t, err)
	_, err = yamlreader.NumberValue(yamlData)
	require.NotNil(t, err)

	vv, vt, err := yamlreader.GetValue(yamlData)
	require.Nil(t, err)
	require.Equal(t, types.STRING, vt)
	require.Equal(t, "I am a longer text", vv)

	vv, vt, err = jsonData.GetValue()
	require.Nil(t, err)
	require.Equal(t, types.STRING, vt)
	require.Equal(t, "I am a longer text", vv)
}
