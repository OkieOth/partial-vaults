package yaml

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/yamlreader"
	"github.com/okieoth/pvault/pkg/types"
	"gopkg.in/yaml.v3"
)

func ProcessYamlFile(inputFile, outputFile string, processor types.ProcessFunc) error {
	root, err := yamlreader.ReadYAML(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading YAML: %v", err)
	}

	if !travers(root, "", processor) {
		return fmt.Errorf("Processing canceled by user")
	}

	if err := yamlreader.WriteYAML(outputFile, root); err != nil {
		return fmt.Errorf("Error writing YAML: %v", err)
	}
	return nil
}

func travers(node *yaml.Node, keyPath string, processor types.ProcessFunc) bool {
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			if !travers(n, keyPath, processor) {
				return false
			}
		}
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			value := node.Content[i+1]
			if !travers(value, keyPath+"."+key.Value, processor) {
				return false
			}
		}
	case yaml.SequenceNode:
		for _, item := range node.Content {
			if !travers(item, keyPath, processor) {
				return false
			}
		}
	case yaml.ScalarNode:
		t, err := getValueType(node.Tag)
		if err != nil {
			fmt.Printf("key: %s, error: %v", keyPath, err)
			return false
		}
		if output, outputType, handling, err := processor([]byte(node.Value), t, keyPath); err == nil {
			if handling == types.HANDLING_PROCESS {
				if err := getProcessingResult(node, output, outputType); err != nil {
					fmt.Println("Error assigning processing results, key:", keyPath, "error: ", err)
				}
			} else if handling == types.HANDLING_CANCEL {
				return false
			}
		} else {
			fmt.Println("Error while process input: ", node.Kind, ", key=", keyPath, "error: ", err)
		}
	default:
		fmt.Println("Error while travers yaml, unknown node type: ", node.Kind, ", key=", keyPath)
	}
	return true
}

func getProcessingResult(node *yaml.Node, output any, outputType types.ValueType) error {
	switch outputType {
	case types.BOOL:
		b, ok := output.(bool)
		if !ok {
			return fmt.Errorf("expected bool, got %T", output)
		}
		node.Value = fmt.Sprintf("%t", b)
		node.Tag = "!!bool"

	case types.STRING:
		s, ok := output.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", output)
		}
		node.Value = s
		node.Tag = "!!str"

	case types.INTEGER:
		i, ok := output.(int)
		if !ok {
			return fmt.Errorf("expected int, got %T", output)
		}
		node.Value = fmt.Sprintf("%d", i)
		node.Tag = "!!int"

	case types.NUMBER:
		f, ok := output.(float64)
		if !ok {
			return fmt.Errorf("expected float64, got %T", output)
		}
		node.Value = fmt.Sprintf("%v", f)
		node.Tag = "!!float"

	case types.NULL:
		node.Value = "null"
		node.Tag = "!!null"

	default:
		return fmt.Errorf("unsupported outputType: %v", outputType)
	}
	return nil
}

func getValueType(nodeTag string) (types.ValueType, error) {
	switch nodeTag {
	case "!!str":
		return types.STRING, nil
	case "!!int":
		return types.INTEGER, nil
	case "!!float":
		return types.NUMBER, nil
	case "!!bool":
		return types.BOOL, nil
	case "!!null":
		return types.NULL, nil
	default:
		return types.NULL, fmt.Errorf("Unknown value type: %s", nodeTag)
	}
}
