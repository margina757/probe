package probe

import (
	"log"
	"net"
	"probe/probe/iface"
	"time"
)

var (
	allActiveInterface []net.Interface
	allActiveIP        []*net.IPAddr
	allDisactiveIP     []*net.IPAddr

	chanSend chan *socketData
	pid      uint16
)

func Start() {
	var err error

	pcapInit()

	allActiveInterface, allActiveIP, allDisactiveIP, err = iface.AllInterface()
	if err != nil {
		log.Println(err)
		return
	}
	err = openAllInterface(allActiveInterface)
	if err != nil {
		log.Println(err)
		return
	}

	err = openSendConn()
	if err != nil {
		log.Println(err)
		return
	}

	err = ping()
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(1 * time.Second)
}
