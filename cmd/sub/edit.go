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
		edit.EditInteractive(input, output, password, keys)
	},
}

func init() {
	initDefaultFlags(EditCmd)
}
