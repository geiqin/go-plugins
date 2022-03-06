// Package etcd provides an etcd v3 service registry
package etcd

import (
	"github.com/geiqin/go-micro/config/cmd"
	"github.com/geiqin/go-micro/registry"
	"github.com/geiqin/go-micro/registry/etcd"
)

func init() {
	cmd.DefaultRegistries["etcd"] = etcd.NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return etcd.NewRegistry(opts...)
}
