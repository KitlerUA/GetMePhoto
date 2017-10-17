// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graber.proto

/*
Package graber is a generated protocol buffer package.

It is generated from these files:
	graber.proto

It has these top-level messages:
	Tag
	Id
	Photo
*/
package graber

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

type Tag struct {
	Tag string `protobuf:"bytes,1,opt,name=Tag,json=tag" json:"Tag,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Id struct {
	Url string `protobuf:"bytes,1,opt,name=Url,json=url" json:"Url,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Photo struct {
	Image []byte `protobuf:"bytes,1,opt,name=Image,json=image,proto3" json:"Image,omitempty"`
}

func (m *Photo) Reset()                    { *m = Photo{} }
func (m *Photo) String() string            { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()               {}
func (*Photo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Photo) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "graber.Tag")
	proto.RegisterType((*Id)(nil), "graber.Id")
	proto.RegisterType((*Photo)(nil), "graber.Photo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GragPhoto service

type GragPhotoClient interface {
	Get(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Photo, error)
}

type gragPhotoClient struct {
	cc *grpc.ClientConn
}

func NewGragPhotoClient(cc *grpc.ClientConn) GragPhotoClient {
	return &gragPhotoClient{cc}
}

func (c *gragPhotoClient) Get(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := grpc.Invoke(ctx, "/graber.GragPhoto/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GragPhoto service

type GragPhotoServer interface {
	Get(context.Context, *Tag) (*Photo, error)
}

func RegisterGragPhotoServer(s *grpc.Server, srv GragPhotoServer) {
	s.RegisterService(&_GragPhoto_serviceDesc, srv)
}

func _GragPhoto_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GragPhotoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graber.GragPhoto/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GragPhotoServer).Get(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

var _GragPhoto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graber.GragPhoto",
	HandlerType: (*GragPhotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GragPhoto_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "graber.proto",
}

// Client API for DownloadById service

type DownloadByIdClient interface {
	Download(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Photo, error)
}

type downloadByIdClient struct {
	cc *grpc.ClientConn
}

func NewDownloadByIdClient(cc *grpc.ClientConn) DownloadByIdClient {
	return &downloadByIdClient{cc}
}

func (c *downloadByIdClient) Download(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := grpc.Invoke(ctx, "/graber.DownloadById/Download", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownloadById service

type DownloadByIdServer interface {
	Download(context.Context, *Id) (*Photo, error)
}

func RegisterDownloadByIdServer(s *grpc.Server, srv DownloadByIdServer) {
	s.RegisterService(&_DownloadById_serviceDesc, srv)
}

func _DownloadById_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadByIdServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graber.DownloadById/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadByIdServer).Download(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownloadById_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graber.DownloadById",
	HandlerType: (*DownloadByIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _DownloadById_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "graber.proto",
}

func init() { proto.RegisterFile("graber.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2f, 0x4a, 0x4c,
	0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xc4, 0xb9, 0x98,
	0x43, 0x12, 0xd3, 0x85, 0x04, 0xc0, 0x94, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x73, 0x49,
	0x62, 0xba, 0x92, 0x18, 0x17, 0x93, 0x67, 0x0a, 0x48, 0x3c, 0xb4, 0x28, 0x07, 0x26, 0x5e, 0x5a,
	0x94, 0xa3, 0x24, 0xcb, 0xc5, 0x1a, 0x90, 0x01, 0x32, 0x41, 0x84, 0x8b, 0xd5, 0x33, 0x37, 0x31,
	0x3d, 0x15, 0x2c, 0xc9, 0x13, 0xc4, 0x9a, 0x09, 0xe2, 0x18, 0x19, 0x70, 0x71, 0xba, 0x17, 0x25,
	0xa6, 0x43, 0x94, 0x28, 0x73, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0x71, 0xeb, 0x41, 0xad, 0x0e, 0x49,
	0x4c, 0x97, 0xe2, 0x85, 0x71, 0xc0, 0x4a, 0x94, 0x18, 0x8c, 0xcc, 0xb9, 0x78, 0x5c, 0xf2, 0xcb,
	0xf3, 0x72, 0xf2, 0x13, 0x53, 0x9c, 0x2a, 0x3d, 0x53, 0x84, 0xd4, 0xb9, 0x38, 0x60, 0x7c, 0x21,
	0x2e, 0x98, 0x62, 0xcf, 0x14, 0x0c, 0x8d, 0x49, 0x6c, 0x60, 0x9f, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0xd0, 0x5b, 0x4d, 0xd9, 0x00, 0x00, 0x00,
}