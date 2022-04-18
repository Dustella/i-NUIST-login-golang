package config

import "fmt"

func proNuist() NUIST_Auth {
	var username, password, isp string
	fmt.Println("Enter your username:")
	fmt.Scan(&username)
	fmt.Println("Enter your password:")
	fmt.Scan(&password)
	fmt.Println("Enter your isp:")
	fmt.Scan(&isp)
	res := NUIST_Auth{Username: username, Password: password, ISP: isp}
	return res
}

func proCFDDNS() Cloudflare_Config {
	var enable string
	var zoneID, APIToken, hostname string
	fmt.Println("Would you like to configure cloudflare DDNS?")
	fmt.Println("Yes (y) / No need (n)")
	fmt.Scan(&enable)
	if enable == "y" {
		fmt.Println("Enter your zone ID:")
		fmt.Scan(&zoneID)
		fmt.Println("Enter your APIToken:")
		fmt.Scan(&APIToken)
		fmt.Println("Enter your hostname:")
		fmt.Scan(&hostname)
		return Cloudflare_Config{
			Cloudflare_ZONEID:    zoneID,
			Cloudflare_API_Token: APIToken,
			Cloudflare_Host:      hostname,
		}
	} else if enable == "n" {
		return Cloudflare_Config{}
	}
	return Cloudflare_Config{}
}
