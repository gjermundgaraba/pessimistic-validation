// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/types/v1/attestation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Attestation struct {
	AttestatorId []byte  `protobuf:"bytes,1,opt,name=attestator_id,json=attestatorId,proto3" json:"attestator_id,omitempty"`
	AttestedData IBCData `protobuf:"bytes,2,opt,name=attested_data,json=attestedData,proto3" json:"attested_data"`
	Signature    []byte  `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_25eb7c0454d2e150, []int{0}
}
func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return m.Size()
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetAttestatorId() []byte {
	if m != nil {
		return m.AttestatorId
	}
	return nil
}

func (m *Attestation) GetAttestedData() IBCData {
	if m != nil {
		return m.AttestedData
	}
	return IBCData{}
}

func (m *Attestation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type IBCData struct {
	ChainId           string       `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ClientId          string       `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientToUpdate    string       `protobuf:"bytes,3,opt,name=client_to_update,json=clientToUpdate,proto3" json:"client_to_update,omitempty"`
	Height            types.Height `protobuf:"bytes,4,opt,name=height,proto3" json:"height"`
	Timestamp         time.Time    `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	PacketCommitments [][]byte     `protobuf:"bytes,6,rep,name=packet_commitments,json=packetCommitments,proto3" json:"packet_commitments,omitempty"`
}

func (m *IBCData) Reset()         { *m = IBCData{} }
func (m *IBCData) String() string { return proto.CompactTextString(m) }
func (*IBCData) ProtoMessage()    {}
func (*IBCData) Descriptor() ([]byte, []int) {
	return fileDescriptor_25eb7c0454d2e150, []int{1}
}
func (m *IBCData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCData.Merge(m, src)
}
func (m *IBCData) XXX_Size() int {
	return m.Size()
}
func (m *IBCData) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCData.DiscardUnknown(m)
}

var xxx_messageInfo_IBCData proto.InternalMessageInfo

func (m *IBCData) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *IBCData) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *IBCData) GetClientToUpdate() string {
	if m != nil {
		return m.ClientToUpdate
	}
	return ""
}

func (m *IBCData) GetHeight() types.Height {
	if m != nil {
		return m.Height
	}
	return types.Height{}
}

func (m *IBCData) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *IBCData) GetPacketCommitments() [][]byte {
	if m != nil {
		return m.PacketCommitments
	}
	return nil
}

func init() {
	proto.RegisterType((*Attestation)(nil), "core.types.v1.Attestation")
	proto.RegisterType((*IBCData)(nil), "core.types.v1.IBCData")
}

func init() { proto.RegisterFile("core/types/v1/attestation.proto", fileDescriptor_25eb7c0454d2e150) }

