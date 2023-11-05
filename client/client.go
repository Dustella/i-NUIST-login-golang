package client

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Status int

const (
	Online Status = iota
	Unauthorized
	Offline
	Unknown
)

type AuthClient struct {
	ApiIp         string
	InterfaceName string
	InterfaceIp   string
	Client        *resty.Client
	Status        Status
}

func NewAuthClient() AuthClient {
	// inter, err := net.InterfaceByName("eth0")
	// count := viper.GetInt("retry")
	// create an empty io that destorys input
	logger := logrus.New()
	logger.Out = io.Discard
	client := resty.New().
		SetBaseURL("http://10.255.255.46/api/v1").
		SetRetryCount(2).
		SetTimeout(10 * time.Second)
		// SetLogger(logger)

	return AuthClient{
		Client: client,
	}

}

func (c *AuthClient) SetInterface(intrf *net.Interface) {
	var ip net.IP
	ip_range, _ := intrf.Addrs()

	for _, addr := range ip_range {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP
				break
			}
		}
	}

	// randPort := 50000 + rand.Intn(10000)
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{
			IP: ip,
			// Port: randPort,
		},
		Timeout: 20 * time.Second,
	}

	c.Client.SetTransport(&http.Transport{
		DialContext: dialer.DialContext,
	})

	c.InterfaceIp = ip.String()
	c.InterfaceName = intrf.Name
}

func (c *AuthClient) MaybeInNuist() bool {
	ip := c.InterfaceIp
	// 10. or DHCP unallocated
	return (strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "169.254"))
}
