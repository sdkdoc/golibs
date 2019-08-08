// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/customer_manager_link_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Request message for [CustomerManagerLinkService.GetCustomerManagerLink][google.ads.googleads.v0.services.CustomerManagerLinkService.GetCustomerManagerLink].
type GetCustomerManagerLinkRequest struct {
	// The resource name of the CustomerManagerLink to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerManagerLinkRequest) Reset()         { *m = GetCustomerManagerLinkRequest{} }
func (m *GetCustomerManagerLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerManagerLinkRequest) ProtoMessage()    {}
func (*GetCustomerManagerLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_manager_link_service_9df054ecaf272672, []int{0}
}
func (m *GetCustomerManagerLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Unmarshal(m, b)
}
func (m *GetCustomerManagerLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerManagerLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerManagerLinkRequest.Merge(dst, src)
}
func (m *GetCustomerManagerLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerManagerLinkRequest.Size(m)
}
func (m *GetCustomerManagerLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerManagerLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerManagerLinkRequest proto.InternalMessageInfo

func (m *GetCustomerManagerLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerManagerLinkRequest)(nil), "google.ads.googleads.v0.services.GetCustomerManagerLinkRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerManagerLinkServiceClient is the client API for CustomerManagerLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerManagerLinkServiceClient interface {
	// Returns the requested CustomerManagerLink in full detail.
	GetCustomerManagerLink(ctx context.Context, in *GetCustomerManagerLinkRequest, opts ...grpc.CallOption) (*resources.CustomerManagerLink, error)
}

type customerManagerLinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerManagerLinkServiceClient(cc *grpc.ClientConn) CustomerManagerLinkServiceClient {
	return &customerManagerLinkServiceClient{cc}
}

func (c *customerManagerLinkServiceClient) GetCustomerManagerLink(ctx context.Context, in *GetCustomerManagerLinkRequest, opts ...grpc.CallOption) (*resources.CustomerManagerLink, error) {
	out := new(resources.CustomerManagerLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerManagerLinkService/GetCustomerManagerLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerManagerLinkServiceServer is the server API for CustomerManagerLinkService service.
type CustomerManagerLinkServiceServer interface {
	// Returns the requested CustomerManagerLink in full detail.
	GetCustomerManagerLink(context.Context, *GetCustomerManagerLinkRequest) (*resources.CustomerManagerLink, error)
}

func RegisterCustomerManagerLinkServiceServer(s *grpc.Server, srv CustomerManagerLinkServiceServer) {
	s.RegisterService(&_CustomerManagerLinkService_serviceDesc, srv)
}

func _CustomerManagerLinkService_GetCustomerManagerLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerManagerLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagerLinkServiceServer).GetCustomerManagerLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerManagerLinkService/GetCustomerManagerLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagerLinkServiceServer).GetCustomerManagerLink(ctx, req.(*GetCustomerManagerLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerManagerLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CustomerManagerLinkService",
	HandlerType: (*CustomerManagerLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerManagerLink",
			Handler:    _CustomerManagerLinkService_GetCustomerManagerLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/customer_manager_link_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/customer_manager_link_service.proto", fileDescriptor_customer_manager_link_service_9df054ecaf272672)
}

var fileDescriptor_customer_manager_link_service_9df054ecaf272672 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0xe4, 0xd2, 0xe2, 0x92, 0xfc, 0xdc, 0xd4, 0xa2, 0xf8,
	0xdc, 0xc4, 0xbc, 0xc4, 0xf4, 0xd4, 0xa2, 0xf8, 0x9c, 0xcc, 0xbc, 0xec, 0x78, 0xa8, 0xb4, 0x5e,
	0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x02, 0x44, 0xab, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0xdc, 0x14,
	0xbd, 0x32, 0x03, 0x3d, 0x98, 0x29, 0x52, 0xb6, 0xb8, 0xec, 0x29, 0x4a, 0x2d, 0xce, 0x2f, 0x2d,
	0xc2, 0x69, 0x11, 0xc4, 0x02, 0x29, 0x19, 0x98, 0xf6, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc,
	0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0xac, 0x92, 0x0b, 0x97, 0xac, 0x7b, 0x6a, 0x89,
	0x33, 0x54, 0xbf, 0x2f, 0x44, 0xbb, 0x4f, 0x66, 0x5e, 0x76, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x90, 0x32, 0x17, 0x2f, 0xcc, 0x9e, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x1e, 0x98, 0xa0, 0x5f, 0x62, 0x6e, 0xaa, 0xd1, 0x07, 0x46, 0x2e, 0x29, 0x2c,
	0x66, 0x04, 0x43, 0xbc, 0x20, 0x74, 0x91, 0x91, 0x4b, 0x0c, 0xbb, 0x2d, 0x42, 0xf6, 0x7a, 0x84,
	0xfc, 0xaf, 0x87, 0xd7, 0x7d, 0x52, 0x66, 0x38, 0x0d, 0x80, 0x07, 0x8f, 0x1e, 0x16, 0xed, 0x4a,
	0x76, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xb2, 0x10, 0x32, 0x03, 0x85, 0x64, 0x35, 0x8a, 0x17, 0x6d,
	0x61, 0xc1, 0x59, 0xac, 0xaf, 0x05, 0x0f, 0x5a, 0x24, 0xbd, 0xc5, 0xfa, 0x5a, 0xb5, 0x4e, 0xf7,
	0x19, 0xb9, 0x54, 0x92, 0xf3, 0x73, 0x09, 0x3a, 0xdf, 0x49, 0x1e, 0x77, 0xc0, 0x04, 0x80, 0xa2,
	0x20, 0x80, 0x31, 0xca, 0x03, 0x6a, 0x48, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51,
	0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0x38, 0x82, 0x60, 0x31, 0x5e, 0x90, 0x59, 0x8c, 0x3b, 0xa1, 0x59,
	0xc3, 0x18, 0x8b, 0x98, 0x98, 0xdd, 0x1d, 0x1d, 0x57, 0x31, 0x29, 0xb8, 0x43, 0x0c, 0x74, 0x4c,
	0x29, 0xd6, 0x83, 0x30, 0x41, 0xac, 0x30, 0x03, 0x3d, 0xa8, 0xc5, 0xc5, 0xa7, 0x60, 0x4a, 0x62,
	0x1c, 0x53, 0x8a, 0x63, 0xe0, 0x4a, 0x62, 0xc2, 0x0c, 0x62, 0x60, 0x4a, 0x92, 0xd8, 0xc0, 0x0e,
	0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xe5, 0xd3, 0xc2, 0xe8, 0x02, 0x00, 0x00,
}
