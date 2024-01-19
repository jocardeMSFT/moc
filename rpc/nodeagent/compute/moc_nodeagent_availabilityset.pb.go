// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_availabilityset.proto

package compute

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

type AvailabilitySetRequest struct {
	AvailabilitySets     []*AvailabilitySet `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	OperationType        common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AvailabilitySetRequest) Reset()         { *m = AvailabilitySetRequest{} }
func (m *AvailabilitySetRequest) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetRequest) ProtoMessage()    {}
func (*AvailabilitySetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{0}
}

func (m *AvailabilitySetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetRequest.Unmarshal(m, b)
}
func (m *AvailabilitySetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetRequest.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetRequest.Merge(m, src)
}
func (m *AvailabilitySetRequest) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetRequest.Size(m)
}
func (m *AvailabilitySetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetRequest proto.InternalMessageInfo

func (m *AvailabilitySetRequest) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type AvailabilitySetResponse struct {
	AvailabilitySets     []*AvailabilitySet  `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AvailabilitySetResponse) Reset()         { *m = AvailabilitySetResponse{} }
func (m *AvailabilitySetResponse) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetResponse) ProtoMessage()    {}
func (*AvailabilitySetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{1}
}

func (m *AvailabilitySetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetResponse.Unmarshal(m, b)
}
func (m *AvailabilitySetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetResponse.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetResponse.Merge(m, src)
}
func (m *AvailabilitySetResponse) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetResponse.Size(m)
}
func (m *AvailabilitySetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetResponse proto.InternalMessageInfo

func (m *AvailabilitySetResponse) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AvailabilitySetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// avset structure is a flattened version of the model in the Azure sdk for go at
// https://github.com/Azure/azure-sdk-for-go/blob/main/sdk/resourcemanager/compute/armcompute/models.go
type AvailabilitySet struct {
	Name                     string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                       string                    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status                   *common.Status            `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Tags                     *common.Tags              `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Entity                   *common.Entity            `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	PlatformFaultDomainCount string                    `protobuf:"bytes,6,opt,name=platformFaultDomainCount,proto3" json:"platformFaultDomainCount,omitempty"`
	VirtualMachines          []*common.NodeSubResource `protobuf:"bytes,7,rep,name=virtualMachines,proto3" json:"virtualMachines,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *AvailabilitySet) Reset()         { *m = AvailabilitySet{} }
func (m *AvailabilitySet) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySet) ProtoMessage()    {}
func (*AvailabilitySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{2}
}

func (m *AvailabilitySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySet.Unmarshal(m, b)
}
func (m *AvailabilitySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySet.Marshal(b, m, deterministic)
}
func (m *AvailabilitySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySet.Merge(m, src)
}
func (m *AvailabilitySet) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySet.Size(m)
}
func (m *AvailabilitySet) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySet.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySet proto.InternalMessageInfo

func (m *AvailabilitySet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AvailabilitySet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AvailabilitySet) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AvailabilitySet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AvailabilitySet) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AvailabilitySet) GetPlatformFaultDomainCount() string {
	if m != nil {
		return m.PlatformFaultDomainCount
	}
	return ""
}