var fileDescriptor_25eb7c0454d2e150 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x6e, 0x74, 0x8d, 0xd7, 0x21, 0xb0, 0x10, 0x0a, 0x05, 0xa5, 0xd5, 0xb8, 0xf4,
	0x32, 0x5b, 0x85, 0x0b, 0x12, 0xa7, 0x65, 0x1c, 0xe8, 0x35, 0xda, 0x2e, 0x5c, 0x2a, 0xc7, 0x36,
	0xae, 0x61, 0x89, 0xa3, 0xe4, 0xa5, 0x12, 0xdf, 0x62, 0x12, 0x17, 0x3e, 0xd2, 0x8e, 0x3d, 0x72,
	0x02, 0xd4, 0x7e, 0x11, 0x14, 0x3b, 0x21, 0xec, 0x66, 0xff, 0xff, 0xbf, 0xe7, 0xf7, 0xde, 0x5f,
	0xc6, 0x33, 0x61, 0x4b, 0xc5, 0xe0, 0x5b, 0xa1, 0x2a, 0xb6, 0x5d, 0x32, 0x0e, 0xa0, 0x2a, 0xe0,
	0x60, 0x6c, 0x4e, 0x8b, 0xd2, 0x82, 0x25, 0x67, 0x0d, 0x40, 0x1d, 0x40, 0xb7, 0xcb, 0xe9, 0x33,
	0x6d, 0xb5, 0x75, 0x0e, 0x6b, 0x4e, 0x1e, 0x9a, 0xce, 0x4c, 0x2a, 0x98, 0x7b, 0x49, 0xdc, 0x1a,
	0x95, 0x43, 0xf3, 0x94, 0x3f, 0x75, 0x80, 0xb6, 0x56, 0xdf, 0x2a, 0xe6, 0x6e, 0x69, 0xfd, 0x99,
	0x81, 0xc9, 0x9a, 0x46, 0x59, 0xe1, 0x81, 0xf3, 0xef, 0x08, 0x9f, 0x5e, 0xf6, 0xcd, 0xc9, 0x6b,
	0x7c, 0xd6, 0xcd, 0x62, 0xcb, 0xb5, 0x91, 0x21, 0x9a, 0xa3, 0xc5, 0x24, 0x99, 0xf4, 0xe2, 0x4a,
	0x92, 0xcb, 0x0e, 0x52, 0x72, 0x2d, 0x39, 0xf0, 0x70, 0x38, 0x47, 0x8b, 0xd3, 0x37, 0xcf, 0xe9,
	0x83, 0x99, 0xe9, 0x2a, 0xbe, 0xfa, 0xc0, 0x81, 0xc7, 0xc7, 0xf7, 0xbf, 0x66, 0x83, 0xee, 0x09,
	0x25, 0x1b, 0x8d, 0xbc, 0xc2, 0x41, 0x65, 0x74, 0xce, 0xa1, 0x2e, 0x55, 0x78, 0xe4, 0x7a, 0xf4,
	0xc2, 0xf9, 0x8f, 0x21, 0x3e, 0x69, 0xab, 0xc9, 0x0b, 0x3c, 0x16, 0x1b, 0x6e, 0xf2, 0x6e, 0x98,
	0x20, 0x39, 0x71, 0xf7, 0x95, 0x24, 0x2f, 0x71, 0xe0, 0xb7, 0x6d, 0xbc, 0xa1, 0xf3, 0xc6, 0x5e,
	0x58, 0x49, 0xb2, 0xc0, 0x4f, 0x5a, 0x13, 0xec, 0xba, 0x2e, 0x24, 0x07, 0xdf, 0x28, 0x48, 0x1e,
	0x7b, 0xfd, 0xda, 0xde, 0x38, 0x95, 0xbc, 0xc3, 0xa3, 0x8d, 0x32, 0x7a, 0x03, 0xe1, 0xb1, 0xdb,
	0x63, 0x4a, 0x4d, 0x2a, 0xfc, 0x2e, 0x6d, 0x98, 0xdb, 0x25, 0xfd, 0xe8, 0x88, 0x76, 0x97, 0x96,
	0x27, 0x31, 0x0e, 0xfe, 0x05, 0x1a, 0x3e, 0x6a, 0x8b, 0x7d, 0xe4, 0xb4, 0x8b, 0x9c, 0x5e, 0x77,
	0x44, 0x3c, 0x6e, 0x8a, 0xef, 0x7e, 0xcf, 0x50, 0xd2, 0x97, 0x91, 0x0b, 0x4c, 0x0a, 0x2e, 0xbe,
	0x2a, 0x58, 0x0b, 0x9b, 0x65, 0x06, 0x32, 0x95, 0x43, 0x15, 0x8e, 0xe6, 0x47, 0x8b, 0x49, 0xf2,
	0xd4, 0x3b, 0x57, 0xbd, 0x11, 0xdf, 0xdc, 0xef, 0x23, 0xb4, 0xdb, 0x47, 0xe8, 0xcf, 0x3e, 0x42,
	0x77, 0x87, 0x68, 0xb0, 0x3b, 0x44, 0x83, 0x9f, 0x87, 0x68, 0xf0, 0xe9, 0xbd, 0x36, 0xb0, 0xa9,
	0x53, 0x2a, 0x6c, 0xc6, 0xf4, 0x17, 0x55, 0x66, 0x75, 0x2e, 0x35, 0x2f, 0x79, 0xca, 0x99, 0xc9,
	0x41, 0x95, 0x2e, 0xb5, 0x8b, 0xff, 0x7e, 0x1a, 0xeb, 0xff, 0x60, 0x3a, 0x72, 0xe3, 0xbe, 0xfd,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x14, 0xb7, 0xa3, 0x19, 0x98, 0x02, 0x00, 0x00,
}

func (m *Attestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.AttestedData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AttestatorId) > 0 {
		i -= len(m.AttestatorId)
		copy(dAtA[i:], m.AttestatorId)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.AttestatorId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketCommitments) > 0 {
		for iNdEx := len(m.PacketCommitments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PacketCommitments[iNdEx])
			copy(dAtA[i:], m.PacketCommitments[iNdEx])
			i = encodeVarintAttestation(dAtA, i, uint64(len(m.PacketCommitments[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAttestation(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ClientToUpdate) > 0 {
		i -= len(m.ClientToUpdate)
		copy(dAtA[i:], m.ClientToUpdate)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.ClientToUpdate)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttestatorId)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = m.AttestedData.Size()
	n += 1 + l + sovAttestation(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *IBCData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = len(m.ClientToUpdate)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = m.Height.Size()
	n += 1 + l + sovAttestation(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovAttestation(uint64(l))
	if len(m.PacketCommitments) > 0 {
		for _, b := range m.PacketCommitments {
			l = len(b)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: Attestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestatorId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestatorId = append(m.AttestatorId[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestatorId == nil {
				m.AttestatorId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestedData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AttestedData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func (m *IBCData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: IBCData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientToUpdate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientToUpdate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketCommitments", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketCommitments = append(m.PacketCommitments, make([]byte, postIndex-iNdEx))
			copy(m.PacketCommitments[len(m.PacketCommitments)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)
