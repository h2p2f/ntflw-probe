package flow

import (
	"bytes"
	"encoding/binary"
)

type Packet struct {
	Header   Header
	Payloads []PayLoads
}

func NewPacket() *Packet {
	return &Packet{}
}

func (p *Packet) Read(buf *bytes.Buffer, tMap *TemplateMap) error {

	err := p.Header.Read(buf)
	if err != nil {
		return err
	}

	count := int(p.Header.Count)

loop:
	for count > 0 {

		var (
			flowSetID uint16
			length    uint16
		)

		err = binary.Read(buf, binary.BigEndian, &flowSetID)
		if err != nil {
			return err
		}

		err = binary.Read(buf, binary.BigEndian, &length)
		if err != nil {
			return err
		}

		if length == 0 {
			continue
		}

		switch {

		case flowSetID == 0:

			template := Template{}
			err = template.Read(buf)
			if err != nil {
				return err
			}

			_, ok := tMap.Get(template.TemplateID)
			if !ok {
				tMap.Add(&template)
			}
			count--

		case flowSetID > 255:

			size := int(length - 4)

			templateFromMap, ok := tMap.Get(flowSetID)
			if !ok {
				break loop
			}

			payLoad := PayLoads{
				Fields: make([]PayLoadFields, len(templateFromMap.Fields)),
			}

			for size > 0 && count > 0 {
				err = payLoad.Read(buf, templateFromMap)
				if err != nil {
					return err
				}
				p.Payloads = append(p.Payloads, payLoad)
				size -= int(templateFromMap.FieldsSize())
				count--
			}

			buf.Reset()
		}
	}
	return nil
}
