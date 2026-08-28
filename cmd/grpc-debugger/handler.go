package main

import (
	"net/http"

	"github.com/stmtc233/bili-grpc-api-go/internal/debugger"
)

func newDebuggerHandler() http.Handler {
	return debugger.NewHandler()
}
