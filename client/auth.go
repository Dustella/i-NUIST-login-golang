package client

import (
	"errors"
	"log"
	"main/config"
)

type IPInfo struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type RespAuthInfo struct {
	Code int `json:"code"`
	Data interface {
	} `json:"data"`
	Message string `json:"message"`
}

func (c *AuthClient) TestConnection() Status {
	// use miui generate 204
	resp, err := c.Client.R().Get("http://connect.rom.miui.com/generate_204")
	if err != nil {
		c.Status = Offline
		return c.Status
	}
	switch resp.StatusCode() {
	case 204:
		c.Status = Online
	case 302:
		c.Status = Unauthorized
	default:
		c.Status = Offline
	}
	return c.Status
}

func (c *AuthClient) GetIp() (string, error) {
	resp, err := c.Client.R().SetResult(&IPInfo{}).Get("/ip")
	if err != nil || resp.IsError() {
		return "", err
	}
	data := resp.Result().(*IPInfo)
	log.Printf("GetIp: %v", data)
	c.ApiIp = data.Data
	return data.Data, nil
}

func (c *AuthClient) DoAuth() error {
	c.GetIp()
	user := config.GetRamdomUser()
	payload := map[string]string{
		"channel":     user.GetChannelCode(),
		"ifautologin": "0",
		"ip":          c.ApiIp,
		"pagesign":    "secondauth",
		"password":    user.Password,
		"username":    user.Username,
		"usripadd":    c.ApiIp,
	}
	resp, err := c.Client.R().SetBody(payload).SetResult(&RespAuthInfo{}).Post("/login")
	if err != nil {
		return err
	}
	data := resp.Result().(*RespAuthInfo)
	if data.Code == 200 {
		c.Status = Online
		return nil
	} else {
		c.Status = Unauthorized
		// return error
		return errors.New("auth failed")
	}
}
