package dns

var QTYPEMap = map[RecordType]uint16{
	A:     1,
	CNAME: 5,
	AAAA:  28,
}
