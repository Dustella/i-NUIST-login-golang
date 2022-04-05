package main

import (
	"fmt"
	"main/auth"
	"main/config"
	"time"
)

func main() {
	username, password, isp := config.GetArgs()
	attempt := 5
	myip := ""
	// ip := ""
	fmt.Printf("username is %s, password is %s, isp is %s\n", username, password, isp)
	for i := 0; i < attempt; i++ {
		ip, err := auth.GetIP()
		if err != nil {
			fmt.Printf("Get IP err, are you in I-NUIST?")
			time.Sleep(time.Second * 1)
			fmt.Printf("Making another getIP request, try time: %v\n", i)
		} else {
			myip = ip
			fmt.Printf("Your IP is %v\n", ip)
			break
		}
	}

	for i := 0; i < attempt; i++ {
		success, err := auth.DoAuth(username, password, isp, myip)
		if success {
			fmt.Println("Login sucess")
			break
		}
		if err != nil {
			fmt.Printf("auth err, message: %v", err)
			time.Sleep(time.Second * 1)
			fmt.Printf("Making another request, try time: %v\n", i)
		}
	}
}
