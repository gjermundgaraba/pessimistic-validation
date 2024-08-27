// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/voteextension/v1/voteextension.proto

package voteextension

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	lightclient "github.com/cosmos/interchain-attestation/core/lightclient"
	types "github.com/cosmos/interchain-attestation/core/types"
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

type VoteExtension struct {
	// one attestation for every chain configured in the sidecar
	Attestations []types.Attestation `protobuf:"bytes,1,rep,name=attestations,proto3" json:"attestations"`
}

func (m *VoteExtension) Reset()         { *m = VoteExtension{} }
func (m *VoteExtension) String() string { return proto.CompactTextString(m) }
func (*VoteExtension) ProtoMessage()    {}
func (*VoteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_b31da80dff3b1d62, []int{0}
}
func (m *VoteExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteExtension.Merge(m, src)
}
func (m *VoteExtension) XXX_Size() int {
	return m.Size()
}
func (m *VoteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_VoteExtension proto.InternalMessageInfo

func (m *VoteExtension) GetAttestations() []types.Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

type ClientUpdate struct {
	ClientToUpdate   string                       `protobuf:"bytes,1,opt,name=client_to_update,json=clientToUpdate,proto3" json:"client_to_update,omitempty"`
	AttestationClaim lightclient.AttestationClaim `protobuf:"bytes,2,opt,name=attestation_claim,json=attestationClaim,proto3" json:"attestation_claim"`
}

func (m *ClientUpdate) Reset()         { *m = ClientUpdate{} }
func (m *ClientUpdate) String() string { return proto.CompactTextString(m) }
func (*ClientUpdate) ProtoMessage()    {}
func (*ClientUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b31da80dff3b1d62, []int{1}
}
func (m *ClientUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientUpdate.Merge(m, src)
}
func (m *ClientUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ClientUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ClientUpdate proto.InternalMessageInfo

func (m *ClientUpdate) GetClientToUpdate() string {
	if m != nil {
		return m.ClientToUpdate
	}
	return ""
}

func (m *ClientUpdate) GetAttestationClaim() lightclient.AttestationClaim {
	if m != nil {
		return m.AttestationClaim
	}
	return lightclient.AttestationClaim{}
}

type ClientUpdates struct {
	ClientUpdates []ClientUpdate `protobuf:"bytes,1,rep,name=client_updates,json=clientUpdates,proto3" json:"client_updates"`
}

func (m *ClientUpdates) Reset()         { *m = ClientUpdates{} }
func (m *ClientUpdates) String() string { return proto.CompactTextString(m) }
func (*ClientUpdates) ProtoMessage()    {}
func (*ClientUpdates) Descriptor() ([]byte, []int) {
	return fileDescriptor_b31da80dff3b1d62, []int{2}
}
func (m *ClientUpdates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientUpdates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientUpdates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientUpdates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientUpdates.Merge(m, src)
}
func (m *ClientUpdates) XXX_Size() int {
	return m.Size()
}
func (m *ClientUpdates) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientUpdates.DiscardUnknown(m)
}

var xxx_messageInfo_ClientUpdates proto.InternalMessageInfo

func (m *ClientUpdates) GetClientUpdates() []ClientUpdate {
	if m != nil {
		return m.ClientUpdates
	}
	return nil
}

func init() {
	proto.RegisterType((*VoteExtension)(nil), "core.voteextension.v1.VoteExtension")
	proto.RegisterType((*ClientUpdate)(nil), "core.voteextension.v1.ClientUpdate")
	proto.RegisterType((*ClientUpdates)(nil), "core.voteextension.v1.ClientUpdates")
}

func init() {
	proto.RegisterFile("core/voteextension/v1/voteextension.proto", fileDescriptor_b31da80dff3b1d62)
}

var fileDescriptor_b31da80dff3b1d62 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0x3b, 0xff, 0x6f, 0x4c, 0x1c, 0xc0, 0x60, 0xa3, 0x09, 0xe9, 0xa2, 0x10, 0x88, 0x49,
	0x5d, 0x38, 0x0d, 0xb8, 0x74, 0x25, 0xe8, 0xde, 0x10, 0x31, 0xc6, 0x0d, 0x19, 0xc6, 0x49, 0x99,
	0x84, 0xf6, 0x12, 0xe6, 0xd2, 0xe8, 0x5b, 0x18, 0x9f, 0x8a, 0x25, 0x4b, 0x57, 0xc6, 0xc0, 0x8b,
	0x98, 0x4e, 0x8b, 0x0e, 0xe8, 0x6e, 0x72, 0xfa, 0xdd, 0x73, 0xce, 0xcd, 0x2d, 0x3d, 0x13, 0x30,
	0x93, 0x61, 0x0a, 0x28, 0xe5, 0x33, 0xca, 0x44, 0x2b, 0x48, 0xc2, 0xb4, 0xbd, 0x2d, 0xb0, 0xe9,
	0x0c, 0x10, 0xdc, 0x93, 0x0c, 0x65, 0xdb, 0x5f, 0xd2, 0xb6, 0x77, 0x1c, 0x41, 0x04, 0x86, 0x08,
	0xb3, 0x57, 0x0e, 0x7b, 0x75, 0xe3, 0x8b, 0x2f, 0x53, 0xa9, 0x33, 0x3f, 0x8e, 0x28, 0x35, 0x72,
	0xfc, 0x76, 0xf3, 0x5a, 0x06, 0x98, 0xa8, 0x68, 0x8c, 0x62, 0xa2, 0x64, 0x82, 0x19, 0x96, 0xbf,
	0x62, 0x1d, 0xe5, 0x50, 0x73, 0x40, 0x2b, 0xf7, 0x80, 0xf2, 0x66, 0x93, 0xe7, 0x5e, 0xd3, 0xb2,
	0x65, 0xa5, 0x6b, 0xa4, 0xf1, 0x3f, 0x28, 0x75, 0x3c, 0x66, 0xaa, 0x99, 0x34, 0x96, 0xb6, 0xd9,
	0xd5, 0x0f, 0xd2, 0xdd, 0x5b, 0x7c, 0xd4, 0x9d, 0xfe, 0xd6, 0x54, 0xf3, 0x8d, 0xd0, 0x72, 0xcf,
	0x44, 0x0d, 0xa6, 0x4f, 0x1c, 0xa5, 0x1b, 0xd0, 0x6a, 0x1e, 0x3d, 0x44, 0x18, 0xce, 0x8d, 0x56,
	0x23, 0x0d, 0x12, 0x1c, 0xf4, 0x0f, 0x73, 0xfd, 0x0e, 0x0a, 0xf2, 0x81, 0x1e, 0x59, 0x56, 0x43,
	0x31, 0xe1, 0x2a, 0xae, 0xfd, 0x6b, 0x90, 0xa0, 0xd4, 0x39, 0xcd, 0x5b, 0x58, 0x2b, 0xed, 0x74,
	0xe9, 0x65, 0x70, 0x51, 0xa8, 0xca, 0x77, 0xf4, 0x26, 0xa7, 0x15, 0xbb, 0x93, 0x76, 0x6f, 0x69,
	0x11, 0x5e, 0x34, 0xda, 0x6c, 0xdb, 0x62, 0x7f, 0x1e, 0x82, 0xd9, 0xd3, 0x45, 0x4a, 0x45, 0xd8,
	0x8e, 0xdd, 0xc1, 0x62, 0xe5, 0x93, 0xe5, 0xca, 0x27, 0x9f, 0x2b, 0x9f, 0xbc, 0xae, 0x7d, 0x67,
	0xb9, 0xf6, 0x9d, 0xf7, 0xb5, 0xef, 0x3c, 0x5e, 0x46, 0x0a, 0xc7, 0xf3, 0x11, 0x13, 0x10, 0x87,
	0x02, 0x74, 0x0c, 0x3a, 0x54, 0x09, 0xca, 0x99, 0x18, 0x73, 0x95, 0x9c, 0x5b, 0x5d, 0xc3, 0xdf,
	0xff, 0xcb, 0x68, 0xdf, 0x1c, 0xeb, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x38, 0x86, 0xe0, 0x50,
	0x4c, 0x02, 0x00, 0x00,
}

