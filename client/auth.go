package client

import (
	"errors"
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
	if resp.StatusCode() == 204 {
		c.Status = Online
	} else if resp.StatusCode() == 302 {
		c.Status = Unauthorized
	} else {
		c.Status = Offline
	}
	return c.Status

}

func (c *AuthClient) GetIp() (string, error) {
	data := IPInfo{}
	resp, err := c.Client.R().SetResult(data).Get("/ip")
	if err != nil || resp.IsError() {
		return "", err
	}
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
	data := RespAuthInfo{}
	_, err := c.Client.R().SetBody(payload).SetResult(data).Post("/login")
	if err != nil {
		return err
	}
	if data.Code == 200 {
		c.Status = Online
		return nil
	} else {
		c.Status = Unauthorized
		// return error
		return errors.New("auth failed")
	}
}
