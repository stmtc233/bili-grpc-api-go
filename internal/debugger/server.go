package debugger

import (
	"context"
	"crypto/tls"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	_ "github.com/stmtc233/bili-grpc-api-go/internal/debugger/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

//go:embed web/index.html
var webFiles embed.FS

const (
	defaultTimeout = 15 * time.Second
	maxTimeout     = 5 * time.Minute
)

// NewHandler returns the standalone debugger HTTP handler.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/api/services", serveServices)
	mux.HandleFunc("/api/describe", serveDescribe)
	mux.HandleFunc("/api/invoke", serveInvoke)
	return mux
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data, err := webFiles.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, "debugger UI is unavailable", http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

type serviceInfo struct {
	Name    string       `json:"name"`
	Methods []methodInfo `json:"methods"`
}

type methodInfo struct {
	Name         string `json:"name"`
	Input        string `json:"input"`
	Output       string `json:"output"`
	ClientStream bool   `json:"clientStream"`
	ServerStream bool   `json:"serverStream"`
}

func registeredServices() []serviceInfo {
	services := make([]serviceInfo, 0)
	files := protoregistry.GlobalFiles
	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		for i := 0; i < file.Services().Len(); i++ {
			svc := file.Services().Get(i)
			info := serviceInfo{Name: string(svc.FullName()), Methods: make([]methodInfo, 0, svc.Methods().Len())}
			for j := 0; j < svc.Methods().Len(); j++ {
				method := svc.Methods().Get(j)
				info.Methods = append(info.Methods, methodInfo{
					Name:         string(method.Name()),
					Input:        string(method.Input().FullName()),
					Output:       string(method.Output().FullName()),
					ClientStream: method.IsStreamingClient(),
					ServerStream: method.IsStreamingServer(),
				})
			}
			sort.Slice(info.Methods, func(a, b int) bool { return info.Methods[a].Name < info.Methods[b].Name })
			services = append(services, info)
		}
		return true
	})
	sort.Slice(services, func(a, b int) bool { return services[a].Name < services[b].Name })
	return services
}

func serveServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	writeJSON(w, http.StatusOK, registeredServices())
}

