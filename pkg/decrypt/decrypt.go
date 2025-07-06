package decrypt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/pkg/types"
)

func Decrypt(inputFile, outputFile, password string, keys []string) error {
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		fmt.Println("Decrypt", "key: ", keyPath, "value: ", string(input))
		return string(input) + "_changed", types.STRING, types.HANDLING_PROCESS, nil
	}
	return cmdbase.CommandBase(inputFile, outputFile, processor)
}

func DecryptInteractive(inputFile, outputFile, password string, keys []string) error {
	first := true
	processor := func(input []byte, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		if first {
			first = false
			fmt.Println("This is the interactive decryption of: ", inputFile)
			fmt.Println("('y' - takes the suggestion, the default, 'n' - rejects the suggestion, 'c' - cancel processing)")
			fmt.Println()
		} else {
			fmt.Println()
		}
		fmt.Printf("key: '%s',\nvalue: %v\n", keyPath, inputFile)
		fmt.Print("Decrypt value? [y|n|(C)ancel]: ")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				return input, vt, types.HANDLING_PROCESS, fmt.Errorf("Error while reading input: %v", err)
			}
			trimmedInput := strings.TrimSpace(input)
			switch trimmedInput {
			case "Y", "y", "yes", "Yes", "":
				// TODO process the value
				return input, vt, types.HANDLING_PROCESS, nil
			case "N", "n", "no", "No":
				// TODO handle skip
				return input, vt, types.HANDLING_SKIP, nil
			case "c", "C":
				// TODO handle cancel
				return input, vt, types.HANDLING_CANCEL, nil
			default:
				fmt.Println("WRONG INPUT! Only [y|n|c] are allowed!")
			}
		}
	}
	return cmdbase.CommandBase(inputFile, outputFile, processor)
}
