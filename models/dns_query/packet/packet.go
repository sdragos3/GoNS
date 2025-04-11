package packet

import (
	"bytes"
	"main/utils"
)

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

func (q *DNSQueryPacket) Encode() []byte {
	var buffer bytes.Buffer

	utils.WriteOrPanic(&buffer, q.ID)

	b2i := func(b bool) int {
		if b {
			return 1
		}

		return 0
	}

	queryParams1 := byte(b2i(q.QR)<<7 | int(q.Opcode)<<3 | b2i(q.AA)<<1 | b2i(q.RD))
	queryParams2 := byte(b2i(q.RA)<<7 | int(q.Z)<<4)

	utils.WriteOrPanic(&buffer, queryParams1)
	utils.WriteOrPanic(&buffer, queryParams2)
	utils.WriteOrPanic(&buffer, q.QDCount)
	utils.WriteOrPanic(&buffer, q.ANCount)
	utils.WriteOrPanic(&buffer, q.NSCount)
	utils.WriteOrPanic(&buffer, q.ARCount)

	for _, question := range q.Questions {
		buffer.Write(question.Encode())
	}

	return buffer.Bytes()
}
