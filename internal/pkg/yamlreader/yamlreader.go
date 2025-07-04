package yamlreader

import (
	"os"

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
