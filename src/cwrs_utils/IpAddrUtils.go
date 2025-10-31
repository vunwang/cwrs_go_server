package cwrs_utils

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"fmt"
	"net"
	"strings"
)

// 获取本地服务器IP地址
func GetLocalIpAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Errorf("IP Error: %v", err)
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// 获取本地局域网IP地址
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	// ip地址 局域网段
	ipAddr := cwrs_viper.GlobalViper.GetString("ip.addr")

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				// 根据你的局域网段过滤，比如你的是 192.168.10.x
				if strings.HasPrefix(ip, ipAddr) {
					return ip
				}
			}
		}
	}
	return "127.0.0.1" // fallback
}
