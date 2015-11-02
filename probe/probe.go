package probe

import (
	"github.com/google/gopacket/layers"
	"log"
	"net"
	"probe/probe/iface"
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
	checkError(err)

	err = openAllInterface(allActiveInterface)
	checkError(err)

	err = openSendThread()
	checkError(err)

	destIP := make([]*net.IPAddr, 0)
	dest, _ := net.ResolveIPAddr("ip", "114.114.114.114")
	destIP = append(destIP, dest)

	go nping(allActiveIP, destIP)

	log.Println("Network connection check finished\n")

	log.Println("Probe started")

	for {
		select {
		case icmp := <-chanICMP:
			onICMP(icmp)
		case synACK := <-chanSYNACK:
			onSYNACK(synACK)
		}
	}
}

func onICMP(data incomePacket) {
	pkt := data.packet
	icmpLayer := pkt.Layer(layers.LayerTypeICMPv4)
	if icmpLayer == nil {
		return
	}
	icmp, _ := icmpLayer.(*layers.ICMPv4)
	onRecvICMP(pkt, icmp, data.ci)
}

func onSYNACK(synACK incomePacket) {

}
