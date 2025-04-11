package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func WriteOrPanic(buf *bytes.Buffer, data interface{}) {
	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		log.Fatalf("binary.Write failed for %v: %v", data, err)
	}
}
