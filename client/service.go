package client

import (
	"time"

	"github.com/spf13/viper"
)

func StartDaemon() {
	PrimeClient = NewAuthClient()

	timeout := viper.GetInt("timeout")
	ticker := time.NewTicker(time.Duration(timeout) * time.Second)
	for range ticker.C {
		FlushRawClients()
		FlushSyncClients()
		MultiDial()
		SingleDial()
	}
}