func (m *VoteExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintVoteextension(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClientUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AttestationClaim.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVoteextension(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ClientToUpdate) > 0 {
		i -= len(m.ClientToUpdate)
		copy(dAtA[i:], m.ClientToUpdate)
		i = encodeVarintVoteextension(dAtA, i, uint64(len(m.ClientToUpdate)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientUpdates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientUpdates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientUpdates) > 0 {
		for iNdEx := len(m.ClientUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVoteextension(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintVoteextension(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoteextension(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovVoteextension(uint64(l))
		}
	}
	return n
}

func (m *ClientUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientToUpdate)
	if l > 0 {
		n += 1 + l + sovVoteextension(uint64(l))
	}
	l = m.AttestationClaim.Size()
	n += 1 + l + sovVoteextension(uint64(l))
	return n
}

func (m *ClientUpdates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ClientUpdates) > 0 {
		for _, e := range m.ClientUpdates {
			l = e.Size()
			n += 1 + l + sovVoteextension(uint64(l))
		}
	}
	return n
}

func sovVoteextension(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoteextension(x uint64) (n int) {
	return sovVoteextension(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteextension
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
			return fmt.Errorf("proto: VoteExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteextension
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
				return ErrInvalidLengthVoteextension
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteextension
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
			skippy, err := skipVoteextension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteextension
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
func (m *ClientUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteextension
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
			return fmt.Errorf("proto: ClientUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientToUpdate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteextension
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
				return ErrInvalidLengthVoteextension
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoteextension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientToUpdate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestationClaim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteextension
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
				return ErrInvalidLengthVoteextension
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteextension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AttestationClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteextension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteextension
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
func (m *ClientUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoteextension
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
			return fmt.Errorf("proto: ClientUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoteextension
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
				return ErrInvalidLengthVoteextension
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVoteextension
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientUpdates = append(m.ClientUpdates, ClientUpdate{})
			if err := m.ClientUpdates[len(m.ClientUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoteextension(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoteextension
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
func skipVoteextension(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoteextension
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
					return 0, ErrIntOverflowVoteextension
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
					return 0, ErrIntOverflowVoteextension
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
				return 0, ErrInvalidLengthVoteextension
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoteextension
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoteextension
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoteextension        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoteextension          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoteextension = fmt.Errorf("proto: unexpected end of group")
)
