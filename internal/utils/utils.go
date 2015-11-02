package utils

import (
	"fmt"
	"net"
)

const (
	a0 = 1000000000
	a1 = 1000000
	a2 = 1000
	a3 = 1
)

func IPToIntString(ip []byte) (v int64) {
	if len(ip) != 4 {
		return 0
	}
	return int64(ip[3]) + int64(ip[2])*a2 + int64(ip[1])*a1 + int64(ip[0])*a0
}

func IntStringToIP(v int64) []byte {
	addr0 := v / a0
	addr1 := v/a1 - addr0*a0
	addr2 := v/a2 - addr0*a1 - addr1*a2
	addr3 := v - addr0*a0 - addr1*a1 - addr2*a2
	return []byte{byte(addr0), byte(addr1), byte(addr2), byte(addr3)}
}

func IntStringToDotString(v int64) string {
	addr0 := v / a0
	addr1 := v/a1 - addr0*a0
	addr2 := v/a2 - addr0*a1 - addr1*a2
	addr3 := v - addr0*a0 - addr1*a1 - addr2*a2
	return fmt.Sprint("%d.%d.%d.%d", addr0, addr1, addr2, addr3)
}

func IPToDotString(ip []byte) string {
	if len(ip) != 4 {
		return "0.0.0.0"
	}
	return fmt.Sprint("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func DotStringToIntString(ip string) int64 {
	addr, e := net.ResolveIPAddr("ip", ip)
	if e != nil {
		return 0
	}
	return IPToIntString(addr.IP.To4())
}
