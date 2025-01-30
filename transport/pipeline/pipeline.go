package pipeline

import (
	"github.com/youngxhui/block-rpc/transport/buffer"
	"github.com/youngxhui/block-rpc/transport/pipeline/handler"
)

type TransportPipeline struct {
	handlers []handler.TransHandler
}

func NewPipeline(handlers ...handler.TransHandler) *TransportPipeline {
	return &TransportPipeline{
		handlers: handlers,
	}
}

// Process pipeline 完成的处理过程
func (p *TransportPipeline) Process(buffer *buffer.ByteBuffer) error {
	for _, h := range p.handlers {
		if err := h.Process(buffer); err != nil {
			return err
		}
	}
	return nil
}
