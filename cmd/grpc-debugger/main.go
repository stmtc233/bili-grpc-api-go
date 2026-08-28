package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:8090", "HTTP listen address")
	flag.Parse()

	server := &http.Server{
		Addr:              *listen,
		Handler:           newDebuggerHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("gRPC debugger listening on http://%s", *listen)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
