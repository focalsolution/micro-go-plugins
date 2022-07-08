// Package memory provides a memory broker
package memory

import (
	"github.com/focalsolution/micro-go-micro/broker"
	"github.com/focalsolution/micro-go-micro/broker/memory"
	"github.com/focalsolution/micro-go-micro/config/cmd"
)

func init() {
	cmd.DefaultBrokers["memory"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return memory.NewBroker(opts...)
}
