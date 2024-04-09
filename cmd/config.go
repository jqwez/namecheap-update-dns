package cmd

import (
	"fmt"
	"log"

	"github.com/jqwez/namecheap-update-dns/internal"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [command]",
	Short: "The 'config' command writes/edits your .env file",
	Long: `The 'config' command is helper tool to set up your .env file
Available Commands:
	show		Displays the current .env configuration
	edit		Edits current .env configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		appConfig, err := internal.GetConfigFromEnv()
		if err != nil {
			log.Println("Config not yet valid")
			log.Println(err)
		}
		if len(args) <= 0 {
			fmt.Println("use 'config --help' for more info")
			return
		}

		switch args[0] {
		case "show":
			show(appConfig)
			return
		case "edit":
			internal.RunEditForm(appConfig)
		default:
			return
		}
	},
}

func show(appConfig *internal.AppConfig) {
	fmt.Printf("Current Configuration:\n")
	appConfig.Print(true)
}

func init() {
	configCmd.PersistentFlags().Bool("hide-api-key", true, "Sets whether API Key is display on 'show'")
}
