package sub

import "github.com/spf13/cobra"

var input string
var output string

var keys []string
var interactive bool
var password string

func initDefaultFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&input, "input", "i", "", "Path to the input file (required)")
	cmd.Flags().StringVarP(&output, "output", "o", "stdout", "Output to create, default is stdout")
	cmd.MarkFlagRequired("input")
	cmd.Flags().StringSliceVarP(&keys, "key", "k", []string{}, "Keys to include in the processing. For nested keys the dot notated styles is to use. Can be given multiple times")
	cmd.Flags().BoolVarP(&interactive, "interactive", "a", false, "Interactive mode")
	cmd.Flags().StringVarP(&password, "passord", "p", "", "Password")
}
