package models

type DNSQuery struct {
	RecordType RecordType
	Domain     Domain
}

func (dq *DNSQuery) String() string {
	return string(dq.RecordType) + "." + dq.Domain.String()
}
