// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/storage/fileset/fileset.proto

package fileset

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	index "github.com/pachyderm/pachyderm/src/server/pkg/storage/fileset/index"
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

type Metadata struct {
	Path                 string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Additive             *index.Index `protobuf:"bytes,2,opt,name=additive,proto3" json:"additive,omitempty"`
	Deletive             *index.Index `protobuf:"bytes,3,opt,name=deletive,proto3" json:"deletive,omitempty"`
	SizeBytes            int64        `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcfbe9461ec0392b, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Metadata) GetAdditive() *index.Index {
	if m != nil {
		return m.Additive
	}
	return nil
}

func (m *Metadata) GetDeletive() *index.Index {
	if m != nil {
		return m.Deletive
	}
	return nil
}

func (m *Metadata) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Metadata)(nil), "fileset.Metadata")
}

func init() {
	proto.RegisterFile("server/pkg/storage/fileset/fileset.proto", fileDescriptor_dcfbe9461ec0392b)
}

var fileDescriptor_dcfbe9461ec0392b = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0xc8, 0x4e, 0xd7, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x4f,
	0xcb, 0xcc, 0x49, 0x2d, 0x4e, 0x2d, 0x81, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec,
	0x50, 0xae, 0x94, 0x0e, 0x1e, 0x2d, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x10, 0x12, 0xa2, 0x4d, 0x69,
	0x22, 0x23, 0x17, 0x87, 0x6f, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x49, 0xa2, 0x90, 0x10, 0x17, 0x4b,
	0x41, 0x62, 0x49, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xa4, 0xc1, 0xc5,
	0x91, 0x98, 0x92, 0x92, 0x59, 0x92, 0x59, 0x96, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4,
	0xa3, 0x07, 0x31, 0xc0, 0x13, 0x44, 0x06, 0xc1, 0x65, 0x41, 0x2a, 0x53, 0x52, 0x73, 0x52, 0xc1,
	0x2a, 0x99, 0xb1, 0xa9, 0x84, 0xc9, 0x0a, 0xc9, 0x72, 0x71, 0x15, 0x67, 0x56, 0xa5, 0xc6, 0x27,
	0x55, 0x96, 0xa4, 0x16, 0x4b, 0xb0, 0x28, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x82, 0x44, 0x9c, 0x40,
	0x02, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63,
	0x94, 0x6d, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x41, 0x62, 0x72,
	0x46, 0x65, 0x4a, 0x6a, 0x11, 0x32, 0xab, 0xb8, 0x28, 0x59, 0x1f, 0xb7, 0x97, 0x93, 0xd8, 0xc0,
	0xfe, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0xba, 0x8a, 0x36, 0x4a, 0x01, 0x00, 0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SizeBytes != 0 {
		i = encodeVarintFileset(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x20
	}
	if m.Deletive != nil {
		{
			size, err := m.Deletive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Additive != nil {
		{
			size, err := m.Additive.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFileset(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintFileset(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFileset(dAtA []byte, offset int, v uint64) int {
	offset -= sovFileset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.Additive != nil {
		l = m.Additive.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.Deletive != nil {
		l = m.Deletive.Size()
		n += 1 + l + sovFileset(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovFileset(uint64(m.SizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFileset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFileset(x uint64) (n int) {
	return sovFileset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFileset
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Additive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Additive == nil {
				m.Additive = &index.Index{}
			}
			if err := m.Additive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deletive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
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
				return ErrInvalidLengthFileset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFileset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Deletive == nil {
				m.Deletive = &index.Index{}
			}
			if err := m.Deletive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFileset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFileset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFileset
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFileset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFileset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFileset
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
					return 0, ErrIntOverflowFileset
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
					return 0, ErrIntOverflowFileset
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
				return 0, ErrInvalidLengthFileset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFileset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFileset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFileset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFileset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFileset = fmt.Errorf("proto: unexpected end of group")
)
