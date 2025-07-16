package encrypt_test

import (
	"os"
	"testing"

	"github.com/okieoth/pvault/internal/pkg/jsonreader"
	"github.com/okieoth/pvault/internal/pkg/yamlreader"
	"github.com/okieoth/pvault/pkg/decrypt"
	"github.com/okieoth/pvault/pkg/encrypt"
	"github.com/okieoth/pvault/pkg/json"
	"github.com/stretchr/testify/require"
)

func TestEncryptYaml(t *testing.T) {
	tests := []struct {
		inputFile     string
		outputFile    string
		referenceFile string
		tmpFile       string
		password      string
		keys          []string
	}{
		{
			inputFile:     "../../resources/tests/example.yaml",
			outputFile:    "../../temp/example_encrypted_01.yaml",
			referenceFile: "../../temp/example_encrypted_tmp_01.yaml",
			tmpFile:       "../../temp/example_encrypted_tmp_01.yaml",
			password:      "test999",
			keys:          []string{},
		},
		{
			inputFile:     "../../resources/tests/example.yaml",
			outputFile:    "../../temp/example_encrypted_03.yaml",
			referenceFile: "../../temp/example_encrypted_tmp_03.yaml",
			tmpFile:       "../../temp/example_encrypted_tmp_03.yaml",
			password:      "test999",
			keys:          []string{"third.carrot"},
		},
		// {
		// 	inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
		// 	outputFile:    "../../temp/partial_encrypted_example_decrypted_04.yaml",
		// 	referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_04.yaml",
		// 	password:      "test999",
		// 	keys:          []string{"first.a", "first.z", "second.b.2", "fourth.list"},
		// },
		// {
		// 	inputFile:     "../../resources/tests/partial_encrypted_example_02.yaml",
		// 	outputFile:    "../../temp/partial_encrypted_example_decrypted_02.yaml",
		// 	referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.yaml",
		// 	password:      "test999",
		// 	keys:          []string{},
		// },
	}
	for _, test := range tests {
		if _, err := os.Stat(test.outputFile); err == nil {
			// file exists
			os.Remove(test.outputFile)
			_, err = os.Stat(test.outputFile)
			require.NotNil(t, err)
		}
		err := encrypt.Encrypt(test.inputFile, test.outputFile, test.password, test.keys)
		require.Nil(t, err)
		_, err = os.Stat(test.outputFile)
		require.Nil(t, err)
		err = decrypt.Decrypt(test.outputFile, test.tmpFile, test.password, test.keys)
		require.Nil(t, err)

		outputYaml, err := yamlreader.ReadYAML(test.tmpFile)
		require.Nil(t, err)
		inputYaml, err := yamlreader.ReadYAML(test.inputFile)
		require.Nil(t, err)
		require.Equal(t, inputYaml, outputYaml, "yaml: encrypted + decrypted doesn't mach input")
	}
}

func TestEncryptJson(t *testing.T) {
	tests := []struct {
		inputFile     string
		outputFile    string
		referenceFile string
		tmpFile       string
		password      string
		keys          []string
	}{
		{
			inputFile:     "../../resources/tests/example.json",
			outputFile:    "../../temp/example_encrypted_01.json",
			referenceFile: "../../resources/tests/example_encrypted_01.json",
			tmpFile:       "../../temp/example_encrypted_tmp_01.json",
			password:      "test999",
			keys:          []string{},
		},
		// {
		// 	inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
		// 	outputFile:    "../../temp/partial_encrypted_example_decrypted_03.yaml",
		// 	referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_03.yaml",
		// 	password:      "test999",
		// 	keys:          []string{"third.carrot"},
		// },
		// {
		// 	inputFile:     "../../resources/tests/partial_encrypted_example.yaml",
		// 	outputFile:    "../../temp/partial_encrypted_example_decrypted_04.yaml",
		// 	referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_04.yaml",
		// 	password:      "test999",
		// 	keys:          []string{"first.a", "first.z", "second.b.2", "fourth.list"},
		// },
		// {
		// 	inputFile:     "../../resources/tests/partial_encrypted_example_02.yaml",
		// 	outputFile:    "../../temp/partial_encrypted_example_decrypted_02.yaml",
		// 	referenceFile: "../../resources/tests/partial_encrypted_example_decrypted_01.yaml",
		// 	password:      "test999",
		// 	keys:          []string{},
		// },
	}
	for _, test := range tests {
		if _, err := os.Stat(test.outputFile); err == nil {
			// file exists
			os.Remove(test.outputFile)
			_, err = os.Stat(test.outputFile)
			require.NotNil(t, err)
		}
		err := encrypt.Encrypt(test.inputFile, test.outputFile, test.password, test.keys)
		require.Nil(t, err)
		_, err = os.Stat(test.outputFile)
		require.Nil(t, err)
		err = decrypt.Decrypt(test.outputFile, test.tmpFile, test.password, test.keys)
		require.Nil(t, err)

		tmpJson, err := jsonreader.ReadJSON(test.tmpFile)
		require.Nil(t, err)
		inputJson, err := jsonreader.ReadJSON(test.inputFile)
		require.Nil(t, err)
		require.Equal(t, inputJson, tmpJson, "json: encrypted + decrypted doesn't mach input")

		outputJson, err := jsonreader.ReadJSON(test.outputFile)
		require.Nil(t, err)
		encryptedKeys1, err := json.GetEncryptedKeys(&outputJson)
		require.Nil(t, err)

		referenceJson, err := jsonreader.ReadJSON(test.referenceFile)
		require.Nil(t, err)
		encryptedKeys2, err := json.GetEncryptedKeys(&referenceJson)
		require.Nil(t, err)
		require.Equal(t, encryptedKeys1, encryptedKeys2, "json: encrypted doesn't mach input")
	}
}
