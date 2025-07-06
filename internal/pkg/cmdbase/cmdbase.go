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

func CommandBase(input, output string, processor types.ProcessFunc) error {
	var e error
	if t := typedetect.IfNotJsonOrYamlThenPanic(input); t == typedetect.INPUT_JSON {
		e = json.ProcessJsonFile(input, output, processor)
	} else {
		e = yaml.ProcessYamlFile(input, output, processor)
	}
	return e
}

type InteractiveProcessor struct {
	introMsg        string
	processQuestion string
	first           bool
	inputFile       string
	processor       types.ProcessFunc
	colored         *color.Color
}

func NewInteractiveProcessor(inputFile, introMsg, processQuestion string, processor types.ProcessFunc) *InteractiveProcessor {
	c := color.New(color.FgGreen)
	c = c.Add(color.Bold)
	return &InteractiveProcessor{
		introMsg:        introMsg,
		processQuestion: processQuestion,
		inputFile:       inputFile,
		processor:       processor,
		first:           true,
		colored:         c,
	}
}

func (ip *InteractiveProcessor) Process(input []byte, inputType types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	if ip.first {
		ip.first = false
		ip.colored.Println(ip.introMsg, ip.inputFile)
		ip.colored.Println("('y' - takes the suggestion (default), 'n' - rejects the suggestion, 'c' - cancel processing)")
	}
	fmt.Println()
	ip.colored.Print("key: ")
	fmt.Print(keyPath)
	ip.colored.Print(", value: ")
	fmt.Printf("%s\n", input)
	ip.colored.Printf("%s [y|n|(C)ancel]: ", ip.processQuestion)
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
			return ip.processor(input, inputType, keyPath)
		case "N", "n", "no", "No":
			return input, inputType, types.HANDLING_SKIP, nil
		case "c", "C":
			return input, inputType, types.HANDLING_CANCEL, nil
		default:
			ip.colored.Println("WRONG INPUT! Only [y|n|c] are allowed!")
		}
	}
}
