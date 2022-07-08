// Package service provides the broker service client
package service

import (
	"github.com/focalsolution/micro-go-micro/broker"
	"github.com/focalsolution/micro-go-micro/broker/service"
	"github.com/focalsolution/micro-go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["service"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return service.NewBroker(opts...)
}
