package config

import (
	"flag"
)

func GetArgs() NUIST_Auth {
	var username string
	var isp string
	var password string
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password")
	flag.StringVar(&isp, "isp", "", "isp")
	flag.Parse()
	res := NUIST_Auth{Username: username, Password: password, ISP: isp}
	return res
}
