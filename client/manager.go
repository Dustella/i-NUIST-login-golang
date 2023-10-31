package client

import (
	"log"
	"main/config"
	"sync"
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
	log.Println(RawClients)

	var wg sync.WaitGroup
	for _, client := range RawClients {
		wg.Add(1)
		go func(client AuthClient) {
			defer wg.Done()
			resp, err := client.GetIp()
			log.Println(client.InterfaceName, resp)
			if err == nil {
				SyncClients = append(SyncClients, client)
			}
		}(client)
	}
	wg.Wait()
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
