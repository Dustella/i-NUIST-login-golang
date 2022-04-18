package ddns

import (
	"fmt"
	"main/config"
)

func DoDDNSCF(ip string, cfconfig config.Cloudflare_Config) {
	if cfconfig.Cloudflare_API_Token != "" && cfconfig.Cloudflare_ZONEID != "" && cfconfig.Cloudflare_Host != "" {
		api, err := NewCloudflareClient(cfconfig.Cloudflare_API_Token, cfconfig.Cloudflare_ZONEID, cfconfig.Cloudflare_Host)
		if err != nil {
			fmt.Print(err)
		}
		records, _ := api.ListDNSRecords("A")
		if len(records) == 0 {
			fmt.Printf("There is no record named %s, \n", cfconfig.Cloudflare_Host)
			fmt.Println("Attempting to add this new record,")
			err := api.CreateDNSRecord(Record{
				Content: ip,
				Name:    cfconfig.Cloudflare_Host,
				Type:    "A",
				Proxied: false,
				TTL:     1})
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		if records[0].Content != ip {
			fmt.Printf("Record not match. \nPrevious record points to %s, \nWhile current IP is %s\n", records[0].Content, ip)
			fmt.Println("Attempting to update Record....")
			err := api.UpdateDNSRecord(Record{
				ID:      records[0].ID,
				Content: ip,
				Name:    cfconfig.Cloudflare_Host,
				Type:    "A",
				Proxied: false,
				TTL:     1})
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Printf("Record update OK")
		} else {
			fmt.Println("Record is match. No need to update record.")
		}
	}
}
