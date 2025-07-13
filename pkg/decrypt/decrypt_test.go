package decrypt_test

import (
	"os"
	"testing"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/internal/pkg/yamlreader"
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

func TestDecryptYaml(t *testing.T) {
	tests := []struct {
		inputFile     string
		outputFile    string
		referenceFile string
		password      string
		keys          []string
	}{
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_01.yaml",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.yaml",
			password:      "test999",
			keys:          []string{},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_03.yaml",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_03.yaml",
			password:      "test999",
			keys:          []string{"third.carrot"},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_04.yaml",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_04.yaml",
			password:      "test999",
			keys:          []string{"first.a", "first.z", "second.b.2", "fourth.list"},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example_02.yaml",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_02.yaml",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.yaml",
			password:      "test999",
			keys:          []string{},
		},
	}
	for _, test := range tests {
		if _, err := os.Stat(test.outputFile); err == nil {
			// file exists
			os.Remove(test.outputFile)
			_, err = os.Stat(test.outputFile)
			require.NotNil(t, err)
		}
		err := decrypt.Decrypt(test.inputFile, test.outputFile, test.password, test.keys)
		require.Nil(t, err)
		_, err = os.Stat(test.outputFile)
		require.Nil(t, err)
		outputYaml, err := yamlreader.ReadYAML(test.outputFile)
		require.Nil(t, err)
		referenceYaml, err := yamlreader.ReadYAML(test.referenceFile)
		require.Equal(t, referenceYaml, outputYaml, "reference and decrypted output are different")

	}
}

func TestDecryptJson(t *testing.T) {
	tests := []struct {
		inputFile     string
		outputFile    string
		referenceFile string
		password      string
		keys          []string
	}{
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.json",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_01.json",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.json",
			password:      "test999",
			keys:          []string{},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.json",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_03.json",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_03.json",
			password:      "test999",
			keys:          []string{"third.carrot"},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example.json",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_04.json",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_04.json",
			password:      "test999",
			keys:          []string{"first.a", "first.z", "second.b.2", "fourth.list"},
		},
		{
			inputFile:     "../../resources/tests/partial_encrypted_example_02.json",
			outputFile:    "../../temp/partial_encrypted_example_decrypted_02.json",
			referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.json",
			password:      "test999",
			keys:          []string{},
		},
	}
	for _, test := range tests {
		if _, err := os.Stat(test.outputFile); err == nil {
			// file exists
			os.Remove(test.outputFile)
			_, err = os.Stat(test.outputFile)
			require.NotNil(t, err)
		}
		err := decrypt.Decrypt(test.inputFile, test.outputFile, test.password, test.keys)
		require.Nil(t, err)
		_, err = os.Stat(test.outputFile)
		require.Nil(t, err)
		outputJson, err := jsonreader.ReadJSON(test.outputFile)
		require.Nil(t, err)
		referenceJson, err := jsonreader.ReadJSON(test.referenceFile)
		require.Equal(t, referenceJson, outputJson, "reference and decrypted output are different")
	}
}
