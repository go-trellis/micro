// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ServiceType int32

const (
	ServiceType_LOCAL ServiceType = 0
	ServiceType_HTTP  ServiceType = 1
	ServiceType_GRPC  ServiceType = 2
)

var ServiceType_name = map[int32]string{
	0: "LOCAL",
	1: "HTTP",
	2: "GRPC",
}

var ServiceType_value = map[string]int32{
	"LOCAL": 0,
	"HTTP":  1,
	"GRPC":  2,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

// Service 基础服务信息
type Service struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type RegistOption struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Protocol             string            `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Domain               string            `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Weight               uint32            `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegistOption) Reset()         { *m = RegistOption{} }
func (m *RegistOption) String() string { return proto.CompactTextString(m) }
func (*RegistOption) ProtoMessage()    {}
func (*RegistOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *RegistOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistOption.Unmarshal(m, b)
}
func (m *RegistOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistOption.Marshal(b, m, deterministic)
}
func (m *RegistOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistOption.Merge(m, src)
}
func (m *RegistOption) XXX_Size() int {
	return xxx_messageInfo_RegistOption.Size(m)
}
func (m *RegistOption) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistOption.DiscardUnknown(m)
}

var xxx_messageInfo_RegistOption proto.InternalMessageInfo

func (m *RegistOption) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegistOption) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegistOption) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *RegistOption) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RegistOption) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *RegistOption) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Payload struct {
	TraceId              string            `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Service              *Service          `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Topic                string            `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Header               map[string]string `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReqBody              []byte            `protobuf:"bytes,7,opt,name=req_body,json=reqBody,proto3" json:"req_body,omitempty"`
	TraceIp              string            `protobuf:"bytes,8,opt,name=trace_ip,json=traceIp,proto3" json:"trace_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Payload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Payload) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Payload) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Payload) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Payload) GetReqBody() []byte {
	if m != nil {
		return m.ReqBody
	}
	return nil
}

