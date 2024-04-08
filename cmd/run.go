package cmd

import (
	"github.com/jqwez/namecheap-update-dns/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "The 'run' command starts your program.",
	Long:  `The 'run' command will run the program according to your env or associated flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.ValidateConfig()
		internal.UpdateRecords()
	},
}

func init() {
}
