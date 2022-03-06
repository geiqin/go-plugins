// Package memory provides a memory broker
package memory

import (
	"github.com/geiqin/go-micro/broker"
	"github.com/geiqin/go-micro/broker/nats"
	"github.com/geiqin/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["nats"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return nats.NewBroker(opts...)
}
