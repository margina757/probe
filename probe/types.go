package probe

import (
	"net"
	"time"
)

const (
	timedOut  = 1000000 // 超时限制：1000000微秒（1秒）
	typeICMP  = 1
	typeTCP   = 2
	typeUDP   = 3
	typeTrace = 4
)

type delays struct {
	delay []int
	stamp time.Time
}

type prober interface {
	probe() (*delays, error)
}

type destAddr struct {
	addr     net.IPAddr
	interval int
	tcpPort  int
	udpPort  int
	lastHops net.IPAddr
}

type srcAddr struct {
	addr net.IPAddr
}
