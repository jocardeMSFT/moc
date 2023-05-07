// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_logicalnetwork.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LogicalNetworkRequest struct {
	LogicalNetworks      []*LogicalNetwork `protobuf:"bytes,1,rep,name=LogicalNetworks,proto3" json:"LogicalNetworks,omitempty"`
	OperationType        common.Operation  `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LogicalNetworkRequest) Reset()         { *m = LogicalNetworkRequest{} }
func (m *LogicalNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*LogicalNetworkRequest) ProtoMessage()    {}
func (*LogicalNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{0}
}

func (m *LogicalNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetworkRequest.Unmarshal(m, b)
}
func (m *LogicalNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetworkRequest.Marshal(b, m, deterministic)
}
func (m *LogicalNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetworkRequest.Merge(m, src)
}
func (m *LogicalNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_LogicalNetworkRequest.Size(m)
}
func (m *LogicalNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetworkRequest proto.InternalMessageInfo

func (m *LogicalNetworkRequest) GetLogicalNetworks() []*LogicalNetwork {
	if m != nil {
		return m.LogicalNetworks
	}
	return nil
}

func (m *LogicalNetworkRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LogicalNetworkResponse struct {
	LogicalNetworks      []*LogicalNetwork   `protobuf:"bytes,1,rep,name=LogicalNetworks,proto3" json:"LogicalNetworks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogicalNetworkResponse) Reset()         { *m = LogicalNetworkResponse{} }
func (m *LogicalNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*LogicalNetworkResponse) ProtoMessage()    {}
func (*LogicalNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{1}
}

func (m *LogicalNetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetworkResponse.Unmarshal(m, b)
}
func (m *LogicalNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetworkResponse.Marshal(b, m, deterministic)
}
func (m *LogicalNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetworkResponse.Merge(m, src)
}
func (m *LogicalNetworkResponse) XXX_Size() int {
	return xxx_messageInfo_LogicalNetworkResponse.Size(m)
}
func (m *LogicalNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetworkResponse proto.InternalMessageInfo

func (m *LogicalNetworkResponse) GetLogicalNetworks() []*LogicalNetwork {
	if m != nil {
		return m.LogicalNetworks
	}
	return nil
}

func (m *LogicalNetworkResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LogicalNetworkResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LogicalNetwork struct {
	Name                         string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                           string              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Subnets                      []*LogicalSubnet    `protobuf:"bytes,3,rep,name=subnets,proto3" json:"subnets,omitempty"`
	NetworkVirtualizationEnabled *wrappers.BoolValue `protobuf:"bytes,4,opt,name=networkVirtualizationEnabled,proto3" json:"networkVirtualizationEnabled,omitempty"`
	Nodefqdn                     string              `protobuf:"bytes,5,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	Status                       *common.Status      `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	LocationName                 string              `protobuf:"bytes,7,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Tags                         *common.Tags        `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}            `json:"-"`
	XXX_unrecognized             []byte              `json:"-"`
	XXX_sizecache                int32               `json:"-"`
}

func (m *LogicalNetwork) Reset()         { *m = LogicalNetwork{} }
func (m *LogicalNetwork) String() string { return proto.CompactTextString(m) }
func (*LogicalNetwork) ProtoMessage()    {}
func (*LogicalNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{2}
}

func (m *LogicalNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalNetwork.Unmarshal(m, b)
}
func (m *LogicalNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalNetwork.Marshal(b, m, deterministic)
}
func (m *LogicalNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalNetwork.Merge(m, src)
}
func (m *LogicalNetwork) XXX_Size() int {
	return xxx_messageInfo_LogicalNetwork.Size(m)
}
func (m *LogicalNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalNetwork proto.InternalMessageInfo

func (m *LogicalNetwork) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalNetwork) GetSubnets() []*LogicalSubnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *LogicalNetwork) GetNetworkVirtualizationEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.NetworkVirtualizationEnabled
	}
	return nil
}

