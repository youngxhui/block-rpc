package rpc

type RpcOptions = string
type Rpc interface {
	Start() error
	Stop() error
}

