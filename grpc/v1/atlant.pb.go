// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atlant.proto

package v1

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListRequest_SortingOption_SortingOptionDirection int32

const (
	ListRequest_SortingOption_SORTING_OPTION_UNSPECIFIED ListRequest_SortingOption_SortingOptionDirection = 0
	ListRequest_SortingOption_SORTING_OPTION_ASC         ListRequest_SortingOption_SortingOptionDirection = 1
	ListRequest_SortingOption_SORTING_OPTION_DESC        ListRequest_SortingOption_SortingOptionDirection = 2
)

var ListRequest_SortingOption_SortingOptionDirection_name = map[int32]string{
	0: "SORTING_OPTION_UNSPECIFIED",
	1: "SORTING_OPTION_ASC",
	2: "SORTING_OPTION_DESC",
}

var ListRequest_SortingOption_SortingOptionDirection_value = map[string]int32{
	"SORTING_OPTION_UNSPECIFIED": 0,
	"SORTING_OPTION_ASC":         1,
	"SORTING_OPTION_DESC":        2,
}

func (x ListRequest_SortingOption_SortingOptionDirection) String() string {
	return proto.EnumName(ListRequest_SortingOption_SortingOptionDirection_name, int32(x))
}

func (ListRequest_SortingOption_SortingOptionDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{1, 0, 0}
}

type FetchRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{0}
}
func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return m.Size()
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ListRequest struct {
	Start   int32                        `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit   int32                        `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Options []*ListRequest_SortingOption `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{1}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOptions() []*ListRequest_SortingOption {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListRequest_SortingOption struct {
	Field     string                                           `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Direction ListRequest_SortingOption_SortingOptionDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=v1.ListRequest_SortingOption_SortingOptionDirection" json:"direction,omitempty"`
}

func (m *ListRequest_SortingOption) Reset()         { *m = ListRequest_SortingOption{} }
func (m *ListRequest_SortingOption) String() string { return proto.CompactTextString(m) }
func (*ListRequest_SortingOption) ProtoMessage()    {}
func (*ListRequest_SortingOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{1, 0}
}
func (m *ListRequest_SortingOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRequest_SortingOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRequest_SortingOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRequest_SortingOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest_SortingOption.Merge(m, src)
}
func (m *ListRequest_SortingOption) XXX_Size() int {
	return m.Size()
}
func (m *ListRequest_SortingOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest_SortingOption.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest_SortingOption proto.InternalMessageInfo

func (m *ListRequest_SortingOption) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ListRequest_SortingOption) GetDirection() ListRequest_SortingOption_SortingOptionDirection {
	if m != nil {
		return m.Direction
	}
	return ListRequest_SortingOption_SORTING_OPTION_UNSPECIFIED
}

type ListResponse struct {
	Products []*ListResponse_Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{2}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetProducts() []*ListResponse_Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type ListResponse_Product struct {
	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price       float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt   int64   `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64   `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdateCount int64   `protobuf:"varint,5,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"`
}

func (m *ListResponse_Product) Reset()         { *m = ListResponse_Product{} }
func (m *ListResponse_Product) String() string { return proto.CompactTextString(m) }
func (*ListResponse_Product) ProtoMessage()    {}
func (*ListResponse_Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_d796a2e9de1ebfea, []int{2, 0}
}
func (m *ListResponse_Product) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResponse_Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResponse_Product.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResponse_Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse_Product.Merge(m, src)
}
func (m *ListResponse_Product) XXX_Size() int {
	return m.Size()
}
func (m *ListResponse_Product) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse_Product.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse_Product proto.InternalMessageInfo

func (m *ListResponse_Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ListResponse_Product) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ListResponse_Product) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ListResponse_Product) GetUpdateCount() int64 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func init() {
	proto.RegisterEnum("v1.ListRequest_SortingOption_SortingOptionDirection", ListRequest_SortingOption_SortingOptionDirection_name, ListRequest_SortingOption_SortingOptionDirection_value)
	proto.RegisterType((*FetchRequest)(nil), "v1.FetchRequest")
	proto.RegisterType((*ListRequest)(nil), "v1.ListRequest")
	proto.RegisterType((*ListRequest_SortingOption)(nil), "v1.ListRequest.SortingOption")
	proto.RegisterType((*ListResponse)(nil), "v1.ListResponse")
	proto.RegisterType((*ListResponse_Product)(nil), "v1.ListResponse.Product")
}

func init() { proto.RegisterFile("atlant.proto", fileDescriptor_d796a2e9de1ebfea) }

