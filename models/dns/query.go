package dns

import (
	"main/models"
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

type Query struct {
	RecordType RecordType
	Domain     models.Domain
}

func (dq *Query) String() string {
	return string(dq.RecordType) + "." + dq.Domain.String()
}

func buildPacketQuestions(query *Query) ([]PacketQuestion, error) {
	qtype, err := query.RecordType.ToQTYPE()
	if err != nil {
		return nil, err
	}
	question := PacketQuestion{
		Domain: query.Domain.Value,
		Type:   qtype,
		Class:  questionClass,
	}
	return []PacketQuestion{question}, nil
}

func buildPacketHeader() Packet {
	return Packet{
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

func (dq *Query) ToPacket() (Packet, error) {
	queryPacket := buildPacketHeader()
	questions, err := buildPacketQuestions(dq)
	if err != nil {
		return queryPacket, err
	}
	queryPacket.Questions = questions
	queryPacket.QDCount = uint16(len(queryPacket.Questions))
	return queryPacket, nil
}
