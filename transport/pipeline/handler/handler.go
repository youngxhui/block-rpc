package handler

import "github.com/youngxhui/block-rpc/transport/buffer"

type TransHandler interface {
	Process(buffer *buffer.ByteBuffer) error
}
