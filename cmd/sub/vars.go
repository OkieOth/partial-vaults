package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

var input string
var output string

var keys []string
var interactive bool
var password string
var overwrite bool
var stdout bool

func CheckForOutput(input, output string, overwrite, stdout bool) (string, bool) {
	if output == "" {
		if overwrite {
			return input, true
		} else if stdout {
			return "stdout", true
		} else {
			fmt.Println("No output file, overwrite flag or stdout given, cancel.")
			return "", false
		}
	}
	return output, true
}

func initDefaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&input, "input", "i", "", "Path to the input file (required)")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file to create, needs to be given or one of 'overwrite' or 'stdout')")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password (required)")
	cmd.Flags().StringSliceVarP(&keys, "key", "k", []string{}, "Keys to include in the processing. For nested keys the dot notated styles is to use. Can be given multiple times")
	cmd.Flags().BoolVarP(&interactive, "interactive", "a", false, "Interactive mode")
	cmd.Flags().BoolVar(&overwrite, "overwrite", false, "if that's given instead of the output parameter, then the input file will be overwritten")
	cmd.Flags().BoolVar(&stdout, "stdout", false, "give it instead of 'output' or 'overwrite' to print the result to stdout")
	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("password")
}
