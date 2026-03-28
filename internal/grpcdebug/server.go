package grpcdebug

import (
	"context"
	"crypto/tls"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

//go:embed static/*
var staticFiles embed.FS

type app struct {
	catalog     Catalog
	methods     map[string]methodEntry
	detailCache map[string]MethodDetail
	mu          sync.RWMutex
}

type methodEntry struct {
	summary MethodSummary
	file    protoreflect.FileDescriptor
	service protoreflect.ServiceDescriptor
	method  protoreflect.MethodDescriptor
}

type Catalog struct {
	Services []ServiceSummary `json:"services"`
}

type ServiceSummary struct {
	Name     string          `json:"name"`
	FullName string          `json:"fullName"`
	File     string          `json:"file"`
	Methods  []MethodSummary `json:"methods"`
}

type MethodSummary struct {
	Name            string `json:"name"`
	FullMethod      string `json:"fullMethod"`
	RequestType     string `json:"requestType"`
	ResponseType    string `json:"responseType"`
	ClientStreaming bool   `json:"clientStreaming"`
	ServerStreaming bool   `json:"serverStreaming"`
}

type MethodDetail struct {
	Summary        MethodSummary  `json:"summary"`
	ServiceName    string         `json:"serviceName"`
	PackageName    string         `json:"packageName"`
	ProtoFile      string         `json:"protoFile"`
	RequestSchema  *MessageSchema `json:"requestSchema"`
	ResponseSchema *MessageSchema `json:"responseSchema"`
	Unary          bool           `json:"unary"`
	StreamingHint  string         `json:"streamingHint,omitempty"`
}

type MessageSchema struct {
	Name      string        `json:"name"`
	FullName  string        `json:"fullName"`
	Fields    []FieldSchema `json:"fields,omitempty"`
	Reference bool          `json:"reference,omitempty"`
	Truncated bool          `json:"truncated,omitempty"`
}

type FieldSchema struct {
	Name        string         `json:"name"`
	JSONName    string         `json:"jsonName"`
	Type        string         `json:"type"`
	TypeLabel   string         `json:"typeLabel"`
	Repeated    bool           `json:"repeated"`
	Map         bool           `json:"map"`
	Required    bool           `json:"required"`
	HasPresence bool           `json:"hasPresence"`
	Oneof       string         `json:"oneof,omitempty"`
	EnumValues  []EnumValue    `json:"enumValues,omitempty"`
	Message     *MessageSchema `json:"message,omitempty"`
	MapSpec     *MapSchema     `json:"mapSpec,omitempty"`
}

type MapSchema struct {
	KeyType        string         `json:"keyType"`
	KeyTypeLabel   string         `json:"keyTypeLabel"`
	ValueType      string         `json:"valueType"`
	ValueTypeLabel string         `json:"valueTypeLabel"`
	EnumValues     []EnumValue    `json:"enumValues,omitempty"`
	Message        *MessageSchema `json:"message,omitempty"`
}

type EnumValue struct {
	Name   string `json:"name"`
	Number int32  `json:"number"`
}

type invokeRequest struct {
	Target             string          `json:"target"`
	UseTLS             bool            `json:"useTls"`
	InsecureSkipVerify bool            `json:"insecureSkipVerify"`
	ServerName         string          `json:"serverName"`
	Authority          string          `json:"authority"`
	TimeoutMS          int             `json:"timeoutMs"`
	FullMethod         string          `json:"fullMethod"`
	Metadata           json.RawMessage `json:"metadata"`
	Payload            json.RawMessage `json:"payload"`
}

type invokeResponse struct {
	FullMethod        string          `json:"fullMethod"`
	DurationMS        int64           `json:"durationMs"`
	NormalizedRequest json.RawMessage `json:"normalizedRequest,omitempty"`
	Response          json.RawMessage `json:"response,omitempty"`
	Error             *invokeError    `json:"error,omitempty"`
}

type invokeError struct {
	Raw     string            `json:"raw"`
	Code    string            `json:"code,omitempty"`
	Message string            `json:"message"`
	Details []json.RawMessage `json:"details,omitempty"`
}

func NewHandler() (http.Handler, error) {
	staticRoot, err := fs.Sub(staticFiles, "static")
	if err != nil {
		return nil, fmt.Errorf("open embedded static files: %w", err)
	}

	app, err := newApp()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticRoot)))
	mux.HandleFunc("/api/catalog", app.handleCatalog)
	mux.HandleFunc("/api/method", app.handleMethod)
	mux.HandleFunc("/api/invoke", app.handleInvoke)
	return mux, nil
}

