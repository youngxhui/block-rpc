package protocol

import (
	"encoding/binary"
	"errors"
)

const (
	MagicNumber = 0xb10cc // BLOCK
	Version     = 0x01
)

type MessageType byte

const (
	TypeRequest  MessageType = 0x01
	TypeResponse MessageType = 0x02
)

// Message 自定义数据类型
//
//	字段  长度（字节） 描述
//	魔数（Magic）| 4 |用于标识协议（如 0xB10CC）
//	版本（Version）| 1 | 协议版本（如 0x01）
//	消息类型（Type）| 1 | 请求（0x01）或响应（0x02）
//	数据长度（Length）| 4 | 数据部分的长度
//	数据（Payload）| 可变 | 实际的数据内容
//	校验和（Checksum）| 4 |用于校验数据的完整性
type Message struct {
	Magic    uint32
	Version  byte
	Type     MessageType
	Length   uint32
	Payload  []byte
	Checksum uint32
}

// Encode 消息编码
func (m *Message) Encode() ([]byte, error) {
	buf := make([]byte, 14+len(m.Payload))
	binary.BigEndian.PutUint32(buf[0:4], MagicNumber)
	buf[4] = Version
	buf[5] = byte(m.Type)
	binary.BigEndian.PutUint32(buf[6:10], uint32(len(m.Payload)))
	copy(buf[10:10+len(m.Payload)], m.Payload)
	m.Checksum = calculateChecksum(m.Payload)
	binary.BigEndian.PutUint32(buf[10+len(m.Payload):14+len(m.Payload)], m.Checksum)
	return buf, nil
}

func (m *Message) Decode(data []byte) error {
	if len(data) < 14 {
		return errors.New("invalid message length")
	}
	m.Magic = binary.BigEndian.Uint32(data[0:4])
	if m.Magic != MagicNumber {
		return errors.New("invalid magic number")
	}
	m.Version = data[4]
	m.Type = MessageType(data[5])
	m.Length = binary.BigEndian.Uint32(data[6:10])
	m.Payload = make([]byte, m.Length)
	copy(m.Payload, data[10:10+m.Length])
	m.Checksum = binary.BigEndian.Uint32(data[10+m.Length : 14+m.Length])
	if m.Checksum != calculateChecksum(m.Payload) {
		return errors.New("checksum mismatch")
	}
	return nil
}

func calculateChecksum(data []byte) uint32 {
	var sum uint32
	for _, b := range data {
		sum += uint32(b)
	}
	return sum
}
