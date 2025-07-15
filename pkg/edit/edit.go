package edit

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/okieoth/pvault/internal/pkg/cmdbase"
	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/decrypt"
	"github.com/okieoth/pvault/pkg/encrypt"
	"github.com/okieoth/pvault/pkg/types"
)

func EditInteractive(inputFile, outputFile, password string, keys []string) error {
	decryptProcessor := decrypt.DecryptProcessor(password)
	encryptProcessor := encrypt.EncryptProcessor(password)

	introMsg := "This is the interactive edit of: "
	processQuestion := "Edit value?"
	interactiveProcessor := InteractiveEditProcessor(introMsg, processQuestion, inputFile, decryptProcessor, encryptProcessor)
	return cmdbase.CommandBase(inputFile, outputFile, interactiveProcessor, keys)
}

func InteractiveEditProcessor(introMsg, processQuestion, inputFile string, decryptProcessor, encryptProcessor types.ProcessFunc) types.ProcessFunc {
	colored := color.New(color.FgGreen)
	colored = colored.Add(color.Bold)
	first := true
	return func(input any, inputType types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		if first {
			first = false
			colored.Println(introMsg, inputFile)
			colored.Println("('y' - takes the suggestion (default), 'n' - rejects the suggestion, 'c' - cancel processing)")
		}
		decryptedInput, valueType, handling, _ := decryptProcessor(input, inputType, keyPath)
		if handling == types.HANDLING_SKIP {
			return input, valueType, types.HANDLING_SKIP, nil
		}
		fmt.Println()
		colored.Print("key: ")
		fmt.Print(keyPath)
		colored.Print(", value: ")
		fmt.Printf("%v\n", decryptedInput)
		colored.Print("New value ⏎: ")
		reader := bufio.NewReader(os.Stdin)
		return readNewValueAndEncryptIt(reader, colored, encryptProcessor, keyPath)
	}
}

func readNewValueAndEncryptIt(reader *bufio.Reader, colored *color.Color, encryptProcessor types.ProcessFunc, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	colored.Print("New value ⏎: ")
	stdinInput, err := reader.ReadString('\n')
	if err != nil {
		return "", types.STRING, types.HANDLING_SKIP, fmt.Errorf("Error while reading input: %v", err)
	}
	trimmedInput := strings.TrimSpace(stdinInput)
	if len(trimmedInput) == 0 {
		colored.Println("No input, original value will stay unchanged")
		return "", types.STRING, types.HANDLING_SKIP, nil
	}
	v, vt := vaultfunc.InputType(trimmedInput)
	return encryptProcessor(v, vt, keyPath)
}