func newApp() (*app, error) {
	catalog, methods := buildCatalog()
	if len(catalog.Services) == 0 {
		return nil, errors.New("no gRPC services were registered; check generated package imports")
	}

	return &app{
		catalog:     catalog,
		methods:     methods,
		detailCache: make(map[string]MethodDetail),
	}, nil
}

func buildCatalog() (Catalog, map[string]methodEntry) {
	services := make([]ServiceSummary, 0)
	methods := make(map[string]methodEntry)

	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if fd.Services().Len() == 0 {
			return true
		}

		for i := 0; i < fd.Services().Len(); i++ {
			serviceDesc := fd.Services().Get(i)
			service := ServiceSummary{
				Name:     string(serviceDesc.Name()),
				FullName: string(serviceDesc.FullName()),
				File:     fd.Path(),
				Methods:  make([]MethodSummary, 0, serviceDesc.Methods().Len()),
			}

			for j := 0; j < serviceDesc.Methods().Len(); j++ {
				methodDesc := serviceDesc.Methods().Get(j)
				fullMethod := fmt.Sprintf("/%s/%s", serviceDesc.FullName(), methodDesc.Name())
				summary := MethodSummary{
					Name:            string(methodDesc.Name()),
					FullMethod:      fullMethod,
					RequestType:     string(methodDesc.Input().FullName()),
					ResponseType:    string(methodDesc.Output().FullName()),
					ClientStreaming: methodDesc.IsStreamingClient(),
					ServerStreaming: methodDesc.IsStreamingServer(),
				}
				service.Methods = append(service.Methods, summary)
				methods[fullMethod] = methodEntry{
					summary: summary,
					file:    fd,
					service: serviceDesc,
					method:  methodDesc,
				}
			}

			sort.Slice(service.Methods, func(i, j int) bool {
				return service.Methods[i].Name < service.Methods[j].Name
			})
			services = append(services, service)
		}
		return true
	})

	sort.Slice(services, func(i, j int) bool {
		return services[i].FullName < services[j].FullName
	})

	return Catalog{Services: services}, methods
}

func (a *app) handleCatalog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, fmt.Errorf("unsupported method %s", r.Method))
		return
	}
	writeJSON(w, http.StatusOK, a.catalog)
}

func (a *app) handleMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, fmt.Errorf("unsupported method %s", r.Method))
		return
	}

	fullMethod := strings.TrimSpace(r.URL.Query().Get("fullMethod"))
	if fullMethod == "" {
		writeJSONError(w, http.StatusBadRequest, errors.New("missing fullMethod"))
		return
	}

	detail, err := a.methodDetail(fullMethod)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err)
		return
	}
	writeJSON(w, http.StatusOK, detail)
}

