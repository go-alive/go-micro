package config

import (
	"github.com/go-alive/go-micro/config/loader"
	"github.com/go-alive/go-micro/config/reader"
	"github.com/go-alive/go-micro/config/source"
)

// WithLoader sets the loader for manager config
func WithLoader(l loader.Loader) Option {
	return func(o *Options) {
		o.Loader = l
	}
}

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}
