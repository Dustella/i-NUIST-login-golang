package config

import "fmt"

type NUIST_Auth struct {
	Username string
	Password string
	ISP      string
}

type Cloudflare_Config struct {
	Cloudflare_API_Token string
	Cloudflare_ZONEID    string
	Cloudflare_Host      string
}

type DNSPod_Config struct {
	// DNSPod_
}

type AliDNS_Config struct {
	// AliDNS
}

func GetConfig() NUIST_Auth {
	args := GetArgs()
	env := GetEnv()
	if args.Username != "" && args.Password != "" && args.ISP != "" {
		fmt.Println("Readed config from args, ")
		fmt.Printf("username is %s, password is %s, isp is %s\n", args.Username, args.Password, args.ISP)
		return args
	}
	if env.Username != "" && env.Password != "" && env.ISP != "" {
		fmt.Println("Readed config from env, ")
		fmt.Printf("username is %s, password is %s, isp is %s\n", env.Username, env.Password, env.ISP)
		return env
	}
	final := proNuist()
	return final
}

func GetCFConfig() Cloudflare_Config {
	args := GetCFArgs()
	env := GetCFEnv()
	if args.Cloudflare_API_Token != "" && args.Cloudflare_Host != "" && args.Cloudflare_ZONEID != "" {
		fmt.Println("Readed Cloudflare DDNS config from args, attempting to Update DNS...")
		return args
	}
	if env.Cloudflare_API_Token != "" && env.Cloudflare_Host != "" && env.Cloudflare_ZONEID != "" {
		fmt.Println("Readed Cloudflare DDNS config from env, attempting to Update DNS...")
		return env
	}
	final := proCFDDNS()
	return final
}
