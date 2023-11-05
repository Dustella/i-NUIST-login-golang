package client

import (
	"log"
	"main/config"
	"strings"
	"sync"
)

var SyncClients []AuthClient
var RawClients []AuthClient
var PrimeClient AuthClient

func InitClients() {
	PrimeClient = NewAuthClient()

	FlushRawClients()
	FlushSyncClients()
}

func FlushRawClients() {
	interfaces := config.ListNetworkInterface()
	RawClients = make([]AuthClient, 0)
	for _, inf := range interfaces {
		client := NewAuthClient()

		client.SetInterface(&inf)
		RawClients = append(RawClients, client)
	}
	log.Println("========= Readed Raw Clients ===========")
	for _, c := range RawClients {
		log.Println("Raw Client:", c)
	}
	log.Println("========================================")
}

func FlushSyncClients() {
	SyncClients = make([]AuthClient, 0)

	var wg sync.WaitGroup
	for _, client := range RawClients {
		wg.Add(1)
		go func(client AuthClient) {
			defer wg.Done()
			_, err := client.GetIp()
			if err == nil && strings.EqualFold(client.ApiIp, client.InterfaceIp) {
				SyncClients = append(SyncClients, client)
			}
		}(client)
	}
	wg.Wait()
	log.Println("========= Readed Sync Clients ===========")
	for _, c := range SyncClients {
		log.Println("Sync Client:", c)
	}
	log.Println("========================================")
}

func MultiDial() {
	wg := sync.WaitGroup{}
	for _, client := range SyncClients {
		wg.Add(1)
		go func(client AuthClient) {
			defer wg.Done()
			res := client.TestConnection()
			if res == Unauthorized {
				client.DoAuth()
			}
		}(client)
	}
	wg.Wait()
}

func SingleDial() {
	res := PrimeClient.TestConnection()
	if res == Unauthorized {
		PrimeClient.DoAuth()
	}
}
