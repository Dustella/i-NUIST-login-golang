package controller

import (
	"log"
	"main/client"
	"main/config"

	"github.com/spf13/viper"
)

func printInfo() {
	mode := config.GetMode()
	userPool := config.GetUserPool()
	interfaces := config.ListNetworkInterface()
	log.Println("retry:", viper.GetInt("retry"))
	log.Println("timeout:", viper.GetInt("timeout"))
	log.Println("mode:", mode)
	log.Println("userPool:", userPool)
	for _, inf := range interfaces {
		ip, _ := inf.Addrs()
		log.Println("interface:", inf.Name)
		for single_ip := range ip {
			log.Print("ip: ", ip[single_ip])
		}
		println("====")
	}

	client.InitClients()
	// for _, client := range client.SyncClients {
	// 	log.Println("client:", client)
	// }
}
