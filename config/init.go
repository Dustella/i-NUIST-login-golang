package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// var userInfo = viper.New()
// var ddnsInfo = viper.New()
var userPool []UserRecord
var usedUser []UserRecord

var mode string

func GetMode() string {
	return mode
}

func GetUserPool() []UserRecord {
	return userPool
}

func InitConfig() {

	pflag.Int("retry", 5, "Retry times")
	pflag.Int("timeout", 30, "Timeout seconds")
	pflag.Bool("syncdial", false, "Force multi dial")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	// userInfo.BindEnv()

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
		records := parseUserRecords(userstring)
		userPool = append(userPool, records...)
	}

}
