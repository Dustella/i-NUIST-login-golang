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
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}

func ReadArgs() {

	args := pflag.Args()

	if len(args) == 0 {
		log.Fatal("No args")
	}

	var userstring string

	if len(args) >= 1 {
		mode = args[0]
	}

	if len(args) == 2 {
		userstring = args[1]
	}

	if len(userstring) > 0 {
		records := ParseUserRecords(userstring)
		userPool = append(userPool, records...)
	}

	ReadRecordFile()
}
