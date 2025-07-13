package decrypt

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/keys"
	"github.com/okieoth/pvault/pkg/types"
)

func decryptImpl(input any, vt types.ValueType, keyPath, password string) (any, types.ValueType, types.ProcessHandling, error) {
	encrypted, valueToDecrypt, err := keys.IsEncrypted(input, vt)
	if !encrypted {
		return input, vt, types.HANDLING_SKIP, err
	}
	v, vtype, err := vaultfunc.Decrypt(valueToDecrypt, password)
	if err != nil {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("error while decrypt, keyPath: %s, err: %v", keyPath, err)
	}
	return v, vtype, types.HANDLING_PROCESS, nil
}

func Decrypt(inputFile, outputFile, password string, keys []string) error {
	return cmdbase.CommandBase(inputFile, outputFile, DecryptProcessor(password), keys)
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
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor, keys)
}
