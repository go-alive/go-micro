// Package service encapsulates the client, server and other interfaces to provide a complete micro service.
package service

import (
	"github.com/go-alive/go-micro/client"
	"github.com/go-alive/go-micro/server"
)

// Service is an interface for a micro service
type Service interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}
