package models

type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

func (p *Protocol) String() string {
	return string(*p)
}
