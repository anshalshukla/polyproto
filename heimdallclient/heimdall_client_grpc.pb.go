// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: heimdallClient/heimdall_client.proto

package heimdallclient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HeimdallClientClient is the client API for HeimdallClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeimdallClientClient interface {
	Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error)
	StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (HeimdallClient_StateSyncEventsClient, error)
	FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error)
}

type heimdallClientClient struct {
	cc grpc.ClientConnInterface
}

func NewHeimdallClientClient(cc grpc.ClientConnInterface) HeimdallClientClient {
	return &heimdallClientClient{cc}
}

func (c *heimdallClientClient) Span(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*SpanResponse, error) {
	out := new(SpanResponse)
	err := c.cc.Invoke(ctx, "/heimdall_client.HeimdallClient/Span", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClientClient) StateSyncEvents(ctx context.Context, in *StateSyncEventsRequest, opts ...grpc.CallOption) (HeimdallClient_StateSyncEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &HeimdallClient_ServiceDesc.Streams[0], "/heimdall_client.HeimdallClient/StateSyncEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &heimdallClientStateSyncEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HeimdallClient_StateSyncEventsClient interface {
	Recv() (*StateSyncEventsResponse, error)
	grpc.ClientStream
}

type heimdallClientStateSyncEventsClient struct {
	grpc.ClientStream
}

func (x *heimdallClientStateSyncEventsClient) Recv() (*StateSyncEventsResponse, error) {
	m := new(StateSyncEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *heimdallClientClient) FetchCheckpoint(ctx context.Context, in *FetchCheckpointRequest, opts ...grpc.CallOption) (*FetchCheckpointResponse, error) {
	out := new(FetchCheckpointResponse)
	err := c.cc.Invoke(ctx, "/heimdall_client.HeimdallClient/FetchCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClientClient) FetchCheckpointCount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FetchCheckpointCountResponse, error) {
	out := new(FetchCheckpointCountResponse)
	err := c.cc.Invoke(ctx, "/heimdall_client.HeimdallClient/FetchCheckpointCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeimdallClientServer is the server API for HeimdallClient service.
// All implementations must embed UnimplementedHeimdallClientServer
// for forward compatibility
type HeimdallClientServer interface {
	Span(context.Context, *SpanRequest) (*SpanResponse, error)
	StateSyncEvents(*StateSyncEventsRequest, HeimdallClient_StateSyncEventsServer) error
	FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error)
	FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error)
	mustEmbedUnimplementedHeimdallClientServer()
}

// UnimplementedHeimdallClientServer must be embedded to have forward compatible implementations.
type UnimplementedHeimdallClientServer struct {
}

func (UnimplementedHeimdallClientServer) Span(context.Context, *SpanRequest) (*SpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Span not implemented")
}
func (UnimplementedHeimdallClientServer) StateSyncEvents(*StateSyncEventsRequest, HeimdallClient_StateSyncEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StateSyncEvents not implemented")
}
func (UnimplementedHeimdallClientServer) FetchCheckpoint(context.Context, *FetchCheckpointRequest) (*FetchCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpoint not implemented")
}
func (UnimplementedHeimdallClientServer) FetchCheckpointCount(context.Context, *emptypb.Empty) (*FetchCheckpointCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCheckpointCount not implemented")
}
func (UnimplementedHeimdallClientServer) mustEmbedUnimplementedHeimdallClientServer() {}

// UnsafeHeimdallClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeimdallClientServer will
// result in compilation errors.
type UnsafeHeimdallClientServer interface {
	mustEmbedUnimplementedHeimdallClientServer()
}

func RegisterHeimdallClientServer(s grpc.ServiceRegistrar, srv HeimdallClientServer) {
	s.RegisterService(&HeimdallClient_ServiceDesc, srv)
}

func _HeimdallClient_Span_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallClientServer).Span(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall_client.HeimdallClient/Span",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallClientServer).Span(ctx, req.(*SpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallClient_StateSyncEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateSyncEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HeimdallClientServer).StateSyncEvents(m, &heimdallClientStateSyncEventsServer{stream})
}

type HeimdallClient_StateSyncEventsServer interface {
	Send(*StateSyncEventsResponse) error
	grpc.ServerStream
}

type heimdallClientStateSyncEventsServer struct {
	grpc.ServerStream
}

func (x *heimdallClientStateSyncEventsServer) Send(m *StateSyncEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HeimdallClient_FetchCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallClientServer).FetchCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall_client.HeimdallClient/FetchCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallClientServer).FetchCheckpoint(ctx, req.(*FetchCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallClient_FetchCheckpointCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallClientServer).FetchCheckpointCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall_client.HeimdallClient/FetchCheckpointCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallClientServer).FetchCheckpointCount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HeimdallClient_ServiceDesc is the grpc.ServiceDesc for HeimdallClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeimdallClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall_client.HeimdallClient",
	HandlerType: (*HeimdallClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Span",
			Handler:    _HeimdallClient_Span_Handler,
		},
		{
			MethodName: "FetchCheckpoint",
			Handler:    _HeimdallClient_FetchCheckpoint_Handler,
		},
		{
			MethodName: "FetchCheckpointCount",
			Handler:    _HeimdallClient_FetchCheckpointCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StateSyncEvents",
			Handler:       _HeimdallClient_StateSyncEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "heimdallClient/heimdall_client.proto",
}
