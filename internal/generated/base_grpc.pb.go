package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SparkConnectService_ExecutePlan_FullMethodName     = "/spark.connect.SparkConnectService/ExecutePlan"
	SparkConnectService_AnalyzePlan_FullMethodName     = "/spark.connect.SparkConnectService/AnalyzePlan"
	SparkConnectService_Config_FullMethodName          = "/spark.connect.SparkConnectService/Config"
	SparkConnectService_AddArtifacts_FullMethodName    = "/spark.connect.SparkConnectService/AddArtifacts"
	SparkConnectService_ArtifactStatus_FullMethodName  = "/spark.connect.SparkConnectService/ArtifactStatus"
	SparkConnectService_Interrupt_FullMethodName       = "/spark.connect.SparkConnectService/Interrupt"
	SparkConnectService_ReattachExecute_FullMethodName = "/spark.connect.SparkConnectService/ReattachExecute"
	SparkConnectService_ReleaseExecute_FullMethodName  = "/spark.connect.SparkConnectService/ReleaseExecute"
)

// SparkConnectServiceClient is the client API for SparkConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SparkConnectServiceClient interface {
	// Executes a request that contains the query and returns a stream of [[Response]].
	//
	// It is guaranteed that there is at least one ARROW batch returned even if the result set is empty.
	ExecutePlan(ctx context.Context, in *ExecutePlanRequest, opts ...grpc.CallOption) (SparkConnectService_ExecutePlanClient, error)
	// Analyzes a query and returns a [[AnalyzeResponse]] containing metadata about the query.
	AnalyzePlan(ctx context.Context, in *AnalyzePlanRequest, opts ...grpc.CallOption) (*AnalyzePlanResponse, error)
	// Update or fetch the configurations and returns a [[ConfigResponse]] containing the result.
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	// Add artifacts to the session and returns a [[AddArtifactsResponse]] containing metadata about
	// the added artifacts.
	AddArtifacts(ctx context.Context, opts ...grpc.CallOption) (SparkConnectService_AddArtifactsClient, error)
	// Check statuses of artifacts in the session and returns them in a [[ArtifactStatusesResponse]]
	ArtifactStatus(ctx context.Context, in *ArtifactStatusesRequest, opts ...grpc.CallOption) (*ArtifactStatusesResponse, error)
	// Interrupts running executions
	Interrupt(ctx context.Context, in *InterruptRequest, opts ...grpc.CallOption) (*InterruptResponse, error)
	// Reattach to an existing reattachable execution.
	// The ExecutePlan must have been started with ReattachOptions.reattachable=true.
	// If the ExecutePlanResponse stream ends without a ResultComplete message, there is more to
	// continue. If there is a ResultComplete, the client should use ReleaseExecute with
	ReattachExecute(ctx context.Context, in *ReattachExecuteRequest, opts ...grpc.CallOption) (SparkConnectService_ReattachExecuteClient, error)
	// Release an reattachable execution, or parts thereof.
	// The ExecutePlan must have been started with ReattachOptions.reattachable=true.
	// Non reattachable executions are released automatically and immediately after the ExecutePlan
	// RPC and ReleaseExecute may not be used.
	ReleaseExecute(ctx context.Context, in *ReleaseExecuteRequest, opts ...grpc.CallOption) (*ReleaseExecuteResponse, error)
}

type sparkConnectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSparkConnectServiceClient(cc grpc.ClientConnInterface) SparkConnectServiceClient {
	return &sparkConnectServiceClient{cc}
}

