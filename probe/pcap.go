package probe

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"io"
	"log"
	"sync"
)

const filter = "tcp[tcpflags] & (tcp-syn|tcp-ack) == (tcp-syn|tcp-ack) ||  icmp[icmptype] != 8"

var (
	ifaceHandles []ifaceHandle_t
	pcapOnce     sync.Once
	chanICMP     chan incomePacket
	chanSYNACK   chan incomePacket
)

type ifaceHandle_t struct {
	handle *pcap.Handle
	iface  string
}

type incomePacket struct {
	packet gopacket.Packet
	ci     *gopacket.CaptureInfo
}

type icmpPacket struct {
	data []byte
	ci   gopacket.CaptureInfo
}

type synackPacket struct {
	packet gopacket.Packet
	ci     gopacket.CaptureInfo
}

func pcapInit() {
	f := func() {
		chanICMP = make(chan incomePacket, 32)
		chanSYNACK = make(chan incomePacket, 32)
	}
	pcapOnce.Do(f)
}

func getIfaceHandle(iface string) (handle *pcap.Handle) {
	handle = nil
	if ifaceHandles == nil || len(ifaceHandles) < 0 {
		return
	}

	for _, h := range ifaceHandles {
		if h.iface == iface {
			return h.handle
		}
	}
	return nil
}
func openAllInterface(ifaces []string) (err error) {
	// for _, iface := range ifaces {
	// 	handle, err := pcap.OpenLive(iface, 128, false, pcap.BlockForever)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	go doReadPacket(handle)
	// }

	iff, _ := pcap.FindAllDevs()
	for _, ifa := range iff {
		log.Println(ifa.Name)
		handle, err := pcap.OpenLive(ifa.Name, 128, false, pcap.BlockForever)
		if err != nil {
			return err
		}
		go doReadPacket(handle)
	}

	return
}

func closeInterface() {
	if ifaceHandles == nil || len(ifaceHandles) <= 0 {
		return
	}
}

func openInterface() (err error) {
	return
}

func doReadPacket(handle *pcap.Handle) {
	var err error
	handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	for {
		data, ci, err := handle.ReadPacketData()

		switch {
		case err == io.EOF:
			return
		case err != nil:
			log.Println(err)
		default:
			packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
			incomePacket := incomePacket{packet: packet, ci: &ci}
			switchPacket(incomePacket)
		}
	}

}

func switchPacket(packet incomePacket) {
	switch {
	case packet.packet.Layer(layers.LayerTypeTCP) != nil:
		chanSYNACK <- packet
	case packet.packet.Layer(layers.LayerTypeICMPv4) != nil:
		chanICMP <- packet
	}
}