func (a *app) handleInvoke(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, fmt.Errorf("unsupported method %s", r.Method))
		return
	}

	defer r.Body.Close()
	body, err := io.ReadAll(io.LimitReader(r.Body, 8<<20))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, fmt.Errorf("read request body: %w", err))
		return
	}

	var req invokeRequest
	if err := json.Unmarshal(body, &req); err != nil {
		writeJSONError(w, http.StatusBadRequest, fmt.Errorf("decode request body: %w", err))
		return
	}

	resp, statusCode, err := a.invoke(r.Context(), req)
	if err != nil {
		writeJSON(w, statusCode, resp)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (a *app) methodDetail(fullMethod string) (MethodDetail, error) {
	a.mu.RLock()
	if detail, ok := a.detailCache[fullMethod]; ok {
		a.mu.RUnlock()
		return detail, nil
	}
	a.mu.RUnlock()

	entry, ok := a.methods[fullMethod]
	if !ok {
		return MethodDetail{}, fmt.Errorf("method %s not found", fullMethod)
	}

	detail := MethodDetail{
		Summary:        entry.summary,
		ServiceName:    string(entry.service.FullName()),
		PackageName:    string(entry.file.Package()),
		ProtoFile:      entry.file.Path(),
		RequestSchema:  buildMessageSchema(entry.method.Input(), make(map[protoreflect.FullName]int), 0),
		ResponseSchema: buildMessageSchema(entry.method.Output(), make(map[protoreflect.FullName]int), 0),
		Unary:          !entry.method.IsStreamingClient() && !entry.method.IsStreamingServer(),
	}
	if !detail.Unary {
		detail.StreamingHint = streamingHint(entry.method)
	}

	a.mu.Lock()
	a.detailCache[fullMethod] = detail
	a.mu.Unlock()
	return detail, nil
}

func (a *app) invoke(parent context.Context, req invokeRequest) (invokeResponse, int, error) {
	entry, ok := a.methods[strings.TrimSpace(req.FullMethod)]
	if !ok {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     "method not found",
				Message: "未找到对应的 gRPC 方法描述符。",
			},
		}
		return resp, http.StatusBadRequest, errors.New("unknown method")
	}

	if entry.method.IsStreamingClient() || entry.method.IsStreamingServer() {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     "streaming methods are not supported yet",
				Message: "当前调试器暂时只支持一元 RPC，流式接口还没有接入。",
			},
		}
		return resp, http.StatusBadRequest, errors.New("streaming rpc is not supported")
	}

	target := strings.TrimSpace(req.Target)
	if target == "" {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     "missing target",
				Message: "请先填写目标 gRPC 地址，例如 grpc.example.com:443。",
			},
		}
		return resp, http.StatusBadRequest, errors.New("missing target")
	}

	timeout := time.Duration(req.TimeoutMS) * time.Millisecond
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	callCtx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	md, err := parseMetadata(req.Metadata)
	if err != nil {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "metadata JSON 解析失败。",
			},
		}
		return resp, http.StatusBadRequest, err
	}
	if len(md) > 0 {
		callCtx = metadata.NewOutgoingContext(callCtx, md)
	}

	input := dynamicpb.NewMessage(entry.method.Input())
	payload := req.Payload
	if len(bytesTrimSpace(payload)) == 0 {
		payload = json.RawMessage(`{}`)
	}

	unmarshalOptions := protojson.UnmarshalOptions{
		DiscardUnknown: false,
		Resolver:       protoregistry.GlobalTypes,
	}
	if err := unmarshalOptions.Unmarshal(payload, input); err != nil {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "请求体 JSON 无法转换成 protobuf 消息。",
			},
		}
		return resp, http.StatusBadRequest, err
	}

	marshalOptions := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: false,
		Resolver:        protoregistry.GlobalTypes,
	}

	normalizedRequest, err := marshalOptions.Marshal(input)
	if err != nil {
		resp := invokeResponse{
			FullMethod: req.FullMethod,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "请求体已经解析成功，但重新序列化时失败了。",
			},
		}
		return resp, http.StatusInternalServerError, err
	}

	dialOptions, err := dialOptionsFor(req)
	if err != nil {
		resp := invokeResponse{
			FullMethod:        req.FullMethod,
			NormalizedRequest: normalizedRequest,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "初始化 gRPC 连接参数失败。",
			},
		}
		return resp, http.StatusBadRequest, err
	}

	conn, err := grpc.DialContext(callCtx, target, dialOptions...)
	if err != nil {
		resp := invokeResponse{
			FullMethod:        req.FullMethod,
			NormalizedRequest: normalizedRequest,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "连接目标 gRPC 服务失败。",
			},
		}
		return resp, http.StatusBadGateway, err
	}
	defer conn.Close()

	output := dynamicpb.NewMessage(entry.method.Output())
	start := time.Now()
	err = conn.Invoke(callCtx, req.FullMethod, input, output)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		invokeErr := &invokeError{
			Raw:     err.Error(),
			Message: err.Error(),
		}
		if st, ok := gstatus.FromError(err); ok {
			invokeErr.Code = st.Code().String()
			invokeErr.Message = st.Message()
			invokeErr.Details = marshalStatusDetails(st.Details(), marshalOptions)
		}

		resp := invokeResponse{
			FullMethod:        req.FullMethod,
			DurationMS:        elapsed,
			NormalizedRequest: normalizedRequest,
			Error:             invokeErr,
		}
		return resp, http.StatusBadGateway, err
	}

	responseJSON, err := marshalOptions.Marshal(output)
	if err != nil {
		resp := invokeResponse{
			FullMethod:        req.FullMethod,
			DurationMS:        elapsed,
			NormalizedRequest: normalizedRequest,
			Error: &invokeError{
				Raw:     err.Error(),
				Message: "接口调用成功，但返回值 JSON 格式化失败。",
			},
		}
		return resp, http.StatusInternalServerError, err
	}

	return invokeResponse{
		FullMethod:        req.FullMethod,
		DurationMS:        elapsed,
		NormalizedRequest: responseOrEmpty(normalizedRequest),
		Response:          responseOrEmpty(responseJSON),
	}, http.StatusOK, nil
}

