package debugger

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	probe "github.com/stmtc233/bili-grpc-api-go/bilibili/api/probe/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestServicesAndDescribeEndpoints(t *testing.T) {
	handler := NewHandler()
	services := httptest.NewRecorder()
	handler.ServeHTTP(services, httptest.NewRequest(http.MethodGet, "/api/services", nil))
	if services.Code != http.StatusOK {
		t.Fatalf("services status = %d, want %d", services.Code, http.StatusOK)
	}
	var serviceList []serviceInfo
	if err := json.Unmarshal(services.Body.Bytes(), &serviceList); err != nil {
		t.Fatalf("decode services: %v", err)
	}
	var found bool
	for _, service := range serviceList {
		if service.Name == "bilibili.api.probe.v1.Probe" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("registered service list does not contain Probe")
	}

	describe := httptest.NewRecorder()
	handler.ServeHTTP(describe, httptest.NewRequest(http.MethodGet, "/api/describe?service=bilibili.api.probe.v1.Probe&method=TestCode", nil))
	if describe.Code != http.StatusOK {
		t.Fatalf("describe status = %d, want %d: %s", describe.Code, http.StatusOK, describe.Body.String())
	}
	if !strings.Contains(describe.Body.String(), "code") {
		t.Fatalf("describe response does not contain request field: %s", describe.Body.String())
	}
}

type probeServer struct {
	probe.UnimplementedProbeServer
	metadata string
}

func (s *probeServer) TestCode(ctx context.Context, _ *probe.CodeReq) (*probe.CodeReply, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		s.metadata = md.Get("x-debugger-test")[0]
	}
	return &probe.CodeReply{}, nil
}

func TestInvokeEndpoint(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	grpcServer := grpc.NewServer()
	service := &probeServer{}
	probe.RegisterProbeServer(grpcServer, service)
	go func() { _ = grpcServer.Serve(listener) }()
	defer grpcServer.Stop()

	payload := fmt.Sprintf(`{"target":%q,"service":"bilibili.api.probe.v1.Probe","method":"TestCode","request":{},"metadata":{"x-debugger-test":"ok"},"timeoutMs":3000}`, listener.Addr().String())
	recorder := httptest.NewRecorder()
	NewHandler().ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "/api/invoke", strings.NewReader(payload)))
	if recorder.Code != http.StatusOK {
		t.Fatalf("invoke status = %d, want %d: %s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	if service.metadata != "ok" {
		t.Fatalf("metadata = %q, want %q", service.metadata, "ok")
	}
	if !strings.Contains(recorder.Body.String(), `"response"`) {
		t.Fatalf("invoke response missing response field: %s", recorder.Body.String())
	}
}
