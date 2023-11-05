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
	log.Println("retry:", viper.GetInt("retry"))
	log.Println("timeout:", viper.GetInt("timeout"))
	log.Println("mode:", mode)
	log.Println("userPool:", userPool)

	client.InitClients()
	client.FlushSyncClients()
	log.Println("sync clients:", client.SyncClients)
	// client.MultiDial()
	// client.SingleDial()
	// for _, client := range client.SyncClients {
	// 	log.Println("client:", client)
	// }
}
