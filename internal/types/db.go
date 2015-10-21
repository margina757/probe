package types

type DestHost struct {
	Src      string
	Dest     string
	TcpPort  int
	Interval int
}

type ProbeResult struct {
	Src   []byte
	Dest  []byte
	Delay int
	Stamp int64
	Type  int
}
