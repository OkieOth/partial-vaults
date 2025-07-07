package yamlreader

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadYAML(filename string) (*yaml.Node, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var root yaml.Node
	err = yaml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}
	return &root, nil
}

func WriteYAML(filename string, root *yaml.Node) error {
	data, err := yaml.Marshal(root)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// since yaml objects are a bit nested, you have to dig deep to reach the value
func traversToValue[T any](node *yaml.Node, errorValue T, fn func(*yaml.Node) (T, error)) (T, error) {
	switch node.Kind {
	case yaml.DocumentNode:
		return traversToValue(node.Content[0], errorValue, fn)
	case yaml.MappingNode:
		if len(node.Content) > 1 {
			return traversToValue(node.Content[1], errorValue, fn)
		}
		return errorValue, fmt.Errorf("too less elements for mapping node")
	case yaml.ScalarNode:
		return fn(node)
	default:
		return errorValue, fmt.Errorf("unknown note type")
	}
}

func StringValue(node *yaml.Node) (string, error) {
	if node == nil {
		return "", fmt.Errorf("no node given")
	}
	castFn := func(node *yaml.Node) (string, error) {
		if node.Tag == "!!str" {
			return node.Value, nil
		} else {
			return "", fmt.Errorf("no string value")
		}
	}
	return traversToValue(node, "", castFn)
}

func IntValue(node *yaml.Node) (int64, error) {
	if node == nil {
		return 0, fmt.Errorf("no node given")
	}
	castFn := func(node *yaml.Node) (int64, error) {
		if node.Tag == "!!int" {
			i, err := strconv.ParseInt(node.Value, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("error while cast to int")
			}
			return i, nil
		} else {
			return 0, fmt.Errorf("no int value")
		}
	}
	return traversToValue(node, 0, castFn)
}

func BoolValue(node *yaml.Node) (bool, error) {
	if node == nil {
		return false, fmt.Errorf("no node given")
	}
	castFn := func(node *yaml.Node) (bool, error) {
		if node.Tag == "!!bool" {
			lowerValue := strings.ToLower(node.Value)
			switch lowerValue {
			case "true":
				return true, nil
			case "false":
				return false, nil
			default:
				return false, fmt.Errorf("error while cast to bool")
			}
		} else {
			return false, fmt.Errorf("no bool value")
		}
	}
	return traversToValue(node, false, castFn)
}

func NumberValue(node *yaml.Node) (float64, error) {
	if node == nil {
		return 0.0, fmt.Errorf("no node given")
	}
	castFn := func(node *yaml.Node) (float64, error) {
		if node.Tag == "!!float" {
			f, err := strconv.ParseFloat(node.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("error while cast to float")
			}
			return f, nil
		} else {
			return 0, fmt.Errorf("no float value")
		}
	}
	return traversToValue(node, 0.0, castFn)
}
