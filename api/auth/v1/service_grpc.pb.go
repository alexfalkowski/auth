// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: auth/v1/service.proto

package v1

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// GeneratePassword that is secure.
	GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratePasswordResponse, error)
	// GenerateKey public and private key based on kind.
	GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...grpc.CallOption) (*GenerateKeyResponse, error)
	// GetPublicKey from kind.
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
	// GenerateAccessToken from RSA keys.
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	// GenerateServiceToken from Ed25519 keys.
	GenerateServiceToken(ctx context.Context, in *GenerateServiceTokenRequest, opts ...grpc.CallOption) (*GenerateServiceTokenResponse, error)
	// VerifyServiceToken based on kind.
	VerifyServiceToken(ctx context.Context, in *VerifyServiceTokenRequest, opts ...grpc.CallOption) (*VerifyServiceTokenResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratePasswordResponse, error) {
	out := new(GeneratePasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/GeneratePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GenerateKey(ctx context.Context, in *GenerateKeyRequest, opts ...grpc.CallOption) (*GenerateKeyResponse, error) {
	out := new(GenerateKeyResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/GenerateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GenerateServiceToken(ctx context.Context, in *GenerateServiceTokenRequest, opts ...grpc.CallOption) (*GenerateServiceTokenResponse, error) {
	out := new(GenerateServiceTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/GenerateServiceToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VerifyServiceToken(ctx context.Context, in *VerifyServiceTokenRequest, opts ...grpc.CallOption) (*VerifyServiceTokenResponse, error) {
	out := new(VerifyServiceTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.v1.Service/VerifyServiceToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// GeneratePassword that is secure.
	GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratePasswordResponse, error)
	// GenerateKey public and private key based on kind.
	GenerateKey(context.Context, *GenerateKeyRequest) (*GenerateKeyResponse, error)
	// GetPublicKey from kind.
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	// GenerateAccessToken from RSA keys.
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	// GenerateServiceToken from Ed25519 keys.
	GenerateServiceToken(context.Context, *GenerateServiceTokenRequest) (*GenerateServiceTokenResponse, error)
	// VerifyServiceToken based on kind.
	VerifyServiceToken(context.Context, *VerifyServiceTokenRequest) (*VerifyServiceTokenResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePassword not implemented")
}
func (UnimplementedServiceServer) GenerateKey(context.Context, *GenerateKeyRequest) (*GenerateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKey not implemented")
}
func (UnimplementedServiceServer) GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedServiceServer) GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAccessToken not implemented")
}
func (UnimplementedServiceServer) GenerateServiceToken(context.Context, *GenerateServiceTokenRequest) (*GenerateServiceTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateServiceToken not implemented")
}
func (UnimplementedServiceServer) VerifyServiceToken(context.Context, *VerifyServiceTokenRequest) (*VerifyServiceTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyServiceToken not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GeneratePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GeneratePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/GeneratePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GeneratePassword(ctx, req.(*GeneratePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/GenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GenerateKey(ctx, req.(*GenerateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GenerateServiceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateServiceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GenerateServiceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/GenerateServiceToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GenerateServiceToken(ctx, req.(*GenerateServiceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_VerifyServiceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyServiceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).VerifyServiceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1.Service/VerifyServiceToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).VerifyServiceToken(ctx, req.(*VerifyServiceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePassword",
			Handler:    _Service_GeneratePassword_Handler,
		},
		{
			MethodName: "GenerateKey",
			Handler:    _Service_GenerateKey_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Service_GetPublicKey_Handler,
		},
		{
			MethodName: "GenerateAccessToken",
			Handler:    _Service_GenerateAccessToken_Handler,
		},
		{
			MethodName: "GenerateServiceToken",
			Handler:    _Service_GenerateServiceToken_Handler,
		},
		{
			MethodName: "VerifyServiceToken",
			Handler:    _Service_VerifyServiceToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/service.proto",
}
