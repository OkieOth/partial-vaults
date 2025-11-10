package sub

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/typedetect"
	"github.com/okieoth/pvault/pkg/encrypt"
	"github.com/spf13/cobra"
)

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Partial encrypts a JSON or YAML file",
	Long:  "Partial encrypts a JSON or YAML file in a Ansible vault compatible way",
	Run: func(cmd *cobra.Command, args []string) {
		var outputFileToUse string
		if overwrite {
			outputFileToUse = input
		} else {
			if output == "" {
				fmt.Println("No output file or overwrite flag given, cancel.")
				return
			}
			outputFileToUse = output
		}
		needTmpIntermediateFile := false
		if ansibleUse {
			// adjust the output file for Ansible
			if t, err := typedetect.DetectFormat(input); err == nil && t == typedetect.INPUT_YAML {
				needTmpIntermediateFile = true
			}
			if needTmpIntermediateFile {
				outputFileToUse = CreateIntermediateFile()
			}
		}

		if interactive {
			encrypt.EncryptInteractive(input, outputFileToUse, password, keys)

		} else {
			encrypt.Encrypt(input, outputFileToUse, password, keys)
		}
		if needTmpIntermediateFile {
			// adjust the output file for Ansible
			if err := CreateOutputFromIntermediate(outputFileToUse, output); err != nil {
				fmt.Println("error while creating ansible usable version from intermediate file", err)
			}
		}
	},
}

var ansibleUse bool

func init() {
	initDefaultFlags(EncryptCmd)
	EncryptCmd.Flags().BoolVar(&ansibleUse, "ansible", false, "Encrypt values in the by ansible used style")

}
