package cmdbase

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/okieoth/pvault/internal/pkg/typedetect"
	"github.com/okieoth/pvault/pkg/json"
	"github.com/okieoth/pvault/pkg/types"
	"github.com/okieoth/pvault/pkg/yaml"
)

func CommandBase(input, output string, processor types.ProcessFunc, keys []string) error {
	var e error
	if t := typedetect.IfNotJsonOrYamlThenPanic(input); t == typedetect.INPUT_JSON {
		e = json.ProcessJsonFile(input, output, processor, keys)
	} else {
		e = yaml.ProcessYamlFile(input, output, processor, keys)
	}
	return e
}

func InteractiveProcessor(introMsg, processQuestion, inputFile string, processor types.ProcessFunc) types.ProcessFunc {
	colored := color.New(color.FgGreen)
	colored = colored.Add(color.Bold)
	first := true
	return func(input any, inputType types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		if first {
			first = false
			colored.Println(introMsg, inputFile)
			colored.Println("('y' - takes the suggestion (default), 'n' - rejects the suggestion, 'c' - cancel processing)")
		}
		fmt.Println()
		colored.Print("key: ")
		fmt.Print(keyPath)
		colored.Print(", value: ")
		fmt.Printf("%s\n", input)
		colored.Printf("%s [y|n|(C)ancel]: ", processQuestion)
		reader := bufio.NewReader(os.Stdin)
		for {
			stdinInput, err := reader.ReadString('\n')
			if err != nil {
				return input, inputType, types.HANDLING_PROCESS, fmt.Errorf("Error while reading input: %v", err)
			}
			trimmedInput := strings.TrimSpace(stdinInput)
			switch trimmedInput {
			case "Y", "y", "yes", "Yes", "":
				// TODO process the value
				return processor(input, inputType, keyPath)
			case "N", "n", "no", "No":
				return input, inputType, types.HANDLING_SKIP, nil
			case "c", "C":
				return input, inputType, types.HANDLING_CANCEL, nil
			default:
				colored.Println("WRONG INPUT! Only [y|n|c] are allowed!")
			}
		}
	}
}
