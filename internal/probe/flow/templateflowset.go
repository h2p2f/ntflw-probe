package flow

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Template struct {
	TemplateID uint16
	FieldCount uint16
	Fields     []Field
}

func (t *Template) Read(b *bytes.Buffer) error {
	err := binary.Read(b, binary.BigEndian, &t.TemplateID)
	if err != nil {
		fmt.Println("Template.Read: ", err)
		return err
	}
	fmt.Println("Template ID: ", t.TemplateID)
	err = binary.Read(b, binary.BigEndian, &t.FieldCount)
	if err != nil {
		fmt.Println("Template.Read: ", err)
		return err
	}
	fmt.Println("Template FieldCount: ", t.FieldCount)
	for i := 0; i < int(t.FieldCount); i++ {
		field := Field{}
		err = field.ReadField(b)
		if err != nil {
			return err
		}
		t.Fields = append(t.Fields, field)
	}
	return nil
}

func (t *Template) FieldsSize() uint16 {
	var n uint16
	for _, field := range t.Fields {
		n += field.FieldLength
	}
	return n
}

type Field struct {
	FieldType   uint16
	FieldLength uint16
}

func (f *Field) ReadField(b *bytes.Buffer) error {
	err := binary.Read(b, binary.BigEndian, &f.FieldType)
	if err != nil {
		fmt.Println("Field.Read: ", err)
		return err
	}
	err = binary.Read(b, binary.BigEndian, &f.FieldLength)
	if err != nil {
		return err
	}
	return nil
}
