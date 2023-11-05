package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func ReadRecordFile() {
	// read record file ./pool.txt and append to userPool
	// format: username:password@isp

	filePath := viper.GetString("userpool")

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		if len(userPool) == 0 {
			panic("No user in userPool")
		}
	}
	users := ParseUserRecords(string(file))

	userPool = append(userPool, users...)
}
