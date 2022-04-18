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
		fmt.Print(cfconfig.Cloudflare_API_Token)
		records, _ := api.ListDNSRecords("A")
		if len(records) == 0 {
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
			err := api.UpdateDNSRecord(Record{
				ID:      records[0].ID,
				Content: ip,
				Name:    cfconfig.Cloudflare_Host,
				Type:    "A",
				Proxied: false,
				TTL:     1})
			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
