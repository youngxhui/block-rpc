package client

import (
	"context"
	"net"

	"github.com/youngxhui/block-rpc/remote"
	"github.com/youngxhui/block-rpc/transport"
	"github.com/youngxhui/block-rpc/transport/buffer"
	"github.com/youngxhui/block-rpc/transport/pipeline"
	"github.com/youngxhui/block-rpc/transport/pipeline/handler"
	"github.com/youngxhui/block-rpc/transport/protocol"
)

type Client interface {
	Call(ctx context.Context, req *remote.Request) (*remote.Response, error)
	Close() error
}

type BlockClient struct {
	transport *transport.TCPTransport
	pipeline  *pipeline.TransportPipeline
	config    *Config
}

func NewClient(opts ...Option) (Client, error) {
	config := defaultConfig()
	for _, opt := range opts {
		opt(config)
	}

	conn, err := net.Dial(config.Network, config.Address)
	if err != nil {
		return nil, err
	}
	p := pipeline.NewPipeline(
		protocol.NewMessageEncodeHandler(protocol.TypeRequest),
		handler.NewEncodeHandler(config.Codec),
	)
	return BlockClient{
		transport: transport.NewTCPTransport(conn),
		pipeline:  p,
		config:    config,
	}, nil
}

func (c BlockClient) Call(ctx context.Context, req *remote.Request) (*remote.Response, error) {
	reqData, err := c.config.Codec.Encode(req)
	if err != nil {
		return nil, err
	}
	buffer := buffer.NewByteBuffer()
	buffer.Write(reqData)
	if err := c.pipeline.Process(buffer); err != nil {
		return nil, err
	}
	buf := buffer.ReadAll()
	msg := protocol.Message{}
	msg.Decode(buf)

	if err := c.transport.Send(&msg); err != nil {
		return nil, err
	}
	respData, err := c.transport.Receive()
	if err != nil {
		return nil, err
	}
	buffer.Reset()
	buf, err = respData.Encode()
	if err != nil {
		return nil, err
	}
	buffer.Write(buf)
	if err := c.pipeline.Process(buffer); err != nil {
		return nil, err
	}
	var resp remote.Response
	if err := c.config.Codec.Decode(buffer.ReadAll(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c BlockClient) Close() error {
	return c.transport.Close()
}