func (m *AvailabilitySet) GetVirtualMachines() []*common.NodeSubResource {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func init() {
	proto.RegisterType((*AvailabilitySetRequest)(nil), "moc.nodeagent.compute.AvailabilitySetRequest")
	proto.RegisterType((*AvailabilitySetResponse)(nil), "moc.nodeagent.compute.AvailabilitySetResponse")
	proto.RegisterType((*AvailabilitySet)(nil), "moc.nodeagent.compute.AvailabilitySet")
}

func init() {
	proto.RegisterFile("moc_nodeagent_availabilityset.proto", fileDescriptor_d615df2311cce965)
}

var fileDescriptor_d615df2311cce965 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0x4e, 0xd3, 0xfc, 0x95, 0x8d, 0x48, 0xd1, 0x2a, 0x50, 0x2b, 0x12, 0x28, 0x4a,
	0x25, 0x94, 0x0b, 0x6b, 0x64, 0x38, 0x71, 0x40, 0x6a, 0xa1, 0x48, 0x1c, 0x00, 0x69, 0x53, 0x71,
	0xe0, 0x52, 0x6d, 0x9c, 0x89, 0xbb, 0xc2, 0xbb, 0x63, 0x76, 0x67, 0x83, 0x72, 0xe4, 0x35, 0x38,
	0xf2, 0x12, 0xbc, 0x1e, 0xca, 0x3a, 0x44, 0xc5, 0x05, 0xa9, 0x17, 0x4e, 0xb6, 0xe7, 0xfb, 0xed,
	0x37, 0x9f, 0x67, 0x87, 0x9d, 0x18, 0x2c, 0x2e, 0x2d, 0x2e, 0x41, 0x95, 0x60, 0xe9, 0x52, 0xad,
	0x95, 0xae, 0xd4, 0x42, 0x57, 0x9a, 0x36, 0x1e, 0x48, 0xd4, 0x0e, 0x09, 0xf9, 0x3d, 0x83, 0x85,
	0xd8, 0x43, 0xa2, 0x40, 0x53, 0x07, 0x82, 0xf1, 0xc3, 0x12, 0xb1, 0xac, 0x20, 0x8b, 0xd0, 0x22,
	0xac, 0xb2, 0x2f, 0x4e, 0xd5, 0x35, 0x38, 0xdf, 0x1c, 0x1b, 0x1f, 0x6f, 0xbd, 0x0b, 0x34, 0x06,
	0xed, 0xee, 0xd1, 0x08, 0xd3, 0xef, 0x09, 0xbb, 0x7f, 0x7a, 0xad, 0xd3, 0x1c, 0x48, 0xc2, 0xe7,
	0x00, 0x9e, 0xb8, 0x64, 0x77, 0x5b, 0x8a, 0x4f, 0x93, 0xc9, 0xc1, 0x6c, 0x90, 0x3f, 0x12, 0x7f,
	0x4c, 0x21, 0xda, 0x46, 0x37, 0xce, 0xf3, 0x67, 0xec, 0xce, 0xfb, 0x1a, 0x9c, 0x22, 0x8d, 0xf6,
	0x62, 0x53, 0x43, 0xda, 0x99, 0x24, 0xb3, 0x61, 0x3e, 0x8c, 0x86, 0x7b, 0x45, 0xfe, 0x0e, 0x4d,
	0x7f, 0x24, 0xec, 0xf8, 0x46, 0x48, 0x5f, 0xa3, 0xf5, 0xf0, 0x4f, 0x52, 0xe6, 0xac, 0x27, 0xc1,
	0x87, 0x8a, 0x62, 0xbc, 0x41, 0x3e, 0x16, 0xcd, 0x78, 0xc5, 0xaf, 0xf1, 0x8a, 0x33, 0xc4, 0xea,
	0x83, 0xaa, 0x02, 0xc8, 0x1d, 0xc9, 0x47, 0xec, 0xf0, 0xdc, 0x39, 0x74, 0xe9, 0xc1, 0x24, 0x99,
	0xf5, 0x65, 0xf3, 0x31, 0xfd, 0xd6, 0x61, 0x47, 0x2d, 0x7b, 0xce, 0x59, 0xd7, 0x2a, 0x03, 0x69,
	0x12, 0xc1, 0xf8, 0xce, 0x87, 0xac, 0xa3, 0x97, 0xb1, 0x5b, 0x5f, 0x76, 0xf4, 0x92, 0x9f, 0xb0,
	0x9e, 0x27, 0x45, 0xc1, 0x47, 0xbb, 0x41, 0x3e, 0x88, 0xff, 0x32, 0x8f, 0x25, 0xb9, 0x93, 0xf8,
	0x03, 0xd6, 0x25, 0x55, 0xfa, 0xb4, 0x1b, 0x91, 0x7e, 0x44, 0x2e, 0x54, 0xe9, 0x65, 0x2c, 0x6f,
	0x3d, 0xc0, 0x92, 0xa6, 0x4d, 0x7a, 0x78, 0xcd, 0xe3, 0x3c, 0x96, 0xe4, 0x4e, 0xe2, 0xcf, 0x59,
	0x5a, 0x57, 0x8a, 0x56, 0xe8, 0xcc, 0x6b, 0x15, 0x2a, 0x7a, 0x85, 0x46, 0x69, 0xfb, 0x12, 0x83,
	0xa5, 0xb4, 0x17, 0xe3, 0xfc, 0x55, 0xe7, 0x2f, 0xd8, 0xd1, 0x5a, 0x3b, 0x0a, 0xaa, 0x7a, 0xab,
	0x8a, 0x2b, 0x6d, 0xc1, 0xa7, 0xff, 0xc7, 0xc9, 0x8f, 0x62, 0xa7, 0x77, 0xb8, 0x84, 0x79, 0x58,
	0x48, 0xf0, 0x18, 0x5c, 0x01, 0xb2, 0x0d, 0xe7, 0x5f, 0x13, 0x36, 0x6a, 0x0d, 0xe7, 0x74, 0x7b,
	0x53, 0x5c, 0xb3, 0xde, 0x1b, 0xbb, 0xc6, 0x4f, 0xc0, 0x1f, 0xdf, 0xf2, 0x0e, 0x9b, 0x95, 0x1d,
	0x8b, 0xdb, 0xe2, 0xcd, 0xf2, 0x4c, 0xff, 0x3b, 0x7b, 0xf2, 0x51, 0x94, 0x9a, 0xae, 0xc2, 0x62,
	0xcb, 0x66, 0x46, 0x17, 0x0e, 0x3d, 0xae, 0x28, 0x33, 0x58, 0x64, 0xae, 0x2e, 0xb2, 0xbd, 0x57,
	0xb6, 0xf3, 0x5a, 0xf4, 0xe2, 0x12, 0x3c, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x7a, 0x6b,
	0x0f, 0xaf, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AvailabilitySetAgentClient is the client API for AvailabilitySetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvailabilitySetAgentClient interface {
	Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error)
}

type availabilitySetAgentClient struct {
	cc *grpc.ClientConn
}

func NewAvailabilitySetAgentClient(cc *grpc.ClientConn) AvailabilitySetAgentClient {
	return &availabilitySetAgentClient{cc}
}

func (c *availabilitySetAgentClient) Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error) {
	out := new(AvailabilitySetResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvailabilitySetAgentServer is the server API for AvailabilitySetAgent service.
type AvailabilitySetAgentServer interface {
	Invoke(context.Context, *AvailabilitySetRequest) (*AvailabilitySetResponse, error)
}

// UnimplementedAvailabilitySetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAvailabilitySetAgentServer struct {
}

func (*UnimplementedAvailabilitySetAgentServer) Invoke(ctx context.Context, req *AvailabilitySetRequest) (*AvailabilitySetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterAvailabilitySetAgentServer(s *grpc.Server, srv AvailabilitySetAgentServer) {
	s.RegisterService(&_AvailabilitySetAgent_serviceDesc, srv)
}

func _AvailabilitySetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailabilitySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, req.(*AvailabilitySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AvailabilitySetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.compute.AvailabilitySetAgent",
	HandlerType: (*AvailabilitySetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _AvailabilitySetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_availabilityset.proto",
}
