package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func GetIP() (string, error) {
	resp, err := http.Get("http://10.255.255.34/api/v1/ip")
	if err != nil {
		fmt.Println("Network err. Not in NUIST?")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Server Internal Error")
	}
	var res IPInfo
	json.Unmarshal(body, &res)
	return res.Data, err
}