func dialOptionsFor(req invokeRequest) ([]grpc.DialOption, error) {
	opts := []grpc.DialOption{grpc.WithBlock()}
	if req.UseTLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: req.InsecureSkipVerify}
		if serverName := strings.TrimSpace(req.ServerName); serverName != "" {
			tlsConfig.ServerName = serverName
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if authority := strings.TrimSpace(req.Authority); authority != "" {
		opts = append(opts, grpc.WithAuthority(authority))
	}
	return opts, nil
}

func buildMessageSchema(desc protoreflect.MessageDescriptor, seen map[protoreflect.FullName]int, depth int) *MessageSchema {
	if desc == nil {
		return nil
	}
	schema := &MessageSchema{Name: string(desc.Name()), FullName: string(desc.FullName())}
	if depth >= 6 {
		schema.Truncated = true
		return schema
	}
	if seen[desc.FullName()] > 0 {
		schema.Reference = true
		return schema
	}

	seen[desc.FullName()]++
	defer func() { seen[desc.FullName()]-- }()

	fields := desc.Fields()
	schema.Fields = make([]FieldSchema, 0, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		schema.Fields = append(schema.Fields, buildFieldSchema(fields.Get(i), seen, depth+1))
	}
	return schema
}

func buildFieldSchema(field protoreflect.FieldDescriptor, seen map[protoreflect.FullName]int, depth int) FieldSchema {
	schema := FieldSchema{
		Name:        string(field.Name()),
		JSONName:    field.JSONName(),
		Repeated:    field.Cardinality() == protoreflect.Repeated && !field.IsMap(),
		Map:         field.IsMap(),
		Required:    field.Cardinality() == protoreflect.Required,
		HasPresence: field.HasPresence(),
	}
	if oneof := field.ContainingOneof(); oneof != nil && !oneof.IsSynthetic() {
		schema.Oneof = string(oneof.Name())
	}

	if field.IsMap() {
		mapValue := field.MapValue()
		schema.Type = "map"
		schema.TypeLabel = fmt.Sprintf("map<%s, %s>", scalarTypeName(field.MapKey().Kind()), descriptorTypeLabel(mapValue))
		schema.MapSpec = &MapSchema{
			KeyType:        scalarTypeName(field.MapKey().Kind()),
			KeyTypeLabel:   scalarTypeName(field.MapKey().Kind()),
			ValueType:      typeName(mapValue),
			ValueTypeLabel: descriptorTypeLabel(mapValue),
		}
		if mapValue.Kind() == protoreflect.EnumKind {
			schema.MapSpec.EnumValues = buildEnumValues(mapValue.Enum())
		}
		if mapValue.Kind() == protoreflect.MessageKind || mapValue.Kind() == protoreflect.GroupKind {
			schema.MapSpec.Message = buildMessageSchema(mapValue.Message(), seen, depth)
		}
		return schema
	}

	schema.Type = typeName(field)
	schema.TypeLabel = descriptorTypeLabel(field)
	switch field.Kind() {
	case protoreflect.EnumKind:
		schema.EnumValues = buildEnumValues(field.Enum())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		schema.Message = buildMessageSchema(field.Message(), seen, depth)
	}
	return schema
}

