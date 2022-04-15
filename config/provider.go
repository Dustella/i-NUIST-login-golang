package config

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
		return args
	}
	if env.Username != "" && env.Password != "" && env.ISP != "" {
		return env
	}
	return NUIST_Auth{}
}

func GetCFConfig() Cloudflare_Config {
	args := GetCFArgs()
	env := GetCFEnv()
	if args.Cloudflare_API_Token != "" && args.Cloudflare_Host != "" && args.Cloudflare_ZONEID != "" {
		return args
	}
	if env.Cloudflare_API_Token != "" && env.Cloudflare_Host != "" && env.Cloudflare_ZONEID != "" {
		return env
	}
	return Cloudflare_Config{}
}
