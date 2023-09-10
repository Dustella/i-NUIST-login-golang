package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type RespAuthInfo struct {
	Code int `json:"code"`
	Data interface {
	} `json:"data"`
	Message string `json:"message"`
}

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
	resp, _ := http.Post("http://10.255.255.46/api/v1/login", "application/json", bytes.NewReader(payloadBytes))
	body, _ := ioutil.ReadAll(resp.Body)
	var res RespAuthInfo
	json.Unmarshal(body, &res)
	if res.Message == "ok" {
		return true, nil
	}
	err := errors.New(res.Message)
	return false, err
}
