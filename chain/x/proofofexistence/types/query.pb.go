// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/proofofexistence/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryGetProofRequest struct {
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryGetProofRequest) Reset()         { *m = QueryGetProofRequest{} }
func (m *QueryGetProofRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetProofRequest) ProtoMessage()    {}
func (*QueryGetProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac74458ad64ae405, []int{0}
}
func (m *QueryGetProofRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProofRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProofRequest.Merge(m, src)
}
func (m *QueryGetProofRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProofRequest proto.InternalMessageInfo

func (m *QueryGetProofRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type QueryGetProofResponse struct {
	Proof Proof `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof"`
}

func (m *QueryGetProofResponse) Reset()         { *m = QueryGetProofResponse{} }
func (m *QueryGetProofResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetProofResponse) ProtoMessage()    {}
func (*QueryGetProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac74458ad64ae405, []int{1}
}
func (m *QueryGetProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProofResponse.Merge(m, src)
}
func (m *QueryGetProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProofResponse proto.InternalMessageInfo

func (m *QueryGetProofResponse) GetProof() Proof {
	if m != nil {
		return m.Proof
	}
	return Proof{}
}

func init() {
	proto.RegisterType((*QueryGetProofRequest)(nil), "empowerchain.empowerchain.proofofexistence.QueryGetProofRequest")
	proto.RegisterType((*QueryGetProofResponse)(nil), "empowerchain.empowerchain.proofofexistence.QueryGetProofResponse")
}

func init() {
	proto.RegisterFile("empowerchain/proofofexistence/query.proto", fileDescriptor_ac74458ad64ae405)
}

var fileDescriptor_ac74458ad64ae405 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xb3, 0x1f, 0xed, 0x07, 0xae, 0xb7, 0xa5, 0x82, 0x14, 0x89, 0xd2, 0x93, 0xf6, 0x90,
	0xa5, 0xf5, 0x22, 0x82, 0x50, 0x7b, 0xf1, 0x24, 0x68, 0x2f, 0x82, 0xb7, 0x4d, 0x98, 0x26, 0x0b,
	0x76, 0x67, 0x9b, 0xdd, 0x6a, 0x8b, 0x78, 0xf1, 0x09, 0x04, 0x5f, 0xaa, 0x27, 0x29, 0x78, 0xe9,
	0x49, 0xa4, 0xf5, 0x41, 0x24, 0x9b, 0x1c, 0x6c, 0x15, 0xad, 0x78, 0x9b, 0x4c, 0x7e, 0xf3, 0x9b,
	0xcc, 0x3f, 0x74, 0x0f, 0x7a, 0x1a, 0x6f, 0x20, 0x8d, 0x12, 0x21, 0x15, 0xd7, 0x29, 0x62, 0x17,
	0xbb, 0x30, 0x94, 0xc6, 0x82, 0x8a, 0x80, 0xf7, 0x07, 0x90, 0x8e, 0x02, 0x9d, 0xa2, 0x45, 0x56,
	0xff, 0x88, 0x06, 0x0b, 0x0f, 0xcb, 0x73, 0xd5, 0x4a, 0x8c, 0x31, 0xba, 0x31, 0x9e, 0x55, 0xb9,
	0xa1, 0xba, 0x15, 0x23, 0xc6, 0x57, 0xc0, 0x85, 0x96, 0x5c, 0x28, 0x85, 0x56, 0x58, 0x89, 0xca,
	0x14, 0x6f, 0xeb, 0x11, 0x9a, 0x1e, 0x1a, 0x1e, 0x0a, 0x53, 0x2c, 0xe6, 0xd7, 0x8d, 0x10, 0xac,
	0x68, 0x70, 0x2d, 0x62, 0xa9, 0x1c, 0x5c, 0xb0, 0x3f, 0x7c, 0xb6, 0x6b, 0xe4, 0x68, 0xad, 0x4e,
	0x2b, 0xe7, 0x99, 0xec, 0x04, 0xec, 0x59, 0xd6, 0xee, 0x40, 0x7f, 0x00, 0xc6, 0x32, 0x46, 0x4b,
	0x89, 0x30, 0xc9, 0x26, 0xd9, 0x21, 0xbb, 0x6b, 0x1d, 0x57, 0xd7, 0xba, 0x74, 0x63, 0x89, 0x35,
	0x1a, 0x95, 0x01, 0x76, 0x4a, 0xcb, 0xce, 0xe9, 0xe8, 0xf5, 0x66, 0x23, 0x58, 0x3d, 0x8b, 0xc0,
	0x99, 0xda, 0xa5, 0xf1, 0xcb, 0xb6, 0xd7, 0xc9, 0x2d, 0xcd, 0x29, 0xa1, 0x65, 0xb7, 0x88, 0x3d,
	0x11, 0x5a, 0x76, 0x00, 0x6b, 0xfd, 0xc6, 0xf9, 0xd5, 0x45, 0xd5, 0xe3, 0x3f, 0x18, 0xf2, 0x3b,
	0x6b, 0xad, 0xfb, 0xe7, 0xb7, 0xc7, 0x7f, 0x87, 0xec, 0x80, 0x2f, 0x04, 0xbc, 0x42, 0xda, 0xfc,
	0x36, 0x4b, 0xf0, 0xae, 0x7d, 0x31, 0x9e, 0xf9, 0x64, 0x32, 0xf3, 0xc9, 0xeb, 0xcc, 0x27, 0x0f,
	0x73, 0xdf, 0x9b, 0xcc, 0x7d, 0x6f, 0x3a, 0xf7, 0xbd, 0xcb, 0xa3, 0x58, 0xda, 0x64, 0x10, 0x06,
	0x11, 0xf6, 0xbe, 0xb1, 0x0f, 0x3f, 0xfb, 0xed, 0x48, 0x83, 0x09, 0xff, 0xbb, 0xdf, 0xb9, 0xff,
	0x1e, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xb2, 0x68, 0x6c, 0xb2, 0x02, 0x00, 0x00,
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
	// Queries a Proof by hash.
	Proof(ctx context.Context, in *QueryGetProofRequest, opts ...grpc.CallOption) (*QueryGetProofResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proof(ctx context.Context, in *QueryGetProofRequest, opts ...grpc.CallOption) (*QueryGetProofResponse, error) {
	out := new(QueryGetProofResponse)
	err := c.cc.Invoke(ctx, "/empowerchain.empowerchain.proofofexistence.Query/Proof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a Proof by hash.
	Proof(context.Context, *QueryGetProofRequest) (*QueryGetProofResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proof(ctx context.Context, req *QueryGetProofRequest) (*QueryGetProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proof not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empowerchain.empowerchain.proofofexistence.Query/Proof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proof(ctx, req.(*QueryGetProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "empowerchain.empowerchain.proofofexistence.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proof",
			Handler:    _Query_Proof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "empowerchain/proofofexistence/query.proto",
}

func (m *QueryGetProofRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProofRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProofRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryGetProofRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Proof.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetProofRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProofRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProofRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetProofResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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