func (m *Payload) GetTraceIp() string {
	if m != nil {
		return m.TraceIp
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.ServiceType", ServiceType_name, ServiceType_value)
	proto.RegisterType((*Service)(nil), "proto.Service")
	proto.RegisterType((*RegistOption)(nil), "proto.RegistOption")
	proto.RegisterMapType((map[string]string)(nil), "proto.RegistOption.MetadataEntry")
	proto.RegisterType((*Payload)(nil), "proto.Payload")
	proto.RegisterMapType((map[string]string)(nil), "proto.Payload.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Payload.MetadataEntry")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xda, 0x34, 0xd9, 0xd7, 0x6d, 0x29, 0x8f, 0x45, 0xc6, 0xe0, 0xa1, 0xf6, 0x20,
	0x41, 0xa4, 0x87, 0x2a, 0xb8, 0x2a, 0x1e, 0xb4, 0x88, 0x2b, 0xac, 0x6c, 0x19, 0x7b, 0x5f, 0x66,
	0x3b, 0x8f, 0xee, 0x60, 0x9a, 0xc9, 0x4e, 0xc7, 0x4a, 0xee, 0xfe, 0x01, 0xfe, 0xc9, 0xd2, 0xc9,
	0x6c, 0xc8, 0xae, 0x78, 0x28, 0x7b, 0xca, 0xfb, 0xde, 0xcb, 0x07, 0xdf, 0xef, 0x1b, 0x18, 0x6c,
	0x68, 0xbb, 0x15, 0x6b, 0x9a, 0x96, 0x46, 0x5b, 0x8d, 0x91, 0xfb, 0x4c, 0xde, 0x40, 0xfc, 0x9d,
	0xcc, 0x4e, 0xad, 0x08, 0x11, 0xba, 0x85, 0xd8, 0x10, 0x0b, 0xc6, 0x41, 0x76, 0xc4, 0xdd, 0x8c,
	0x0c, 0xe2, 0x1d, 0x99, 0xad, 0xd2, 0x05, 0x0b, 0xdd, 0xfa, 0x56, 0x4e, 0x7e, 0x87, 0x70, 0xcc,
	0x69, 0xad, 0xb6, 0xf6, 0xa2, 0xb4, 0x4a, 0x17, 0x87, 0xd9, 0x31, 0x85, 0xc4, 0x05, 0x58, 0xe9,
	0x9c, 0x75, 0xdc, 0xa9, 0xd1, 0xf8, 0x18, 0x7a, 0x52, 0x6f, 0x84, 0x2a, 0x58, 0xd7, 0x5d, 0xbc,
	0xda, 0xef, 0x7f, 0x91, 0x5a, 0x5f, 0x5b, 0x16, 0x8d, 0x83, 0x6c, 0xc0, 0xbd, 0xc2, 0x0f, 0x90,
	0x6c, 0xc8, 0x0a, 0x29, 0xac, 0x60, 0xbd, 0x71, 0x27, 0xeb, 0xcf, 0x9e, 0xd5, 0x90, 0xd3, 0x76,
	0xc0, 0xe9, 0x37, 0xff, 0xcf, 0xe7, 0xc2, 0x9a, 0x8a, 0x37, 0x96, 0xf4, 0x3d, 0x0c, 0xee, 0x9c,
	0x70, 0x04, 0x9d, 0x1f, 0x54, 0x79, 0x90, 0xfd, 0x88, 0x27, 0x10, 0xed, 0x44, 0xfe, 0x93, 0x3c,
	0x45, 0x2d, 0xde, 0x85, 0xa7, 0xc1, 0xe4, 0x4f, 0x07, 0xe2, 0x85, 0xa8, 0x72, 0x2d, 0x24, 0x3e,
	0x81, 0xc4, 0x1a, 0xb1, 0xa2, 0x4b, 0x25, 0xbd, 0x39, 0x76, 0xfa, 0xab, 0xc4, 0x21, 0x84, 0x4a,
	0x7a, 0x77, 0xa8, 0x24, 0x66, 0x4d, 0xed, 0x8e, 0xbe, 0x3f, 0x1b, 0xfa, 0xc4, 0x7e, 0xcb, 0x9b,
	0x57, 0x39, 0x81, 0xc8, 0xea, 0x52, 0xad, 0x7c, 0x17, 0xb5, 0xc0, 0x19, 0xf4, 0xae, 0x49, 0x48,
	0x32, 0x2c, 0x72, 0xc0, 0xa9, 0xb7, 0xfb, 0x28, 0xd3, 0x33, 0x77, 0xac, 0x49, 0xfd, 0x9f, 0x78,
	0xfa, 0x4f, 0x4d, 0x4f, 0xef, 0xb9, 0xfe, 0xd3, 0xd0, 0x1e, 0xcc, 0xd0, 0xcd, 0xe5, 0x95, 0x96,
	0x15, 0x8b, 0xc7, 0x41, 0x76, 0xcc, 0x63, 0x43, 0x37, 0x9f, 0xb4, 0xac, 0x5a, 0xcc, 0x25, 0x4b,
	0xda, 0xcc, 0x65, 0xfa, 0x16, 0xfa, 0xad, 0x18, 0x87, 0xb4, 0xfa, 0xa0, 0x27, 0x79, 0xf1, 0x12,
	0xfa, 0xbe, 0xbc, 0x65, 0x55, 0x12, 0x1e, 0x41, 0x74, 0x7e, 0x31, 0xff, 0x78, 0x3e, 0x7a, 0x84,
	0x09, 0x74, 0xcf, 0x96, 0xcb, 0xc5, 0x28, 0xd8, 0x4f, 0x5f, 0xf8, 0x62, 0x3e, 0x0a, 0x67, 0xaf,
	0x01, 0xf8, 0x62, 0x7e, 0xdb, 0xf6, 0x73, 0xe8, 0xce, 0x45, 0x9e, 0xe3, 0xf0, 0x6e, 0x33, 0xe9,
	0x3d, 0x7d, 0xd5, 0x73, 0xf2, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xdf, 0xe1, 0x24,
	0x55, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCServiceClient is the client API for RPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCServiceClient interface {
	Call(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
}

type rPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewRPCServiceClient(cc *grpc.ClientConn) RPCServiceClient {
	return &rPCServiceClient{cc}
}

func (c *rPCServiceClient) Call(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := c.cc.Invoke(ctx, "/proto.RPCService/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServiceServer is the server API for RPCService service.
type RPCServiceServer interface {
	Call(context.Context, *Payload) (*Payload, error)
}

// UnimplementedRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServiceServer struct {
}

func (*UnimplementedRPCServiceServer) Call(ctx context.Context, req *Payload) (*Payload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterRPCServiceServer(s *grpc.Server, srv RPCServiceServer) {
	s.RegisterService(&_RPCService_serviceDesc, srv)
}

func _RPCService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RPCService/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServiceServer).Call(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RPCService",
	HandlerType: (*RPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _RPCService_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
