package cmd

import (
	"log"

	"github.com/jqwez/namecheap-update-dns/internal"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "The 'config' command edits your .env file",
	Long:  `The 'edit' command is helper tool to set up your .env file`,
	Run: func(cmd *cobra.Command, args []string) {
		valid, err := internal.ValidateConfig()
		if err != nil {
			log.Println("Config not yet valid")
		}
		log.Println(valid)

		internal.UpdateRecords()
	},
}

func init() {
}
