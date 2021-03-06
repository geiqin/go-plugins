package router

import (
	"github.com/geiqin/go-micro/config"
)

type Options struct {
	Config config.Config
}

func Config(c config.Config) Option {
	return func(o *Options) {
		o.Config = c
	}
}
