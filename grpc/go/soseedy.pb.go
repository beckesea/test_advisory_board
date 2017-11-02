// Code generated by protoc-gen-go. DO NOT EDIT.
// source: soseedy.proto

/*
Package soseedy is a generated protocol buffer package.

It is generated from these files:
	soseedy.proto

It has these top-level messages:
	CreateTeacherRequest
	Teacher
*/
package soseedy

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

type CreateTeacherRequest struct {
}

func (m *CreateTeacherRequest) Reset()                    { *m = CreateTeacherRequest{} }
func (m *CreateTeacherRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTeacherRequest) ProtoMessage()               {}
func (*CreateTeacherRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Teacher struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Domain    string `protobuf:"bytes,4,opt,name=domain" json:"domain,omitempty"`
	Token     string `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	Name      string `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	ShortName string `protobuf:"bytes,7,opt,name=short_name,json=shortName" json:"short_name,omitempty"`
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
}

func (m *Teacher) Reset()                    { *m = Teacher{} }
func (m *Teacher) String() string            { return proto.CompactTextString(m) }
func (*Teacher) ProtoMessage()               {}
func (*Teacher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Teacher) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Teacher) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Teacher) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Teacher) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Teacher) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Teacher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Teacher) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Teacher) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTeacherRequest)(nil), "soseedy.CreateTeacherRequest")
	proto.RegisterType((*Teacher)(nil), "soseedy.Teacher")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SoSeedy service

type SoSeedyClient interface {
	CreateTeacher(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*Teacher, error)
}

type soSeedyClient struct {
	cc *grpc.ClientConn
}

func NewSoSeedyClient(cc *grpc.ClientConn) SoSeedyClient {
	return &soSeedyClient{cc}
}

func (c *soSeedyClient) CreateTeacher(ctx context.Context, in *CreateTeacherRequest, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := grpc.Invoke(ctx, "/soseedy.SoSeedy/CreateTeacher", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SoSeedy service

type SoSeedyServer interface {
	CreateTeacher(context.Context, *CreateTeacherRequest) (*Teacher, error)
}

func RegisterSoSeedyServer(s *grpc.Server, srv SoSeedyServer) {
	s.RegisterService(&_SoSeedy_serviceDesc, srv)
}

func _SoSeedy_CreateTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoSeedyServer).CreateTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soseedy.SoSeedy/CreateTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoSeedyServer).CreateTeacher(ctx, req.(*CreateTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SoSeedy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soseedy.SoSeedy",
	HandlerType: (*SoSeedyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeacher",
			Handler:    _SoSeedy_CreateTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "soseedy.proto",
}

func init() { proto.RegisterFile("soseedy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xed, 0x5c, 0xdb, 0xed, 0xe0, 0x44, 0x0e, 0x63, 0x86, 0xc1, 0x60, 0xf4, 0x6a, 0xde,
	0xf4, 0x42, 0xdf, 0xa0, 0x5e, 0x2b, 0xa3, 0xd3, 0xeb, 0x11, 0x9b, 0x03, 0x2b, 0xae, 0xcd, 0x3c,
	0x49, 0x14, 0x1f, 0xd3, 0x37, 0x92, 0xa6, 0xe9, 0x40, 0xf0, 0x2e, 0xff, 0xf7, 0x05, 0x0e, 0xff,
	0x0f, 0x33, 0xa3, 0x0d, 0x91, 0xfa, 0xce, 0x4f, 0xac, 0xad, 0xc6, 0x34, 0xc4, 0x6c, 0x01, 0xf3,
	0x47, 0x26, 0x69, 0xe9, 0x85, 0x64, 0x75, 0x20, 0x2e, 0xe9, 0xc3, 0x91, 0xb1, 0xd9, 0x4f, 0x04,
	0x69, 0x40, 0x78, 0x0d, 0xa3, 0x5a, 0x89, 0x68, 0x1d, 0x6d, 0xe2, 0x72, 0x54, 0x2b, 0x5c, 0xc2,
	0xc4, 0x19, 0xe2, 0x56, 0x36, 0x24, 0x46, 0xeb, 0x68, 0x33, 0x2d, 0xcf, 0xb9, 0x73, 0x27, 0x69,
	0xcc, 0x97, 0x66, 0x25, 0x2e, 0x7b, 0x37, 0x64, 0x5c, 0x40, 0xa2, 0x74, 0x23, 0xeb, 0x56, 0x8c,
	0xbd, 0x09, 0x09, 0xe7, 0x10, 0x5b, 0xfd, 0x4e, 0xad, 0x88, 0x3d, 0xee, 0x03, 0x22, 0x8c, 0xfd,
	0x85, 0xc4, 0x43, 0xff, 0xc6, 0x15, 0x80, 0x39, 0x68, 0xb6, 0x7b, 0x6f, 0x52, 0x6f, 0xa6, 0x9e,
	0x3c, 0x07, 0x2d, 0x3f, 0xa5, 0x95, 0xbc, 0x77, 0x7c, 0x14, 0x93, 0x5e, 0xf7, 0xe4, 0x95, 0x8f,
	0xf7, 0x4f, 0x90, 0xee, 0xf4, 0xae, 0xab, 0x8d, 0x05, 0xcc, 0xfe, 0xd4, 0xc6, 0x55, 0x3e, 0x0c,
	0xf4, 0xdf, 0x1c, 0xcb, 0x9b, 0xb3, 0x0e, 0x22, 0xbb, 0x28, 0xee, 0xe0, 0xb6, 0xd2, 0x4d, 0x5e,
	0xb7, 0xc6, 0xb2, 0xab, 0xac, 0x63, 0x1a, 0x3e, 0x15, 0x57, 0xe1, 0xce, 0xb6, 0x1b, 0x7b, 0x1b,
	0xbd, 0x25, 0x7e, 0xf5, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x69, 0x0f, 0x89, 0x86,
	0x01, 0x00, 0x00,
}
