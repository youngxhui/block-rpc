package protocol

import "github.com/youngxhui/block-rpc/transport/buffer"

type MessageDecodeHandler struct{}

func NewMessageDecodeHandler() *MessageDecodeHandler {
	return &MessageDecodeHandler{}
}

// Process 自定义 Message 处理
func (h *MessageDecodeHandler) Process(buffer *buffer.ByteBuffer) error {
	data, _ := buffer.Read(buffer.Len())
	var msg Message
	if err := msg.Decode(data); err != nil {
		return err
	}

	buffer.Reset()
	buffer.Write(msg.Payload)
	return nil
}
