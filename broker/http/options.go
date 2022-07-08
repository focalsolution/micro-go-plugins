package http

import (
	"net/http"

	"github.com/focalsolution/micro-go-micro/broker"
	thttp "github.com/focalsolution/micro-go-micro/broker/http"
)

// Handle registers the handler for the given pattern.
func Handle(pattern string, handler http.Handler) broker.Option {
	return thttp.Handle(pattern, handler)
}
