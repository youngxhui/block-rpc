package handler

import (
	"github.com/youngxhui/block-rpc/transport/buffer"
	"github.com/youngxhui/block-rpc/transport/codec"
)

type EncodeHandler struct {
	codec codec.PayloadCodec
}

func NewEncodeHandler(codec codec.PayloadCodec) *EncodeHandler {
	return &EncodeHandler{codec: codec}
}

func (h *EncodeHandler) Process(buffer *buffer.ByteBuffer) error {
	// 从buffer读取原始数据并编码
	rawData, _ := buffer.Read(buffer.Len())
	encoded, err := h.codec.Encode(rawData)
	if err != nil {
		return err
	}
	buffer.Write(encoded)
	return nil
}
