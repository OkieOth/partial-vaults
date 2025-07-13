package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/okieoth/pvault/cmd/sub"
)

var rootCmd = &cobra.Command{
	Use:   "pvault",
	Short: "Tool to partial encrypt and decrypt JSON or YAML files",
	Long:  `Tool to partial encrypt and decrypt JSON or YAML files, compatible to Ansible Vault spec.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please call this with one of the provided sub-commands")
	},
}

func init() {
	rootCmd.AddCommand(sub.EncryptCmd)
	rootCmd.AddCommand(sub.DecryptCmd)
	rootCmd.AddCommand(sub.EditCmd)
	rootCmd.AddCommand(sub.VersionCmd)
}

func main() {
	rootCmd.Execute()
}
