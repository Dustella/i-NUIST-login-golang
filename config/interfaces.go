package config

import "net"

func ListNetworkInterface() []net.Interface {
	infs, _ := net.Interfaces()
	// filter loopback interface
	for i := 0; i < len(infs); i++ {
		if infs[i].Flags&net.FlagLoopback != 0 {
			infs = append(infs[:i], infs[i+1:]...)
			i--
		}
	}
	return infs
}
