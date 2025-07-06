package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/pkg/types"
)

func ProcessJsonFile(inputFile, outputFile string, processor types.ProcessFunc) error {
	inputBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error while reading input file: %v", err)
	}

	var root jsonreader.OrderedValue
	if err := json.Unmarshal(inputBytes, &root); err != nil {
		return fmt.Errorf("Error while unmarshal input: %v", err)
	}

	if !travers(&root, "", processor) {
		return fmt.Errorf("Processing canceled by user")
	}

	outputBytes, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshal output: %v", err)
	}
	if err := os.WriteFile(outputFile, outputBytes, 0644); err != nil {
		return fmt.Errorf("Error while writing output file: %v", err)
	}

	return nil
}

func travers(val *jsonreader.OrderedValue, keyPath string, processor types.ProcessFunc) bool {
	switch val.Type {
	case types.OBJECT:
		for _, pair := range val.Value.(jsonreader.OrderedObject) {
			if !travers(pair.Value, keyPath+"."+pair.Key, processor) {
				return false
			}
		}
	case types.ARRAY:
		for _, v := range val.Value.(jsonreader.OrderedArray) {
			if !travers(v, keyPath, processor) {
				return false
			}
		}
	default:
		out, _ := json.Marshal(val.Value)
		if val.Type == types.STRING {
			last := len(out) - 1
			out = out[1:last]
		}
		if output, t, h, err := processor(out, val.Type, keyPath); err == nil {
			if h == types.HANDLING_PROCESS {
				val.Value = output
				val.Type = t
			} else if h == types.HANDLING_CANCEL {
				return false
			}
		} else {
			fmt.Println("Error while process input: ", err, ", key=", keyPath)
			return false
		}
	}
	return true
}
