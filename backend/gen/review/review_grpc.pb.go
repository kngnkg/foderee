// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: review.proto

package review

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

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	ListReviews(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ReviewListReply, error)
	GetReviewById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*ReviewReply, error)
	CreateReview(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ReviewReply, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) ListReviews(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ReviewListReply, error) {
	out := new(ReviewListReply)
	err := c.cc.Invoke(ctx, "/review.ReviewService/ListReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetReviewById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*ReviewReply, error) {
	out := new(ReviewReply)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetReviewById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) CreateReview(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*ReviewReply, error) {
	out := new(ReviewReply)
	err := c.cc.Invoke(ctx, "/review.ReviewService/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	ListReviews(context.Context, *ListReviewsRequest) (*ReviewListReply, error)
	GetReviewById(context.Context, *GetByIdRequest) (*ReviewReply, error)
	CreateReview(context.Context, *CreateRequest) (*ReviewReply, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) ListReviews(context.Context, *ListReviewsRequest) (*ReviewListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviews not implemented")
}
func (UnimplementedReviewServiceServer) GetReviewById(context.Context, *GetByIdRequest) (*ReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviewById not implemented")
}
func (UnimplementedReviewServiceServer) CreateReview(context.Context, *CreateRequest) (*ReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_ListReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/ListReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListReviews(ctx, req.(*ListReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReviewById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReviewById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetReviewById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReviewById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReview(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListReviews",
			Handler:    _ReviewService_ListReviews_Handler,
		},
		{
			MethodName: "GetReviewById",
			Handler:    _ReviewService_GetReviewById_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _ReviewService_CreateReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}