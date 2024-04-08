package internal

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"
)

func init() {
	godotenv.Load()
}

func UpdateRecords() {
	client := newClientFromEnv()
	config := &ReplaceRecordsConfig{
		DomainName: os.Getenv("NM_DOMAIN"),
		Targets:    strings.Split(os.Getenv("NM_TARGETS"), ";"),
	}
	ReplaceRecords(client, config)
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
			r.Address = &client.ClientOptions.ClientIp
		}
		records = append(records, r)
	}

	setHostsResp, err := client.DomainsDNS.SetHosts(&namecheap.DomainsDNSSetHostsArgs{
		Domain:    &domainName,
		Records:   &records,
		EmailType: response.DomainDNSGetHostsResult.EmailType,
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(setHostsResp)
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
