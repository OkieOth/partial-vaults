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

	travers(root, "", processor)

	if err := yamlreader.WriteYAML(outputFile, root); err != nil {
		return fmt.Errorf("Error writing YAML: %v", err)
	}
	return nil
}

func travers(node *yaml.Node, keyPath string, processor types.ProcessFunc) {
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			travers(n, keyPath, processor)
		}
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			value := node.Content[i+1]
			travers(value, keyPath+"."+key.Value, processor)
		}
	case yaml.SequenceNode:
		for _, item := range node.Content {
			travers(item, keyPath, processor)
		}
	case yaml.ScalarNode:
		if output, _, err := processor([]byte(node.Value), types.STRING, keyPath); err == nil {
			node.Value = fmt.Sprintf("%v", output)
		}
	default:
		fmt.Println("Error while process input: ", node.Kind, ", key=", keyPath)
	}

}
