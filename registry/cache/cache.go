// Package cache provides a registry cache
package cache

import (
	"github.com/geiqin/go-micro/registry"
	"github.com/geiqin/go-micro/registry/cache"
)

// New returns a new cache
func New(r registry.Registry, opts ...cache.Option) cache.Cache {
	return cache.New(r, opts...)
}
