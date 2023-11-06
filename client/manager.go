package client

import (
	"main/config"
	"strings"
	"sync"
)

var SyncClients []AuthClient
var RawClients []AuthClient
var PrimeClient AuthClient

func FlushClients() {
	PrimeClient = NewAuthClient()

	FlushRawClients()
	FlushSyncClients()
}

func FlushRawClients() {
	interfaces := config.ListNetworkInterface()
	RawClients = make([]AuthClient, 0)
	for _, inf := range interfaces {
		client := NewAuthClient()
		client.SetInterface(inf)
		client.TestConnection()
		RawClients = append(RawClients, client)
	}
	PrintRawClients()
}

func FlushSyncClients() {
	SyncClients = make([]AuthClient, 0)

	var wg sync.WaitGroup
	for _, client := range RawClients {
		wg.Add(1)
		go func(client AuthClient) {
			defer wg.Done()
			_, err := client.GetIp()
			strings.EqualFold(client.ApiIp, client.InterfaceIp)
			if err == nil {
				SyncClients = append(SyncClients, client)
			}
		}(client)
	}
	wg.Wait()
	PrintSyncClients()
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
		PrimeClient.GetIp()
		PrimeClient.DoAuth()
	}
}
