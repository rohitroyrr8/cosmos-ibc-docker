// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSendIbcTopRank struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	PlayerId         string `protobuf:"bytes,5,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Rank             uint64 `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
	Score            string `protobuf:"bytes,7,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *MsgSendIbcTopRank) Reset()         { *m = MsgSendIbcTopRank{} }
func (m *MsgSendIbcTopRank) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcTopRank) ProtoMessage()    {}
func (*MsgSendIbcTopRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{0}
}
func (m *MsgSendIbcTopRank) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcTopRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcTopRank.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcTopRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcTopRank.Merge(m, src)
}
func (m *MsgSendIbcTopRank) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcTopRank) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcTopRank.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcTopRank proto.InternalMessageInfo

func (m *MsgSendIbcTopRank) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendIbcTopRank) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetRank() uint64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *MsgSendIbcTopRank) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

type MsgSendIbcTopRankResponse struct {
}

func (m *MsgSendIbcTopRankResponse) Reset()         { *m = MsgSendIbcTopRankResponse{} }
func (m *MsgSendIbcTopRankResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcTopRankResponse) ProtoMessage()    {}
func (*MsgSendIbcTopRankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{1}
}
func (m *MsgSendIbcTopRankResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcTopRankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcTopRankResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcTopRankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcTopRankResponse.Merge(m, src)
}
func (m *MsgSendIbcTopRankResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcTopRankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcTopRankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcTopRankResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendIbcTopRank)(nil), "tmsdkeys.leaderboard.leaderboard.MsgSendIbcTopRank")
	proto.RegisterType((*MsgSendIbcTopRankResponse)(nil), "tmsdkeys.leaderboard.leaderboard.MsgSendIbcTopRankResponse")
}

func init() { proto.RegisterFile("leaderboard/tx.proto", fileDescriptor_abcbe4eb090e075c) }

var fileDescriptor_abcbe4eb090e075c = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0xeb, 0xb7, 0xff, 0xde, 0x7a, 0x40, 0x60, 0x75, 0x30, 0x05, 0x59, 0x55, 0xa7, 0x8a,
	0x21, 0x95, 0xe8, 0xc8, 0x86, 0x58, 0x3a, 0x74, 0x09, 0x9d, 0xd8, 0x9c, 0xe4, 0x94, 0x46, 0x49,
	0x6c, 0xcb, 0x76, 0xa5, 0x66, 0x65, 0x64, 0xe2, 0x63, 0x31, 0x96, 0x8d, 0x11, 0x25, 0x5f, 0x04,
	0x61, 0x14, 0x08, 0x64, 0x40, 0x62, 0x7b, 0x7e, 0xcf, 0xf9, 0xee, 0x7c, 0x77, 0x78, 0x9c, 0x01,
	0x8f, 0x40, 0x07, 0x92, 0xeb, 0x68, 0x61, 0xf7, 0x9e, 0xd2, 0xd2, 0x4a, 0x32, 0xb5, 0xb9, 0x89,
	0x52, 0x28, 0x8c, 0xd7, 0x08, 0x37, 0xf5, 0xec, 0x19, 0xe1, 0x93, 0xb5, 0x89, 0x6f, 0x41, 0x44,
	0xab, 0x20, 0xdc, 0x48, 0xe5, 0x73, 0x91, 0x12, 0x8a, 0x87, 0xa1, 0x06, 0x6e, 0xa5, 0xa6, 0x68,
	0x8a, 0xe6, 0x23, 0xbf, 0x46, 0x42, 0x70, 0x4f, 0x49, 0x6d, 0xe9, 0x3f, 0x67, 0x3b, 0x4d, 0xce,
	0xf1, 0x28, 0xdc, 0x72, 0x21, 0x20, 0x5b, 0xdd, 0xd0, 0xae, 0x0b, 0x7c, 0x19, 0xe4, 0x02, 0x1f,
	0xdb, 0x24, 0x07, 0xb9, 0xb3, 0x9b, 0x24, 0x07, 0x63, 0x79, 0xae, 0x68, 0x6f, 0x8a, 0xe6, 0x3d,
	0xbf, 0xe5, 0x93, 0x09, 0xfe, 0xaf, 0x32, 0x5e, 0x80, 0x5e, 0x45, 0xb4, 0xef, 0x0a, 0x7d, 0xf2,
	0x7b, 0x67, 0xcd, 0x45, 0x4a, 0x07, 0x2e, 0xd7, 0x69, 0x32, 0xc6, 0x7d, 0x13, 0x4a, 0x0d, 0x74,
	0xe8, 0x1e, 0x7f, 0xc0, 0xec, 0x0c, 0x9f, 0xb6, 0x46, 0xf2, 0xc1, 0x28, 0x29, 0x0c, 0x5c, 0x3e,
	0x20, 0xdc, 0x5d, 0x9b, 0x98, 0xdc, 0x23, 0x7c, 0xf4, 0x63, 0xea, 0xa5, 0xf7, 0xdb, 0xba, 0xbc,
	0x56, 0xdd, 0xc9, 0xd5, 0x1f, 0x92, 0xea, 0xcf, 0x5c, 0xaf, 0x9f, 0x4a, 0x86, 0x0e, 0x25, 0x43,
	0xaf, 0x25, 0x43, 0x8f, 0x15, 0xeb, 0x1c, 0x2a, 0xd6, 0x79, 0xa9, 0x58, 0xe7, 0x6e, 0x19, 0x27,
	0x76, 0xbb, 0x0b, 0xbc, 0x50, 0xe6, 0x8b, 0xba, 0xc1, 0xa2, 0x79, 0xe3, 0xfd, 0x37, 0xb2, 0x85,
	0x02, 0x13, 0x0c, 0xdc, 0xd5, 0x97, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x15, 0x08, 0xfd,
	0x0d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendIbcTopRank(ctx context.Context, in *MsgSendIbcTopRank, opts ...grpc.CallOption) (*MsgSendIbcTopRankResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendIbcTopRank(ctx context.Context, in *MsgSendIbcTopRank, opts ...grpc.CallOption) (*MsgSendIbcTopRankResponse, error) {
	out := new(MsgSendIbcTopRankResponse)
	err := c.cc.Invoke(ctx, "/tmsdkeys.leaderboard.leaderboard.Msg/SendIbcTopRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendIbcTopRank(context.Context, *MsgSendIbcTopRank) (*MsgSendIbcTopRankResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendIbcTopRank(ctx context.Context, req *MsgSendIbcTopRank) (*MsgSendIbcTopRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIbcTopRank not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendIbcTopRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendIbcTopRank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendIbcTopRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tmsdkeys.leaderboard.leaderboard.Msg/SendIbcTopRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendIbcTopRank(ctx, req.(*MsgSendIbcTopRank))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tmsdkeys.leaderboard.leaderboard.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIbcTopRank",
			Handler:    _Msg_SendIbcTopRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaderboard/tx.proto",
}

func (m *MsgSendIbcTopRank) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcTopRank) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcTopRank) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Rank != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PlayerId) > 0 {
		i -= len(m.PlayerId)
		copy(dAtA[i:], m.PlayerId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlayerId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendIbcTopRankResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcTopRankResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcTopRankResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendIbcTopRank) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.PlayerId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovTx(uint64(m.Rank))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendIbcTopRankResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendIbcTopRank) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendIbcTopRank: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcTopRank: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSendIbcTopRankResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendIbcTopRankResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcTopRankResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
