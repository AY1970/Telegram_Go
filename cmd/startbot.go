package cmd

import (
	"fmt"
	"main/bot"
	"os"

	"github.com/spf13/cobra"
)

var startBot = &cobra.Command{
	Use:   "start_bot",
	Short: "This command will start the telegram bot",
	Long:  "This command will check the TOKEN in env variable and start the bot",
	Run: func(cmd *cobra.Command, args []string) {
		var token = os.Getenv("TELE_TOKEN")
		if len(token) > 0 {
			bot.CreateBot(token)
		} else {
			fmt.Println("TOKEN is not correct or is not provided")
		}
	},
}
