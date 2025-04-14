package main

import (
	"net"
	"strings"
)

func getInterfaceAddrs() (infAddrs []string, err error) {
	inf, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var virtualMACPrefixes = []string{
		"00:50:56", // VMware virtual adapters
		"00:15:5D", // Hyper-V virtual adapters
		"02:42",    // Docker virtual adapters (Linux-specific)
		"00:1A:4A", // Wi-Fi Direct adapters (common pattern)
	}
	for _, i := range inf {
		isV := false
		for _, pre := range virtualMACPrefixes {
			if strings.HasPrefix(i.HardwareAddr.String(), pre) {
				isV = true
				break
			}
		}
		if !isV && i.Flags&net.FlagUp == 1 && i.Flags&net.FlagLoopback == 0 {
			addrs, err := i.Addrs()
			if err != nil {
				return nil, err
			}
			for _, addr := range addrs {
				var ip net.IP

				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				// Filter IPv4 only
				if ip != nil && ip.To4() != nil {
					infAddrs = append(infAddrs, ip.String())
				}
			}
		}

	}
	return infAddrs, nil
}