var fileDescriptor_d796a2e9de1ebfea = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0xf9, 0x43, 0xc9, 0x24, 0x05, 0x6b, 0x41, 0xc1, 0x32, 0xaa, 0x15, 0x72, 0x0a,
	0x17, 0x57, 0x09, 0x95, 0x38, 0x70, 0x0a, 0x49, 0x8a, 0x22, 0xa1, 0x24, 0x5a, 0x97, 0x0b, 0x97,
	0xc8, 0xb5, 0xb7, 0xc1, 0x92, 0xe3, 0x35, 0xf6, 0x38, 0x12, 0x6f, 0xc1, 0x81, 0x87, 0xe2, 0xd8,
	0x23, 0x88, 0x0b, 0x4a, 0x78, 0x10, 0xb4, 0xbb, 0x36, 0x6d, 0x23, 0xd4, 0xdb, 0x7e, 0xdf, 0xfc,
	0xbc, 0xb3, 0x33, 0xfe, 0xa0, 0xed, 0x61, 0xe4, 0xc5, 0xe8, 0x24, 0xa9, 0x40, 0x41, 0xab, 0xdb,
	0x81, 0xf5, 0x7c, 0x2d, 0xc4, 0x3a, 0xe2, 0xa7, 0xca, 0xb9, 0xcc, 0xaf, 0x4e, 0xf9, 0x26, 0xc1,
	0x2f, 0x1a, 0xe8, 0x75, 0xa1, 0x7d, 0xce, 0xd1, 0xff, 0xc4, 0xf8, 0xe7, 0x9c, 0x67, 0x48, 0x0d,
	0xa8, 0xe5, 0x69, 0x64, 0x92, 0x2e, 0xe9, 0x37, 0x99, 0x3c, 0xf6, 0x7e, 0x56, 0xa1, 0xf5, 0x3e,
	0xcc, 0xb0, 0x24, 0x9e, 0x42, 0x23, 0x43, 0x2f, 0x45, 0xc5, 0x34, 0x98, 0x16, 0xd2, 0x8d, 0xc2,
	0x4d, 0x88, 0x66, 0x55, 0xbb, 0x4a, 0xd0, 0xd7, 0x70, 0x24, 0x12, 0x0c, 0x45, 0x9c, 0x99, 0xb5,
	0x6e, 0xad, 0xdf, 0x1a, 0x9e, 0x38, 0xdb, 0x81, 0x73, 0xeb, 0x36, 0xc7, 0x15, 0x29, 0x86, 0xf1,
	0x7a, 0xa1, 0x28, 0x56, 0xd2, 0xd6, 0x1f, 0x02, 0xc7, 0x77, 0x4a, 0xb2, 0xc1, 0x55, 0xc8, 0xa3,
	0xa0, 0x78, 0x9a, 0x16, 0x94, 0x41, 0x33, 0x08, 0x53, 0xee, 0x4b, 0x44, 0xb5, 0x7e, 0x34, 0x3c,
	0xbb, 0xb7, 0xc5, 0x5d, 0x35, 0x29, 0xbf, 0x65, 0x37, 0xd7, 0xf4, 0x42, 0xe8, 0xfc, 0x1f, 0xa2,
	0x36, 0x58, 0xee, 0x82, 0x5d, 0xcc, 0xe6, 0xef, 0x56, 0x8b, 0xe5, 0xc5, 0x6c, 0x31, 0x5f, 0x7d,
	0x98, 0xbb, 0xcb, 0xe9, 0x78, 0x76, 0x3e, 0x9b, 0x4e, 0x8c, 0x0a, 0xed, 0x00, 0x3d, 0xa8, 0x8f,
	0xdc, 0xb1, 0x41, 0xe8, 0x33, 0x78, 0x72, 0xe0, 0x4f, 0xa6, 0xee, 0xd8, 0xa8, 0xf6, 0x7e, 0x11,
	0x68, 0xeb, 0xa7, 0x66, 0x89, 0x88, 0x33, 0x4e, 0xcf, 0xe0, 0x61, 0x92, 0x8a, 0x20, 0xf7, 0x31,
	0x33, 0x89, 0xda, 0x98, 0x79, 0x33, 0x8e, 0x66, 0x9c, 0xa5, 0x06, 0xd8, 0x3f, 0xd2, 0xfa, 0x46,
	0xe0, 0xa8, 0x70, 0x29, 0x85, 0x7a, 0xec, 0x6d, 0x78, 0xb1, 0x26, 0x75, 0x96, 0xbb, 0x4b, 0xd2,
	0xd0, 0xe7, 0x6a, 0x43, 0x84, 0x69, 0x41, 0x4f, 0x00, 0xfc, 0x94, 0x7b, 0xc8, 0x83, 0x95, 0x87,
	0x66, 0xad, 0x4b, 0xfa, 0x35, 0xd6, 0x2c, 0x9c, 0x11, 0xca, 0x72, 0x9e, 0x04, 0x65, 0xb9, 0xae,
	0xcb, 0x85, 0x33, 0x42, 0xfa, 0x02, 0xda, 0x5a, 0xac, 0x7c, 0x91, 0xc7, 0x68, 0x36, 0x14, 0xd0,
	0xd2, 0xde, 0x58, 0x5a, 0xc3, 0x0d, 0x1c, 0x8f, 0x54, 0x18, 0x5d, 0x9e, 0x6e, 0x65, 0xc7, 0x01,
	0x34, 0x54, 0xd8, 0xa8, 0x21, 0x87, 0xba, 0x9d, 0x3b, 0xab, 0xe3, 0xe8, 0x94, 0x3a, 0x65, 0x4a,
	0x9d, 0xa9, 0x4c, 0x29, 0x7d, 0x09, 0x75, 0x39, 0x3c, 0x7d, 0x7c, 0xf0, 0x57, 0x2d, 0xe3, 0x70,
	0x2f, 0x6f, 0xed, 0xef, 0x3b, 0x9b, 0x5c, 0xef, 0x6c, 0xf2, 0x7b, 0x67, 0x93, 0xaf, 0x7b, 0xbb,
	0x72, 0xbd, 0xb7, 0x2b, 0x3f, 0xf6, 0x76, 0xe5, 0x63, 0xdd, 0x79, 0xb3, 0x1d, 0x5c, 0x3e, 0x50,
	0x57, 0xbf, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x82, 0x69, 0xe4, 0x80, 0x22, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AtlantServiceClient is the client API for AtlantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AtlantServiceClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type atlantServiceClient struct {
	cc *grpc.ClientConn
}

