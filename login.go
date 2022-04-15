package main

import (
	"fmt"
	"main/auth"
	"main/config"
	"main/ddns"
)

func main() {
	nuist := config.GetConfig()
	attempt := 5
	// ip := ""
	fmt.Printf("username is %s, password is %s, isp is %s\n", nuist.Username, nuist.Password, nuist.ISP)
	ip, err := auth.DoLogin(nuist, attempt)
	fmt.Printf("ip is %s, isp is %s\n", ip, err)
	cfddns := config.GetCFConfig()
	fmt.Printf("cfddns is %s", cfddns)
	api, err := ddns.NewCloudflareClient(cfddns.Cloudflare_API_Token, cfddns.Cloudflare_ZONEID, cfddns.Cloudflare_Host)
	records, _ := api.ListDNSRecords("A")
	fmt.Printf("records%v", records)
}
