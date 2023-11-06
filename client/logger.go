package client

import (
	"log"
	"main/config"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

func PrintInfo() {
	mode := config.GetMode()
	userPool := config.GetUserPool()
	log.Println("retry:", viper.GetInt("retry"))
	log.Println("timeout:", viper.GetInt("timeout"))
	log.Println("mode:", mode)
	log.Println("userPool:", userPool)

	PrimeClient = NewAuthClient()
	FlushRawClients()
	FlushSyncClients()
	log.Println("sync clients:", SyncClients)
}

func PrintSyncClients() {
	data := [][]string{}

	log.Printf("Sync Clients: %d\n", len(SyncClients))
	for _, c := range SyncClients {
		data = append(data, []string{c.InterfaceName, c.InterfaceIp, c.ApiIp, c.Status.String()})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Interface", "Interface IP", "API IP", "Status"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func PrintRawClients() {
	data := [][]string{}

	log.Printf("Raw Clients: %d\n", len(RawClients))
	for _, c := range RawClients {
		data = append(data, []string{c.InterfaceName, c.InterfaceIp, c.ApiIp, c.Status.String()})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Interface", "Interface IP", "API IP", "Status"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func PrintPrimeClient() {
	log.Println("IP: ", PrimeClient.ApiIp, ", Status: ", PrimeClient.Status.String())
}
