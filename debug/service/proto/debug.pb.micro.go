// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: debug/service/proto/debug.proto

package debug

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/go-alive/go-micro/api"
	client "github.com/go-alive/go-micro/client"
	server "github.com/go-alive/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Debug service

func NewDebugEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Debug service

type DebugService interface {
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error)
	Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error)
	Cache(ctx context.Context, in *CacheRequest, opts ...client.CallOption) (*CacheResponse, error)
}

type debugService struct {
	c    client.Client
	name string
}

func NewDebugService(name string, c client.Client) DebugService {
	return &debugService{
		c:    c,
		name: name,
	}
}

func (c *debugService) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error) {
	req := c.c.NewRequest(c.name, "Debug.Log", &LogRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &debugServiceLog{stream}, nil
}

type Debug_LogService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Record, error)
}

type debugServiceLog struct {
	stream client.Stream
}

func (x *debugServiceLog) Close() error {
	return x.stream.Close()
}

func (x *debugServiceLog) Context() context.Context {
	return x.stream.Context()
}

func (x *debugServiceLog) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugServiceLog) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugServiceLog) Recv() (*Record, error) {
	m := new(Record)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugService) Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Health", in)
	out := new(HealthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Trace", in)
	out := new(TraceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Cache(ctx context.Context, in *CacheRequest, opts ...client.CallOption) (*CacheResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Cache", in)
	out := new(CacheResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Debug service

type DebugHandler interface {
	Log(context.Context, *LogRequest, Debug_LogStream) error
	Health(context.Context, *HealthRequest, *HealthResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
	Trace(context.Context, *TraceRequest, *TraceResponse) error
	Cache(context.Context, *CacheRequest, *CacheResponse) error
}

func RegisterDebugHandler(s server.Server, hdlr DebugHandler, opts ...server.HandlerOption) error {
	type debug interface {
		Log(ctx context.Context, stream server.Stream) error
		Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error
		Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error
		Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error
		Cache(ctx context.Context, in *CacheRequest, out *CacheResponse) error
	}
	type Debug struct {
		debug
	}
	h := &debugHandler{hdlr}
	return s.Handle(s.NewHandler(&Debug{h}, opts...))
}

type debugHandler struct {
	DebugHandler
}

func (h *debugHandler) Log(ctx context.Context, stream server.Stream) error {
	m := new(LogRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Log(ctx, m, &debugLogStream{stream})
}

type Debug_LogStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Record) error
}

type debugLogStream struct {
	stream server.Stream
}

func (x *debugLogStream) Close() error {
	return x.stream.Close()
}

func (x *debugLogStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugLogStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugLogStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugLogStream) Send(m *Record) error {
	return x.stream.Send(m)
}

func (h *debugHandler) Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error {
	return h.DebugHandler.Health(ctx, in, out)
}

func (h *debugHandler) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.DebugHandler.Stats(ctx, in, out)
}

func (h *debugHandler) Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error {
	return h.DebugHandler.Trace(ctx, in, out)
}

func (h *debugHandler) Cache(ctx context.Context, in *CacheRequest, out *CacheResponse) error {
	return h.DebugHandler.Cache(ctx, in, out)
}
