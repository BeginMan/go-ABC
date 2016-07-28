package netcom

//协议集

//ICMP报文
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Indentifier uint16
	SequenceNum uint16
}
