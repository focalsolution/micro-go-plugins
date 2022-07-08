package cache

import (
	"time"

	"github.com/focalsolution/micro-go-micro/registry/cache"
)

// WithTTL sets the cache TTL
func WithTTL(t time.Duration) cache.Option {
	return cache.WithTTL(t)
}
