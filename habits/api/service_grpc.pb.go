
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: service.proto

package api

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
	Habits_CreateHabit_FullMethodName = "/habits.Habits/CreateHabit"
)

// HabitsClient is the client API for Habits service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HabitsClient interface {
	// CreateHabit is the endpoint that registers a habit.
	CreateHabit(ctx context.Context, in *CreateHabitRequest, opts ...grpc.CallOption) (*CreateHabitResponse, error)
}

type habitsClient struct {
	cc grpc.ClientConnInterface
}

func NewHabitsClient(cc grpc.ClientConnInterface) HabitsClient {
	return &habitsClient{cc}
}

func (c *habitsClient) CreateHabit(ctx context.Context, in *CreateHabitRequest, opts ...grpc.CallOption) (*CreateHabitResponse, error) {
	out := new(CreateHabitResponse)
	err := c.cc.Invoke(ctx, Habits_CreateHabit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HabitsServer is the server API for Habits service.
// All implementations must embed UnimplementedHabitsServer
// for forward compatibility
type HabitsServer interface {
	// CreateHabit is the endpoint that registers a habit.
	CreateHabit(context.Context, *CreateHabitRequest) (*CreateHabitResponse, error)
	mustEmbedUnimplementedHabitsServer()
}

// UnimplementedHabitsServer must be embedded to have forward compatible implementations.
type UnimplementedHabitsServer struct {
}

func (UnimplementedHabitsServer) CreateHabit(context.Context, *CreateHabitRequest) (*CreateHabitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHabit not implemented")
}
func (UnimplementedHabitsServer) mustEmbedUnimplementedHabitsServer() {}

// UnsafeHabitsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HabitsServer will
// result in compilation errors.
type UnsafeHabitsServer interface {
	mustEmbedUnimplementedHabitsServer()
}

func RegisterHabitsServer(s grpc.ServiceRegistrar, srv HabitsServer) {
	s.RegisterService(&Habits_ServiceDesc, srv)
}

func _Habits_CreateHabit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHabitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HabitsServer).CreateHabit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Habits_CreateHabit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HabitsServer).CreateHabit(ctx, req.(*CreateHabitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Habits_ServiceDesc is the grpc.ServiceDesc for Habits service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Habits_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "habits.Habits",
	HandlerType: (*HabitsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHabit",
			Handler:    _Habits_CreateHabit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
