package client

import (
	"log"
	"main/config"
)

var SyncClients []AuthClient
var RawClients []AuthClient
var PrimeClient *AuthClient

func InitClients() {
	interfaces := config.ListNetworkInterface()

	RawClients = make([]AuthClient, 0)
	SyncClients = make([]AuthClient, 0)

	c, _ := NewAuthClient()
	PrimeClient = c
	log.Println("init clients")

	for _, inf := range interfaces {
		client, err := NewAuthClient()
		if err != nil {
			continue
		}
		client.SetInterface(&inf)
		log.Println("client:", client)
		RawClients = append(RawClients, *client)
	}

	// if (len(SyncClients)) == 0 {
	// 	// not in nuist

	// }

	// if (len(SyncClients)) == 1 {
	// 	// disable syncdial, use single dial
	// }
	// if len(SyncClients) > 1 {
	// 	// syncdial
	// }
}

func PrepareSyncClients() {
	for _, client := range RawClients {
		go func(client AuthClient) {
			_, err := client.GetIp()
			if err == nil {
				SyncClients = append(SyncClients, client)
			}
		}(client)
	}
}

func MultiDial() {
	for _, client := range SyncClients {
		go func(client AuthClient) {
			res := client.TestConnection()
			if res != Online {
				client.DoAuth()
			}
		}(client)
	}
}

func SingleDial() {
	res := PrimeClient.TestConnection()
	if res != Online {
		PrimeClient.DoAuth()
	}

}
