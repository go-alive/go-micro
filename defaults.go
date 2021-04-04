package micro

import (
	"github.com/go-alive/go-micro/client"
	"github.com/go-alive/go-micro/debug/trace"
	"github.com/go-alive/go-micro/server"
	"github.com/go-alive/go-micro/store"

	// set defaults
	gcli "github.com/go-alive/go-micro/client/grpc"
	memTrace "github.com/go-alive/go-micro/debug/trace/memory"
	gsrv "github.com/go-alive/go-micro/server/grpc"
	memoryStore "github.com/go-alive/go-micro/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
