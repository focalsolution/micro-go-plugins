module github.com/asim/nitro-plugins/wrapper/ratelimiter/ratelimit/v3

go 1.13

require (
	github.com/juju/ratelimit v1.0.2-0.20191002062651-f60b32039441
	github.com/asim/nitro/v3 v3.3.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
