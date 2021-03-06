// Code generated by protoc-gen-go. DO NOT EDIT.
// source: btdxcx.com/micro/account-srv/proto/account/account.proto

/*
Package com_btdxcx_micro_srv_account is a generated protocol buffer package.

It is generated from these files:
	btdxcx.com/micro/account-srv/proto/account/account.proto

It has these top-level messages:
	Record
	ReadRequest
	ReadResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
*/
package com_btdxcx_micro_srv_account

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

type Record struct {
	// uuid
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// service or user
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// service name, username, etc
	ClientId string `protobuf:"bytes,3,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// we leave this blank in responses
	// used for update and create
	ClientSecret string            `protobuf:"bytes,4,opt,name=client_secret,json=clientSecret" json:"client_secret,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,5,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// unix timestamp
	Created int64 `protobuf:"varint,6,opt,name=created" json:"created,omitempty"`
	// unix timestamp
	Updated int64 `protobuf:"varint,7,opt,name=updated" json:"updated,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Record) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Record) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Record) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Record) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type ReadRequest struct {
	ShopId   string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *ReadRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type ReadResponse struct {
	Account *Record `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadResponse) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateRequest struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	// If id is blank, one will be generated
	Account *Record `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *CreateRequest) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UpdateRequest struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	// Id and client_id cannot be changed
	Account *Record `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *UpdateRequest) GetAccount() *Record {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteRequest struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type SearchRequest struct {
	ShopId   string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Limit    int64  `protobuf:"varint,4,opt,name=limit" json:"limit"`
	Offset   int64  `protobuf:"varint,5,opt,name=offset" json:"offset"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SearchRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *SearchRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *SearchRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SearchResponse struct {
	Accounts []*Record `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SearchResponse) GetAccounts() []*Record {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "com.btdxcx.micro.srv.account.Record")
	proto.RegisterType((*ReadRequest)(nil), "com.btdxcx.micro.srv.account.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "com.btdxcx.micro.srv.account.ReadResponse")
	proto.RegisterType((*CreateRequest)(nil), "com.btdxcx.micro.srv.account.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "com.btdxcx.micro.srv.account.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "com.btdxcx.micro.srv.account.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "com.btdxcx.micro.srv.account.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "com.btdxcx.micro.srv.account.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "com.btdxcx.micro.srv.account.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "com.btdxcx.micro.srv.account.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "com.btdxcx.micro.srv.account.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.btdxcx.micro.srv.account"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Account{hdlr}, opts...))
}

type Account struct {
	AccountHandler
}

func (h *Account) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *Account) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *Account) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *Account) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.AccountHandler.Search(ctx, in, out)
}

