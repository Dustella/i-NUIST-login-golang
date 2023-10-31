package controller

import "main/config"

func Init() {
	mode := config.GetMode()
	if mode == "info" {
		printInfo()
	}
	if mode == "login" {

	}
}
