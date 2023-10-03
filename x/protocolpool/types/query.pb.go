// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryCommunityPoolRequest is the request type for the Query/CommunityPool RPC
// method.
type QueryCommunityPoolRequest struct {
}

func (m *QueryCommunityPoolRequest) Reset()         { *m = QueryCommunityPoolRequest{} }
func (m *QueryCommunityPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCommunityPoolRequest) ProtoMessage()    {}
func (*QueryCommunityPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51500a0a77d57843, []int{0}
}
func (m *QueryCommunityPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommunityPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommunityPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommunityPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommunityPoolRequest.Merge(m, src)
}
func (m *QueryCommunityPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommunityPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommunityPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommunityPoolRequest proto.InternalMessageInfo

// QueryCommunityPoolResponse is the response type for the Query/CommunityPool
// RPC method.
type QueryCommunityPoolResponse struct {
	// pool defines community pool's coins.
	Pool github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=pool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"pool"`
}

func (m *QueryCommunityPoolResponse) Reset()         { *m = QueryCommunityPoolResponse{} }
func (m *QueryCommunityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCommunityPoolResponse) ProtoMessage()    {}
func (*QueryCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51500a0a77d57843, []int{1}
}
func (m *QueryCommunityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCommunityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCommunityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCommunityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCommunityPoolResponse.Merge(m, src)
}
func (m *QueryCommunityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCommunityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCommunityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCommunityPoolResponse proto.InternalMessageInfo

func (m *QueryCommunityPoolResponse) GetPool() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Pool
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCommunityPoolRequest)(nil), "cosmos.protocolpool.v1.QueryCommunityPoolRequest")
	proto.RegisterType((*QueryCommunityPoolResponse)(nil), "cosmos.protocolpool.v1.QueryCommunityPoolResponse")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/query.proto", fileDescriptor_51500a0a77d57843)
}

var fileDescriptor_51500a0a77d57843 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3f, 0x4e, 0xf3, 0x30,
	0x14, 0x8f, 0xbf, 0x0f, 0x18, 0x82, 0x58, 0x22, 0x84, 0x20, 0x54, 0x2e, 0x64, 0x40, 0x95, 0x10,
	0x7e, 0x4a, 0x3b, 0xb2, 0xb5, 0x1c, 0x00, 0x3a, 0xb2, 0xa0, 0x24, 0xb5, 0x82, 0xd5, 0xc4, 0x2f,
	0xad, 0x9d, 0x8a, 0xae, 0x70, 0x01, 0x24, 0xee, 0xc0, 0xc0, 0xc4, 0x31, 0x3a, 0x56, 0x62, 0x61,
	0x02, 0xd4, 0x72, 0x10, 0x94, 0x38, 0x95, 0xa8, 0xd4, 0x0e, 0x4c, 0xb6, 0x7e, 0xef, 0xfd, 0xfe,
	0xd9, 0xb6, 0x17, 0xa1, 0x4a, 0x51, 0x41, 0x36, 0x44, 0x8d, 0x11, 0x26, 0x19, 0x62, 0x02, 0x23,
	0x1f, 0x06, 0x39, 0x1f, 0x8e, 0x59, 0x89, 0x3a, 0x7b, 0x66, 0x87, 0xfd, 0xde, 0x61, 0x23, 0xdf,
	0xdd, 0x8d, 0x31, 0xc6, 0x12, 0x84, 0xe2, 0x66, 0xe6, 0x6e, 0x2d, 0x46, 0x8c, 0x13, 0x0e, 0x41,
	0x26, 0x20, 0x90, 0x12, 0x75, 0xa0, 0x05, 0xca, 0x8a, 0xed, 0xd2, 0xca, 0x2f, 0x0c, 0x14, 0x87,
	0x91, 0x1f, 0x72, 0x1d, 0xf8, 0x10, 0xa1, 0x90, 0x66, 0xee, 0x1d, 0xda, 0x07, 0x57, 0x85, 0x75,
	0x07, 0xd3, 0x34, 0x97, 0x42, 0x8f, 0x2f, 0x11, 0x93, 0x2e, 0x1f, 0xe4, 0x5c, 0x69, 0xef, 0x81,
	0xd8, 0xee, 0xaa, 0xa9, 0xca, 0x50, 0x2a, 0xee, 0x70, 0x7b, 0xa3, 0x88, 0xb6, 0x4f, 0x8e, 0xfe,
	0x37, 0xb6, 0x9b, 0x35, 0x56, 0xc5, 0x2e, 0xac, 0x58, 0x65, 0xc5, 0x2e, 0x78, 0xd4, 0x41, 0x21,
	0xdb, 0xad, 0xc9, 0x47, 0xdd, 0x7a, 0xf9, 0xac, 0x9f, 0xc6, 0x42, 0xdf, 0xe6, 0x21, 0x8b, 0x30,
	0x85, 0x2a, 0x9a, 0x39, 0xce, 0x54, 0xaf, 0x0f, 0x7a, 0x9c, 0x71, 0xb5, 0xe0, 0xa8, 0x6e, 0x29,
	0xdf, 0x7c, 0x25, 0xf6, 0x66, 0x99, 0xc2, 0x79, 0x26, 0xf6, 0xce, 0x52, 0x14, 0xc7, 0x67, 0xab,
	0xdf, 0x8a, 0xad, 0x2d, 0xe5, 0x36, 0xff, 0x42, 0x31, 0x4d, 0x3d, 0x76, 0xff, 0xf6, 0xfd, 0xf4,
	0xaf, 0xe1, 0x9c, 0xc0, 0x9a, 0xef, 0x8b, 0x16, 0xb4, 0x9b, 0x02, 0x69, 0x9f, 0x4f, 0x66, 0x94,
	0x4c, 0x67, 0x94, 0x7c, 0xcd, 0x28, 0x79, 0x9c, 0x53, 0x6b, 0x3a, 0xa7, 0xd6, 0xfb, 0x9c, 0x5a,
	0xd7, 0xc7, 0x46, 0x40, 0xf5, 0xfa, 0x4c, 0x20, 0xdc, 0x2d, 0x0b, 0x95, 0xf5, 0xc3, 0xad, 0x12,
	0x6b, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xc7, 0xae, 0xa2, 0x2b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CommunityPool queries the community pool coins.
	CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CommunityPool(ctx context.Context, in *QueryCommunityPoolRequest, opts ...grpc.CallOption) (*QueryCommunityPoolResponse, error) {
	out := new(QueryCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.protocolpool.v1.Query/CommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CommunityPool queries the community pool coins.
	CommunityPool(context.Context, *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CommunityPool(ctx context.Context, req *QueryCommunityPoolRequest) (*QueryCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommunityPool not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.protocolpool.v1.Query/CommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.protocolpool.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommunityPool",
			Handler:    _Query_CommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/protocolpool/v1/query.proto",
}

func (m *QueryCommunityPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommunityPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommunityPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCommunityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCommunityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCommunityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for iNdEx := len(m.Pool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCommunityPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCommunityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pool) > 0 {
		for _, e := range m.Pool {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCommunityPoolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommunityPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommunityPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCommunityPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCommunityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pool = append(m.Pool, types.DecCoin{})
			if err := m.Pool[len(m.Pool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)