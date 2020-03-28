// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smartpark.proto

/*
Package smartpark is a generated protocol buffer package.

It is generated from these files:
	smartpark.proto

It has these top-level messages:
	ParkingLotEvent
	LicensePlateEvent
	Ack
*/
package smartpark

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type LicensePlateEvent_Direction int32

const (
	LicensePlateEvent_ENTRY LicensePlateEvent_Direction = 0
	LicensePlateEvent_EXIT  LicensePlateEvent_Direction = 1
)

var LicensePlateEvent_Direction_name = map[int32]string{
	0: "ENTRY",
	1: "EXIT",
}
var LicensePlateEvent_Direction_value = map[string]int32{
	"ENTRY": 0,
	"EXIT":  1,
}

func (x LicensePlateEvent_Direction) String() string {
	return proto.EnumName(LicensePlateEvent_Direction_name, int32(x))
}
func (LicensePlateEvent_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ParkingLotEvent struct {
	ParkingLotId int32  `protobuf:"varint,1,opt,name=parking_lot_id,json=parkingLotId" json:"parking_lot_id,omitempty"`
	Plate        string `protobuf:"bytes,2,opt,name=plate" json:"plate,omitempty"`
}

func (m *ParkingLotEvent) Reset()                    { *m = ParkingLotEvent{} }
func (m *ParkingLotEvent) String() string            { return proto.CompactTextString(m) }
func (*ParkingLotEvent) ProtoMessage()               {}
func (*ParkingLotEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ParkingLotEvent) GetParkingLotId() int32 {
	if m != nil {
		return m.ParkingLotId
	}
	return 0
}

func (m *ParkingLotEvent) GetPlate() string {
	if m != nil {
		return m.Plate
	}
	return ""
}

type LicensePlateEvent struct {
	Direction LicensePlateEvent_Direction `protobuf:"varint,1,opt,name=direction,enum=smartpark.LicensePlateEvent_Direction" json:"direction,omitempty"`
	Plate     string                      `protobuf:"bytes,2,opt,name=plate" json:"plate,omitempty"`
}

func (m *LicensePlateEvent) Reset()                    { *m = LicensePlateEvent{} }
func (m *LicensePlateEvent) String() string            { return proto.CompactTextString(m) }
func (*LicensePlateEvent) ProtoMessage()               {}
func (*LicensePlateEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LicensePlateEvent) GetDirection() LicensePlateEvent_Direction {
	if m != nil {
		return m.Direction
	}
	return LicensePlateEvent_ENTRY
}

func (m *LicensePlateEvent) GetPlate() string {
	if m != nil {
		return m.Plate
	}
	return ""
}

type Ack struct {
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ParkingLotEvent)(nil), "smartpark.ParkingLotEvent")
	proto.RegisterType((*LicensePlateEvent)(nil), "smartpark.LicensePlateEvent")
	proto.RegisterType((*Ack)(nil), "smartpark.Ack")
	proto.RegisterEnum("smartpark.LicensePlateEvent_Direction", LicensePlateEvent_Direction_name, LicensePlateEvent_Direction_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SmartParkProto service

type SmartParkProtoClient interface {
	SendParkingLotEvent(ctx context.Context, in *ParkingLotEvent, opts ...grpc.CallOption) (*Ack, error)
	SendLicensePlateEvent(ctx context.Context, in *LicensePlateEvent, opts ...grpc.CallOption) (*Ack, error)
}

type smartParkProtoClient struct {
	cc *grpc.ClientConn
}

func NewSmartParkProtoClient(cc *grpc.ClientConn) SmartParkProtoClient {
	return &smartParkProtoClient{cc}
}

func (c *smartParkProtoClient) SendParkingLotEvent(ctx context.Context, in *ParkingLotEvent, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/smartpark.SmartParkProto/SendParkingLotEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartParkProtoClient) SendLicensePlateEvent(ctx context.Context, in *LicensePlateEvent, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/smartpark.SmartParkProto/SendLicensePlateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SmartParkProto service

type SmartParkProtoServer interface {
	SendParkingLotEvent(context.Context, *ParkingLotEvent) (*Ack, error)
	SendLicensePlateEvent(context.Context, *LicensePlateEvent) (*Ack, error)
}

func RegisterSmartParkProtoServer(s *grpc.Server, srv SmartParkProtoServer) {
	s.RegisterService(&_SmartParkProto_serviceDesc, srv)
}

func _SmartParkProto_SendParkingLotEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParkingLotEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartParkProtoServer).SendParkingLotEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartpark.SmartParkProto/SendParkingLotEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartParkProtoServer).SendParkingLotEvent(ctx, req.(*ParkingLotEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartParkProto_SendLicensePlateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LicensePlateEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartParkProtoServer).SendLicensePlateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartpark.SmartParkProto/SendLicensePlateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartParkProtoServer).SendLicensePlateEvent(ctx, req.(*LicensePlateEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartParkProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartpark.SmartParkProto",
	HandlerType: (*SmartParkProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendParkingLotEvent",
			Handler:    _SmartParkProto_SendParkingLotEvent_Handler,
		},
		{
			MethodName: "SendLicensePlateEvent",
			Handler:    _SmartParkProto_SendLicensePlateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smartpark.proto",
}

func init() { proto.RegisterFile("smartpark.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x5d, 0xd4, 0x8a, 0xbd, 0x48, 0x37, 0xa3, 0xc2, 0x18, 0x3e, 0x94, 0x22, 0xb2, 0xa7, 0x3c,
	0xcc, 0x2f, 0xd8, 0x5c, 0x91, 0xc1, 0x94, 0xd2, 0xed, 0x41, 0x9f, 0x46, 0x4c, 0x82, 0x84, 0x76,
	0x49, 0x48, 0xa3, 0xff, 0xe1, 0x0f, 0xf8, 0xad, 0xd2, 0x16, 0xbb, 0xd2, 0xa2, 0x8f, 0x39, 0xf7,
	0xdc, 0x73, 0xcf, 0x39, 0x81, 0x61, 0xb1, 0xa7, 0xd6, 0x19, 0x6a, 0x33, 0x62, 0xac, 0x76, 0x1a,
	0xfb, 0x0d, 0x10, 0x3d, 0xc1, 0x30, 0xa1, 0x36, 0x93, 0xea, 0x7d, 0xad, 0x5d, 0xfc, 0x29, 0x94,
	0xc3, 0xb7, 0x10, 0x98, 0x1a, 0xda, 0xe5, 0xda, 0xed, 0x24, 0x1f, 0xa3, 0x10, 0x4d, 0xbd, 0xf4,
	0xdc, 0x34, 0xc4, 0x15, 0xc7, 0x57, 0xe0, 0x99, 0x9c, 0x3a, 0x31, 0x3e, 0x0a, 0xd1, 0xd4, 0x4f,
	0xeb, 0x47, 0xf4, 0x85, 0xe0, 0x62, 0x2d, 0x99, 0x50, 0x85, 0x48, 0x4a, 0xa0, 0x56, 0x5c, 0x82,
	0xcf, 0xa5, 0x15, 0xcc, 0x49, 0xad, 0x2a, 0xb1, 0x60, 0x76, 0x47, 0x0e, 0xa6, 0x7a, 0x0b, 0x64,
	0xf9, 0xcb, 0x4e, 0x0f, 0x8b, 0x7f, 0x5c, 0x0c, 0xc1, 0x6f, 0xd8, 0xd8, 0x07, 0x2f, 0x7e, 0xde,
	0xa6, 0xaf, 0xa3, 0x01, 0x3e, 0x83, 0x93, 0xf8, 0x65, 0xb5, 0x1d, 0xa1, 0xc8, 0x83, 0xe3, 0x39,
	0xcb, 0x66, 0xdf, 0x08, 0x82, 0x4d, 0x79, 0xb3, 0xcc, 0x9b, 0x54, 0x3d, 0x3c, 0xc0, 0xe5, 0x46,
	0x28, 0xde, 0x2d, 0x60, 0xd2, 0xf2, 0xd6, 0x99, 0x4d, 0x82, 0xd6, 0x6c, 0xce, 0xb2, 0x68, 0x80,
	0x1f, 0xe1, 0xba, 0x14, 0xe9, 0xa7, 0xbe, 0xf9, 0x2f, 0x62, 0x5f, 0x68, 0x41, 0x60, 0xc2, 0x72,
	0xfd, 0xc1, 0x09, 0xcd, 0x9d, 0xd8, 0xcb, 0xc2, 0xd1, 0x16, 0x65, 0xd1, 0xf1, 0x9e, 0xa0, 0xb7,
	0xd3, 0xea, 0x33, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0x90, 0xb2, 0x2c, 0xdf, 0x01,
	0x00, 0x00,
}
