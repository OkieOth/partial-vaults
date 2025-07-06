package decrypt

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/pkg/types"
)

func decryptImpl(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	// switch vt {
	// case types.STRING:
	// 	s := string(input)
	// 	return s, vt, types.HANDLING_PROCESS, nil
	// case types.BOOL:
	// 	s := string(input)
	// 	return s, vt, types.HANDLING_PROCESS, nil
	// case types.INTEGER:
	//     s := fmt.Sprintf(format string, a ...any)
	// case types.NUMBER:
	// case types.NULL:
	// 	return input, vt, types.HANDLING_PROCESS, nil
	// }
	return input, vt, types.HANDLING_PROCESS, nil // TODO
}

func Decrypt(inputFile, outputFile, password string, keys []string) error {
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		fmt.Println("Decrypt", "key: ", keyPath, "value: ", string(input))
		return string(input) + "_changed", types.STRING, types.HANDLING_PROCESS, nil
	}
	return cmdbase.CommandBase(inputFile, outputFile, processor)
}

func DecryptInteractive(inputFile, outputFile, password string, keys []string) error {
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		fmt.Println("Decrypt Interactive variant", "key: ", keyPath, "value: ", string(input))
		return input, vt, types.HANDLING_PROCESS, nil
	}

	introMsg := "This is the interactive decryption of: "
	processQuestion := "Decrypt value?"
	interactiveProcessor := cmdbase.NewInteractiveProcessor(inputFile, introMsg, processQuestion, processor)
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor.Process)
}
