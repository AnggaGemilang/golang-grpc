// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: proto/vehicles.proto

package vehicle

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
	VehiclesCRUD_CreateVehicle_FullMethodName = "/grpc.VehiclesCRUD/CreateVehicle"
	VehiclesCRUD_DeleteVehicle_FullMethodName = "/grpc.VehiclesCRUD/DeleteVehicle"
	VehiclesCRUD_ListVehicles_FullMethodName  = "/grpc.VehiclesCRUD/ListVehicles"
	VehiclesCRUD_MostVehicles_FullMethodName  = "/grpc.VehiclesCRUD/MostVehicles"
)

// VehiclesCRUDClient is the client API for VehiclesCRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehiclesCRUDClient interface {
	CreateVehicle(ctx context.Context, in *VehicleInput, opts ...grpc.CallOption) (*VehicleResponse, error)
	DeleteVehicle(ctx context.Context, in *ID, opts ...grpc.CallOption) (*VehicleResponse, error)
	ListVehicles(ctx context.Context, in *ListVehicleReq, opts ...grpc.CallOption) (*ListVehicleRes, error)
	MostVehicles(ctx context.Context, in *MostVehicleReq, opts ...grpc.CallOption) (*ListMostVehicleRes, error)
}

type vehiclesCRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewVehiclesCRUDClient(cc grpc.ClientConnInterface) VehiclesCRUDClient {
	return &vehiclesCRUDClient{cc}
}

func (c *vehiclesCRUDClient) CreateVehicle(ctx context.Context, in *VehicleInput, opts ...grpc.CallOption) (*VehicleResponse, error) {
	out := new(VehicleResponse)
	err := c.cc.Invoke(ctx, VehiclesCRUD_CreateVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehiclesCRUDClient) DeleteVehicle(ctx context.Context, in *ID, opts ...grpc.CallOption) (*VehicleResponse, error) {
	out := new(VehicleResponse)
	err := c.cc.Invoke(ctx, VehiclesCRUD_DeleteVehicle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehiclesCRUDClient) ListVehicles(ctx context.Context, in *ListVehicleReq, opts ...grpc.CallOption) (*ListVehicleRes, error) {
	out := new(ListVehicleRes)
	err := c.cc.Invoke(ctx, VehiclesCRUD_ListVehicles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehiclesCRUDClient) MostVehicles(ctx context.Context, in *MostVehicleReq, opts ...grpc.CallOption) (*ListMostVehicleRes, error) {
	out := new(ListMostVehicleRes)
	err := c.cc.Invoke(ctx, VehiclesCRUD_MostVehicles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehiclesCRUDServer is the server API for VehiclesCRUD service.
// All implementations must embed UnimplementedVehiclesCRUDServer
// for forward compatibility
type VehiclesCRUDServer interface {
	CreateVehicle(context.Context, *VehicleInput) (*VehicleResponse, error)
	DeleteVehicle(context.Context, *ID) (*VehicleResponse, error)
	ListVehicles(context.Context, *ListVehicleReq) (*ListVehicleRes, error)
	MostVehicles(context.Context, *MostVehicleReq) (*ListMostVehicleRes, error)
	mustEmbedUnimplementedVehiclesCRUDServer()
}

// UnimplementedVehiclesCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedVehiclesCRUDServer struct {
}

func (UnimplementedVehiclesCRUDServer) CreateVehicle(context.Context, *VehicleInput) (*VehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehiclesCRUDServer) DeleteVehicle(context.Context, *ID) (*VehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVehicle not implemented")
}
func (UnimplementedVehiclesCRUDServer) ListVehicles(context.Context, *ListVehicleReq) (*ListVehicleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicles not implemented")
}
func (UnimplementedVehiclesCRUDServer) MostVehicles(context.Context, *MostVehicleReq) (*ListMostVehicleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostVehicles not implemented")
}
func (UnimplementedVehiclesCRUDServer) mustEmbedUnimplementedVehiclesCRUDServer() {}

// UnsafeVehiclesCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehiclesCRUDServer will
// result in compilation errors.
type UnsafeVehiclesCRUDServer interface {
	mustEmbedUnimplementedVehiclesCRUDServer()
}

func RegisterVehiclesCRUDServer(s grpc.ServiceRegistrar, srv VehiclesCRUDServer) {
	s.RegisterService(&VehiclesCRUD_ServiceDesc, srv)
}

func _VehiclesCRUD_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VehicleInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehiclesCRUDServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehiclesCRUD_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehiclesCRUDServer).CreateVehicle(ctx, req.(*VehicleInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehiclesCRUD_DeleteVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehiclesCRUDServer).DeleteVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehiclesCRUD_DeleteVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehiclesCRUDServer).DeleteVehicle(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehiclesCRUD_ListVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehiclesCRUDServer).ListVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehiclesCRUD_ListVehicles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehiclesCRUDServer).ListVehicles(ctx, req.(*ListVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehiclesCRUD_MostVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MostVehicleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehiclesCRUDServer).MostVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehiclesCRUD_MostVehicles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehiclesCRUDServer).MostVehicles(ctx, req.(*MostVehicleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VehiclesCRUD_ServiceDesc is the grpc.ServiceDesc for VehiclesCRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehiclesCRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.VehiclesCRUD",
	HandlerType: (*VehiclesCRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVehicle",
			Handler:    _VehiclesCRUD_CreateVehicle_Handler,
		},
		{
			MethodName: "DeleteVehicle",
			Handler:    _VehiclesCRUD_DeleteVehicle_Handler,
		},
		{
			MethodName: "ListVehicles",
			Handler:    _VehiclesCRUD_ListVehicles_Handler,
		},
		{
			MethodName: "MostVehicles",
			Handler:    _VehiclesCRUD_MostVehicles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vehicles.proto",
}
