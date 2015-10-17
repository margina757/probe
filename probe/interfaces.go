package probe

import (
	"log"
	"net"
)

/*
  检测本机网卡和IP的情况，并检测哪些IP可以连接到互联网。
*/

const (
	// 这两个IP都不能连通的本机IP视为没有互联网连接
	outIP1 = "114.114.114.114"
	outIP2 = "8.8.8.8"
)

func AllInterface() {
	ifaces, _ := allInterface()
	for _, iface := range ifaces {
		addrs := addrsOfOneInterface(iface)
		for _, addr := range addrs {
			log.Println(addr)
		}
	}
}

/* 所有网卡
 */
func allInterface() (ifaces []net.Interface, err error) {
	ifaces, err = net.Interfaces()
	return
}

// 一个接口上的所有ip4地址
func addrsOfOneInterface(iface net.Interface) (addrs []*net.IPAddr) {
	ifaddrs, err := iface.Addrs()

	if (err != nil) || (len(ifaddrs) == 0) {
		return
	}

	for _, ifaddr := range ifaddrs {
		var ip net.IP
		switch v := ifaddr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}
		if ip.IsLoopback() {
			return
		}
		ip = ip.To4()
		if ip != nil {
			addr, _ := net.ResolveIPAddr("ip", ip.String())
			addrs = append(addrs, addr)
		}
	}
	return
}

// 检查单独一个网卡的情况
func checkOneInterface(ifa net.Interface) {
	addrs, err := ifa.Addrs()
	if err != nil {
		log.Println(err)
		return
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:

			ip = v.IP
			log.Println("IPNet", ip)
		case *net.IPAddr:

			ip = v.IP
			// checkOneIP(ip)
			log.Println("IPAddr", ip)
		}
	}
}

// 检查IP地址是否能连接到互联网
func checkOneIP(ip *net.IPAddr) bool {
	return true
}
