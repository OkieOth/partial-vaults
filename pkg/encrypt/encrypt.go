package encrypt

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/types"
)

func encryptImpl(input any, vt types.ValueType, keyPath, password string) (any, types.ValueType, types.ProcessHandling, error) {
	v, err := vaultfunc.Encrypt(input, vt, password)
	if err != nil {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("error while encrypt, keyPath: %s, err: %v", keyPath, err)
	}
	return v, types.STRING, types.HANDLING_PROCESS, nil
}

func Encrypt(inputFile, outputFile, password string, keys []string) error {
	processor := func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return encryptImpl(input, vt, keyPath, password)
	}
	return cmdbase.CommandBase(inputFile, outputFile, processor)
}

func EncryptInteractive(inputFile, outputFile, password string, keys []string) error {
	processor := func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return encryptImpl(input, vt, keyPath, password)
	}

	introMsg := "This is the interactive encryption of: "
	processQuestion := "Encrypt value?"
	interactiveProcessor := cmdbase.NewInteractiveProcessor(inputFile, introMsg, processQuestion, processor)
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor.Process)
}
