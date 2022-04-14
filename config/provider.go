package config

type NUIST_Auth struct {
	Username string
	Password string
	ISP      string
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
