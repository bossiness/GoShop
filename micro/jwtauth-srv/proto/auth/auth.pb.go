// Code generated by protoc-gen-go. DO NOT EDIT.
// source: btdxcx.com/micro/jwtauth-srv/proto/auth/auth.proto

/*
Package com_btdxcx_micro_srv_jwtauth is a generated protocol buffer package.

It is generated from these files:
	btdxcx.com/micro/jwtauth-srv/proto/auth/auth.proto

It has these top-level messages:
	Token
	Record
	TokenRequest
	TokenResponse
	RevokeRequest
	RevokeResponse
	IntrospectRequest
	IntrospectResponse
	RefreshRequest
	RefreshResponse
*/
package com_btdxcx_micro_srv_jwtauth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	// jwt
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	// jwt
	RefreshToken string            `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ExpiresAt    int64             `protobuf:"varint,3,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	Scopes       []string          `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Token) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Token) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Record struct {
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Jti          string `protobuf:"bytes,2,opt,name=jti" json:"jti,omitempty"`
	AccessToken  string `protobuf:"bytes,3,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Cipher       string `protobuf:"bytes,5,opt,name=cipher" json:"cipher,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Record) GetJti() string {
	if m != nil {
		return m.Jti
	}
	return ""
}

func (m *Record) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Record) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Record) GetCipher() string {
	if m != nil {
		return m.Cipher
	}
	return ""
}

type TokenRequest struct {
	ClientId      string            `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientSecrent string            `protobuf:"bytes,2,opt,name=client_secrent,json=clientSecrent" json:"client_secrent,omitempty"`
	Scopes        []string          `protobuf:"bytes,3,rep,name=scopes" json:"scopes,omitempty"`
	Metadata      map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ShopId        string            `protobuf:"bytes,5,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *TokenRequest) GetClientSecrent() string {
	if m != nil {
		return m.ClientSecrent
	}
	return ""
}

