package packet

type DNSQueryPacket struct {
	ID        uint16
	QR        bool
	Opcode    uint8
	AA        bool
	TC        bool
	RD        bool
	RA        bool
	Z         uint8
	Rcode     uint8
	QDCount   uint16
	ANCount   uint16
	NSCount   uint16
	ARCount   uint16
	Questions []DNSPacketQuestion
}

type DNSPacketQuestion struct {
	Domain string
	Type   uint16
	Class  uint16
}
