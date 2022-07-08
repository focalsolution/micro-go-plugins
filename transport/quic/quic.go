// Package quic provides a QUIC based transport
package quic

import (
	"github.com/focalsolution/micro-go-micro/config/cmd"
	"github.com/focalsolution/micro-go-micro/transport"
	"github.com/focalsolution/micro-go-micro/transport/quic"
)

func init() {
	cmd.DefaultTransports["quic"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return quic.NewTransport(opts...)
}
