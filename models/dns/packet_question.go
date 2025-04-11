package dns

import (
	"bytes"
	"encoding/binary"
	"log"
	"main/utils"
	"strings"
)

type PacketQuestion struct {
	Domain string
	Type   uint16
	Class  uint16
}

func (q *PacketQuestion) Encode() []byte {
	var buffer bytes.Buffer

	domainParts := strings.Split(q.Domain, ".")
	for _, part := range domainParts {
		if err := binary.Write(&buffer, binary.BigEndian, byte(len(part))); err != nil {
			log.Fatalf("Error binary.Write(..) for '%s': '%s'", part, err)
		}

		for _, c := range part {
			if err := binary.Write(&buffer, binary.BigEndian, uint8(c)); err != nil {
				log.Fatalf("Error binary.Write(..) for '%s'; '%c': '%s'", part, c, err)
			}
		}
	}

	utils.WriteOrPanic(&buffer, uint8(0))
	utils.WriteOrPanic(&buffer, q.Type)
	utils.WriteOrPanic(&buffer, q.Class)

	return buffer.Bytes()
}
