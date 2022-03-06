// Package quic provides a QUIC based transport
package quic

import (
	"github.com/geiqin/go-micro/config/cmd"
	"github.com/geiqin/go-micro/transport"
	"github.com/geiqin/go-micro/transport/quic"
)

func init() {
	cmd.DefaultTransports["quic"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return quic.NewTransport(opts...)
}
