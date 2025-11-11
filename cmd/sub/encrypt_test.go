package sub_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/okieoth/pvault/cmd/sub"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		inputFile      string
		outputFile     string
		passwordToUse  string
		keysToEncrypt  []string
		overwriteInput bool
		stdoutOutput   bool
		forAnsible     bool
	}{
		{
			inputFile:      "../../resources/tests/example.yaml",
			outputFile:     "../../temp/t_example_01.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{"third.orange"},
			overwriteInput: false,
			stdoutOutput:   false,
			forAnsible:     false,
		},
		{
			inputFile:      "../../resources/tests/example.yaml",
			outputFile:     "../../temp/t_example_02.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{},
			overwriteInput: false,
			stdoutOutput:   false,
			forAnsible:     false,
		},
		{
			inputFile:      "../../resources/tests/example.yaml",
			outputFile:     "../../temp/t_example_03.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{"third.orange"},
			overwriteInput: false,
			stdoutOutput:   false,
			forAnsible:     true,
		},
		{
			inputFile:      "../../resources/tests/example.yaml",
			outputFile:     "../../temp/t_example_04.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{},
			overwriteInput: false,
			stdoutOutput:   false,
			forAnsible:     true,
		},
		{
			inputFile:      "../../temp/t_example_03.yaml",
			outputFile:     "../../temp/t_example_03.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{"second.b.3"},
			overwriteInput: true,
			stdoutOutput:   false,
			forAnsible:     false,
		},
		{
			inputFile:      "../../temp/t_example_03.yaml",
			outputFile:     "../../temp/t_example_03.yaml",
			passwordToUse:  "test999",
			keysToEncrypt:  []string{"second.b.1"},
			overwriteInput: true,
			stdoutOutput:   false,
			forAnsible:     true,
		},
		{
			inputFile:      "../../temp/t_example_03.yaml",
			outputFile:     "../../temp/t_example_03.yaml",
			passwordToUse:  "test999",
			overwriteInput: true,
			stdoutOutput:   false,
			forAnsible:     true,
		},
	}
	for i, test := range tests {
		if test.inputFile != test.outputFile {
			if _, err := os.Stat(test.outputFile); err == nil {
				os.Remove(test.outputFile)
			}
			require.NoFileExists(t, test.outputFile, fmt.Sprintf("[Test: %d] output file already exist", i))
		}
		err := sub.EncryptImpl(test.inputFile, test.outputFile, test.passwordToUse,
			test.keysToEncrypt, test.overwriteInput, test.stdoutOutput, test.forAnsible,
			false)
		require.Nil(t, err, fmt.Sprintf("[Test: %d] error while encrypting", i))
		require.FileExists(t, test.outputFile, fmt.Sprintf("[Test: %d] cant't find output file", i))
	}
}
