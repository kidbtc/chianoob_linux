package myfunc

import (
	"fmt"
	"net"
	"strings"
)

//获取内网ip

func LocalIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						if strings.Index(ipnet.IP.String(), "10.") != -1 || strings.Index(ipnet.IP.String(), "192.") != -1 || strings.Index(ipnet.IP.String(), "172.") != -1 {
							return ipnet.IP.String()
						}
					}
				}
			}
		}
	}
	return ""
}
