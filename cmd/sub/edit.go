package sub

import (
	"github.com/okieoth/pvault/pkg/edit"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Enables in place changes",
	Long:  "Allows interactive editing of partial encrypts a JSON or YAML file in a Ansible vault compatible way",
	Run: func(cmd *cobra.Command, args []string) {
		if newOutput, cont := CheckForOutput(input, output, overwrite, stdout); cont {
			output = newOutput
		} else {
			return
		}
		edit.EditInteractive(input, output, password, keys)
	},
}

func init() {
	initDefaultFlags(EditCmd)
}
