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

func GetCFArgs() Cloudflare_Config {
	var cfapitoken string
	var cfzoneid string
	var cfhostname string
	flag.StringVar(&cfapitoken, "cfapitoken", "", "cfapitoken")
	flag.StringVar(&cfzoneid, "cfzoneid", "", " cfzoneid")
	flag.StringVar(&cfhostname, "cfhostname", "", " cfhostname")
	flag.Parse()
	res := Cloudflare_Config{Cloudflare_API_Token: cfapitoken,
		Cloudflare_ZONEID: cfzoneid,
		Cloudflare_Host:   cfhostname}
	return res
}
