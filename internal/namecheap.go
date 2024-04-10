package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"
)

func init() {
	godotenv.Load()
}

func UpdateRecords(ac *AppConfig) {
	client := newClientFromConfig(ac)
	config := &ReplaceRecordsConfig{
		DomainName: ac.DomainName,
		Targets:    ac.Targets,
	}
	err := ReplaceRecords(client, config)
	if err != nil {
		log.Println(err)
		return
	}
	ac.WriteToDotEnv()
	fmt.Println("Succesfully updated at Namecheap")
}

func newClientFromConfig(ac *AppConfig) *namecheap.Client {
	return namecheap.NewClient(&namecheap.ClientOptions{
		UserName:   ac.Username,
		ApiUser:    ac.ApiUser,
		ApiKey:     ac.ApiKey,
		ClientIp:   ac.ClientIp,
		UseSandbox: false,
	})
}

func newClientFromEnv() *namecheap.Client {
	return namecheap.NewClient(&namecheap.ClientOptions{
		UserName:   os.Getenv("NM_USERNAME"),
		ApiUser:    os.Getenv("NM_API_USER"),
		ApiKey:     os.Getenv("NM_API_KEY"),
		ClientIp:   os.Getenv("NM_CLIENT_IP"),
		UseSandbox: false,
	})
}

type ReplaceRecordsConfig struct {
	DomainName string
	Targets    []string
}

func ReplaceRecords(client *namecheap.Client, config *ReplaceRecordsConfig) error {
	domainName := config.DomainName
	targets := config.Targets
	if domainName == "" {
		log.Fatal("Did not find domain name in .env")
	}
	response, err := client.DomainsDNS.GetHosts(domainName)
	if err != nil {
		log.Fatal(err)
	}

	var didChange bool = false
	var records []namecheap.DomainsDNSHostRecord
	for _, old := range *response.DomainDNSGetHostsResult.Hosts {
		mxPref := (uint8)(*old.MXPref)

		r := namecheap.DomainsDNSHostRecord{
			HostName:   old.Name,
			RecordType: old.Type,
			Address:    old.Address,
			MXPref:     &mxPref,
			TTL:        old.TTL,
		}
		if isTargetName(targets, *r.HostName) {
			if *old.Address != client.ClientOptions.ClientIp {
				didChange = true
			}
			r.Address = &client.ClientOptions.ClientIp
		}
		records = append(records, r)
	}

	if !didChange {
		fmt.Println("IP Checked and did not change!")
		return nil
	}
	_, err = client.DomainsDNS.SetHosts(&namecheap.DomainsDNSSetHostsArgs{
		Domain:    &domainName,
		Records:   &records,
		EmailType: response.DomainDNSGetHostsResult.EmailType,
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func isTargetName(targets []string, compare string) bool {
	for _, target := range targets {
		if target == compare {
			return true
		}
	}
	return false
}
