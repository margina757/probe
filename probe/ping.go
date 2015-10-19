package probe

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
)

func ping() error {

	ip := &layers.IPv4{}
	ip.SrcIP = net.IP{192, 168, 1, 1}
	ip.DstIP = net.IP{8, 8, 8, 8}
	ip.Version = 4
	ip.Protocol = layers.IPProtocolTCP
	ip.Length = 20

	buf := gopacket.NewSerializeBuffer()

	opts := gopacket.SerializeOptions{}
	opts.ComputeChecksums = true
	// opts.FixLengths = true

	eth := &layers.Ethernet{}
	eth.Length = 14
	eth.SerializeTo(buf, opts)

	// gopacket.SerializableLayer(buf, opts, &layers.Ethernet{})
	gopacket.SerializeLayers(buf, opts, ip)
	// chanSend <- &socketData{buf.Bytes(), net.IP{8, 8, 8, 8}}

	pkt := gopacket.NewPacket(buf.Bytes(), layers.LayerTypeIPv4, gopacket.Default)
	iplayer := pkt.Layer(layers.LayerTypeIPv4)
	ipp, _ := iplayer.(*layers.IPv4)
	fmt.Println(ip)
	fmt.Printf("%+v\n", ipp)
	// addr, _ := net.ResolveIPAddr("ip4", "8.8.8.8")
	data := buf.Bytes()
	writebyte(data)

	return nil

}
