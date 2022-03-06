// Package memory provides a memory broker
package memory

import (
	"github.com/geiqin/go-micro/broker"
	"github.com/geiqin/go-micro/broker/memory"
	"github.com/geiqin/go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["memory"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return memory.NewBroker(opts...)
}
