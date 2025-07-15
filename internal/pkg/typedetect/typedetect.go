package typedetect

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type InputType int

const (
	INPUT_JSON InputType = iota
	INPUT_YAML
	INPUT_UNKNOWN
)

func (t InputType) String() string {
	switch t {
	case INPUT_JSON:
		return "JSON"
	case INPUT_YAML:
		return "YAML"
	default:
		return "INPUT_UNKNOWN"
	}
}

func DetectFormat(file string) (InputType, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return INPUT_UNKNOWN, fmt.Errorf("Error while testing input file for type: %v", err)
	}
	var js map[string]interface{}
	if json.Unmarshal(data, &js) == nil {
		return INPUT_JSON, nil
	}

	var yml map[string]interface{}
	if yaml.Unmarshal(data, &yml) == nil {
		return INPUT_YAML, nil
	}

	return INPUT_UNKNOWN, nil
}

func IfNotJsonOrYamlThenPanic(inputFile string) InputType {
	t, err := DetectFormat(inputFile)
	if err != nil {
		panic(fmt.Sprintf("error while detecting input format: %v", err))
	}
	if t == INPUT_UNKNOWN {
		panic("Input is neither JSON nor YAML input. Other types are not supported")
	}
	return t
}
