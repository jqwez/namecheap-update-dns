package cmd

import (
	"fmt"

	"github.com/jqwez/namecheap-update-dns/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "The 'run' command starts your program.",
	Long:  `The 'run' command will run the program according to your env or associated flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		appConfig, err := internal.GetConfigFromEnv()
		if err != nil {
			fmt.Println("Failed to load from config, please run 'config edit'")
		}
		internal.UpdateRecords(appConfig)
	},
}

func init() {
}