type fieldInfo struct {
	Name        string      `json:"name"`
	JSONName    string      `json:"jsonName"`
	Kind        string      `json:"kind"`
	Cardinality string      `json:"cardinality"`
	Optional    bool        `json:"optional"`
	Message     string      `json:"message,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
	Fields      []fieldInfo `json:"fields,omitempty"`
	Recursive   bool        `json:"recursive,omitempty"`
	MapKey      string      `json:"mapKey,omitempty"`
	MapValue    string      `json:"mapValue,omitempty"`
}

type messageInfo struct {
	Name   string      `json:"name"`
	Fields []fieldInfo `json:"fields"`
}

type describeResponse struct {
	Service string      `json:"service"`
	Method  string      `json:"method"`
	Input   messageInfo `json:"input"`
	Output  string      `json:"output"`
}

func serveDescribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	service, method, err := findMethod(r.URL.Query().Get("service"), r.URL.Query().Get("method"))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, describeResponse{
		Service: string(service.FullName()),
		Method:  string(method.Name()),
		Input: messageInfo{
			Name:   string(method.Input().FullName()),
			Fields: describeFields(method.Input(), map[protoreflect.FullName]bool{}),
		},
		Output: string(method.Output().FullName()),
	})
}

func describeFields(message protoreflect.MessageDescriptor, stack map[protoreflect.FullName]bool) []fieldInfo {
	if stack[message.FullName()] {
		return nil
	}
	stack[message.FullName()] = true
	defer delete(stack, message.FullName())
	fields := make([]fieldInfo, 0, message.Fields().Len())
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		item := fieldInfo{
			Name:        string(field.Name()),
			JSONName:    field.JSONName(),
			Kind:        field.Kind().String(),
			Cardinality: field.Cardinality().String(),
			Optional:    field.HasOptionalKeyword(),
		}
		if field.IsList() {
			item.Cardinality = "repeated"
		}
		if field.IsMap() {
			item.Kind = "map"
			item.MapKey = field.MapKey().Kind().String()
			item.MapValue = field.MapValue().Kind().String()
		}
		if field.Enum() != nil {
			item.Enum = make([]string, 0, field.Enum().Values().Len())
			for j := 0; j < field.Enum().Values().Len(); j++ {
				item.Enum = append(item.Enum, string(field.Enum().Values().Get(j).Name()))
			}
		}
		if field.Message() != nil {
			item.Message = string(field.Message().FullName())
			if !stack[field.Message().FullName()] {
				item.Fields = describeFields(field.Message(), stack)
			} else {
				item.Recursive = true
			}
		}
		fields = append(fields, item)
	}
	return fields
}

type invokeRequest struct {
	Target     string                     `json:"target"`
	Service    string                     `json:"service"`
	Method     string                     `json:"method"`
	Request    json.RawMessage            `json:"request"`
	TLS        bool                       `json:"tls"`
	Authority  string                     `json:"authority"`
	ServerName string                     `json:"serverName"`
	Metadata   map[string]json.RawMessage `json:"metadata"`
	TimeoutMS  int                        `json:"timeoutMs"`
}

func serveInvoke(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var request invokeRequest
	decoder := json.NewDecoder(io.LimitReader(r.Body, 2<<20))
	if err := decoder.Decode(&request); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid request: "+err.Error())
		return
	}
	service, method, err := findMethod(request.Service, request.Method)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	if method.IsStreamingClient() || method.IsStreamingServer() {
		writeJSONError(w, http.StatusBadRequest, "streaming methods are not supported by the debugger")
		return
	}
	if strings.TrimSpace(request.Target) == "" {
		writeJSONError(w, http.StatusBadRequest, "target is required")
		return
	}
	metadataValues, err := parseMetadata(request.Metadata)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	requestBody := request.Request
	if len(requestBody) == 0 || string(requestBody) == "null" {
		requestBody = []byte("{}")
	}
	in := dynamicpb.NewMessage(method.Input())
	if err := (protojson.UnmarshalOptions{DiscardUnknown: false}).Unmarshal(requestBody, in); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid protobuf request: "+err.Error())
		return
	}

	timeout := defaultTimeout
	if request.TimeoutMS > 0 {
		timeout = time.Duration(request.TimeoutMS) * time.Millisecond
	}
	if timeout > maxTimeout {
		timeout = maxTimeout
	}
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	options := []grpc.DialOption{grpc.WithBlock()}
	if request.TLS {
		options = append(options, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: request.ServerName,
		})))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if request.Authority != "" {
		options = append(options, grpc.WithAuthority(request.Authority))
	}
	conn, err := grpc.DialContext(ctx, request.Target, options...)
	if err != nil {
		writeJSONError(w, http.StatusBadGateway, "dial failed: "+err.Error())
		return
	}
	defer conn.Close()
	if metadataValues != nil {
		ctx = metadata.NewOutgoingContext(ctx, metadataValues)
	}
	out := dynamicpb.NewMessage(method.Output())
	fullMethod := "/" + string(service.FullName()) + "/" + string(method.Name())
	if err := conn.Invoke(ctx, fullMethod, in, out); err != nil {
		writeJSONError(w, http.StatusBadGateway, "RPC failed: "+err.Error())
		return
	}
	response, err := (protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true, Indent: "  "}).Marshal(out)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to encode response: "+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]json.RawMessage{
		"request":  json.RawMessage(requestBody),
		"response": response,
	})
}

func findMethod(serviceName, methodName string) (protoreflect.ServiceDescriptor, protoreflect.MethodDescriptor, error) {
	serviceName = strings.TrimPrefix(strings.TrimSpace(serviceName), "/")
	methodName = strings.TrimSpace(methodName)
	if serviceName == "" || methodName == "" {
		return nil, nil, errors.New("service and method are required")
	}
	descriptor, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(serviceName))
	if err != nil {
		return nil, nil, fmt.Errorf("service %q was not found", serviceName)
	}
	service, ok := descriptor.(protoreflect.ServiceDescriptor)
	if !ok {
		return nil, nil, fmt.Errorf("%q is not a service", serviceName)
	}
	method := service.Methods().ByName(protoreflect.Name(methodName))
	if method == nil {
		return nil, nil, fmt.Errorf("method %q was not found in service %q", methodName, serviceName)
	}
	return service, method, nil
}

func parseMetadata(raw map[string]json.RawMessage) (metadata.MD, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	md := metadata.MD{}
	for key, value := range raw {
		if strings.TrimSpace(key) == "" {
			return nil, errors.New("metadata keys cannot be empty")
		}
		var single string
		if err := json.Unmarshal(value, &single); err == nil {
			md.Set(key, single)
			continue
		}
		var multiple []string
		if err := json.Unmarshal(value, &multiple); err != nil {
			return nil, fmt.Errorf("metadata %q must be a string or string array", key)
		}
		for _, item := range multiple {
			md.Append(key, item)
		}
	}
	return md, nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
