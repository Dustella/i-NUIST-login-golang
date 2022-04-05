package config

import (
	"flag"
)

func GetArgs() (string, string, string) {
	var username string
	var isp string
	var password string
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password")
	flag.StringVar(&isp, "isp", "", "isp")
	flag.Parse()
	return username, password, isp
}
