// Package http returns a http2 transport using net/http
package http

import (
	"github.com/go-alive/go-micro/transport"
)

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
