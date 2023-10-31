package config

import (
	"math/rand"
	"strings"
	"time"
)

type UserRecord struct {
	Username string
	Password string
	ISP      string
}

func (u UserRecord) GetChannelCode() string {
	// channelMap := map[string]string{"I-NUIST": "1", "CMCC": "2", "ChinaNet": "3", "UNION": "4"}
	switch u.ISP {
	case "I-NUIST":
		return "1"
	case "CMCC":
		return "2"
	case "ChinaNet":
		return "3"
	case "UNION":
		return "4"
	default:
		return "0"
	}
}

func parseUserString(str string) UserRecord {
	// str format: username:password@isp
	userpass := strings.Split(str, "@")[0]
	isp := strings.Split(str, "@")[1]
	username := strings.Split(userpass, ":")[0]
	password := strings.Split(userpass, ":")[1]
	return UserRecord{username, password, isp}
}

func parseUserRecords(str string) []UserRecord {
	// str format: username:password@isp,username:password@isp
	strs := strings.Split(str, ",")
	var records []UserRecord
	for _, v := range strs {
		records = append(records, parseUserString(v))
	}
	return records
}

func GetRamdomUser() UserRecord {
	// random choose one user from userPool

	rand.Seed(time.Now().UnixNano())
	len := len(userPool)
	if len == 0 {
		// error
		panic("No user in userPool")
	}
	return userPool[rand.Intn(len)]
}
