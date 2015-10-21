package utils

import (
// "log"
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
