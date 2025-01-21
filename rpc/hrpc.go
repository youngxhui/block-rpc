package rpc

type HRPC struct {
}

func New(opts ...RpcOptions) Rpc {
	return HRPC{}
}

func (h HRPC) Start() error {
	return nil
}

func (h HRPC) Stop() error {
	return nil
}
