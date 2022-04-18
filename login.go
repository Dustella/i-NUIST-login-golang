package main

import (
	"fmt"
	"main/auth"
	"main/config"
	"main/ddns"
)

func main() {
	fmt.Println("==================NUIST LOGIN GOLANG =================")
	fmt.Println("Doing Auth......")
	nuist := config.GetConfig()
	attempt := 5
	ip, _ := auth.DoLogin(nuist, attempt)
	fmt.Println("==================LOGIN IS OK =================")
	fmt.Println("Doing DDNS......")
	cfddns := config.GetCFConfig()
	ddns.DoDDNSCF(ip, cfddns)
}
