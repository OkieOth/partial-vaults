package json

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/pkg/keys"
	"github.com/okieoth/pvault/pkg/types"
)

func ProcessJsonFile(inputFile, outputFile string, processor types.ProcessFunc, keys []string) error {
	root, err := jsonreader.ReadJSON(inputFile)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling input: %v", err)
	}

	if !travers(&root, "", processor, keys) {
		return fmt.Errorf("Processing canceled by user")
	}

	outputBytes, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		return fmt.Errorf("Error while marshal output: %v", err)
	}
	if outputFile == "stdout" {
		fmt.Println(string(outputBytes))
	} else {
		if err := os.WriteFile(outputFile, outputBytes, 0644); err != nil {
			return fmt.Errorf("Error while writing output file: %v", err)
		}
	}
	return nil
}

func travers(val *jsonreader.OrderedValue, keyPath string, processor types.ProcessFunc, keys []string) bool {
	switch val.Type {
	case types.OBJECT:
		for _, pair := range val.Value.(jsonreader.OrderedObject) {
			if !travers(pair.Value, types.NewKeyPath(keyPath, pair.Key), processor, keys) {
				return false
			}
		}
	case types.ARRAY:
		for _, v := range val.Value.(jsonreader.OrderedArray) {
			if !travers(v, keyPath, processor, keys) {
				return false
			}
		}
	default:
		if len(keys) > 0 && (!slices.Contains(keys, keyPath)) {
			break
		}
		value, t, err := val.GetValue()
		if err != nil {
			fmt.Printf("key: %s, error: %v", keyPath, err)
			return false
		}
		if output, t, h, err := processor(value, t, keyPath); err == nil {
			switch h {
			case types.HANDLING_PROCESS:
				val.Value = output
				val.Type = t
			case types.HANDLING_CANCEL:
				return false
			}
		} else {
			fmt.Println("Error while process input: ", err, ", key=", keyPath)
			return false
		}
	}
	return true
}

func GetEncryptedKeys(root *jsonreader.OrderedValue) ([]string, error) {
	ret := make([]string, 0)

	testForEncrypted := keys.TestEncryptedProcessor()
	f := func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		_, _, handling, _ := testForEncrypted(input, vt, keyPath)
		if handling == types.HANDLING_PROCESS {
			ret = append(ret, keyPath)
		}
		return input, vt, handling, nil
	}

	if !travers(root, "", f, []string{}) {
		return ret, fmt.Errorf("Processing canceled by user")
	}

	return ret, nil
}
