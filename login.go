package main

import (
	"fmt"
	"main/auth"
	"main/config"
	"main/ddns"
)

func main() {
	nuist := config.GetConfig()
	cfddns := config.GetCFConfig()
	attempt := 5
	// ip := ""
	fmt.Printf("username is %s, password is %s, isp is %s\n", nuist.Username, nuist.Password, nuist.ISP)
	ip, err := auth.DoLogin(nuist, attempt)
	fmt.Printf("ip is %s, isp is %s\n", ip, err)
	ddns.DoDDNSCF(ip, cfddns)
}
