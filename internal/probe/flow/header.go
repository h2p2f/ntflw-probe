package flow

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	Version        uint16
	Count          uint16
	SystemUptime   uint32
	UnixSeconds    uint32
	SequenceNumber uint32
	SourceID       uint32
}

func (h *Header) Read(b *bytes.Buffer) error {
	return binary.Read(b, binary.BigEndian, h)
}
