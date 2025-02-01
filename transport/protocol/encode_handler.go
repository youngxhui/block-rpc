package protocol

import "github.com/youngxhui/block-rpc/transport/buffer"

type MessageEncodeHandler struct {
	msgType MessageType
}

func NewMessageEncodeHandler(msgType MessageType) *MessageEncodeHandler {
	return &MessageEncodeHandler{msgType: msgType}
}

func (h *MessageEncodeHandler) Process(buffer *buffer.ByteBuffer) error {
	// 读取处理后的 Payload
	payload, _ := buffer.Read(buffer.Len())

	// 封装为 Message
	msg := &Message{
		Type:    h.msgType,
		Payload: payload,
	}

	// 编码为字节流
	data, err := msg.Encode()
	if err != nil {
		return err
	}

	// 清空 buffer 并写入完整的报文
	buffer.Reset()
	buffer.Write(data)
	return nil
}
