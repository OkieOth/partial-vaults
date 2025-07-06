package encrypt

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/pkg/types"
)

func Encrypt(input, output, password string, keys []string) error {
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		fmt.Println("Encrypt", "key: ", keyPath, "value: ", string(input))
		return string(input) + "_changed", types.STRING, types.HANDLING_PROCESS, nil
	}
	return cmdbase.CommandBase(input, output, processor)
}

func EncryptInteractive(inputFile, outputFile, password string, keys []string) error {
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		fmt.Println("Encrypt Interactive variant", "key: ", keyPath, "value: ", string(input))
		return input, vt, types.HANDLING_PROCESS, nil
	}

	introMsg := "This is the interactive encryption of: "
	processQuestion := "Encrypt value?"
	interactiveProcessor := cmdbase.NewInteractiveProcessor(inputFile, introMsg, processQuestion, processor)
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor.Process)
}
