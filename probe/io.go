package probe

import (
	"fmt"
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
	ip   net.IP
}

func openSendConn() error {

	pid = uint16(os.Getegid() & 0xffff)

	s, e := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if e != nil {
		return e
	}
	fd = s
	e = syscall.SetsockoptInt(s, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if e != nil {
		return e
	}

	f := os.NewFile(uintptr(s), "xxxx")
	sendf = f

	packetSend, e = net.FilePacketConn(f)
	packetSend.LocalAddr()

	chanSend = make(chan *socketData, 1024)
	go doSend(s)
	return nil

}

func doSend(conn int) {
	for data := range chanSend {

		if len(data.ip) < 4 {
			continue
		}
		addr := syscall.SockaddrInet4{}
		addr.Port = 0
		// addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = data.ip[0], data.ip[1], data.ip[2], data.ip[3]

		addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = 192, 168, 199, 138
		err := syscall.Sendto(conn, data.data, 0, &addr)

		if err != nil {
			fmt.Println("send err:", err, data.data, addr.Addr)
			log.Fatal(err)
			return
		}
		fmt.Println("send")
	}
}

func sendbyte(p []byte) error {
	to := syscall.SockaddrInet4{}
	to.Port = 0
	to.Addr[0] = 8
	to.Addr[1] = 8
	to.Addr[2] = 8
	to.Addr[3] = 8
	_, e := syscall.Write(fd, p)
	return e

}

func sendto(data []byte) {
	addr := syscall.SockaddrInet4{}
	addr.Port = 0
	addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = 8, 8, 8, 8
	e := syscall.Sendto(fd, data, 0, &addr)
	if e != nil {
		fmt.Println(e)
	}

}
