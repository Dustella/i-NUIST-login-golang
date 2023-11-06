package config

import "net"

func ListNetworkInterface() []net.Interface {
	infs, _ := net.Interfaces()
	// filter loopback interface, and those without valid ipv4 address
	for i := 0; i < len(infs); i++ {

		flag := infs[i].Flags&net.FlagLoopback != 0            // loopback
		flag = flag || infs[i].Flags&net.FlagUp == 0           // down
		flag = flag || infs[i].Flags&net.FlagPointToPoint != 0 // point to point
		flag = flag || infs[i].Flags&net.FlagMulticast == 0    // not multicast
		// has no ipv4 address
		addrs, _ := infs[i].Addrs()
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					flag = false
					break
				}
			}
		}
		if flag {
			infs = append(infs[:i], infs[i+1:]...)
			i--
		}
	}
	return infs
}
