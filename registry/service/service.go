// Package service uses the registry service
package service

import (
	"github.com/geiqin/go-micro/config/cmd"
	"github.com/geiqin/go-micro/registry"
	"github.com/geiqin/go-micro/registry/service"
)

func init() {
	cmd.DefaultRegistries["service"] = NewRegistry
}

// NewRegistry returns a new registry service client
func NewRegistry(opts ...registry.Option) registry.Registry {
	return service.NewRegistry(opts...)
}
