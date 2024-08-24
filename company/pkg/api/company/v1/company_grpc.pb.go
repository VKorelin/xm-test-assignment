// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.14.0
// source: company.proto

package company

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CompanyService_Get_FullMethodName   = "/xm.api.company.v1.CompanyService/Get"
	CompanyService_Patch_FullMethodName = "/xm.api.company.v1.CompanyService/Patch"
)

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	Get(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	Patch(ctx context.Context, in *PatchCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) Get(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Patch(ctx context.Context, in *PatchCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CompanyService_Patch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility.
type CompanyServiceServer interface {
	Get(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	Patch(context.Context, *PatchCompanyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyServiceServer struct{}

func (UnimplementedCompanyServiceServer) Get(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCompanyServiceServer) Patch(context.Context, *PatchCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Patch not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}
func (UnimplementedCompanyServiceServer) testEmbeddedByValue()                        {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Get(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_Patch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Patch(ctx, req.(*PatchCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xm.api.company.v1.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CompanyService_Get_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _CompanyService_Patch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company.proto",
}
