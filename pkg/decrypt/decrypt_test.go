package decrypt_test

import (
	"testing"

	"github.com/okieoth/pvault/pkg/decrypt"
	"github.com/okieoth/pvault/pkg/types"
	"github.com/okieoth/pvault/pkg/yaml"
	"github.com/stretchr/testify/require"
)

func TestBlocks(t *testing.T) {
	inputFile := "../../resources/tests/partial_encrypted_example.yaml"
	outputFile := "../../temp/partial_encrypted_example_decrypted.yaml"
	password := "test999"

	processor := func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return decrypt.DecryptImpl(input, vt, keyPath, password)
	}

	err := yaml.ProcessYamlFile(inputFile, outputFile, processor, []string{})
	require.Nil(t, err)
}
