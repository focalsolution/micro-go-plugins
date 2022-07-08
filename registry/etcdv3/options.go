package etcdv3

import (
	"github.com/focalsolution/micro-go-micro/registry"
	"github.com/focalsolution/micro-go-micro/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
