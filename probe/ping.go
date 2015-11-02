package probe

import (
	"encoding/binary"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
	"probe/db"
	"probe/internal/types"
	"sync"
	"time"
)

var (
	pingLock sync.Mutex
)

const (
	icmpSeqOffset = 0xff00
	icmpSeqPing   = 0x0
	icmpSeqTrace  = 0x0100
)

func nping(srcIP, destIP []*net.IPAddr) error {

	pingLock.Lock()
	defer pingLock.Unlock()

	for _, src := range srcIP {
		for _, dest := range destIP {
			for {
				data := newIcmpData(src, dest, layers.ICMPv4TypeEchoRequest, 0, 64)
				chanSend <- &socketData{data, dest}
				time.Sleep(1 * time.Second)
			}
		}
	}

	return nil
}

func newIcmpData(src, dest *net.IPAddr, typeCode, offSet, ttl int) (data []byte) {
	ip := &layers.IPv4{}
	ip.Version = 4
	ip.Protocol = layers.IPProtocolICMPv4
	ip.SrcIP = src.IP
	ip.DstIP = dest.IP
	ip.Length = 20
	ip.TTL = uint8(ttl)

	icmp := &layers.ICMPv4{}
	icmp.TypeCode = layers.ICMPv4TypeCode(uint16(typeCode) << 8)
	icmp.Id = pid
	icmp.Seq = 1
	icmp.Checksum = 0

	opts := gopacket.SerializeOptions{}
	opts.ComputeChecksums = true
	opts.FixLengths = true

	now := time.Now().UnixNano()
	var payload = make([]byte, 8)
	binary.LittleEndian.PutUint64(payload, uint64(now))

	buf := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buf, opts, ip, icmp, gopacket.Payload(payload))

	return buf.Bytes()
}

func onRecvICMP(pkt gopacket.Packet, icmp *layers.ICMPv4, ci *gopacket.CaptureInfo) {
	offset := icmp.Seq & icmpSeqOffset
	switch {
	case icmp.Id == pid && offset == 0:
		onRecvPing(pkt, icmp, ci)
	default:
	}
}

func onRecvPing(pkt gopacket.Packet, icmp *layers.ICMPv4, ci *gopacket.CaptureInfo) {
	payload := icmp.Payload
	if payload == nil || len(payload) <= 0 {
		return
	}
	sendStamp := binary.LittleEndian.Uint64(payload)
	if sendStamp < 1000000 {
		return
	}
	delay := ci.Timestamp.UnixNano() - int64(sendStamp)

	probeResult := &types.ProbeResult{}
	probeResult.Src = pkt.NetworkLayer().NetworkFlow().Dst().Raw()
	probeResult.Dest = pkt.NetworkLayer().NetworkFlow().Src().Raw()
	probeResult.Delay = int(delay / 1000)
	probeResult.Stamp = ci.Timestamp.Unix()
	probeResult.Type = types.ProbeTypePing
	db.InsertProbeResult(probeResult)
}
