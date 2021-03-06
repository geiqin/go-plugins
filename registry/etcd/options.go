package etcd

import (
	"github.com/geiqin/go-micro/registry"
	"github.com/geiqin/go-micro/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
