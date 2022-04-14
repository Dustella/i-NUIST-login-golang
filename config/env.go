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
