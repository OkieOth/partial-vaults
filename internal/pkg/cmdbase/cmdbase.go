package cmdbase

import (
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
