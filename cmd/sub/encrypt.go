package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Partial encrypts a JSON or YAML file",
	Long:  "Partial encrypts a JSON or YAML file in a Ansible vault compatible way",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Encrypt - TODO")
	},
}

func init() {
	initDefaultFlags(EncryptCmd)
}
