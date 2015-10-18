package iface

import (
	"fmt"
	"net"
)

const (
	// 这两个IP都不能连通的本机IP视为没有互联网连接
	outIP1  = "1.114.114.114"
	outIP2  = "2.8.8.8"
	timeout = 1000
)

var (
	allIface  []net.Interface
	allIPAddr []*net.IPAddr
)

func AllInterface() (ifaces []net.Interface, active []*net.IPAddr, disactive []*net.IPAddr, err error) {

	allifaces, err := allInterface()
	if err != nil {
		return ifaces, active, disactive, err
	}

	for _, iface := range allifaces {
		addThisInterface := false
		addrs := addrsOfOneInterface(iface)
		if len(addrs) <= 0 {
			continue
		}

		for _, addr := range addrs {

			alive := checkConnection(addr.String())
			if alive {
				addThisInterface = true
				allIPAddr = append(allIPAddr, addr)
			} else {
				disactive = append(disactive, addr)
			}

		}
		if addThisInterface == true {
			allIface = append(allIface, iface)
		}
	}
	return allIface, allIPAddr, disactive, nil
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

func checkConnection(laddr string) bool {

	fmt.Print("Testing ", laddr, ": ")

	ping1, e := Ping(laddr, outIP1, timeout)
	if ping1 {
		return true
	} else {
		if ne, ok := e.(net.Error); ok && ne.Timeout() {
		} else {
			fmt.Println("Error", e)
			return false
		}
	}

	ping2, e := Ping(laddr, outIP2, timeout)
	if ping2 {
		return true
	} else {
		if ne, ok := e.(net.Error); ok && ne.Timeout() {
			fmt.Println("Time Out")
			return false
		} else {
			fmt.Println("Error", e)
			return false
		}
	}

}
