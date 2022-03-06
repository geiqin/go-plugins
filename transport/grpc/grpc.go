// Package grpc provides a grpc transport
package grpc

import (
	"github.com/geiqin/go-micro/config/cmd"
	"github.com/geiqin/go-micro/transport"
	"github.com/geiqin/go-micro/transport/grpc"
)

func init() {
	cmd.DefaultTransports["grpc"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return grpc.NewTransport(opts...)
}
