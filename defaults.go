package micro

import (
	"github.com/go-alive/go-micro/client"
	"github.com/go-alive/go-micro/debug/trace"
	"github.com/go-alive/go-micro/server"

	// set defaults
	gcli "github.com/go-alive/go-micro/client/grpc"
	memTrace "github.com/go-alive/go-micro/debug/trace/memory"
	gsrv "github.com/go-alive/go-micro/server/grpc"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
