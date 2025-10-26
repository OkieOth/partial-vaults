package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/okieoth/ic0bra"
	"github.com/okieoth/pvault/cmd/sub"
)

var rootCmd = &cobra.Command{
	Use:   "pvault",
	Short: "Tool to partial encrypt and decrypt JSON or YAML files",
	Long:  `Tool to partial encrypt and decrypt JSON or YAML files, compatible to Ansible Vault spec.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmdToCall, err := ic0bra.RunInteractiveWithHistory(cmd, "pvault"); err == nil {
			if cmdToCall != nil {
				cmdToCall.Run(cmdToCall, args)
			}
		} else {
			fmt.Println("error while running in interactive mode:", err)
		}
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
