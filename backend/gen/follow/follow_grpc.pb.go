// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: follow.proto

package follow

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
	FollowService_LookupRelationships_FullMethodName = "/follow.FollowService/LookupRelationships"
)

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	LookupRelationships(ctx context.Context, in *LookupRelationshipRequest, opts ...grpc.CallOption) (*RelationshipResponseList, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) LookupRelationships(ctx context.Context, in *LookupRelationshipRequest, opts ...grpc.CallOption) (*RelationshipResponseList, error) {
	out := new(RelationshipResponseList)
	err := c.cc.Invoke(ctx, FollowService_LookupRelationships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	LookupRelationships(context.Context, *LookupRelationshipRequest) (*RelationshipResponseList, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) LookupRelationships(context.Context, *LookupRelationshipRequest) (*RelationshipResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupRelationships not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_LookupRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).LookupRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_LookupRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).LookupRelationships(ctx, req.(*LookupRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupRelationships",
			Handler:    _FollowService_LookupRelationships_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
