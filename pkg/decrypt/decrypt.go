package decrypt

import (
	"fmt"
	"strings"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/types"
)

func decryptImpl(input any, vt types.ValueType, keyPath, password string) (any, types.ValueType, types.ProcessHandling, error) {
	if vt != types.STRING {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("skip decrypt because it's no string, keyPath: %s", keyPath)
	}

	valueToDecrypt, ok := input.(string)
	if !ok {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("error while casting value to decrypt to string, keyPath: %s", keyPath)
	}

	seperator := "$ANSIBLE_VAULT;"
	index := strings.Index(valueToDecrypt, seperator)
	if index == -1 {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("value to decrypt doesn't contain Ansible vault prefix, keyPath: %s", keyPath)
	}
	valueToDecrypt = valueToDecrypt[index:]
	v, vtype, err := vaultfunc.Decrypt(valueToDecrypt, password)
	if err != nil {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("error while decrypt, keyPath: %s, err: %v", keyPath, err)
	}
	return v, vtype, types.HANDLING_PROCESS, nil
}

func Decrypt(inputFile, outputFile, password string, keys []string) error {
	return cmdbase.CommandBase(inputFile, outputFile, DecryptProcessor(password))
}

func DecryptProcessor(password string) types.ProcessFunc {
	return func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return decryptImpl(input, vt, keyPath, password)
	}
}

func DecryptInteractive(inputFile, outputFile, password string, keys []string) error {
	introMsg := "This is the interactive decryption of: "
	processQuestion := "Decrypt value?"
	interactiveProcessor := cmdbase.InteractiveProcessor(inputFile, introMsg, processQuestion, DecryptProcessor(password))
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor)
}
