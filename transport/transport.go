package transport

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/youngxhui/block-rpc/transport/protocol"
)

// TCPTransport 是一个基于 TCP 的自定义传输层
type TCPTransport struct {
	conn net.Conn
}

// NewTCPTransport 创建一个新的 TCPTransport
func NewTCPTransport(conn net.Conn) *TCPTransport {
	return &TCPTransport{conn: conn}
}

// Send 发送数据
func (t *TCPTransport) Send(message *protocol.Message) error {
	data, err := message.Encode()
	if err != nil {
		return err
	}
	// 先发送数据长度
	length := uint32(len(data))
	if err := binary.Write(t.conn, binary.BigEndian, length); err != nil {
		return err
	}

	// 发送实际数据
	_, err = t.conn.Write(data)
	return err
}

// Receive 接收数据
func (t *TCPTransport) Receive() (*protocol.Message, error) {
	// 先读取数据长度
	var length uint32
	if err := binary.Read(t.conn, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	// 读取实际数据
	data := make([]byte, length)
	_, err := io.ReadFull(t.conn, data)
	if err != nil {
		return nil, err
	}

	var message protocol.Message
	if err := message.Decode(data); err != nil {
		return nil, err
	}
	return &message, nil
}

// Close 关闭连接
func (t *TCPTransport) Close() error {
	return t.conn.Close()
}
