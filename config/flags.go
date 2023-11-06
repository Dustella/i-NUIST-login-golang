package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BindFlags() {
	pflag.Int("retry", 5, "Retry times")
	pflag.Int("timeout", 30, "Timeout seconds")
	pflag.Bool("syncdial", false, "Force multi dial")
	pflag.Bool("single", false, "Force single dial")
	pflag.String("userpool", "./pool.txt", "User record file")
	pflag.Bool("verbose", false, "Verbose mode")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func ReadArgs() {
	userstring := ""

	switch length := len(pflag.Args()); length {
	case 0:
		log.Fatal("No args")
	case 1:
		mode = pflag.Arg(0)
	case 2:
		mode = pflag.Arg(0)
		userstring = pflag.Arg(1)
	default:
		log.Fatal("Too many args")
	}

	if len(userstring) > 0 {
		userPool = append(userPool, ParseUserRecords(userstring)...)
	}

}
