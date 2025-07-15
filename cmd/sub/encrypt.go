package sub

import (
	"github.com/okieoth/pvault/pkg/encrypt"
	"github.com/spf13/cobra"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Partial encrypts a JSON or YAML file",
	Long:  "Partial encrypts a JSON or YAML file in a Ansible vault compatible way",
	Run: func(cmd *cobra.Command, args []string) {
		if interactive {
			encrypt.EncryptInteractive(input, output, password, keys)

		} else {
			encrypt.Encrypt(input, output, password, keys)
		}
	},
}

func init() {
	initDefaultFlags(EncryptCmd)
}