func (c *sparkConnectServiceClient) ExecutePlan(ctx context.Context, in *ExecutePlanRequest, opts ...grpc.CallOption) (SparkConnectService_ExecutePlanClient, error) {
	stream, err := c.cc.NewStream(ctx, &SparkConnectService_ServiceDesc.Streams[0], SparkConnectService_ExecutePlan_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sparkConnectServiceExecutePlanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SparkConnectService_ExecutePlanClient interface {
	Recv() (*ExecutePlanResponse, error)
	grpc.ClientStream
}

type sparkConnectServiceExecutePlanClient struct {
	grpc.ClientStream
}

// 1
func (x *sparkConnectServiceExecutePlanClient) Recv() (*ExecutePlanResponse, error) {
	m := new(ExecutePlanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sparkConnectServiceClient) AnalyzePlan(ctx context.Context, in *AnalyzePlanRequest, opts ...grpc.CallOption) (*AnalyzePlanResponse, error) {
	out := new(AnalyzePlanResponse)
	err := c.cc.Invoke(ctx, SparkConnectService_AnalyzePlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkConnectServiceClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, SparkConnectService_Config_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkConnectServiceClient) AddArtifacts(ctx context.Context, opts ...grpc.CallOption) (SparkConnectService_AddArtifactsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SparkConnectService_ServiceDesc.Streams[1], SparkConnectService_AddArtifacts_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sparkConnectServiceAddArtifactsClient{stream}
	return x, nil
}

type SparkConnectService_AddArtifactsClient interface {
	Send(*AddArtifactsRequest) error
	CloseAndRecv() (*AddArtifactsResponse, error)
	grpc.ClientStream
}

type sparkConnectServiceAddArtifactsClient struct {
	grpc.ClientStream
}

func (x *sparkConnectServiceAddArtifactsClient) Send(m *AddArtifactsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sparkConnectServiceAddArtifactsClient) CloseAndRecv() (*AddArtifactsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddArtifactsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sparkConnectServiceClient) ArtifactStatus(ctx context.Context, in *ArtifactStatusesRequest, opts ...grpc.CallOption) (*ArtifactStatusesResponse, error) {
	out := new(ArtifactStatusesResponse)
	err := c.cc.Invoke(ctx, SparkConnectService_ArtifactStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkConnectServiceClient) Interrupt(ctx context.Context, in *InterruptRequest, opts ...grpc.CallOption) (*InterruptResponse, error) {
	out := new(InterruptResponse)
	err := c.cc.Invoke(ctx, SparkConnectService_Interrupt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkConnectServiceClient) ReattachExecute(ctx context.Context, in *ReattachExecuteRequest, opts ...grpc.CallOption) (SparkConnectService_ReattachExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &SparkConnectService_ServiceDesc.Streams[2], SparkConnectService_ReattachExecute_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sparkConnectServiceReattachExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SparkConnectService_ReattachExecuteClient interface {
	Recv() (*ExecutePlanResponse, error)
	grpc.ClientStream
}

type sparkConnectServiceReattachExecuteClient struct {
	grpc.ClientStream
}

func (x *sparkConnectServiceReattachExecuteClient) Recv() (*ExecutePlanResponse, error) {
	m := new(ExecutePlanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sparkConnectServiceClient) ReleaseExecute(ctx context.Context, in *ReleaseExecuteRequest, opts ...grpc.CallOption) (*ReleaseExecuteResponse, error) {
	out := new(ReleaseExecuteResponse)
	err := c.cc.Invoke(ctx, SparkConnectService_ReleaseExecute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SparkConnectServiceServer is the server API for SparkConnectService service.
// All implementations must embed UnimplementedSparkConnectServiceServer
// for forward compatibility
type SparkConnectServiceServer interface {
	// Executes a request that contains the query and returns a stream of [[Response]].
	//
	// It is guaranteed that there is at least one ARROW batch returned even if the result set is empty.
	ExecutePlan(*ExecutePlanRequest, SparkConnectService_ExecutePlanServer) error
	// Analyzes a query and returns a [[AnalyzeResponse]] containing metadata about the query.
	AnalyzePlan(context.Context, *AnalyzePlanRequest) (*AnalyzePlanResponse, error)
	// Update or fetch the configurations and returns a [[ConfigResponse]] containing the result.
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
	// Add artifacts to the session and returns a [[AddArtifactsResponse]] containing metadata about
	// the added artifacts.
	AddArtifacts(SparkConnectService_AddArtifactsServer) error
	// Check statuses of artifacts in the session and returns them in a [[ArtifactStatusesResponse]]
	ArtifactStatus(context.Context, *ArtifactStatusesRequest) (*ArtifactStatusesResponse, error)
	// Interrupts running executions
	Interrupt(context.Context, *InterruptRequest) (*InterruptResponse, error)
	// Reattach to an existing reattachable execution.
	// The ExecutePlan must have been started with ReattachOptions.reattachable=true.
	// If the ExecutePlanResponse stream ends without a ResultComplete message, there is more to
	// continue. If there is a ResultComplete, the client should use ReleaseExecute with
	ReattachExecute(*ReattachExecuteRequest, SparkConnectService_ReattachExecuteServer) error
	// Release an reattachable execution, or parts thereof.
	// The ExecutePlan must have been started with ReattachOptions.reattachable=true.
	// Non reattachable executions are released automatically and immediately after the ExecutePlan
	// RPC and ReleaseExecute may not be used.
	ReleaseExecute(context.Context, *ReleaseExecuteRequest) (*ReleaseExecuteResponse, error)
	mustEmbedUnimplementedSparkConnectServiceServer()
}

// UnimplementedSparkConnectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSparkConnectServiceServer struct {
}

func (UnimplementedSparkConnectServiceServer) ExecutePlan(*ExecutePlanRequest, SparkConnectService_ExecutePlanServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecutePlan not implemented")
}
func (UnimplementedSparkConnectServiceServer) AnalyzePlan(context.Context, *AnalyzePlanRequest) (*AnalyzePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzePlan not implemented")
}
func (UnimplementedSparkConnectServiceServer) Config(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedSparkConnectServiceServer) AddArtifacts(SparkConnectService_AddArtifactsServer) error {
	return status.Errorf(codes.Unimplemented, "method AddArtifacts not implemented")
}
func (UnimplementedSparkConnectServiceServer) ArtifactStatus(context.Context, *ArtifactStatusesRequest) (*ArtifactStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArtifactStatus not implemented")
}
func (UnimplementedSparkConnectServiceServer) Interrupt(context.Context, *InterruptRequest) (*InterruptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interrupt not implemented")
}
func (UnimplementedSparkConnectServiceServer) ReattachExecute(*ReattachExecuteRequest, SparkConnectService_ReattachExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method ReattachExecute not implemented")
}
func (UnimplementedSparkConnectServiceServer) ReleaseExecute(context.Context, *ReleaseExecuteRequest) (*ReleaseExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseExecute not implemented")
}
func (UnimplementedSparkConnectServiceServer) mustEmbedUnimplementedSparkConnectServiceServer() {}

// UnsafeSparkConnectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SparkConnectServiceServer will
// result in compilation errors.
type UnsafeSparkConnectServiceServer interface {
	mustEmbedUnimplementedSparkConnectServiceServer()
}

func RegisterSparkConnectServiceServer(s grpc.ServiceRegistrar, srv SparkConnectServiceServer) {
	s.RegisterService(&SparkConnectService_ServiceDesc, srv)
}

func _SparkConnectService_ExecutePlan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecutePlanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SparkConnectServiceServer).ExecutePlan(m, &sparkConnectServiceExecutePlanServer{stream})
}

type SparkConnectService_ExecutePlanServer interface {
	Send(*ExecutePlanResponse) error
	grpc.ServerStream
}

type sparkConnectServiceExecutePlanServer struct {
	grpc.ServerStream
}

func (x *sparkConnectServiceExecutePlanServer) Send(m *ExecutePlanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SparkConnectService_AnalyzePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkConnectServiceServer).AnalyzePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkConnectService_AnalyzePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkConnectServiceServer).AnalyzePlan(ctx, req.(*AnalyzePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparkConnectService_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkConnectServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkConnectService_Config_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkConnectServiceServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparkConnectService_AddArtifacts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SparkConnectServiceServer).AddArtifacts(&sparkConnectServiceAddArtifactsServer{stream})
}

type SparkConnectService_AddArtifactsServer interface {
	SendAndClose(*AddArtifactsResponse) error
	Recv() (*AddArtifactsRequest, error)
	grpc.ServerStream
}

type sparkConnectServiceAddArtifactsServer struct {
	grpc.ServerStream
}

func (x *sparkConnectServiceAddArtifactsServer) SendAndClose(m *AddArtifactsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sparkConnectServiceAddArtifactsServer) Recv() (*AddArtifactsRequest, error) {
	m := new(AddArtifactsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SparkConnectService_ArtifactStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkConnectServiceServer).ArtifactStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkConnectService_ArtifactStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkConnectServiceServer).ArtifactStatus(ctx, req.(*ArtifactStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparkConnectService_Interrupt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterruptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkConnectServiceServer).Interrupt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkConnectService_Interrupt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkConnectServiceServer).Interrupt(ctx, req.(*InterruptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SparkConnectService_ReattachExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReattachExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SparkConnectServiceServer).ReattachExecute(m, &sparkConnectServiceReattachExecuteServer{stream})
}

type SparkConnectService_ReattachExecuteServer interface {
	Send(*ExecutePlanResponse) error
	grpc.ServerStream
}

type sparkConnectServiceReattachExecuteServer struct {
	grpc.ServerStream
}

func (x *sparkConnectServiceReattachExecuteServer) Send(m *ExecutePlanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SparkConnectService_ReleaseExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkConnectServiceServer).ReleaseExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SparkConnectService_ReleaseExecute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkConnectServiceServer).ReleaseExecute(ctx, req.(*ReleaseExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SparkConnectService_ServiceDesc is the grpc.ServiceDesc for SparkConnectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SparkConnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spark.connect.SparkConnectService",
	HandlerType: (*SparkConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyzePlan",
			Handler:    _SparkConnectService_AnalyzePlan_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _SparkConnectService_Config_Handler,
		},
		{
			MethodName: "ArtifactStatus",
			Handler:    _SparkConnectService_ArtifactStatus_Handler,
		},
		{
			MethodName: "Interrupt",
			Handler:    _SparkConnectService_Interrupt_Handler,
		},
		{
			MethodName: "ReleaseExecute",
			Handler:    _SparkConnectService_ReleaseExecute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecutePlan",
			Handler:       _SparkConnectService_ExecutePlan_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddArtifacts",
			Handler:       _SparkConnectService_AddArtifacts_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReattachExecute",
			Handler:       _SparkConnectService_ReattachExecute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spark/connect/base.proto",
}
