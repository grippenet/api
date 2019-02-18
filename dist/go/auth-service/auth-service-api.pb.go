// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth-service-api.proto

package auth_service

import (
	context "context"
	fmt "fmt"
	_go "github.com/Influenzanet/api/dist/go"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EncodedToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodedToken) Reset()         { *m = EncodedToken{} }
func (m *EncodedToken) String() string { return proto.CompactTextString(m) }
func (*EncodedToken) ProtoMessage()    {}
func (*EncodedToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8e86308594dc230, []int{0}
}

func (m *EncodedToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodedToken.Unmarshal(m, b)
}
func (m *EncodedToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodedToken.Marshal(b, m, deterministic)
}
func (m *EncodedToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodedToken.Merge(m, src)
}
func (m *EncodedToken) XXX_Size() int {
	return xxx_messageInfo_EncodedToken.Size(m)
}
func (m *EncodedToken) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodedToken.DiscardUnknown(m)
}

var xxx_messageInfo_EncodedToken proto.InternalMessageInfo

func (m *EncodedToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*EncodedToken)(nil), "influenzanet.auth_service_api.EncodedToken")
}

func init() { proto.RegisterFile("auth-service-api.proto", fileDescriptor_e8e86308594dc230) }

var fileDescriptor_e8e86308594dc230 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x4b, 0x42, 0x31,
	0x14, 0xc7, 0x91, 0x50, 0x6a, 0x45, 0x0f, 0x43, 0xa4, 0x6e, 0x08, 0x11, 0x3d, 0x04, 0xe1, 0x16,
	0x06, 0xbd, 0x5b, 0xf8, 0x50, 0xf4, 0x10, 0x5e, 0x2d, 0xe8, 0x45, 0x76, 0xbd, 0xc7, 0x79, 0x68,
	0x6e, 0xe3, 0xee, 0xac, 0xb0, 0xbf, 0xaa, 0x3f, 0x31, 0xf4, 0x2a, 0x78, 0x5f, 0x02, 0xa1, 0xb7,
	0x1d, 0xce, 0xe7, 0xfb, 0x63, 0x6c, 0xac, 0xa5, 0x22, 0xcd, 0x3a, 0x01, 0x8a, 0x4f, 0x9c, 0x40,
	0x47, 0x79, 0x14, 0xbe, 0x70, 0xe4, 0x78, 0x1b, 0xed, 0xd4, 0x44, 0xb0, 0xdf, 0xca, 0x02, 0x89,
	0x25, 0x34, 0x5e, 0x43, 0x63, 0xe5, 0x31, 0xe1, 0xda, 0xb8, 0x4c, 0x99, 0x0e, 0x2d, 0x3c, 0x84,
	0x52, 0x92, 0x9c, 0x69, 0xe7, 0xb4, 0x01, 0xb9, 0x9a, 0xb2, 0x38, 0x95, 0x30, 0xf7, 0xb4, 0x28,
	0x97, 0x17, 0x97, 0xec, 0xa8, 0x6f, 0x27, 0x2e, 0x87, 0x7c, 0xe8, 0x3e, 0xc0, 0xf2, 0x26, 0xab,
	0xd3, 0xf2, 0x70, 0x52, 0x3b, 0xaf, 0x5d, 0x1d, 0x0c, 0xca, 0xa1, 0xfb, 0xb3, 0xc7, 0x8e, 0x7b,
	0x91, 0x66, 0x69, 0x19, 0xd5, 0xf3, 0xc8, 0xef, 0x58, 0x23, 0x25, 0x45, 0x31, 0xf0, 0x96, 0x28,
	0x03, 0xc4, 0x26, 0x40, 0xf4, 0x97, 0x01, 0x49, 0x53, 0x54, 0xba, 0xae, 0xe9, 0x94, 0xd5, 0x9f,
	0x9d, 0x46, 0xcb, 0xdb, 0xd5, 0xf5, 0x28, 0x40, 0xf1, 0x50, 0x40, 0x0e, 0x96, 0x50, 0x99, 0x90,
	0x5c, 0x8b, 0x3f, 0x6f, 0x2a, 0x2a, 0xad, 0x87, 0xac, 0x91, 0xa2, 0xb6, 0xd1, 0xff, 0xab, 0xeb,
	0x88, 0x1d, 0xbe, 0x2a, 0x83, 0xb9, 0x22, 0x78, 0x7a, 0x1b, 0xf2, 0x5d, 0xb4, 0xc9, 0x69, 0x15,
	0x7e, 0x51, 0x45, 0xd8, 0xd8, 0xe6, 0x6c, 0x7f, 0x00, 0x16, 0xbe, 0x76, 0xf6, 0xdc, 0x05, 0xbe,
	0xef, 0xbe, 0xdf, 0x68, 0xa4, 0x59, 0xcc, 0xc4, 0xc4, 0xcd, 0xe5, 0xe3, 0x96, 0x50, 0x2a, 0x8f,
	0x32, 0xc7, 0x40, 0x52, 0x3b, 0xb9, 0xfd, 0xcd, 0xb2, 0xc6, 0xea, 0x05, 0x6f, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x9b, 0x2a, 0x0a, 0xe0, 0x7d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceApiClient is the client API for AuthServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error)
	// Auth:
	Login(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error)
	Signup(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error)
	RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error)
}

type authServiceApiClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceApiClient(cc *grpc.ClientConn) AuthServiceApiClient {
	return &authServiceApiClient{cc}
}

func (c *authServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*_go.Status, error) {
	out := new(_go.Status)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) Login(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) Signup(ctx context.Context, in *_go.UserCredentials, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) ValidateJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*_go.ParsedToken, error) {
	out := new(_go.ParsedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceApiClient) RenewJWT(ctx context.Context, in *EncodedToken, opts ...grpc.CallOption) (*EncodedToken, error) {
	out := new(EncodedToken)
	err := c.cc.Invoke(ctx, "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceApiServer is the server API for AuthServiceApi service.
type AuthServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*_go.Status, error)
	// Auth:
	Login(context.Context, *_go.UserCredentials) (*EncodedToken, error)
	Signup(context.Context, *_go.UserCredentials) (*EncodedToken, error)
	// Token handling:
	ValidateJWT(context.Context, *EncodedToken) (*_go.ParsedToken, error)
	RenewJWT(context.Context, *EncodedToken) (*EncodedToken, error)
}

func RegisterAuthServiceApiServer(s *grpc.Server, srv AuthServiceApiServer) {
	s.RegisterService(&_AuthServiceApi_serviceDesc, srv)
}

func _AuthServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).Login(ctx, req.(*_go.UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(_go.UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).Signup(ctx, req.(*_go.UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_ValidateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/ValidateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).ValidateJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceApi_RenewJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodedToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/influenzanet.auth_service_api.AuthServiceApi/RenewJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceApiServer).RenewJWT(ctx, req.(*EncodedToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "influenzanet.auth_service_api.AuthServiceApi",
	HandlerType: (*AuthServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AuthServiceApi_Status_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthServiceApi_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _AuthServiceApi_Signup_Handler,
		},
		{
			MethodName: "ValidateJWT",
			Handler:    _AuthServiceApi_ValidateJWT_Handler,
		},
		{
			MethodName: "RenewJWT",
			Handler:    _AuthServiceApi_RenewJWT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-service-api.proto",
}
