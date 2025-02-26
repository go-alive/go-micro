package reader

import (
	"github.com/go-alive/go-micro/config/encoder"
	"github.com/go-alive/go-micro/config/encoder/hcl"
	"github.com/go-alive/go-micro/config/encoder/json"
	"github.com/go-alive/go-micro/config/encoder/toml"
	"github.com/go-alive/go-micro/config/encoder/xml"
	"github.com/go-alive/go-micro/config/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"hcl":  hcl.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
