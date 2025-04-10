package record_type

import "fmt"

type RecordType string

const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	CNAME RecordType = "CNAME"
)

func (p *RecordType) Set(s string) error {
	switch s {
	case "A", "AAAA", "CNAME":
		*p = RecordType(s)
		return nil
	default:
		return fmt.Errorf("invalid RecordType: %s", s)
	}
}

func (p *RecordType) String() string {
	return string(*p)
}

func (p *RecordType) Type() string {
	return "RecordType"
}

func (p *RecordType) ToQTYPE() (uint16, error) {
	value, exists := QTYPEMap[*p]
	if exists {
		return value, nil
	}
	return 0, fmt.Errorf("invalid QType %s", *p)
}
