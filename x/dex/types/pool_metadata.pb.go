// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/pool_metadata.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type PoolMetadata struct {
	Id     uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tick   int64   `protobuf:"varint,2,opt,name=tick,proto3" json:"tick,omitempty"`
	Fee    uint64  `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	PairId *PairID `protobuf:"bytes,4,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
}

func (m *PoolMetadata) Reset()         { *m = PoolMetadata{} }
func (m *PoolMetadata) String() string { return proto.CompactTextString(m) }
func (*PoolMetadata) ProtoMessage()    {}
func (*PoolMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ee8eaeac9c06d8, []int{0}
}
func (m *PoolMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolMetadata.Merge(m, src)
}
func (m *PoolMetadata) XXX_Size() int {
	return m.Size()
}
func (m *PoolMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PoolMetadata proto.InternalMessageInfo

func (m *PoolMetadata) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PoolMetadata) GetTick() int64 {
	if m != nil {
		return m.Tick
	}
	return 0
}

func (m *PoolMetadata) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *PoolMetadata) GetPairId() *PairID {
	if m != nil {
		return m.PairId
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolMetadata)(nil), "neutron.dex.PoolMetadata")
}

func init() { proto.RegisterFile("neutron/dex/pool_metadata.proto", fileDescriptor_c2ee8eaeac9c06d8) }

var fileDescriptor_c2ee8eaeac9c06d8 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4b, 0x2d, 0x2d,
	0x29, 0xca, 0xcf, 0xd3, 0x4f, 0x49, 0xad, 0xd0, 0x2f, 0xc8, 0xcf, 0xcf, 0x89, 0xcf, 0x4d, 0x2d,
	0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x2a, 0xd0,
	0x4b, 0x49, 0xad, 0x90, 0x92, 0x44, 0x51, 0x9d, 0x98, 0x59, 0x14, 0x9f, 0x99, 0x02, 0x51, 0xa7,
	0x54, 0xc4, 0xc5, 0x13, 0x90, 0x9f, 0x9f, 0xe3, 0x0b, 0xd5, 0x2d, 0xc4, 0xc7, 0xc5, 0x94, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x52, 0x92,
	0x99, 0x9c, 0x2d, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0x66, 0x0b, 0x09, 0x70, 0x31, 0xa7,
	0xa5, 0xa6, 0x4a, 0x30, 0x83, 0x15, 0x81, 0x98, 0x42, 0x3a, 0x5c, 0xec, 0x50, 0x63, 0x25, 0x58,
	0x14, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0x90, 0xec, 0xd7, 0x0b, 0x48, 0xcc, 0x2c, 0xf2, 0x74,
	0x09, 0x62, 0x03, 0xa9, 0xf1, 0x4c, 0x71, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xa8,
	0x01, 0xba, 0xf9, 0x45, 0xe9, 0x30, 0xb6, 0x7e, 0x99, 0x91, 0x7e, 0x05, 0xd8, 0x13, 0x25, 0x95,
	0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x3f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0x50,
	0x66, 0x98, 0x0e, 0x01, 0x00, 0x00,
}

func (m *PoolMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairId != nil {
		{
			size, err := m.PairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Fee != 0 {
		i = encodeVarintPoolMetadata(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if m.Tick != 0 {
		i = encodeVarintPoolMetadata(dAtA, i, uint64(m.Tick))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintPoolMetadata(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPoolMetadata(uint64(m.Id))
	}
	if m.Tick != 0 {
		n += 1 + sovPoolMetadata(uint64(m.Tick))
	}
	if m.Fee != 0 {
		n += 1 + sovPoolMetadata(uint64(m.Fee))
	}
	if m.PairId != nil {
		l = m.PairId.Size()
		n += 1 + l + sovPoolMetadata(uint64(l))
	}
	return n
}

func sovPoolMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolMetadata(x uint64) (n int) {
	return sovPoolMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolMetadata
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
			return fmt.Errorf("proto: PoolMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tick", wireType)
			}
			m.Tick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolMetadata
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
				return ErrInvalidLengthPoolMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairId == nil {
				m.PairId = &PairID{}
			}
			if err := m.PairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolMetadata
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
func skipPoolMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolMetadata
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
					return 0, ErrIntOverflowPoolMetadata
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
					return 0, ErrIntOverflowPoolMetadata
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
				return 0, ErrInvalidLengthPoolMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolMetadata = fmt.Errorf("proto: unexpected end of group")
)