func buildEnumValues(enum protoreflect.EnumDescriptor) []EnumValue {
	values := make([]EnumValue, 0, enum.Values().Len())
	for i := 0; i < enum.Values().Len(); i++ {
		value := enum.Values().Get(i)
		values = append(values, EnumValue{Name: string(value.Name()), Number: int32(value.Number())})
	}
	return values
}

func typeName(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "bytes"
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return "float"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "int64"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "int32"
	case protoreflect.EnumKind:
		return "enum"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return "message"
	default:
		return strings.ToLower(field.Kind().String())
	}
}

func scalarTypeName(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "bytes"
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return "float"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "int64"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "int32"
	default:
		return strings.ToLower(kind.String())
	}
}

func descriptorTypeLabel(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.EnumKind:
		return string(field.Enum().FullName())
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return string(field.Message().FullName())
	default:
		return scalarTypeName(field.Kind())
	}
}

func streamingHint(method protoreflect.MethodDescriptor) string {
	switch {
	case method.IsStreamingClient() && method.IsStreamingServer():
		return "双向流"
	case method.IsStreamingClient():
		return "客户端流"
	case method.IsStreamingServer():
		return "服务端流"
	default:
		return ""
	}
}

func parseMetadata(raw json.RawMessage) (metadata.MD, error) {
	md := metadata.MD{}
	if len(bytesTrimSpace(raw)) == 0 {
		return md, nil
	}

	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}
	for key, value := range payload {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}
		switch typed := value.(type) {
		case string:
			md.Append(key, typed)
		case []any:
			for _, item := range typed {
				md.Append(key, fmt.Sprint(item))
			}
		case nil:
		default:
			md.Append(key, fmt.Sprint(typed))
		}
	}
	return md, nil
}

func marshalStatusDetails(details []any, marshalOptions protojson.MarshalOptions) []json.RawMessage {
	out := make([]json.RawMessage, 0, len(details))
	for _, detail := range details {
		message, ok := detail.(proto.Message)
		if !ok {
			raw, err := json.Marshal(map[string]any{"type": fmt.Sprintf("%T", detail), "value": fmt.Sprint(detail)})
			if err == nil {
				out = append(out, raw)
			}
			continue
		}
		raw, err := marshalOptions.Marshal(message)
		if err != nil {
			fallback, marshalErr := json.Marshal(map[string]any{
				"type":  string(message.ProtoReflect().Descriptor().FullName()),
				"error": err.Error(),
			})
			if marshalErr == nil {
				out = append(out, fallback)
			}
			continue
		}
		out = append(out, raw)
	}
	return out
}

func bytesTrimSpace(raw []byte) []byte {
	return []byte(strings.TrimSpace(string(raw)))
}

func responseOrEmpty(raw []byte) json.RawMessage {
	if len(raw) == 0 {
		return json.RawMessage(`{}`)
	}
	return json.RawMessage(raw)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeJSONError(w http.ResponseWriter, statusCode int, err error) {
	writeJSON(w, statusCode, map[string]any{"error": err.Error()})
}
