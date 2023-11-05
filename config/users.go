package config

import (
	"log"
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

func ParseUserString(str string) (UserRecord, error) {
	// str format: username:password@isp
	log.Println(str)
	if strings.Count(str, ":") != 1 || strings.Count(str, "@") != 1 {
		return UserRecord{}, nil
	}
	userpass := strings.Split(str, "@")[0]
	isp := strings.Split(str, "@")[1]
	username := strings.Split(userpass, ":")[0]
	password := strings.Split(userpass, ":")[1]
	if len(username) == 0 || len(password) == 0 || len(isp) == 0 {
		return UserRecord{}, nil
	}
	return UserRecord{username, password, isp}, nil
}

func ParseUserRecords(str string) []UserRecord {
	// str format: username:password@isp,username:password@isp
	formatted := strings.ReplaceAll(str, "\n", "")
	strs := strings.Split(formatted, ",")
	var records []UserRecord
	for _, line := range strs {
		line := strings.Trim(line, " ")
		record, err := ParseUserString(line)
		if err != nil {
			continue
		}
		records = append(records, record)
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
