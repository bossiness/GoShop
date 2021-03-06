// Code generated by protoc-gen-go. DO NOT EDIT.
// source: btdxcx.com/micro/shop-srv/proto/shop/details/shop_details.proto

/*
Package com_btdxcx_micro_srv_shop_details is a generated protocol buffer package.

It is generated from these files:
	btdxcx.com/micro/shop-srv/proto/shop/details/shop_details.proto

It has these top-level messages:
	ShopDetails
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
	UpdateRequest
	UpdateResponse
*/
package com_btdxcx_micro_srv_shop_details

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

type State int32

const (
	State_untreated State = 0
	State_reviewing State = 1
	State_completed State = 2
)

var State_name = map[int32]string{
	0: "untreated",
	1: "reviewing",
	2: "completed",
}
var State_value = map[string]int32{
	"untreated": 0,
	"reviewing": 1,
	"completed": 2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ShopDetails struct {
	Name      string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Logo      string                     `protobuf:"bytes,2,opt,name=logo" json:"logo,omitempty"`
	Introduce string                     `protobuf:"bytes,3,opt,name=introduce" json:"introduce,omitempty"`
	Owner     *ShopDetails_ShopOwner     `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	Weixin    *ShopDetails_WeiXin        `protobuf:"bytes,5,opt,name=weixin" json:"weixin,omitempty"`
	Physical  *ShopDetails_PhysicalStore `protobuf:"bytes,6,opt,name=physical" json:"physical,omitempty"`
}

func (m *ShopDetails) Reset()                    { *m = ShopDetails{} }
func (m *ShopDetails) String() string            { return proto.CompactTextString(m) }
func (*ShopDetails) ProtoMessage()               {}
func (*ShopDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ShopDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShopDetails) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *ShopDetails) GetIntroduce() string {
	if m != nil {
		return m.Introduce
	}
	return ""
}

func (m *ShopDetails) GetOwner() *ShopDetails_ShopOwner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ShopDetails) GetWeixin() *ShopDetails_WeiXin {
	if m != nil {
		return m.Weixin
	}
	return nil
}

func (m *ShopDetails) GetPhysical() *ShopDetails_PhysicalStore {
	if m != nil {
		return m.Physical
	}
	return nil
}

type ShopDetails_ShopOwner struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
}

func (m *ShopDetails_ShopOwner) Reset()                    { *m = ShopDetails_ShopOwner{} }
func (m *ShopDetails_ShopOwner) String() string            { return proto.CompactTextString(m) }
func (*ShopDetails_ShopOwner) ProtoMessage()               {}
func (*ShopDetails_ShopOwner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ShopDetails_ShopOwner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ShopDetails_ShopOwner) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShopDetails_ShopOwner) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ShopDetails_ShopOwner) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ShopDetails_WeiXin struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	WechatId   string `protobuf:"bytes,2,opt,name=wechat_id,json=wechatId" json:"wechat_id,omitempty"`
	Appid      string `protobuf:"bytes,3,opt,name=appid" json:"appid,omitempty"`
	AppSecret  string `protobuf:"bytes,4,opt,name=app_secret,json=appSecret" json:"app_secret,omitempty"`
	PartnerId  string `protobuf:"bytes,5,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	PartnerKey string `protobuf:"bytes,6,opt,name=partner_key,json=partnerKey" json:"partner_key,omitempty"`
}

func (m *ShopDetails_WeiXin) Reset()                    { *m = ShopDetails_WeiXin{} }
func (m *ShopDetails_WeiXin) String() string            { return proto.CompactTextString(m) }
func (*ShopDetails_WeiXin) ProtoMessage()               {}
func (*ShopDetails_WeiXin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *ShopDetails_WeiXin) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ShopDetails_WeiXin) GetWechatId() string {
	if m != nil {
		return m.WechatId
	}
	return ""
}

func (m *ShopDetails_WeiXin) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *ShopDetails_WeiXin) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *ShopDetails_WeiXin) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *ShopDetails_WeiXin) GetPartnerKey() string {
	if m != nil {
		return m.PartnerKey
	}
	return ""
}

type ShopDetails_PhysicalStore struct {
	Id       string                              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string                              `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Contact  string                              `protobuf:"bytes,3,opt,name=contact" json:"contact,omitempty"`
	Email    string                              `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	ZipCode  string                              `protobuf:"bytes,5,opt,name=zipCode" json:"zipCode,omitempty"`
	Address  string                              `protobuf:"bytes,6,opt,name=address" json:"address,omitempty"`
	Location *ShopDetails_PhysicalStore_Location `protobuf:"bytes,7,opt,name=location" json:"location,omitempty"`
}

