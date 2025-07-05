package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts a partial encoded JSON or YAML file",
	Long:  "Decrypts a partial Ansible vault encoded JSON or YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Decrypt - TODO")
	},
}

func init() {
	initDefaultFlags(DecryptCmd)
}