func NewAtlantServiceClient(cc *grpc.ClientConn) AtlantServiceClient {
	return &atlantServiceClient{cc}
}

func (c *atlantServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.AtlantService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *atlantServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/v1.AtlantService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AtlantServiceServer is the server API for AtlantService service.
type AtlantServiceServer interface {
	Fetch(context.Context, *FetchRequest) (*empty.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedAtlantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAtlantServiceServer struct {
}

func (*UnimplementedAtlantServiceServer) Fetch(ctx context.Context, req *FetchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (*UnimplementedAtlantServiceServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterAtlantServiceServer(s *grpc.Server, srv AtlantServiceServer) {
	s.RegisterService(&_AtlantService_serviceDesc, srv)
}

func _AtlantService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtlantServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AtlantService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtlantServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AtlantService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtlantServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AtlantService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtlantServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AtlantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AtlantService",
	HandlerType: (*AtlantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _AtlantService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AtlantService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atlant.proto",
}

func (m *FetchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintAtlant(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Options[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAtlant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Limit != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ListRequest_SortingOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRequest_SortingOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRequest_SortingOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Direction != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintAtlant(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Products) > 0 {
		for iNdEx := len(m.Products) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Products[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAtlant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListResponse_Product) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResponse_Product) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResponse_Product) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateCount != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.UpdateCount))
		i--
		dAtA[i] = 0x28
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x20
	}
	if m.CreatedAt != 0 {
		i = encodeVarintAtlant(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x18
	}
	if m.Price != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Price))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAtlant(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAtlant(dAtA []byte, offset int, v uint64) int {
	offset -= sovAtlant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FetchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovAtlant(uint64(l))
	}
	return n
}

func (m *ListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovAtlant(uint64(m.Start))
	}
	if m.Limit != 0 {
		n += 1 + sovAtlant(uint64(m.Limit))
	}
	if len(m.Options) > 0 {
		for _, e := range m.Options {
			l = e.Size()
			n += 1 + l + sovAtlant(uint64(l))
		}
	}
	return n
}

func (m *ListRequest_SortingOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovAtlant(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovAtlant(uint64(m.Direction))
	}
	return n
}

func (m *ListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Products) > 0 {
		for _, e := range m.Products {
			l = e.Size()
			n += 1 + l + sovAtlant(uint64(l))
		}
	}
	return n
}

func (m *ListResponse_Product) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAtlant(uint64(l))
	}
	if m.Price != 0 {
		n += 9
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAtlant(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovAtlant(uint64(m.UpdatedAt))
	}
	if m.UpdateCount != 0 {
		n += 1 + sovAtlant(uint64(m.UpdateCount))
	}
	return n
}

func sovAtlant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAtlant(x uint64) (n int) {
	return sovAtlant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FetchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FetchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAtlant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAtlant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtlant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtlant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAtlant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, &ListRequest_SortingOption{})
			if err := m.Options[len(m.Options)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtlant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListRequest_SortingOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SortingOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortingOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAtlant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAtlant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= ListRequest_SortingOption_SortingOptionDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAtlant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Products", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAtlant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAtlant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Products = append(m.Products, &ListResponse_Product{})
			if err := m.Products[len(m.Products)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAtlant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListResponse_Product) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Product: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Product: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAtlant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAtlant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Price = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateCount", wireType)
			}
			m.UpdateCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAtlant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAtlant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAtlant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAtlant
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAtlant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAtlant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAtlant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAtlant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAtlant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAtlant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAtlant = fmt.Errorf("proto: unexpected end of group")
)
