package packet

type DNSQueryPacket struct {
	Header    DNSPacketHeader
	Questions []DNSPacketQuestion
}

type DNSPacketHeader struct {
	ID           uint16
	QR           bool
	Opcode       uint8
	AA           bool
	TC           bool
	RD           bool
	RA           bool
	Z            uint8
	ResponseCode uint8
	QDCount      uint16
	ANCount      uint16
	NSCount      uint16
	ARCount      uint16
}

type DNSPacketQuestion struct {
	Domain string
	Type   uint16
	Class  uint16
}
