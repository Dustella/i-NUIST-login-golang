package main

import (
	"main/config"
	"main/controller"
)

func main() {
	// fmt.Println("==================NUIST LOGIN GOLANG =================")
	// fmt.Println("Doing Auth......")
	// nuist := config.GetConfig()
	// attempt := 5
	// auth.DoLogin(nuist, attempt)
	// fmt.Println("==================LOGIN IS OK =================")
	// fmt.Println("Doing DDNS......")

	config.InitConfig()
	controller.Init()
}
