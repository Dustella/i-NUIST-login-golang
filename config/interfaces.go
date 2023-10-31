package config

import "net"

func ListNetworkInterface() []net.Interface {
	infs, _ := net.Interfaces()
	return infs
}
