package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nm_updatedns",
	Short: "Namecheap Update DNS (nm_updateip) Updates your dynamic address to your DNS config",
	Long:  `Namecheap Upate DNS. A util written in go that updates your DNS config to reflect your current IP address.`,
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
