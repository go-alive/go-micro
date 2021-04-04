package mock

import (
	"github.com/go-alive/go-micro/registry"
	"github.com/go-alive/go-micro/server"
)

type MockSubscriber struct {
	Id   string
	Opts server.SubscriberOptions
	Sub  interface{}
}

func (m *MockSubscriber) Topic() string {
	return m.Id
}

func (m *MockSubscriber) Subscriber() interface{} {
	return m.Sub
}

func (m *MockSubscriber) Endpoints() []*registry.Endpoint {
	return []*registry.Endpoint{}
}

func (m *MockSubscriber) Options() server.SubscriberOptions {
	return m.Opts
}
