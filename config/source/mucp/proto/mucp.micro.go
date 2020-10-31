// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mucp.proto

package mucp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/asim/nitro/v3/client"
	server "github.com/asim/nitro/v3/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Source service

type SourceService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Source_WatchService, error)
}

type sourceService struct {
	c    client.Client
	name string
}

func NewSourceService(name string, c client.Client) SourceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "source"
	}
	return &sourceService{
		c:    c,
		name: name,
	}
}

func (c *sourceService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Source.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceService) Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Source_WatchService, error) {
	req := c.c.NewRequest(c.name, "Source.Watch", &WatchRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &sourceServiceWatch{stream}, nil
}

type Source_WatchService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WatchResponse, error)
}

type sourceServiceWatch struct {
	stream client.Stream
}

func (x *sourceServiceWatch) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceWatch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceWatch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceWatch) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Source service

type SourceHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Watch(context.Context, *WatchRequest, Source_WatchStream) error
}

func RegisterSourceHandler(s server.Server, hdlr SourceHandler, opts ...server.HandlerOption) error {
	type source interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Watch(ctx context.Context, stream server.Stream) error
	}
	type Source struct {
		source
	}
	h := &sourceHandler{hdlr}
	return s.Handle(s.NewHandler(&Source{h}, opts...))
}

type sourceHandler struct {
	SourceHandler
}

func (h *sourceHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.SourceHandler.Read(ctx, in, out)
}

func (h *sourceHandler) Watch(ctx context.Context, stream server.Stream) error {
	m := new(WatchRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SourceHandler.Watch(ctx, m, &sourceWatchStream{stream})
}

type Source_WatchStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WatchResponse) error
}

type sourceWatchStream struct {
	stream server.Stream
}

func (x *sourceWatchStream) Close() error {
	return x.stream.Close()
}

func (x *sourceWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceWatchStream) Send(m *WatchResponse) error {
	return x.stream.Send(m)
}
