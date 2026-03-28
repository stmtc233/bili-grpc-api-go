package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/stmtc233/bili-grpc-api-go/internal/grpcdebug"
)

func main() {
	listenAddr := flag.String("listen", ":8090", "HTTP listen address for the debug UI")
	flag.Parse()

	handler, err := grpcdebug.NewHandler()
	if err != nil {
		log.Fatalf("init debug handler: %v", err)
	}

	server := &http.Server{
		Addr:              *listenAddr,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	displayAddr := *listenAddr
	if strings.HasPrefix(displayAddr, ":") {
		displayAddr = "127.0.0.1" + displayAddr
	}

	fmt.Printf("bili-grpc-api-go debug UI is running at http://%s\n", displayAddr)
	fmt.Println("Press Ctrl+C to stop.")

	log.Fatal(server.ListenAndServe())
}
