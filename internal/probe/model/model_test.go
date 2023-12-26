package model

import "testing"

func TestPacket_SetField(t *testing.T) {
	type fields struct {
		SourceID string
		InBytes  string
		InPkts   string
	}
	type args struct {
		key   int
		value []uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				SourceID: "SourceID",
				InBytes:  "InBytes",
				InPkts:   "InPkts",
			},
			args: args{
				key:   1,
				value: []uint8{1, 2, 3},
			},
		},
		{
			name: "test2",
			fields: fields{
				SourceID: "SourceID",
				InBytes:  "InBytes",
				InPkts:   "InPkts",
			},
			args: args{
				key:   2,
				value: []uint8{1, 2, 3},
			},
		},
		{
			name: "test3",
			fields: fields{
				SourceID: "SourceID",
				InBytes:  "InBytes",
				InPkts:   "InPkts",
			},
			args: args{
				key:   3,
				value: []uint8{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Packet{
				SourceID: tt.fields.SourceID,
				InBytes:  tt.fields.InBytes,
				InPkts:   tt.fields.InPkts,
			}
			p.SetField(tt.args.key, tt.args.value)
		})
	}
}
