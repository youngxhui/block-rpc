package client

import (
	"time"

	"github.com/youngxhui/block-rpc/transport/codec"
)

type Option func(*Config)

type Config struct {
	Network string             // 网络协议，例如 tcp
	Address string             // 网络服务器
	Codec   codec.PayloadCodec // 编解码器
	Timeout time.Duration      // 超时时间
}

func defaultConfig() *Config {
	return &Config{
		Network: "tcp",
		Codec:   &codec.JSONCodec{},
		Timeout: 10 * time.Second,
	}
}

func WithAddress(addr string) Option {
	return func(c *Config) {
		c.Address = addr
	}
}

func WithCodec(codec codec.PayloadCodec) Option {
	return func(c *Config) {
		c.Codec = codec
	}
}

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithNetwork 设置网络类型
func WithNetwork(network string) Option {
	return func(c *Config) {
		c.Network = network
	}
}
