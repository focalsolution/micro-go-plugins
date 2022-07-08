package main

import (
	"net"
	"net/http"

	"github.com/focalsolution/micro-go-micro/config"
	"github.com/focalsolution/micro-go-micro/config/source/file"
	"github.com/focalsolution/micro-go-micro/util/log"
	"github.com/focalsolution/micro-go-plugins/focalsolution/micro-router"
)

func main() {
	// Create listener
	l, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Create Config Source
	f := file.NewSource(file.WithPath("routes.json"))
	conf := config.NewConfig(config.WithSource(f))

	// Create Router
	r := router.NewRouter(router.Config(conf))

	// Setup Handler
	wr := r.Handler()
	h := wr(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", 404)
	}))

	// Start Server
	if err := http.Serve(l, h); err != nil {
		log.Fatal(err)
	}
}
