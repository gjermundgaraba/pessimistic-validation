// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/lightclient/v1/clientmsg.proto

package lightclient

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	types "github.com/gjermundgaraba/pessimistic-validation/core/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// AttestationClaim is the clientMsg that is sent to the light client to update
// the consensus state All the claims need to be for the same height and have
// the same packet_commitments
type AttestationClaim struct {
	Attestations []types.Attestation `protobuf:"bytes,1,rep,name=attestations,proto3" json:"attestations"`
}

func (m *AttestationClaim) Reset()         { *m = AttestationClaim{} }
func (m *AttestationClaim) String() string { return proto.CompactTextString(m) }
func (*AttestationClaim) ProtoMessage()    {}
func (*AttestationClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_8256c76801ad19ea, []int{0}
}
func (m *AttestationClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttestationClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttestationClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestationClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationClaim.Merge(m, src)
}
func (m *AttestationClaim) XXX_Size() int {
	return m.Size()
}
func (m *AttestationClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationClaim.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationClaim proto.InternalMessageInfo

func (m *AttestationClaim) GetAttestations() []types.Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestationClaim)(nil), "core.lightclient.v1.AttestationClaim")
}

func init() {
	proto.RegisterFile("core/lightclient/v1/clientmsg.proto", fileDescriptor_8256c76801ad19ea)
}

var fileDescriptor_8256c76801ad19ea = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x13, 0x81, 0x18, 0x02, 0x03, 0x2a, 0x0c, 0x28, 0x83, 0x8b, 0x60, 0x61, 0xc1, 0xa7,
	0xc0, 0x13, 0xb4, 0xf0, 0x04, 0x4c, 0x08, 0x26, 0xdb, 0x35, 0xee, 0x21, 0x3b, 0x8e, 0xe2, 0x4b,
	0x24, 0xde, 0x82, 0xc7, 0xea, 0xd8, 0x91, 0x09, 0xa1, 0xe4, 0x45, 0x50, 0x1c, 0x55, 0x84, 0x6e,
	0x67, 0xdf, 0xf7, 0x7f, 0xfa, 0x75, 0xd9, 0xb5, 0xf2, 0xb5, 0x06, 0x8b, 0x66, 0x4d, 0xca, 0xa2,
	0x2e, 0x09, 0xda, 0x02, 0xc6, 0xc9, 0x05, 0xc3, 0xab, 0xda, 0x93, 0x9f, 0x9d, 0x0d, 0x10, 0x9f,
	0x40, 0xbc, 0x2d, 0xf2, 0x73, 0xe3, 0x8d, 0x8f, 0x7b, 0x18, 0xa6, 0x11, 0xcd, 0xe7, 0x28, 0x15,
	0x44, 0xe7, 0xbe, 0x6e, 0x07, 0x18, 0xef, 0x8d, 0xd5, 0x10, 0x5f, 0xb2, 0x79, 0x03, 0x42, 0xa7,
	0x03, 0x09, 0x57, 0xed, 0x80, 0x98, 0xa6, 0x8f, 0x4a, 0x87, 0x21, 0x2c, 0x88, 0x86, 0x35, 0xa1,
	0x2f, 0x47, 0xe0, 0xea, 0x39, 0x3b, 0x5d, 0xfc, 0x7d, 0x3e, 0x58, 0x81, 0x6e, 0xf6, 0x98, 0x9d,
	0x4c, 0xc0, 0x70, 0x91, 0x5e, 0x1e, 0xdc, 0x1c, 0xdf, 0xe5, 0x3c, 0x16, 0x8f, 0x2e, 0xde, 0x16,
	0x7c, 0x12, 0x5b, 0x1e, 0x6e, 0xbe, 0xe7, 0xc9, 0xd3, 0xbf, 0xd4, 0xf2, 0x75, 0xd3, 0xb1, 0x74,
	0xdb, 0xb1, 0xf4, 0xa7, 0x63, 0xe9, 0x67, 0xcf, 0x92, 0x6d, 0xcf, 0x92, 0xaf, 0x9e, 0x25, 0x2f,
	0x0b, 0x83, 0xb4, 0x6e, 0x24, 0x57, 0xde, 0x81, 0x79, 0xd7, 0xb5, 0x6b, 0xca, 0x95, 0x11, 0xb5,
	0x90, 0x02, 0x2a, 0x1d, 0x02, 0x3a, 0x0c, 0x84, 0xea, 0xb6, 0x15, 0x16, 0x57, 0x51, 0x06, 0xfb,
	0x77, 0x95, 0x47, 0xb1, 0xfd, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xb1, 0x16, 0x78,
	0x72, 0x01, 0x00, 0x00,
}

func (m *AttestationClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestationClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttestationClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attestations) > 0 {
		for iNdEx := len(m.Attestations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attestations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClientmsg(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintClientmsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovClientmsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AttestationClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovClientmsg(uint64(l))
		}
	}
	return n
}

func sovClientmsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClientmsg(x uint64) (n int) {
	return sovClientmsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AttestationClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientmsg
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
			return fmt.Errorf("proto: AttestationClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttestationClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientmsg
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
				return ErrInvalidLengthClientmsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientmsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attestations = append(m.Attestations, types.Attestation{})
			if err := m.Attestations[len(m.Attestations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientmsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClientmsg
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
func skipClientmsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientmsg
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
					return 0, ErrIntOverflowClientmsg
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
					return 0, ErrIntOverflowClientmsg
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
				return 0, ErrInvalidLengthClientmsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClientmsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClientmsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClientmsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientmsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClientmsg = fmt.Errorf("proto: unexpected end of group")
)