func (m *LogicalNetwork) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *LogicalNetwork) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LogicalNetwork) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *LogicalNetwork) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type LogicalSubnet struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	AddressPrefix        string                    `protobuf:"bytes,3,opt,name=addressPrefix,proto3" json:"addressPrefix,omitempty"`
	Routes               []*common.Route           `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
	Allocation           common.IPAllocationMethod `protobuf:"varint,5,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Vlan                 uint32                    `protobuf:"varint,6,opt,name=vlan,proto3" json:"vlan,omitempty"`
	IpPools              []*common.IPPool          `protobuf:"bytes,7,rep,name=ipPools,proto3" json:"ipPools,omitempty"`
	Dns                  *common.Dns               `protobuf:"bytes,8,opt,name=dns,proto3" json:"dns,omitempty"`
	IsPublic             *wrappers.BoolValue       `protobuf:"bytes,9,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LogicalSubnet) Reset()         { *m = LogicalSubnet{} }
func (m *LogicalSubnet) String() string { return proto.CompactTextString(m) }
func (*LogicalSubnet) ProtoMessage()    {}
func (*LogicalSubnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_37e4212e5c247b0a, []int{3}
}

func (m *LogicalSubnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalSubnet.Unmarshal(m, b)
}
func (m *LogicalSubnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalSubnet.Marshal(b, m, deterministic)
}
func (m *LogicalSubnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalSubnet.Merge(m, src)
}
func (m *LogicalSubnet) XXX_Size() int {
	return xxx_messageInfo_LogicalSubnet.Size(m)
}
func (m *LogicalSubnet) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalSubnet.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalSubnet proto.InternalMessageInfo

func (m *LogicalSubnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalSubnet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalSubnet) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *LogicalSubnet) GetRoutes() []*common.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *LogicalSubnet) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *LogicalSubnet) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *LogicalSubnet) GetIpPools() []*common.IPPool {
	if m != nil {
		return m.IpPools
	}
	return nil
}

func (m *LogicalSubnet) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *LogicalSubnet) GetIsPublic() *wrappers.BoolValue {
	if m != nil {
		return m.IsPublic
	}
	return nil
}

func init() {
	proto.RegisterType((*LogicalNetworkRequest)(nil), "moc.cloudagent.network.LogicalNetworkRequest")
	proto.RegisterType((*LogicalNetworkResponse)(nil), "moc.cloudagent.network.LogicalNetworkResponse")
	proto.RegisterType((*LogicalNetwork)(nil), "moc.cloudagent.network.LogicalNetwork")
	proto.RegisterType((*LogicalSubnet)(nil), "moc.cloudagent.network.LogicalSubnet")
}

func init() {
	proto.RegisterFile("moc_cloudagent_logicalnetwork.proto", fileDescriptor_37e4212e5c247b0a)
}

var fileDescriptor_37e4212e5c247b0a = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xf2, 0xd3, 0xfc, 0xdc, 0x7c, 0x09, 0xd2, 0x00, 0xad, 0x15, 0x41, 0x55, 0xa5, 0x14,
	0x75, 0x83, 0x2d, 0x02, 0x82, 0x25, 0x6a, 0x45, 0x17, 0x95, 0xa0, 0x44, 0xd3, 0xaa, 0x0b, 0x16,
	0x54, 0x13, 0x7b, 0xe2, 0x5a, 0x1d, 0xcf, 0x75, 0x67, 0xc6, 0x2d, 0xb0, 0xe2, 0x4d, 0x78, 0x09,
	0x5e, 0x81, 0x87, 0xe0, 0x6d, 0x50, 0xae, 0x9d, 0xb4, 0xae, 0x10, 0xed, 0x86, 0x55, 0xec, 0x7b,
	0xce, 0x3d, 0xf7, 0xdc, 0xe3, 0x99, 0xc0, 0x66, 0x8a, 0xe1, 0x49, 0xa8, 0x30, 0x8f, 0x44, 0x2c,
	0xb5, 0x3b, 0x51, 0x18, 0x27, 0xa1, 0x50, 0x5a, 0xba, 0x4b, 0x34, 0x67, 0x7e, 0x66, 0xd0, 0x21,
	0x5b, 0x4d, 0x31, 0xf4, 0xaf, 0x48, 0x7e, 0x89, 0x0e, 0xd7, 0x63, 0xc4, 0x58, 0xc9, 0x80, 0x58,
	0xd3, 0x7c, 0x16, 0x5c, 0x1a, 0x91, 0x65, 0xd2, 0xd8, 0xa2, 0x6f, 0xb8, 0x46, 0xe2, 0x98, 0xa6,
	0xa8, 0xcb, 0x9f, 0x12, 0x58, 0xbf, 0x06, 0x94, 0x62, 0xd7, 0xf1, 0xd1, 0xf7, 0x1a, 0x3c, 0x7c,
	0x57, 0x38, 0x39, 0x28, 0x60, 0x2e, 0xcf, 0x73, 0x69, 0x1d, 0x9b, 0xc0, 0xbd, 0x2a, 0x60, 0xbd,
	0xda, 0x46, 0x63, 0xbb, 0x37, 0x7e, 0xea, 0xff, 0xd9, 0xa4, 0x7f, 0x43, 0xe7, 0x66, 0x3b, 0x7b,
	0x09, 0xfd, 0x0f, 0x99, 0x34, 0xc2, 0x25, 0xa8, 0x8f, 0xbe, 0x64, 0xd2, 0xab, 0x6f, 0xd4, 0xb6,
	0x07, 0xe3, 0x01, 0xe9, 0x2d, 0x11, 0x5e, 0x25, 0x8d, 0x7e, 0xd4, 0x60, 0xf5, 0xa6, 0x43, 0x9b,
	0xa1, 0xb6, 0xf2, 0x1f, 0x58, 0x1c, 0x43, 0x8b, 0x4b, 0x9b, 0x2b, 0x47, 0xde, 0x7a, 0xe3, 0xa1,
	0x5f, 0x04, 0xef, 0x2f, 0x82, 0xf7, 0x77, 0x11, 0xd5, 0xb1, 0x50, 0xb9, 0xe4, 0x25, 0x93, 0x3d,
	0x80, 0x95, 0x3d, 0x63, 0xd0, 0x78, 0x8d, 0x8d, 0xda, 0x76, 0x97, 0x17, 0x2f, 0xa3, 0x5f, 0x75,
	0x18, 0x54, 0xd5, 0x19, 0x83, 0xa6, 0x16, 0xa9, 0xf4, 0x6a, 0xc4, 0xa3, 0x67, 0x36, 0x80, 0x7a,
	0x12, 0xd1, 0xb0, 0x2e, 0xaf, 0x27, 0x11, 0x7b, 0x03, 0x6d, 0x9b, 0x4f, 0xb5, 0x74, 0xd6, 0x6b,
	0xd0, 0x2a, 0x5b, 0xb7, 0xac, 0x72, 0x48, 0x6c, 0xbe, 0xe8, 0x62, 0x9f, 0xe0, 0x51, 0xc9, 0x38,
	0x4e, 0x8c, 0xcb, 0x85, 0x4a, 0xbe, 0x52, 0x96, 0x7b, 0x5a, 0x4c, 0x95, 0x8c, 0xbc, 0xe6, 0xad,
	0x7b, 0xfd, 0xb5, 0x9f, 0x0d, 0xa1, 0xa3, 0x31, 0x92, 0xb3, 0xf3, 0x48, 0x7b, 0x2b, 0x64, 0x7b,
	0xf9, 0xce, 0x36, 0xa1, 0x65, 0x9d, 0x70, 0xb9, 0xf5, 0x5a, 0x34, 0xa5, 0x47, 0xde, 0x0f, 0xa9,
	0xc4, 0x4b, 0x88, 0x8d, 0xe0, 0x7f, 0x85, 0x21, 0x69, 0x1e, 0xcc, 0xd3, 0x68, 0x93, 0x48, 0xa5,
	0xc6, 0x1e, 0x43, 0xd3, 0x89, 0xd8, 0x7a, 0x1d, 0x92, 0xe9, 0x92, 0xcc, 0x91, 0x88, 0x2d, 0xa7,
	0xf2, 0xe8, 0x67, 0x1d, 0xfa, 0x95, 0xf5, 0xef, 0x14, 0xed, 0x13, 0xe8, 0x8b, 0x28, 0x32, 0xd2,
	0xda, 0x89, 0x91, 0xb3, 0xe4, 0x73, 0xf9, 0xbd, 0xaa, 0x45, 0x36, 0x82, 0x96, 0xc1, 0xdc, 0x49,
	0xeb, 0x35, 0x29, 0x7f, 0xa0, 0xe1, 0x7c, 0x5e, 0xe2, 0x25, 0xc2, 0x5e, 0x03, 0x08, 0xb5, 0x30,
	0x4c, 0x29, 0x0c, 0xc6, 0x6b, 0xc4, 0xdb, 0x9f, 0xec, 0x2c, 0x81, 0xf7, 0xd2, 0x9d, 0x62, 0xc4,
	0xaf, 0x51, 0xe7, 0x36, 0x2f, 0x94, 0xd0, 0x14, 0x4f, 0x9f, 0xd3, 0x33, 0xdb, 0x82, 0x76, 0x92,
	0x4d, 0x10, 0x95, 0xf5, 0xda, 0x34, 0xb1, 0x57, 0x2a, 0xcd, 0x6b, 0x7c, 0x81, 0xb1, 0x21, 0x34,
	0x22, 0xbd, 0x48, 0xa4, 0x43, 0x94, 0xb7, 0xda, 0xf2, 0x79, 0x91, 0xbd, 0x82, 0x4e, 0x62, 0x27,
	0xf9, 0x54, 0x25, 0xa1, 0xd7, 0xbd, 0xf5, 0xfb, 0x2e, 0xb9, 0xe3, 0x6f, 0x35, 0xb8, 0x5f, 0x3d,
	0xa3, 0x3b, 0xf3, 0x23, 0xc6, 0x12, 0x68, 0xed, 0xeb, 0x0b, 0x3c, 0x93, 0xec, 0xd9, 0x1d, 0x2f,
	0x52, 0xf1, 0x9f, 0x31, 0xf4, 0xef, 0x4a, 0x2f, 0x2e, 0xf0, 0xe8, 0xbf, 0xdd, 0xe7, 0x1f, 0x83,
	0x38, 0x71, 0xa7, 0xf9, 0xd4, 0x0f, 0x31, 0x0d, 0xd2, 0x24, 0x34, 0x68, 0x71, 0xe6, 0x82, 0x14,
	0xc3, 0xc0, 0x64, 0x61, 0x70, 0xa5, 0x15, 0x94, 0x5a, 0xd3, 0x16, 0xed, 0xf4, 0xe2, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0xc8, 0xf6, 0x53, 0x51, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicalNetworkAgentClient is the client API for LogicalNetworkAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicalNetworkAgentClient interface {
	Invoke(ctx context.Context, in *LogicalNetworkRequest, opts ...grpc.CallOption) (*LogicalNetworkResponse, error)
}

type logicalNetworkAgentClient struct {
	cc *grpc.ClientConn
}

func NewLogicalNetworkAgentClient(cc *grpc.ClientConn) LogicalNetworkAgentClient {
	return &logicalNetworkAgentClient{cc}
}

func (c *logicalNetworkAgentClient) Invoke(ctx context.Context, in *LogicalNetworkRequest, opts ...grpc.CallOption) (*LogicalNetworkResponse, error) {
	out := new(LogicalNetworkResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.LogicalNetworkAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicalNetworkAgentServer is the server API for LogicalNetworkAgent service.
type LogicalNetworkAgentServer interface {
	Invoke(context.Context, *LogicalNetworkRequest) (*LogicalNetworkResponse, error)
}

// UnimplementedLogicalNetworkAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLogicalNetworkAgentServer struct {
}

func (*UnimplementedLogicalNetworkAgentServer) Invoke(ctx context.Context, req *LogicalNetworkRequest) (*LogicalNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterLogicalNetworkAgentServer(s *grpc.Server, srv LogicalNetworkAgentServer) {
	s.RegisterService(&_LogicalNetworkAgent_serviceDesc, srv)
}

func _LogicalNetworkAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogicalNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicalNetworkAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.LogicalNetworkAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicalNetworkAgentServer).Invoke(ctx, req.(*LogicalNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicalNetworkAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.LogicalNetworkAgent",
	HandlerType: (*LogicalNetworkAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LogicalNetworkAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_logicalnetwork.proto",
}
