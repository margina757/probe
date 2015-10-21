package probe

import (
	"log"
	"net"
	"os"
	"syscall"
)

var packetSend net.PacketConn
var sendf *os.File
var fd int

type socketData struct {
	data []byte
	addr *net.IPAddr
}

func openSendThread() error {
	pid = uint16(os.Getpid() & 0xffff)

	s, e := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if e != nil {
		return e
	}
	fd = s
	e = syscall.SetsockoptInt(s, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if e != nil {
		return e
	}

	chanSend = make(chan *socketData, 1024)
	go doSendThread(s)
	return nil
}

func doSendThread(conn int) {

	for data := range chanSend {

		dstIP := data.addr.IP.To4()
		if dstIP == nil {
			log.Println("send fail: no dest ip")
			continue
		}

		addr := syscall.SockaddrInet4{}
		addr.Port = 0
		addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = dstIP[0], dstIP[1], dstIP[2], dstIP[3]

		err := syscall.Sendto(conn, data.data, 0, &addr)

		if err != nil {
			log.Println("Send Thread ", err, data.addr.String())
			continue
		}
	}
}
