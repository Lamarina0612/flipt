// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: data/data.proto

package data

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
	DataService_EvaluationSnapshotNamespace_FullMethodName = "/flipt.data.DataService/EvaluationSnapshotNamespace"
)

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	EvaluationSnapshotNamespace(ctx context.Context, in *EvaluationNamespaceSnapshotRequest, opts ...grpc.CallOption) (*EvaluationNamespaceSnapshot, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) EvaluationSnapshotNamespace(ctx context.Context, in *EvaluationNamespaceSnapshotRequest, opts ...grpc.CallOption) (*EvaluationNamespaceSnapshot, error) {
	out := new(EvaluationNamespaceSnapshot)
	err := c.cc.Invoke(ctx, DataService_EvaluationSnapshotNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	EvaluationSnapshotNamespace(context.Context, *EvaluationNamespaceSnapshotRequest) (*EvaluationNamespaceSnapshot, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) EvaluationSnapshotNamespace(context.Context, *EvaluationNamespaceSnapshotRequest) (*EvaluationNamespaceSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluationSnapshotNamespace not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_EvaluationSnapshotNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationNamespaceSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).EvaluationSnapshotNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_EvaluationSnapshotNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).EvaluationSnapshotNamespace(ctx, req.(*EvaluationNamespaceSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.data.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EvaluationSnapshotNamespace",
			Handler:    _DataService_EvaluationSnapshotNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data/data.proto",
}
