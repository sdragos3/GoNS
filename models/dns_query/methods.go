package dns_query

import "main/models/dns_query/packet"

const headerId uint16 = 0xAAAA
const recursive bool = true
const questionClass uint16 = 0x0001

func (dq *DNSQuery) String() string {
	return string(dq.RecordType) + "." + dq.Domain.String()
}

func buildPacketHeader() packet.DNSPacketHeader {
	return packet.DNSPacketHeader{
		ID: headerId,
		RD: recursive,
	}
}

func buildPacketQuestions(query *DNSQuery) ([]packet.DNSPacketQuestion, error) {
	qtype, err := query.RecordType.ToQTYPE()
	if err != nil {
		return nil, err
	}
	question := packet.DNSPacketQuestion{
		Domain: query.Domain.Value,
		Type:   qtype,
		Class:  questionClass,
	}
	return []packet.DNSPacketQuestion{question}, nil
}

func (dq *DNSQuery) ToPacket() (packet.DNSQueryPacket, error) {
	header := buildPacketHeader()
	questions, err := buildPacketQuestions(dq)
	if err != nil {
		return packet.DNSQueryPacket{}, err
	}
	return packet.DNSQueryPacket{Header: header, Questions: questions}, nil
}
