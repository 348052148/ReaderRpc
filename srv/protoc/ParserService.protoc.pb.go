// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoc/ParserService.protoc

package srv

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

//标注+类型+属性名+属性顺序号+[默认值]
type ChapterRequest struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterRequest) Reset()         { *m = ChapterRequest{} }
func (m *ChapterRequest) String() string { return proto.CompactTextString(m) }
func (*ChapterRequest) ProtoMessage()    {}
func (*ChapterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f162124cde3fd4d1, []int{0}
}

func (m *ChapterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterRequest.Unmarshal(m, b)
}
func (m *ChapterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterRequest.Marshal(b, m, deterministic)
}
func (m *ChapterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterRequest.Merge(m, src)
}
func (m *ChapterRequest) XXX_Size() int {
	return xxx_messageInfo_ChapterRequest.Size(m)
}
func (m *ChapterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterRequest proto.InternalMessageInfo

func (m *ChapterRequest) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type ChapterResponse struct {
	Chapters             []*ChapterResponse_Chapter `protobuf:"bytes,1,rep,name=chapters,proto3" json:"chapters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ChapterResponse) Reset()         { *m = ChapterResponse{} }
func (m *ChapterResponse) String() string { return proto.CompactTextString(m) }
func (*ChapterResponse) ProtoMessage()    {}
func (*ChapterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f162124cde3fd4d1, []int{1}
}

func (m *ChapterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterResponse.Unmarshal(m, b)
}
func (m *ChapterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterResponse.Marshal(b, m, deterministic)
}
func (m *ChapterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterResponse.Merge(m, src)
}
func (m *ChapterResponse) XXX_Size() int {
	return xxx_messageInfo_ChapterResponse.Size(m)
}
func (m *ChapterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterResponse proto.InternalMessageInfo

func (m *ChapterResponse) GetChapters() []*ChapterResponse_Chapter {
	if m != nil {
		return m.Chapters
	}
	return nil
}

type ChapterResponse_Chapter struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	ContentsLink         string   `protobuf:"bytes,3,opt,name=contentsLink,proto3" json:"contentsLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterResponse_Chapter) Reset()         { *m = ChapterResponse_Chapter{} }
func (m *ChapterResponse_Chapter) String() string { return proto.CompactTextString(m) }
func (*ChapterResponse_Chapter) ProtoMessage()    {}
func (*ChapterResponse_Chapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f162124cde3fd4d1, []int{1, 0}
}

func (m *ChapterResponse_Chapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterResponse_Chapter.Unmarshal(m, b)
}
func (m *ChapterResponse_Chapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterResponse_Chapter.Marshal(b, m, deterministic)
}
func (m *ChapterResponse_Chapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterResponse_Chapter.Merge(m, src)
}
func (m *ChapterResponse_Chapter) XXX_Size() int {
	return xxx_messageInfo_ChapterResponse_Chapter.Size(m)
}
func (m *ChapterResponse_Chapter) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterResponse_Chapter.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterResponse_Chapter proto.InternalMessageInfo

func (m *ChapterResponse_Chapter) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ChapterResponse_Chapter) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ChapterResponse_Chapter) GetContentsLink() string {
	if m != nil {
		return m.ContentsLink
	}
	return ""
}

func init() {
	proto.RegisterType((*ChapterRequest)(nil), "srv.ChapterRequest")
	proto.RegisterType((*ChapterResponse)(nil), "srv.ChapterResponse")
	proto.RegisterType((*ChapterResponse_Chapter)(nil), "srv.ChapterResponse.Chapter")
}

func init() { proto.RegisterFile("protoc/ParserService.protoc", fileDescriptor_f162124cde3fd4d1) }

var fileDescriptor_f162124cde3fd4d1 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xd6, 0x0f, 0x48, 0x2c, 0x2a, 0x4e, 0x2d, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x83, 0x08, 0x0a, 0x31, 0x17, 0x17, 0x95, 0x29, 0xa9, 0x70, 0xf1, 0x39, 0x67, 0x24, 0x16,
	0x94, 0xa4, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x64,
	0xe6, 0x65, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xcb, 0x18, 0xb9, 0xf8,
	0xe1, 0xca, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x2c, 0xb8, 0x38, 0x92, 0x21, 0x42, 0xc5,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x32, 0x7a, 0xc5, 0x45, 0x65, 0x7a, 0x68, 0xea, 0xe0,
	0x7c, 0xb8, 0x6a, 0xa9, 0x48, 0x2e, 0x76, 0xa8, 0xa0, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49,
	0x4e, 0x2a, 0xd4, 0x36, 0x08, 0x07, 0x24, 0x9a, 0x99, 0x97, 0x92, 0x5a, 0x21, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x1a, 0x04, 0xe1, 0x08, 0x29, 0x71, 0xf1, 0x24, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95,
	0x14, 0xfb, 0x80, 0x1c, 0xc8, 0x0c, 0xd6, 0x82, 0x22, 0x66, 0xe4, 0xc7, 0xc5, 0x8b, 0xe2, 0x57,
	0x21, 0x5b, 0x2e, 0x3e, 0x88, 0x00, 0xd4, 0xc6, 0x62, 0x21, 0x61, 0x54, 0x57, 0x82, 0x3d, 0x2d,
	0x25, 0x82, 0xcd, 0xe9, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xb0, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0xf8, 0xd6, 0x4e, 0x49, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ParserServiceClient is the client API for ParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParserServiceClient interface {
	ParserChapters(ctx context.Context, in *ChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error)
}

type parserServiceClient struct {
	cc *grpc.ClientConn
}

func NewParserServiceClient(cc *grpc.ClientConn) ParserServiceClient {
	return &parserServiceClient{cc}
}

func (c *parserServiceClient) ParserChapters(ctx context.Context, in *ChapterRequest, opts ...grpc.CallOption) (*ChapterResponse, error) {
	out := new(ChapterResponse)
	err := c.cc.Invoke(ctx, "/srv.ParserService/ParserChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServiceServer is the server API for ParserService service.
type ParserServiceServer interface {
	ParserChapters(context.Context, *ChapterRequest) (*ChapterResponse, error)
}

// UnimplementedParserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParserServiceServer struct {
}

func (*UnimplementedParserServiceServer) ParserChapters(ctx context.Context, req *ChapterRequest) (*ChapterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParserChapters not implemented")
}

func RegisterParserServiceServer(s *grpc.Server, srv ParserServiceServer) {
	s.RegisterService(&_ParserService_serviceDesc, srv)
}

func _ParserService_ParserChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServiceServer).ParserChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srv.ParserService/ParserChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServiceServer).ParserChapters(ctx, req.(*ChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "srv.ParserService",
	HandlerType: (*ParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParserChapters",
			Handler:    _ParserService_ParserChapters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/ParserService.protoc",
}
