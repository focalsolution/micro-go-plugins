// Package service provides the broker service client
package service

import (
	"github.com/asim/nitro/v3/broker"
	"github.com/asim/nitro/v3/broker/service"
	"github.com/asim/nitro/v3/cmd"
)

func init() {
	cmd.DefaultBrokers["service"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return service.NewBroker(opts...)
}