func (m *ShopDetails_PhysicalStore) Reset()                    { *m = ShopDetails_PhysicalStore{} }
func (m *ShopDetails_PhysicalStore) String() string            { return proto.CompactTextString(m) }
func (*ShopDetails_PhysicalStore) ProtoMessage()               {}
func (*ShopDetails_PhysicalStore) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *ShopDetails_PhysicalStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ShopDetails_PhysicalStore) GetLocation() *ShopDetails_PhysicalStore_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type ShopDetails_PhysicalStore_Location struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *ShopDetails_PhysicalStore_Location) Reset()         { *m = ShopDetails_PhysicalStore_Location{} }
func (m *ShopDetails_PhysicalStore_Location) String() string { return proto.CompactTextString(m) }
func (*ShopDetails_PhysicalStore_Location) ProtoMessage()    {}
func (*ShopDetails_PhysicalStore_Location) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0}
}

func (m *ShopDetails_PhysicalStore_Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ShopDetails_PhysicalStore_Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type CreateRequest struct {
	Details *ShopDetails `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetDetails() *ShopDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type CreateResponse struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type ReadRequest struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type ReadResponse struct {
	ShopId   string       `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	CreateAt int64        `protobuf:"varint,2,opt,name=create_at,json=createAt" json:"create_at,omitempty"`
	UpdateAt int64        `protobuf:"varint,3,opt,name=update_at,json=updateAt" json:"update_at,omitempty"`
	PeriodAt int64        `protobuf:"varint,4,opt,name=period_at,json=periodAt" json:"period_at,omitempty"`
	SubmitAt int64        `protobuf:"varint,5,opt,name=submit_at,json=submitAt" json:"submit_at,omitempty"`
	State    State        `protobuf:"varint,6,opt,name=state,enum=com.btdxcx.micro.srv.shop.details.State" json:"state,omitempty"`
	Details  *ShopDetails `protobuf:"bytes,7,opt,name=details" json:"details,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *ReadResponse) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *ReadResponse) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *ReadResponse) GetPeriodAt() int64 {
	if m != nil {
		return m.PeriodAt
	}
	return 0
}

func (m *ReadResponse) GetSubmitAt() int64 {
	if m != nil {
		return m.SubmitAt
	}
	return 0
}

func (m *ReadResponse) GetState() State {
	if m != nil {
		return m.State
	}
	return State_untreated
}

func (m *ReadResponse) GetDetails() *ShopDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type DeleteRequest struct {
	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ListRequest struct {
	Limit int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Start int32 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

type ListResponse struct {
	Start int32           `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	Limit int32           `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Total int32           `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Items []*ReadResponse `protobuf:"bytes,4,rep,name=items" json:"items,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListResponse) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListResponse) GetItems() []*ReadResponse {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateRequest struct {
	ShopId  string       `protobuf:"bytes,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	Details *ShopDetails `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
	State   State        `protobuf:"varint,3,opt,name=state,enum=com.btdxcx.micro.srv.shop.details.State" json:"state,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateRequest) GetShopId() string {
	if m != nil {
		return m.ShopId
	}
	return ""
}

func (m *UpdateRequest) GetDetails() *ShopDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *UpdateRequest) GetState() State {
	if m != nil {
		return m.State
	}
	return State_untreated
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*ShopDetails)(nil), "com.btdxcx.micro.srv.shop.details.ShopDetails")
	proto.RegisterType((*ShopDetails_ShopOwner)(nil), "com.btdxcx.micro.srv.shop.details.ShopDetails.ShopOwner")
	proto.RegisterType((*ShopDetails_WeiXin)(nil), "com.btdxcx.micro.srv.shop.details.ShopDetails.WeiXin")
	proto.RegisterType((*ShopDetails_PhysicalStore)(nil), "com.btdxcx.micro.srv.shop.details.ShopDetails.PhysicalStore")
	proto.RegisterType((*ShopDetails_PhysicalStore_Location)(nil), "com.btdxcx.micro.srv.shop.details.ShopDetails.PhysicalStore.Location")
	proto.RegisterType((*CreateRequest)(nil), "com.btdxcx.micro.srv.shop.details.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "com.btdxcx.micro.srv.shop.details.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "com.btdxcx.micro.srv.shop.details.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "com.btdxcx.micro.srv.shop.details.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "com.btdxcx.micro.srv.shop.details.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "com.btdxcx.micro.srv.shop.details.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "com.btdxcx.micro.srv.shop.details.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "com.btdxcx.micro.srv.shop.details.ListResponse")
	proto.RegisterType((*UpdateRequest)(nil), "com.btdxcx.micro.srv.shop.details.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "com.btdxcx.micro.srv.shop.details.UpdateResponse")
	proto.RegisterEnum("com.btdxcx.micro.srv.shop.details.State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Shop service

type ShopClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type shopClient struct {
	c           client.Client
	serviceName string
}

func NewShopClient(serviceName string, c client.Client) ShopClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.btdxcx.micro.srv.shop.details"
	}
	return &shopClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shopClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shop.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shop.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shop.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shop.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Shop.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shop service

type ShopHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterShopHandler(s server.Server, hdlr ShopHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Shop{hdlr}, opts...))
}

type Shop struct {
	ShopHandler
}

func (h *Shop) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ShopHandler.Create(ctx, in, out)
}

func (h *Shop) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ShopHandler.Read(ctx, in, out)
}

func (h *Shop) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ShopHandler.Delete(ctx, in, out)
}

func (h *Shop) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.ShopHandler.List(ctx, in, out)
}

func (h *Shop) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ShopHandler.Update(ctx, in, out)
}

func init() {
	proto.RegisterFile("btdxcx.com/micro/shop-srv/proto/shop/details/shop_details.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x8f, 0x23, 0x35,
	0x10, 0x9e, 0xee, 0xa4, 0xf3, 0xa8, 0x6c, 0x46, 0x23, 0x6b, 0x25, 0x5a, 0x01, 0xc4, 0x90, 0x03,
	0x0a, 0x48, 0x74, 0x20, 0x80, 0x04, 0x12, 0x02, 0xad, 0x76, 0x56, 0x62, 0xc4, 0xf2, 0x90, 0x47,
	0x88, 0xe5, 0x34, 0xf2, 0xb4, 0xad, 0x89, 0xb5, 0xdd, 0x6d, 0xd3, 0x76, 0xe6, 0xc1, 0x1f, 0xe1,
	0xc8, 0x89, 0x0b, 0xdc, 0xf8, 0x05, 0xfc, 0x34, 0x64, 0x97, 0x3b, 0x8f, 0x45, 0xa3, 0xe9, 0xec,
	0xdc, 0xfa, 0xab, 0xcf, 0xf5, 0xb9, 0xea, 0xab, 0x72, 0x14, 0xf8, 0xfa, 0xc2, 0xf2, 0x9b, 0xfc,
	0x26, 0xcb, 0x55, 0x39, 0x2f, 0x65, 0x5e, 0xab, 0xb9, 0x59, 0x2a, 0xfd, 0xa1, 0xa9, 0xaf, 0xe6,
	0xba, 0x56, 0x16, 0xe1, 0x9c, 0x0b, 0xcb, 0x64, 0x61, 0x3c, 0x38, 0x0f, 0x20, 0xf3, 0x3c, 0x79,
	0x37, 0x57, 0x65, 0x16, 0x44, 0xbc, 0x40, 0x66, 0xea, 0xab, 0xcc, 0x1d, 0xcc, 0xc2, 0xc1, 0xe9,
	0x1f, 0x7d, 0x18, 0x9d, 0x2d, 0x95, 0x3e, 0x41, 0x4c, 0x08, 0x74, 0x2b, 0x56, 0x8a, 0x34, 0x3a,
	0x8e, 0x66, 0x43, 0xea, 0xbf, 0x5d, 0xac, 0x50, 0x97, 0x2a, 0x8d, 0x31, 0xe6, 0xbe, 0xc9, 0x5b,
	0x30, 0x94, 0x95, 0xad, 0x15, 0x5f, 0xe5, 0x22, 0xed, 0x78, 0x62, 0x13, 0x20, 0xdf, 0x43, 0xa2,
	0xae, 0x2b, 0x51, 0xa7, 0xdd, 0xe3, 0x68, 0x36, 0x5a, 0x7c, 0x9e, 0xdd, 0x5b, 0x48, 0xb6, 0x55,
	0x84, 0xff, 0xfe, 0xc1, 0xe5, 0x53, 0x94, 0x21, 0xdf, 0x41, 0xef, 0x5a, 0xc8, 0x1b, 0x59, 0xa5,
	0x89, 0x17, 0xfc, 0x6c, 0x4f, 0xc1, 0x9f, 0x85, 0x7c, 0x21, 0x2b, 0x1a, 0x44, 0xc8, 0x0b, 0x18,
	0xe8, 0xe5, 0xad, 0x91, 0x39, 0x2b, 0xd2, 0x9e, 0x17, 0xfc, 0x72, 0x4f, 0xc1, 0x1f, 0x43, 0xfa,
	0x99, 0x55, 0xb5, 0xa0, 0x6b, 0xb5, 0x09, 0x83, 0xe1, 0xba, 0x78, 0x72, 0x08, 0xb1, 0xe4, 0xc1,
	0xc9, 0x58, 0xf2, 0xb5, 0xb7, 0xf1, 0x96, 0xb7, 0x13, 0x18, 0x54, 0x32, 0x7f, 0xe9, 0xe3, 0x68,
	0xe3, 0x1a, 0x93, 0xc7, 0x90, 0xe8, 0xa5, 0xaa, 0x84, 0x77, 0x71, 0x48, 0x11, 0x4c, 0xfe, 0x8a,
	0xa0, 0x87, 0xfd, 0xfc, 0xef, 0x82, 0x37, 0x61, 0x78, 0x2d, 0xf2, 0x25, 0xb3, 0xe7, 0x92, 0x87,
	0x5b, 0x06, 0x18, 0x38, 0xe5, 0x4e, 0x8d, 0x69, 0x2d, 0x79, 0xb8, 0x06, 0x01, 0x79, 0x1b, 0x80,
	0x69, 0x7d, 0x6e, 0x44, 0x5e, 0x0b, 0x1b, 0x2e, 0x1a, 0x32, 0xad, 0xcf, 0x7c, 0xc0, 0xd1, 0x9a,
	0xd5, 0xb6, 0x12, 0xb5, 0x93, 0x4c, 0x90, 0x0e, 0x91, 0x53, 0x4e, 0xde, 0x81, 0x51, 0x43, 0xbf,
	0x14, 0xb7, 0xde, 0xcb, 0x21, 0x6d, 0x32, 0xbe, 0x15, 0xb7, 0x93, 0x7f, 0x63, 0x18, 0xef, 0x78,
	0xd5, 0xca, 0x94, 0x14, 0xfa, 0xb9, 0xaa, 0x2c, 0xcb, 0x6d, 0x28, 0xb6, 0x81, 0xae, 0x09, 0x51,
	0x32, 0x59, 0x34, 0x96, 0x78, 0xe0, 0xce, 0xff, 0x26, 0xf5, 0x53, 0xc5, 0x45, 0x28, 0xb1, 0x81,
	0x8e, 0x61, 0x9c, 0xd7, 0xc2, 0x98, 0x50, 0x5c, 0x03, 0x09, 0x83, 0x41, 0xa1, 0x72, 0x66, 0xa5,
	0xaa, 0xd2, 0xbe, 0xdf, 0x81, 0x67, 0x0f, 0xd9, 0x81, 0xec, 0x79, 0x10, 0xa3, 0x6b, 0xd9, 0xc9,
	0x09, 0x0c, 0x9a, 0xa8, 0x9b, 0x73, 0xc1, 0xac, 0xb4, 0x2b, 0x8e, 0x6f, 0x2b, 0xa2, 0x6b, 0xec,
	0xde, 0x52, 0xa1, 0xaa, 0x4b, 0x24, 0x63, 0x4f, 0x6e, 0x02, 0xd3, 0x5f, 0x60, 0xfc, 0xb4, 0x16,
	0xcc, 0x0a, 0x2a, 0x7e, 0x5d, 0x09, 0x63, 0xc9, 0x37, 0xd0, 0x0f, 0xe5, 0x78, 0xa5, 0xd1, 0x22,
	0xdb, 0xaf, 0x70, 0xda, 0xa4, 0x4f, 0xdf, 0x87, 0xc3, 0x46, 0xda, 0x68, 0x55, 0x19, 0x41, 0xde,
	0x80, 0xbe, 0xff, 0x1d, 0x59, 0x8f, 0xa8, 0xe7, 0xe0, 0x29, 0x9f, 0xbe, 0x07, 0x23, 0x2a, 0x18,
	0x6f, 0x6a, 0xb8, 0xf3, 0xdc, 0x9f, 0x31, 0x3c, 0xc2, 0x83, 0xf7, 0x28, 0xba, 0x65, 0xcd, 0xfd,
	0xe5, 0xe7, 0xcc, 0xfa, 0xae, 0x3b, 0x74, 0x80, 0x81, 0x27, 0xd6, 0x91, 0x2b, 0xcd, 0x03, 0xd9,
	0x41, 0x12, 0x03, 0x48, 0x6a, 0x51, 0x4b, 0xc5, 0x1d, 0xd9, 0x45, 0x12, 0x03, 0x48, 0x9a, 0xd5,
	0x45, 0x29, 0xad, 0x23, 0x13, 0x24, 0x31, 0xf0, 0xc4, 0x92, 0xaf, 0x20, 0x31, 0x96, 0x59, 0xe1,
	0x97, 0xe1, 0x70, 0x31, 0x6b, 0x63, 0x9c, 0x3b, 0x4f, 0x31, 0x6d, 0xdb, 0xfa, 0xfe, 0xc3, 0xac,
	0x9f, 0xc1, 0xf8, 0x44, 0x14, 0x62, 0x33, 0xd5, 0x3b, 0x1d, 0x3d, 0x82, 0xc3, 0xe6, 0x24, 0x5a,
	0x3a, 0xfd, 0x02, 0x46, 0xcf, 0xa5, 0xb1, 0x4d, 0xe6, 0x63, 0x48, 0x0a, 0x59, 0x4a, 0xeb, 0xf3,
	0x12, 0x8a, 0xc0, 0x45, 0x8d, 0x65, 0x35, 0x5a, 0x9b, 0x50, 0x04, 0xd3, 0xdf, 0x23, 0x78, 0x84,
	0xb9, 0x61, 0x3c, 0xeb, 0x63, 0xd1, 0xd6, 0xb1, 0x8d, 0x64, 0xfc, 0x8a, 0xa4, 0x55, 0x96, 0x15,
	0x7e, 0x20, 0x09, 0x45, 0x40, 0x9e, 0x41, 0x22, 0xad, 0x28, 0x4d, 0xda, 0x3d, 0xee, 0xcc, 0x46,
	0x8b, 0x79, 0x0b, 0x47, 0xb6, 0x17, 0x84, 0x62, 0xf6, 0xf4, 0x9f, 0x08, 0xc6, 0x3f, 0xf9, 0x09,
	0xdf, 0xe7, 0xc8, 0xf6, 0x14, 0xe2, 0x07, 0x4d, 0x61, 0xb3, 0x0f, 0x9d, 0xd7, 0xda, 0x07, 0x37,
	0x9b, 0xa6, 0x66, 0xec, 0xe6, 0x83, 0x4f, 0x21, 0xf1, 0x27, 0xc8, 0x18, 0x86, 0xab, 0xca, 0xfa,
	0x7d, 0xe6, 0x47, 0x07, 0x0e, 0xd6, 0xe2, 0x4a, 0x8a, 0x6b, 0x59, 0x5d, 0x1e, 0x45, 0x0e, 0xe6,
	0xaa, 0xd4, 0x6e, 0xac, 0xfc, 0x28, 0x5e, 0xfc, 0xdd, 0x85, 0xae, 0x2b, 0x90, 0x28, 0xe8, 0xe1,
	0x8b, 0x24, 0x1f, 0xb5, 0xa8, 0x65, 0xe7, 0x77, 0x61, 0xf2, 0xf1, 0x1e, 0x19, 0x61, 0x93, 0x0e,
	0x88, 0x84, 0xae, 0x9b, 0x06, 0xc9, 0x5a, 0x8f, 0x0d, 0x2f, 0xdb, 0x77, 0xcc, 0xd3, 0x03, 0xd7,
	0x1b, 0x2e, 0x72, 0xab, 0xde, 0x76, 0x5e, 0x47, 0xab, 0xde, 0x5e, 0x79, 0x25, 0x07, 0xe4, 0x12,
	0xba, 0x6e, 0xd7, 0x5b, 0xf5, 0xb6, 0xf5, 0xa0, 0x5a, 0xf5, 0xb6, 0xf3, 0x88, 0x4a, 0xe8, 0xe1,
	0x1a, 0xb4, 0xea, 0x6c, 0x67, 0xcb, 0x5b, 0x75, 0xb6, 0xbb, 0x63, 0x17, 0x3d, 0xff, 0xef, 0xee,
	0x93, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xcc, 0x8c, 0x9d, 0x20, 0x0a, 0x00, 0x00,
}
