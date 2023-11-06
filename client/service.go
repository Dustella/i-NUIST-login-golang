package client

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

func StartDaemon() {
	timeout := viper.GetInt("timeout")
	ticker := time.NewTicker(time.Duration(timeout) * time.Second)
	for range ticker.C {
		Dial()
	}
}

func Dial() {
	go FlushClients()
	if len(SyncClients) <= 1 {
		log.Println("No sync client, dialing single client")
		PrimeClient.GetIp()
		SingleDial()
		PrintPrimeClient()
	} else {
		log.Println("Detected Multiple Clients, Dialing multi clients")
		MultiDial()
		PrintSyncClients()
	}
}
