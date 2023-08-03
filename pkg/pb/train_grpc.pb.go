// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// TrainManagementClient is the client API for TrainManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainManagementClient interface {
	AddTrain(ctx context.Context, in *AddTrainRequest, opts ...grpc.CallOption) (*AddTrainResponse, error)
	AddStation(ctx context.Context, in *AddStationRequest, opts ...grpc.CallOption) (*AddStationResponse, error)
	AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*AddRouteResponse, error)
}

type trainManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainManagementClient(cc grpc.ClientConnInterface) TrainManagementClient {
	return &trainManagementClient{cc}
}

func (c *trainManagementClient) AddTrain(ctx context.Context, in *AddTrainRequest, opts ...grpc.CallOption) (*AddTrainResponse, error) {
	out := new(AddTrainResponse)
	err := c.cc.Invoke(ctx, "/Train.TrainManagement/AddTrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainManagementClient) AddStation(ctx context.Context, in *AddStationRequest, opts ...grpc.CallOption) (*AddStationResponse, error) {
	out := new(AddStationResponse)
	err := c.cc.Invoke(ctx, "/Train.TrainManagement/AddStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainManagementClient) AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*AddRouteResponse, error) {
	out := new(AddRouteResponse)
	err := c.cc.Invoke(ctx, "/Train.TrainManagement/AddRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainManagementServer is the server API for TrainManagement service.
// All implementations must embed UnimplementedTrainManagementServer
// for forward compatibility
type TrainManagementServer interface {
	AddTrain(context.Context, *AddTrainRequest) (*AddTrainResponse, error)
	AddStation(context.Context, *AddStationRequest) (*AddStationResponse, error)
	AddRoute(context.Context, *AddRouteRequest) (*AddRouteResponse, error)
	mustEmbedUnimplementedTrainManagementServer()
}

// UnimplementedTrainManagementServer must be embedded to have forward compatible implementations.
type UnimplementedTrainManagementServer struct {
}

func (UnimplementedTrainManagementServer) AddTrain(context.Context, *AddTrainRequest) (*AddTrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrain not implemented")
}
func (UnimplementedTrainManagementServer) AddStation(context.Context, *AddStationRequest) (*AddStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStation not implemented")
}
func (UnimplementedTrainManagementServer) AddRoute(context.Context, *AddRouteRequest) (*AddRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute not implemented")
}
func (UnimplementedTrainManagementServer) mustEmbedUnimplementedTrainManagementServer() {}

// UnsafeTrainManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainManagementServer will
// result in compilation errors.
type UnsafeTrainManagementServer interface {
	mustEmbedUnimplementedTrainManagementServer()
}

func RegisterTrainManagementServer(s grpc.ServiceRegistrar, srv TrainManagementServer) {
	s.RegisterService(&TrainManagement_ServiceDesc, srv)
}

func _TrainManagement_AddTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainManagementServer).AddTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Train.TrainManagement/AddTrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainManagementServer).AddTrain(ctx, req.(*AddTrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainManagement_AddStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainManagementServer).AddStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Train.TrainManagement/AddStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainManagementServer).AddStation(ctx, req.(*AddStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainManagement_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainManagementServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Train.TrainManagement/AddRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainManagementServer).AddRoute(ctx, req.(*AddRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainManagement_ServiceDesc is the grpc.ServiceDesc for TrainManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Train.TrainManagement",
	HandlerType: (*TrainManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTrain",
			Handler:    _TrainManagement_AddTrain_Handler,
		},
		{
			MethodName: "AddStation",
			Handler:    _TrainManagement_AddStation_Handler,
		},
		{
			MethodName: "AddRoute",
			Handler:    _TrainManagement_AddRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/train.proto",
}
