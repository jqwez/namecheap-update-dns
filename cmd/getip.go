package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jqwez/namecheap-update-dns/internal"
	"github.com/spf13/cobra"
)

var getIpCmd = &cobra.Command{
	Use:   "get-ip",
	Short: "The 'get-ip' command prints your public ip to the cli",
	Long:  "The 'get-ip' command prints your public ip to the cli",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := internal.GetMyIp()
		if err != nil {
			log.Println("Failed to fetch IP. Try sudo, firewall settings, or else manually add IP to .env")
			log.Fatal(err)
		}
		fmt.Println("Public IP: " + ip)
		fmt.Println(".env IP : " + os.Getenv("NM_CLIENT_IP"))

	},
}

func init() {
}
