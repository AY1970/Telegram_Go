package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [string to print]",
	Short: "Version of the application",
	Long:  "Version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Telegram Bot. Version 1.0.1")
		fmt.Println("Hello from my point of view")
	},
}
