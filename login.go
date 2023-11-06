package main

import (
	"main/client"
	"main/config"
)

func main() {
	BootStrap()
}

func BootStrap() {
	config.InitConfig()
	mode := config.GetMode()

	switch mode {
	case "info":
		client.PrintInfo()
	case "login":
		client.Dial()
	case "daemon":
		client.StartDaemon()
	}
}
