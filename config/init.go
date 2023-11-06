package config

// var userInfo = viper.New()
// var ddnsInfo = viper.New()
var userPool []UserRecord

var mode string

func GetMode() string {
	return mode
}

func GetUserPool() []UserRecord {
	return userPool
}

func InitConfig() {

	BindFlags()
	ReadArgs()
	ReadRecordFile()

}
