package probe

import (
	"errors"
	"log"
	"net"
	"syscall"
)

/*
 读写数据包
*/

var (
	icmpSocket = 0
)

func socketAddr() (addr [4]byte, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) == net.IPv4len {
				copy(addr[:], ipnet.IP.To4())
				return
			}
		}
	}
	err = errors.New("You do not appear to be connected to the Internet")
	return
}

func CreateSocket() (e error) {
	recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return err
	}
	socketAddr, err := socketAddr()
	if err != nil {
		return
	}
	syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 55555, Addr: socketAddr})
	p := make([]byte, 1024)
	size, _, err := syscall.Recvfrom(recvSocket, p, 0)
	if err != nil {
		return err
	}

	log.Println(size)

	return
}
