package controller

import (
	"main/client"
	"main/config"
)

func BootStrap() {
	mode := config.GetMode()
	client.InitClients()

	switch mode {
	case "info":
		printInfo()
	case "login":
		client.SingleDial()
	case "daemon":
		client.StartDaemon()
	}
}
