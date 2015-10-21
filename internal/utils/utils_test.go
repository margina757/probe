package utils

import (
	"testing"
)

func TestIPtoIntString(t *testing.T) {
	ip := []byte{192, 168, 0, 1}
	ipInt := IPToIntString(ip)
	t.Log(ipInt)
}

func TestIntStringToIP(t *testing.T) {
	var ipString int64
	ipString = 192168000001
	ip := IntStringToIP(ipString)
	t.Log(ip)
}
