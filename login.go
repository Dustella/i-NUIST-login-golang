package main

import (
	"main/config"
	"main/controller"
)

func main() {
	config.InitConfig()
	controller.BootStrap()
}
