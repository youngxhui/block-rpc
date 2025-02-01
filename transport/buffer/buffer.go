package buffer

import "bytes"

// ByteBuffer 提供高效的字节操作
type ByteBuffer struct {
	buf *bytes.Buffer
}

func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{buf: bytes.NewBuffer(nil)}
}

func (b *ByteBuffer) Write(data []byte) (int, error) {
	return b.buf.Write(data)
}

func (b *ByteBuffer) Read(n int) ([]byte, error) {
	data := make([]byte, n)
	_, err := b.buf.Read(data)
	return data, err
}

func (b *ByteBuffer) Reset() {
	b.buf.Reset()
}

func (b *ByteBuffer) Len() int {
	return b.buf.Len()
}

func (b *ByteBuffer) ReadAll() []byte {
	data := b.buf.Bytes()
	b.Reset()
	return data
}
