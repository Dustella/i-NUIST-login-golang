package auth

import (
	"fmt"
	"main/config"
	"time"
)

func DoLogin(config config.NUIST_Auth, attempt int) (string, error) {
	var myip string
	for i := 0; i < attempt; i++ {
		ip, err := GetIP()
		if err != nil {
			fmt.Printf("Get IP err, are you in I-NUIST?\n")
			time.Sleep(time.Second * 1)
			fmt.Printf("Making another getIP request, try time: %v\n", i)
			if i == attempt-1 {
				return "", err
			}
		} else {
			myip = ip
			fmt.Printf("Your IP is %v\n", ip)
			break
		}
	}

	for i := 0; i < attempt; i++ {
		success, err := DoAuth(config.Username, config.Password, config.ISP, myip)
		if success {
			fmt.Println("Login sucess")
			break
		}
		if err != nil {
			fmt.Printf("auth err, message: %v", err)
			time.Sleep(time.Second * 1)
			fmt.Printf("Making another request, try time: %v\n", i)
			if i == attempt-1 {
				return "", err
			}
		}
	}
	return myip, nil
}