func init() {
	proto.RegisterFile("btdxcx.com/micro/account-srv/proto/account/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0x23, 0xc9, 0x96, 0x9d, 0x71, 0x64, 0xc2, 0x12, 0x5a, 0xe1, 0xf6, 0x60, 0xd4, 0x1e,
	0xdc, 0x2f, 0x19, 0xdc, 0x8b, 0x69, 0xa1, 0xb4, 0xa4, 0x3d, 0xe4, 0xd0, 0x1c, 0x14, 0x7a, 0x2c,
	0x41, 0xd9, 0x9d, 0x60, 0x51, 0xcb, 0x52, 0x77, 0xd7, 0x26, 0xbe, 0x17, 0xfa, 0x38, 0x7d, 0xc5,
	0xa2, 0xfd, 0x10, 0x51, 0x0e, 0x91, 0x68, 0x21, 0x27, 0x6b, 0x66, 0x76, 0xe6, 0x37, 0x3b, 0xfb,
	0x1f, 0x0c, 0xcb, 0x2b, 0xc9, 0x6e, 0xe8, 0x4d, 0x4c, 0x8b, 0x7c, 0x9e, 0x67, 0x94, 0x17, 0xf3,
	0x94, 0xd2, 0x62, 0xbb, 0x91, 0x6f, 0x04, 0xdf, 0xcd, 0x4b, 0x5e, 0xc8, 0xda, 0x63, 0x7f, 0x63,
	0xe5, 0x25, 0x4f, 0x69, 0x91, 0xc7, 0x26, 0x5b, 0x65, 0xc6, 0x82, 0xef, 0x62, 0x73, 0x26, 0xfa,
	0xe3, 0x82, 0x9f, 0x20, 0x2d, 0x38, 0x23, 0x63, 0x70, 0x33, 0x16, 0x3a, 0x53, 0x67, 0x76, 0x98,
	0xb8, 0x19, 0x23, 0x04, 0x7a, 0x72, 0x5f, 0x62, 0xe8, 0x2a, 0x8f, 0xfa, 0x26, 0x4f, 0xe0, 0x90,
	0xae, 0x33, 0xdc, 0xc8, 0xcb, 0x8c, 0x85, 0x9e, 0x0a, 0x0c, 0xb5, 0xe3, 0x8c, 0x91, 0x67, 0x10,
	0x98, 0xa0, 0x40, 0xca, 0x51, 0x86, 0x3d, 0x75, 0xe0, 0x48, 0x3b, 0x2f, 0x94, 0x8f, 0x9c, 0xc3,
	0x30, 0x47, 0x99, 0xb2, 0x54, 0xa6, 0x61, 0x7f, 0xea, 0xcd, 0x46, 0x8b, 0x45, 0x7c, 0x5f, 0x87,
	0xb1, 0xee, 0x2e, 0xfe, 0x6a, 0x92, 0xbe, 0x6c, 0x24, 0xdf, 0x27, 0x75, 0x0d, 0x12, 0xc2, 0x80,
	0x72, 0x4c, 0x25, 0xb2, 0xd0, 0x9f, 0x3a, 0x33, 0x2f, 0xb1, 0x66, 0x15, 0xd9, 0x96, 0x4c, 0x45,
	0x06, 0x3a, 0x62, 0xcc, 0xc9, 0x7b, 0x08, 0x1a, 0xe5, 0xc8, 0x31, 0x78, 0x3f, 0x70, 0x6f, 0xee,
	0x5e, 0x7d, 0x92, 0x13, 0xe8, 0xef, 0xd2, 0xf5, 0xd6, 0xde, 0x5e, 0x1b, 0xef, 0xdc, 0xa5, 0x13,
	0x9d, 0xc2, 0x28, 0xc1, 0x94, 0x25, 0xf8, 0x73, 0x8b, 0x42, 0x92, 0xc7, 0x30, 0x10, 0xab, 0xa2,
	0xbc, 0xac, 0x47, 0xe7, 0x57, 0xe6, 0x19, 0x6b, 0x8e, 0xca, 0x6d, 0x8e, 0x2a, 0x3a, 0x87, 0x23,
	0x5d, 0x44, 0x94, 0xc5, 0x46, 0x20, 0xf9, 0x00, 0x03, 0x73, 0x5f, 0x55, 0x65, 0xb4, 0x78, 0xde,
	0x65, 0x28, 0x89, 0x4d, 0x8a, 0x56, 0x10, 0x9c, 0xaa, 0x6b, 0xb7, 0xb6, 0x75, 0x8b, 0xe4, 0xfe,
	0x0b, 0xe9, 0x18, 0xc6, 0x96, 0xa4, 0x7b, 0xaf, 0xd8, 0xdf, 0xd4, 0x60, 0x1f, 0x82, 0x6d, 0x49,
	0x86, 0xbd, 0x84, 0xe0, 0x33, 0xae, 0xb1, 0x03, 0x5b, 0xab, 0xdb, 0xb5, 0xea, 0xae, 0x6a, 0xd9,
	0x4c, 0x53, 0xeb, 0xb7, 0x03, 0xc1, 0x05, 0xa6, 0x9c, 0xae, 0xfe, 0xeb, 0x6d, 0xeb, 0xbd, 0xf1,
	0x6e, 0xed, 0xcd, 0x09, 0xf4, 0xd7, 0x59, 0x9e, 0xe9, 0x95, 0xf0, 0x12, 0x6d, 0x90, 0x47, 0xe0,
	0x17, 0xd7, 0xd7, 0x02, 0x65, 0xd8, 0x57, 0x6e, 0x63, 0x45, 0x09, 0x8c, 0x6d, 0x23, 0x46, 0x1f,
	0x1f, 0x61, 0x68, 0x86, 0x20, 0x42, 0x47, 0x6d, 0x4d, 0xb7, 0xd1, 0xd5, 0x59, 0x8b, 0x5f, 0x3d,
	0x18, 0x7c, 0xd2, 0x06, 0xf9, 0x0e, 0xbd, 0x4a, 0x7d, 0xe4, 0x45, 0x5b, 0x8d, 0x5a, 0xe6, 0x93,
	0x97, 0x5d, 0x8e, 0x9a, 0x66, 0x11, 0x7c, 0x2d, 0x11, 0xf2, 0xea, 0xfe, 0xac, 0x86, 0x64, 0x27,
	0xaf, 0xbb, 0x1d, 0x36, 0xaf, 0x75, 0x50, 0x61, 0xb4, 0x1a, 0xda, 0x30, 0x0d, 0x75, 0xb6, 0x61,
	0xee, 0x08, 0x4c, 0x61, 0xb4, 0x50, 0xda, 0x30, 0x0d, 0x21, 0xb6, 0x61, 0xee, 0x68, 0x4f, 0x61,
	0xf4, 0x9b, 0xb7, 0x61, 0x1a, 0x12, 0x6d, 0xc3, 0x34, 0x65, 0x14, 0x1d, 0x5c, 0xf9, 0xea, 0x4f,
	0xe1, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x89, 0xca, 0x5b, 0x50, 0x06, 0x00, 0x00,
}