func (m *TokenRequest) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *TokenRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TokenRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type TokenResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type RevokeRequest struct {
	// revoke access token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	// revoke via refresh token
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ShopId       string `protobuf:"bytes,3,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *RevokeRequest) Reset()                    { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()               {}
func (*RevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RevokeRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RevokeRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RevokeRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type RevokeResponse struct {
}

func (m *RevokeResponse) Reset()                    { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string            { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()               {}
func (*RevokeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type IntrospectRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	ShopId      string `protobuf:"bytes,2,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *IntrospectRequest) Reset()                    { *m = IntrospectRequest{} }
func (m *IntrospectRequest) String() string            { return proto.CompactTextString(m) }
func (*IntrospectRequest) ProtoMessage()               {}
func (*IntrospectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IntrospectRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *IntrospectRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type IntrospectResponse struct {
	Token  *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Active bool   `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
}

func (m *IntrospectResponse) Reset()                    { *m = IntrospectResponse{} }
func (m *IntrospectResponse) String() string            { return proto.CompactTextString(m) }
func (*IntrospectResponse) ProtoMessage()               {}
func (*IntrospectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IntrospectResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *IntrospectResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type RefreshRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ShopId       string `protobuf:"bytes,2,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *RefreshRequest) Reset()                    { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()               {}
func (*RefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RefreshRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type RefreshResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *RefreshResponse) Reset()                    { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string            { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()               {}
func (*RefreshResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RefreshResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "com.btdxcx.micro.srv.jwtauth.Token")
	proto.RegisterType((*Record)(nil), "com.btdxcx.micro.srv.jwtauth.Record")
	proto.RegisterType((*TokenRequest)(nil), "com.btdxcx.micro.srv.jwtauth.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "com.btdxcx.micro.srv.jwtauth.TokenResponse")
	proto.RegisterType((*RevokeRequest)(nil), "com.btdxcx.micro.srv.jwtauth.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "com.btdxcx.micro.srv.jwtauth.RevokeResponse")
	proto.RegisterType((*IntrospectRequest)(nil), "com.btdxcx.micro.srv.jwtauth.IntrospectRequest")
	proto.RegisterType((*IntrospectResponse)(nil), "com.btdxcx.micro.srv.jwtauth.IntrospectResponse")
	proto.RegisterType((*RefreshRequest)(nil), "com.btdxcx.micro.srv.jwtauth.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "com.btdxcx.micro.srv.jwtauth.RefreshResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for JwtAuth service

type JwtAuthClient interface {
	Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error)
	Revoke(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error)
	Introspect(ctx context.Context, in *IntrospectRequest, opts ...client.CallOption) (*IntrospectResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...client.CallOption) (*RefreshResponse, error)
}

type jwtAuthClient struct {
	c           client.Client
	serviceName string
}

func NewJwtAuthClient(serviceName string, c client.Client) JwtAuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.btdxcx.micro.srv.jwtauth"
	}
	return &jwtAuthClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *jwtAuthClient) Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JwtAuth.Token", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtAuthClient) Revoke(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JwtAuth.Revoke", in)
	out := new(RevokeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtAuthClient) Introspect(ctx context.Context, in *IntrospectRequest, opts ...client.CallOption) (*IntrospectResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JwtAuth.Introspect", in)
	out := new(IntrospectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtAuthClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...client.CallOption) (*RefreshResponse, error) {
	req := c.c.NewRequest(c.serviceName, "JwtAuth.Refresh", in)
	out := new(RefreshResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JwtAuth service

type JwtAuthHandler interface {
	Token(context.Context, *TokenRequest, *TokenResponse) error
	Revoke(context.Context, *RevokeRequest, *RevokeResponse) error
	Introspect(context.Context, *IntrospectRequest, *IntrospectResponse) error
	Refresh(context.Context, *RefreshRequest, *RefreshResponse) error
}

func RegisterJwtAuthHandler(s server.Server, hdlr JwtAuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&JwtAuth{hdlr}, opts...))
}

type JwtAuth struct {
	JwtAuthHandler
}

func (h *JwtAuth) Token(ctx context.Context, in *TokenRequest, out *TokenResponse) error {
	return h.JwtAuthHandler.Token(ctx, in, out)
}

func (h *JwtAuth) Revoke(ctx context.Context, in *RevokeRequest, out *RevokeResponse) error {
	return h.JwtAuthHandler.Revoke(ctx, in, out)
}

func (h *JwtAuth) Introspect(ctx context.Context, in *IntrospectRequest, out *IntrospectResponse) error {
	return h.JwtAuthHandler.Introspect(ctx, in, out)
}

func (h *JwtAuth) Refresh(ctx context.Context, in *RefreshRequest, out *RefreshResponse) error {
	return h.JwtAuthHandler.Refresh(ctx, in, out)
}

func init() { proto.RegisterFile("btdxcx.com/micro/jwtauth-srv/proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0xf3, 0x35, 0x49, 0x4a, 0x59, 0xa1, 0x12, 0x05, 0x90, 0x82, 0x2b, 0xa4, 0x08,
	0xa8, 0x03, 0xe1, 0x52, 0xe0, 0xd4, 0x03, 0x87, 0x54, 0x14, 0x24, 0xd3, 0x7b, 0xe4, 0xac, 0xa7,
	0xc4, 0x4d, 0xe3, 0x75, 0x77, 0x37, 0x69, 0xfa, 0x37, 0x38, 0xc0, 0xcf, 0xe0, 0x2f, 0xa2, 0xfd,
	0x48, 0xea, 0x28, 0x90, 0x1a, 0x85, 0x4b, 0xe4, 0x99, 0x9d, 0xb7, 0xef, 0xcd, 0xbc, 0x89, 0x0d,
	0xbd, 0xa1, 0x8c, 0xe6, 0x74, 0xee, 0x53, 0x36, 0xe9, 0x4e, 0x62, 0xca, 0x59, 0xf7, 0xe2, 0x5a,
	0x86, 0x53, 0x39, 0x3a, 0x14, 0x7c, 0xd6, 0x4d, 0x39, 0x93, 0xac, 0xab, 0x42, 0xfd, 0xe3, 0xeb,
	0x98, 0x3c, 0xa1, 0x6c, 0xe2, 0x5b, 0x9c, 0xc6, 0xf8, 0x82, 0xcf, 0x7c, 0x8b, 0xf3, 0xbe, 0x17,
	0xa0, 0x78, 0xc6, 0xc6, 0x98, 0x90, 0x67, 0x50, 0x0f, 0x29, 0x45, 0x21, 0x06, 0x52, 0xc5, 0x4d,
	0xa7, 0xed, 0x74, 0xaa, 0x41, 0xcd, 0xe4, 0x4c, 0xc9, 0x01, 0x34, 0x38, 0x9e, 0x73, 0x14, 0x23,
	0x5b, 0x53, 0xd0, 0x35, 0x75, 0x9b, 0x34, 0x45, 0x4f, 0x01, 0x70, 0x9e, 0xc6, 0x1c, 0xc5, 0x20,
	0x94, 0x4d, 0xb7, 0xed, 0x74, 0xdc, 0xa0, 0x6a, 0x33, 0xc7, 0x92, 0xec, 0x43, 0x49, 0x50, 0x96,
	0xa2, 0x68, 0xee, 0xb4, 0xdd, 0x4e, 0x35, 0xb0, 0x11, 0x39, 0x85, 0xca, 0x04, 0x65, 0x18, 0x85,
	0x32, 0x6c, 0x16, 0xdb, 0x6e, 0xa7, 0xd6, 0x7b, 0xe3, 0x6f, 0x52, 0xee, 0x6b, 0x36, 0xff, 0xd4,
	0x62, 0x3e, 0x26, 0x92, 0xdf, 0x04, 0xcb, 0x2b, 0x5a, 0x1f, 0xa0, 0xb1, 0x72, 0x44, 0xf6, 0xc0,
	0x1d, 0xe3, 0x8d, 0xed, 0x4a, 0x3d, 0x92, 0x87, 0x50, 0x9c, 0x85, 0x97, 0x53, 0xb4, 0x5d, 0x98,
	0xe0, 0x7d, 0xe1, 0xc8, 0xf1, 0x7e, 0x3a, 0x50, 0x0a, 0x90, 0x32, 0x1e, 0x91, 0xc7, 0x50, 0xa5,
	0x97, 0x31, 0x26, 0x72, 0x10, 0x47, 0x16, 0x5c, 0x31, 0x89, 0x7e, 0xa4, 0xee, 0xbc, 0x90, 0xb1,
	0xc5, 0xab, 0xc7, 0xb5, 0x21, 0xba, 0x39, 0x86, 0xb8, 0xf3, 0x87, 0x21, 0xee, 0x43, 0x89, 0xc6,
	0xe9, 0x08, 0x79, 0xb3, 0xa8, 0x4f, 0x6d, 0xe4, 0xfd, 0x28, 0x40, 0x5d, 0x57, 0x04, 0x78, 0x35,
	0x45, 0x21, 0x37, 0xeb, 0x7b, 0x0e, 0xbb, 0xf6, 0x50, 0x20, 0xe5, 0x98, 0x48, 0x2b, 0xb5, 0x61,
	0xb2, 0x5f, 0x4d, 0x32, 0x63, 0x89, 0xbb, 0x62, 0xc9, 0x59, 0xc6, 0x92, 0x1d, 0x6d, 0xc9, 0x51,
	0x0e, 0x4b, 0xac, 0xb2, 0xbf, 0x39, 0x43, 0x1e, 0x41, 0x59, 0x8c, 0x58, 0xaa, 0xf4, 0xda, 0xde,
	0x54, 0xd8, 0x8f, 0xb6, 0xb3, 0xec, 0x04, 0x1a, 0x96, 0x5d, 0xa4, 0x2c, 0x11, 0x48, 0xde, 0x41,
	0xf1, 0x76, 0x8f, 0x6b, 0xbd, 0x83, 0x3c, 0xca, 0x0d, 0xc2, 0x4b, 0xa1, 0x11, 0xe0, 0x8c, 0x8d,
	0x71, 0x31, 0xe4, 0xff, 0xf5, 0xd7, 0xc8, 0xb4, 0xee, 0x66, 0x5b, 0xf7, 0xf6, 0x60, 0x77, 0xc1,
	0x68, 0xe4, 0x7b, 0x5f, 0xe0, 0x41, 0x3f, 0x91, 0x9c, 0x89, 0x14, 0xa9, 0xfc, 0x07, 0x1d, 0x19,
	0x8a, 0xc2, 0x0a, 0xc5, 0x37, 0x20, 0xd9, 0x0b, 0xb7, 0x9e, 0x92, 0xda, 0x9a, 0x90, 0xca, 0x78,
	0x66, 0xcc, 0xa8, 0x04, 0x36, 0xf2, 0x3e, 0xab, 0x5e, 0x74, 0xd3, 0x0b, 0xd9, 0x6b, 0xb3, 0x71,
	0x36, 0xcf, 0x66, 0x55, 0xf8, 0x27, 0xb8, 0xbf, 0xbc, 0x6f, 0x6b, 0xd5, 0xbd, 0x5f, 0x2e, 0x94,
	0x4f, 0xae, 0xe5, 0xf1, 0x54, 0x8e, 0xc8, 0x70, 0xf1, 0xea, 0x7b, 0x91, 0x7f, 0xad, 0x5b, 0x2f,
	0x73, 0xd5, 0x5a, 0x17, 0xef, 0x11, 0x54, 0x6f, 0x12, 0xe5, 0x2c, 0xb9, 0x03, 0xb8, 0xb2, 0x71,
	0xad, 0x57, 0xf9, 0x8a, 0x97, 0x34, 0x57, 0x00, 0xb7, 0xee, 0x92, 0xee, 0x66, 0xf4, 0xda, 0x62,
	0xb5, 0x5e, 0xe7, 0x07, 0x2c, 0x29, 0xcf, 0xa1, 0x6c, 0x7d, 0x21, 0x77, 0xaa, 0xcd, 0xae, 0x43,
	0xeb, 0x30, 0x67, 0xb5, 0x61, 0x1a, 0x96, 0xf4, 0x67, 0xec, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x0a, 0xd9, 0x39, 0xfc, 0x06, 0x00, 0x00,
}
