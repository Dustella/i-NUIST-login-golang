package config

import "os"

func GetEnv() NUIST_Auth {
	var username string
	var isp string
	var password string

	if os.Getenv("INUIST_USERNAME") != "" {
		username = os.Getenv("INUIST_USERNAME")
	}
	if os.Getenv("INUIST_PASSWORD") != "" {
		password = os.Getenv("INUIST_PASSWORD")
	}
	if os.Getenv("INUIST_ISP") != "" {
		isp = os.Getenv("INUIST_ISP")
	}

	res := NUIST_Auth{Username: username, Password: password, ISP: isp}
	return res
}

func GetCFEnv() Cloudflare_Config {
	var cfhostname string
	var cfzoneid string
	var cfapitoken string
	if os.Getenv("cfhostname") != "" {
		cfhostname = os.Getenv("cfhostname")
	}
	if os.Getenv("cfapitoken") != "" {
		cfapitoken = os.Getenv("cfapitoken")
	}
	if os.Getenv("cfzoneid") != "" {
		cfzoneid = os.Getenv("cfzoneid")
	}
	res := Cloudflare_Config{
		Cloudflare_API_Token: cfapitoken,
		Cloudflare_ZONEID:    cfzoneid,
		Cloudflare_Host:      cfhostname,
	}
	return res
}
