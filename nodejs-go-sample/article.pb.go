// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package article

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

type Article_Status int32

const (
	Article_DRAFT   Article_Status = 0
	Article_PUBLISH Article_Status = 1
)

var Article_Status_name = map[int32]string{
	0: "DRAFT",
	1: "PUBLISH",
}

var Article_Status_value = map[string]int32{
	"DRAFT":   0,
	"PUBLISH": 1,
}

func (x Article_Status) String() string {
	return proto.EnumName(Article_Status_name, int32(x))
}

func (Article_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{1, 0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Article struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string         `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status               Article_Status `protobuf:"varint,4,opt,name=status,proto3,enum=article.Article_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{1}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetStatus() Article_Status {
	if m != nil {
		return m.Status
	}
	return Article_DRAFT
}

func init() {
	proto.RegisterEnum("article.Article_Status", Article_Status_name, Article_Status_value)
	proto.RegisterType((*Empty)(nil), "article.Empty")
	proto.RegisterType((*Article)(nil), "article.Article")
}

func init() {
	proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2)
}

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x37, 0x6b, 0xb3, 0xa1, 0x23, 0x2e, 0xcb, 0x20, 0x18, 0x7a, 0x0a, 0x39, 0x2d, 0x0a,
	0x2b, 0xd4, 0x27, 0xa8, 0x68, 0x51, 0xf0, 0x50, 0xb2, 0xfa, 0x00, 0x75, 0x9b, 0x43, 0xa0, 0x36,
	0x25, 0x19, 0x05, 0xdf, 0xc5, 0x87, 0x15, 0xb2, 0xa9, 0x87, 0xf5, 0xf8, 0xff, 0xff, 0xfc, 0xf3,
	0x31, 0x03, 0x17, 0xdb, 0x40, 0x6e, 0xd8, 0xdb, 0xee, 0x18, 0x3c, 0x79, 0x14, 0x59, 0x6a, 0x01,
	0xfc, 0xf1, 0xe3, 0x48, 0xdf, 0xfa, 0x87, 0x81, 0x58, 0x8d, 0x26, 0xd6, 0x50, 0xba, 0x9d, 0x64,
	0x8a, 0xb5, 0xdc, 0x94, 0x6e, 0x87, 0x97, 0xc0, 0xc9, 0xd1, 0xde, 0xca, 0x52, 0xb1, 0x76, 0x6e,
	0x46, 0x81, 0x12, 0xc4, 0xe0, 0x0f, 0x64, 0x0f, 0x24, 0xcf, 0x92, 0x7f, 0x92, 0x78, 0x0b, 0x55,
	0xa4, 0x2d, 0x7d, 0x46, 0x39, 0x53, 0xac, 0xad, 0x97, 0x57, 0xdd, 0x89, 0x9e, 0x09, 0x5d, 0x9f,
	0x62, 0x93, 0xc7, 0xb4, 0x82, 0x6a, 0x74, 0x70, 0x0e, 0xfc, 0xc1, 0xac, 0xd6, 0xaf, 0x4d, 0x81,
	0xe7, 0x20, 0x36, 0x6f, 0xf7, 0x2f, 0xcf, 0xfd, 0x53, 0xc3, 0x96, 0x0e, 0xea, 0xdc, 0xed, 0x6d,
	0xf8, 0x72, 0x83, 0xc5, 0x1b, 0xe0, 0x6b, 0x17, 0x22, 0x61, 0xfd, 0xb7, 0x3d, 0x5d, 0xb2, 0x68,
	0xa6, 0x34, 0x5d, 0xe0, 0x35, 0xcc, 0x36, 0x3e, 0x12, 0xfe, 0xcb, 0x16, 0x93, 0xb6, 0x2e, 0xde,
	0xab, 0xf4, 0xa2, 0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xc3, 0x79, 0xd8, 0x33, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleServiceClient interface {
	First(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Article, error)
	Post(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Empty, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) First(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/article.ArticleService/First", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) Post(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/article.ArticleService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
type ArticleServiceServer interface {
	First(context.Context, *Empty) (*Article, error)
	Post(context.Context, *Article) (*Empty, error)
}

// UnimplementedArticleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (*UnimplementedArticleServiceServer) First(ctx context.Context, req *Empty) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method First not implemented")
}
func (*UnimplementedArticleServiceServer) Post(ctx context.Context, req *Article) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}

func RegisterArticleServiceServer(s *grpc.Server, srv ArticleServiceServer) {
	s.RegisterService(&_ArticleService_serviceDesc, srv)
}

func _ArticleService_First_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).First(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/First",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).First(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.ArticleService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).Post(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "First",
			Handler:    _ArticleService_First_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _ArticleService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
