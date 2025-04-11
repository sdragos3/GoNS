package query

import (
	dns_packet2 "main/models/dns_packet"
)

const queryIdentifier uint16 = 0xAAAA // query identifier
const recursionDesired bool = true
const questionClass uint16 = 0x0001
const packetType bool = false  // false for query, true for response
const queryType uint8 = 0x0000 // represents a standard query
const authoritativeAnswer bool = false
const recursionAvailable bool = false // mainly used for dns response packets
const futureUse uint8 = 0x0000        // reserved for future use
const responseCode uint8 = 0x0000     //  mainly used for dns response packets
const answerCount uint16 = 0x0000     // mainly used for dns response packets
const nameServerCount uint16 = 0x0000
const additionalRecordCount uint16 = 0x0000

func (dq *Query) String() string {
	return string(dq.RecordType) + "." + dq.Domain.String()
}

func buildPacketQuestions(query *Query) ([]dns_packet2.DNSPacketQuestion, error) {
	qtype, err := query.RecordType.ToQTYPE()
	if err != nil {
		return nil, err
	}
	question := dns_packet2.DNSPacketQuestion{
		Domain: query.Domain.Value,
		Type:   qtype,
		Class:  questionClass,
	}
	return []dns_packet2.DNSPacketQuestion{question}, nil
}

func buildPacketHeader() dns_packet2.DNSPacket {
	return dns_packet2.DNSPacket{
		ID:      queryIdentifier,
		RD:      recursionDesired,
		QR:      packetType,
		Opcode:  queryType,
		AA:      authoritativeAnswer,
		RA:      recursionAvailable,
		Z:       futureUse,
		Rcode:   responseCode,
		ANCount: answerCount,
		NSCount: nameServerCount,
		ARCount: additionalRecordCount,
	}
}

func (dq *Query) ToDNSPacket() (dns_packet2.DNSPacket, error) {
	queryPacket := buildPacketHeader()
	questions, err := buildPacketQuestions(dq)
	if err != nil {
		return queryPacket, err
	}
	queryPacket.Questions = questions
	queryPacket.QDCount = uint16(len(queryPacket.Questions))
	return queryPacket, nil
}
