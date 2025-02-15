// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/pep_nonce.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type PepNonce struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Nonce   uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *PepNonce) Reset()         { *m = PepNonce{} }
func (m *PepNonce) String() string { return proto.CompactTextString(m) }
func (*PepNonce) ProtoMessage()    {}
func (*PepNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_5760221b7253d15f, []int{0}
}
func (m *PepNonce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PepNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PepNonce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PepNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PepNonce.Merge(m, src)
}
func (m *PepNonce) XXX_Size() int {
	return m.Size()
}
func (m *PepNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_PepNonce.DiscardUnknown(m)
}

var xxx_messageInfo_PepNonce proto.InternalMessageInfo

func (m *PepNonce) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PepNonce) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*PepNonce)(nil), "fairyring.pep.PepNonce")
}

func init() { proto.RegisterFile("fairyring/pep/pep_nonce.proto", fileDescriptor_5760221b7253d15f) }

var fileDescriptor_5760221b7253d15f = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x48, 0x2d, 0x00, 0xe1, 0xf8, 0xbc, 0xfc, 0xbc, 0xe4,
	0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x5e, 0xb8, 0xb4, 0x5e, 0x41, 0x6a, 0x81, 0x92,
	0x15, 0x17, 0x47, 0x40, 0x6a, 0x81, 0x1f, 0x48, 0x81, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a,
	0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc2, 0xc5,
	0x0a, 0x36, 0x43, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc2, 0x71, 0xd2, 0x3f, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x51, 0x84, 0x1b, 0x2a, 0xc0, 0xae, 0x28, 0xa9,
	0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x90, 0xaa,
	0x45, 0x10, 0xa3, 0x00, 0x00, 0x00,
}

func (m *PepNonce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PepNonce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PepNonce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintPepNonce(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPepNonce(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPepNonce(dAtA []byte, offset int, v uint64) int {
	offset -= sovPepNonce(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PepNonce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPepNonce(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovPepNonce(uint64(m.Nonce))
	}
	return n
}

func sovPepNonce(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPepNonce(x uint64) (n int) {
	return sovPepNonce(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PepNonce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPepNonce
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
			return fmt.Errorf("proto: PepNonce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PepNonce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPepNonce
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
				return ErrInvalidLengthPepNonce
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPepNonce
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPepNonce
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPepNonce(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPepNonce
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
func skipPepNonce(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPepNonce
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
					return 0, ErrIntOverflowPepNonce
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
					return 0, ErrIntOverflowPepNonce
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
				return 0, ErrInvalidLengthPepNonce
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPepNonce
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPepNonce
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPepNonce        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPepNonce          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPepNonce = fmt.Errorf("proto: unexpected end of group")
)
