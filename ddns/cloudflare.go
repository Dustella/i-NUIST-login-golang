package ddns

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DoAuth(username string, password string, isp string, ip string) (bool, error) {
	channelMap := map[string]string{"I-NUIST": "1", "CMCC": "2", "ChinaNet": "3", "UNION": "4"}
	channel := channelMap[isp]
	payload := map[string]string{
		"channel":     channel,
		"ifautologin": "0",
		"ip":          ip,
		"pagesign":    "secondauth",
		"password":    password,
		"username":    username,
		"usripadd":    ip,
	}
	payloadBytes, _ := json.Marshal(payload)
	resp, _ := http.Post("http://10.255.255.34/api/v1/login", "application/json", bytes.NewReader(payloadBytes))
	body, _ := ioutil.ReadAll(resp.Body)
	// var res AuthInfo
	// json.Unmarshal(body, &res)
	// if res.Message == "ok" {
	// 	return true, nil
	// }
	// err := errors.New(res.Message)
	// return false, err
}
