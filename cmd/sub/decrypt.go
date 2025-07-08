package sub

import (
	"fmt"

	"github.com/okieoth/pvault/pkg/decrypt"
	"github.com/spf13/cobra"
)

var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts a partial encoded JSON or YAML file",
	Long:  "Decrypts a partial Ansible vault encoded JSON or YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		if interactive {
			decrypt.DecryptInteractive(input, output, password, keys)

		} else {
			if err := decrypt.Decrypt(input, output, password, keys); err != nil {
				fmt.Println("an error occured: ", err)
			}
		}
	},
}

func init() {
	initDefaultFlags(DecryptCmd)
}
