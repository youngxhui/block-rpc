package rpc

import "time"

type RpcOptions = string

type Rpc interface {
	Start() error
	Stop() error
}

type Option struct {
}

type defaultServer struct {
	Timeout time.Duration
}

// NewServer 服务
func NewServer(opt ...Option) Rpc {
	return defaultServer{
		Timeout: 5 * time.Second,
	}
}

func (s defaultServer) Start() error {
	return nil
}

func (s defaultServer) Stop() error {
	return nil
}
