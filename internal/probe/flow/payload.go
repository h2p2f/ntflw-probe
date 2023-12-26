package flow

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PayLoads struct {
	Length   uint16
	SourceID uint32
	Fields   []PayLoadFields
}

type PayLoadFields struct {
	FieldType  uint16
	FieldValue []uint8
}

func (p *PayLoads) Read(buf *bytes.Buffer, template *Template) error {
	var err error

	p.Length = template.FieldsSize()

	for i, f := range template.Fields {
		p.Fields[i].FieldType = f.FieldType
		p.Fields[i].FieldValue = make([]uint8, 0)

		for j := 0; j < int(f.FieldLength); j++ {
			var b uint8
			err := binary.Read(buf, binary.BigEndian, &b)
			if err != nil {
				fmt.Println("cannot read flow")
			}
			p.Fields[i].FieldValue = append(p.Fields[i].FieldValue, b)
		}
	}
	return err
}
