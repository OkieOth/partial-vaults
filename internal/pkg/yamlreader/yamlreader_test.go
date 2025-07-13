package yamlreader_test

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func readYAML(filename string) (*yaml.Node, error) {
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

// writeYAML writes a *yaml.Node back to file, preserving structure and order
func writeYAML(filename string, root *yaml.Node) error {
	data, err := yaml.Marshal(root)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// printYAMLNode recursively prints a yaml.Node for debugging or inspection
func printYAMLNode(node *yaml.Node, indent int) {
	ind := ""
	for i := 0; i < indent; i++ {
		ind += "  "
	}
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			printYAMLNode(n, indent)
		}
	case yaml.MappingNode:
		fmt.Println(ind + "{")
		for i := 0; i < len(node.Content); i += 2 {
			k := node.Content[i]
			v := node.Content[i+1]
			fmt.Printf("%s  %s:\n", ind, k.Value)
			printYAMLNode(v, indent+2)
		}
		fmt.Println(ind + "}")
	case yaml.SequenceNode:
		fmt.Println(ind + "[")
		for _, item := range node.Content {
			printYAMLNode(item, indent+1)
		}
		fmt.Println(ind + "]")
	case yaml.ScalarNode:
		fmt.Printf("%s%s\n", ind, node.Value)
	default:
		fmt.Printf("%s<unknown kind %d>\n", ind, node.Kind)
	}
}

func TestYamlRead(t *testing.T) {
	inputFile := "../../../resources/tests/partial.yaml"
	outputFile := "../../../temp/partial_output.yaml"
	root, err := readYAML(inputFile)
	if err != nil {
		fmt.Println("Error reading YAML:", err)
		os.Exit(1)
	}

	// Optional: inspect structure
	fmt.Println("Parsed YAML structure:")
	printYAMLNode(root, 0)

	// Write YAML back (preserves order)
	if err := writeYAML(outputFile, root); err != nil {
		fmt.Println("Error writing YAML:", err)
		os.Exit(1)
	}
}

func TestYamlBlock(t *testing.T) {
	inputFile := "../../../resources/tests/block.yaml"
	//outputFile := "../../../temp/block_output.yaml"
	_, err := readYAML(inputFile)
	if err != nil {
		fmt.Println("Error reading YAML:", err)
		os.Exit(1)
	}
}
