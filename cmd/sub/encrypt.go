package sub

import (
	"fmt"

	"github.com/okieoth/pvault/internal/pkg/typedetect"
	"github.com/okieoth/pvault/pkg/encrypt"
	"github.com/spf13/cobra"
)

func EncryptImpl(inputFile, outputFile, passwordToUse string, keysToEncrypt []string, overwriteInput, stdoutOutput, forAnsible, runInteractive bool) error {
	var outputFileToUse string
	if newOutput, cont := CheckForOutput(inputFile, outputFile, overwriteInput, stdoutOutput); cont {
		outputFile = newOutput
	} else {
		return fmt.Errorf("given command switches are not sufficient")
	}
	needTmpIntermediateFile := false
	if forAnsible {
		// adjust the output file for Ansible
		if t, err := typedetect.DetectFormat(inputFile); err == nil && t == typedetect.INPUT_YAML {
			needTmpIntermediateFile = true
		}
		if needTmpIntermediateFile {
			outputFileToUse = CreateIntermediateFile()
		}
	} else {
		outputFileToUse = outputFile
	}

	if runInteractive {
		encrypt.EncryptInteractive(inputFile, outputFileToUse, passwordToUse, keysToEncrypt)
	} else {
		encrypt.Encrypt(inputFile, outputFileToUse, passwordToUse, keysToEncrypt)
	}
	if needTmpIntermediateFile {
		// adjust the output file for Ansible
		if err := CreateOutputFromIntermediate(outputFileToUse, outputFile); err != nil {
			fmt.Println("error while creating ansible usable version from intermediate file:", err)
			return err
		}
	}
	return nil
}

var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Partial encrypts a JSON or YAML file",
	Long:  "Partial encrypts a JSON or YAML file in a Ansible vault compatible way",
	Run: func(cmd *cobra.Command, args []string) {
		EncryptImpl(input, output, password, keys, overwrite, stdout, ansibleUse, interactive)
	},
}

var ansibleUse bool

func init() {
	initDefaultFlags(EncryptCmd)
	EncryptCmd.Flags().BoolVar(&ansibleUse, "ansible", false, "Encrypt values in the by ansible used style")

}
