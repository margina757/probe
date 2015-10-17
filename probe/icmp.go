package probe

/*
 使用ICMP协议探测延迟
*/

type icmpPkt struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

type icmpProber struct {
}

func (icmp *icmpProber) probe() (d *delays, e error) {
	return
}

func (icmp *icmpProber) checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
