package flow

import (
	"bytes"
	"context"
	"github.com/h2p2f/ntflw-probe/internal/probe/model"
	"go.uber.org/zap"
	"strconv"
)

func ParsePacket(ctx context.Context,
	templateMap *TemplateMap,
	bufRawFlowChan chan []byte,
	bufParsingResultChan chan model.Packet,
	log *zap.Logger) {
	for {
		select {
		case <-ctx.Done():
			return
		case buf := <-bufRawFlowChan:
			netFlowPacket := NewPacket()
			err := netFlowPacket.Read(bytes.NewBuffer(buf), templateMap)
			if err != nil {
				log.Error("Error parsing packet", zap.Error(err))
				continue
			}
			for _, payload := range netFlowPacket.Payloads {
				packet := model.Packet{}
				packet.SourceID = strconv.Itoa(int(payload.SourceID))
				for _, field := range payload.Fields {
					packet.SetField(int(field.FieldType), field.FieldValue)
				}
				bufParsingResultChan <- packet
			}
		}
	}

}
