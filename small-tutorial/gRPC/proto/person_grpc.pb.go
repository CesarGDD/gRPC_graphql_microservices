// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: person.proto

package tutorial

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PersonService_GetPerson_FullMethodName = "/tutorial.PersonService/GetPerson"
)

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonServiceClient interface {
	GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonResponse)
	err := c.cc.Invoke(ctx, PersonService_GetPerson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServiceServer is the server API for PersonService service.
// All implementations must embed UnimplementedPersonServiceServer
// for forward compatibility.
type PersonServiceServer interface {
	GetPerson(context.Context, *PersonRequest) (*PersonResponse, error)
	mustEmbedUnimplementedPersonServiceServer()
}

// UnimplementedPersonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPersonServiceServer struct{}

func (UnimplementedPersonServiceServer) GetPerson(context.Context, *PersonRequest) (*PersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedPersonServiceServer) mustEmbedUnimplementedPersonServiceServer() {}
func (UnimplementedPersonServiceServer) testEmbeddedByValue()                       {}

// UnsafePersonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonServiceServer will
// result in compilation errors.
type UnsafePersonServiceServer interface {
	mustEmbedUnimplementedPersonServiceServer()
}

func RegisterPersonServiceServer(s grpc.ServiceRegistrar, srv PersonServiceServer) {
	// If the following call pancis, it indicates UnimplementedPersonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PersonService_ServiceDesc, srv)
}

func _PersonService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonService_GetPerson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).GetPerson(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PersonService_ServiceDesc is the grpc.ServiceDesc for PersonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tutorial.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _PersonService_GetPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
