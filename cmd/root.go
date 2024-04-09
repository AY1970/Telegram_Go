package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tebot",
	Short: "Test Boot for Learning Go",
	Long:  "This is telegram bot written in Go and used only for learning porposes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startBot)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
