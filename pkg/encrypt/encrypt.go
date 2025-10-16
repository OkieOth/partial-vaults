package encrypt

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/types"
)

func encryptImpl(input any, vt types.ValueType, keyPath, password string) (string, types.ValueType, types.ProcessHandling, error) {
	v, err := vaultfunc.Encrypt(input, vt, password)
	if err != nil {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("error while encrypt, keyPath: %s, err: %v", keyPath, err)
	}
	return v, types.STRING, types.HANDLING_PROCESS, nil
}

func Encrypt(inputFile, outputFile, password string, keys []string) error {
	return cmdbase.CommandBase(inputFile, outputFile, EncryptProcessor(password), keys)
}

func EncryptProcessor(password string) types.ProcessFunc {
	return func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return encryptImpl(input, vt, keyPath, password)
	}
}

func EncryptInteractive(inputFile, outputFile, password string, keys []string) error {
	introMsg := "This is the interactive encryption of: "
	processQuestion := "Encrypt value?"
	interactiveProcessor := cmdbase.InteractiveProcessor(inputFile, introMsg, processQuestion, EncryptProcessor(password)) // TODO
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor, keys)
}
