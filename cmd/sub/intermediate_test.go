package sub_test

import (
	"fmt"
	"testing"

	"github.com/okieoth/pvault/cmd/sub"
	"github.com/okieoth/pvault/internal/pkg/yamlreader"
	"github.com/stretchr/testify/require"
)

func TestCreateIntermediateFile(t *testing.T) {
	f1 := sub.CreateIntermediateFile()
	require.NotEmpty(t, f1, "intermediate file (f1) wasn't created")
	fmt.Println("f1: ", f1)
	f2 := sub.CreateIntermediateFile()
	require.NotEmpty(t, f2, "intermediate file (f2) wasn't created")
	fmt.Println("f2: ", f2)
	require.NotEqual(t, f1, f2)
	f3 := sub.CreateIntermediateFile()
	require.NotEmpty(t, f3, "intermediate file (f3) wasn't created")
	fmt.Println("f3: ", f3)
	require.NotEqual(t, f1, f3)
	require.NotEqual(t, f2, f3)
	f4 := sub.CreateIntermediateFile()
	require.NotEmpty(t, f4, "intermediate file (f4) wasn't created")
	fmt.Println("f4: ", f4)
	require.NotEqual(t, f1, f4)
	require.NotEqual(t, f2, f4)
	require.NotEqual(t, f3, f4)
	f5 := sub.CreateIntermediateFile()
	require.NotEmpty(t, f5, "intermediate file (f5) wasn't created")
	fmt.Println("f5: ", f5)
	require.NotEqual(t, f1, f5)
	require.NotEqual(t, f2, f5)
	require.NotEqual(t, f3, f5)
	require.NotEqual(t, f4, f5)
}

func TestCreateOutputFromIntermediate(t *testing.T) {
	tests := []struct {
		input     string
		output    string
		reference string
	}{
		{
			input:     "../../resources/tests/example_changes_no.yaml",
			output:    "../../temp/example_changes_no.yaml",
			reference: "../../resources/tests/example_changes_no.yaml",
		},
		{
			input:     "../../resources/tests/example_changes.yaml",
			output:    "../../temp/example_changes.yaml",
			reference: "../../resources/tests/example_changes_ref.yaml",
		},
		{
			input:     "../../temp/example_changes.yaml",
			output:    "../../temp/example_changes_02.yaml",
			reference: "../../resources/tests/example_changes_ref.yaml",
		},
	}

	for i, test := range tests {
		err := sub.CreateOutputFromIntermediate(test.input, test.output)
		require.Nil(t, err, "got error in func call", i)

		outputYaml, err := yamlreader.ReadYAML(test.output)
		require.Nil(t, err)
		referenceYaml, err := yamlreader.ReadYAML(test.reference)
		require.Nil(t, err)
		require.Equal(t, referenceYaml, outputYaml, "created output doesn't have the expected output")
	}
}
