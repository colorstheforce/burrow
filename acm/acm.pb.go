// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acm.proto

package acm

import (
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	permission "github.com/hyperledger/burrow/permission"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address"`
	PublicKey            crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey"`
	Sequence             uint64                                       `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Balance              uint64                                       `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
	EVMCode              Bytecode                                     `protobuf:"bytes,5,opt,name=EVMCode,proto3,customtype=Bytecode" json:"EVMCode"`
	Permissions          permission.AccountPermissions                `protobuf:"bytes,6,opt,name=Permissions,proto3" json:"Permissions"`
	WASMCode             Bytecode                                     `protobuf:"bytes,7,opt,name=WASMCode,proto3,customtype=Bytecode" json:",omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Account) Reset()      { *m = Account{} }
func (*Account) ProtoMessage() {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ed775bc0a6adf6, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func (m *Account) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetPermissions() permission.AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return permission.AccountPermissions{}
}

func (*Account) XXX_MessageName() string {
	return "acm.Account"
}
func init() {
	proto.RegisterType((*Account)(nil), "acm.Account")
	golang_proto.RegisterType((*Account)(nil), "acm.Account")
}

func init() { proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }
func init() { golang_proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }

var fileDescriptor_49ed775bc0a6adf6 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbf, 0x4f, 0xfa, 0x40,
	0x1c, 0xe5, 0xa0, 0x5f, 0x0a, 0x07, 0x03, 0xdf, 0x9b, 0x1a, 0x86, 0x16, 0x9d, 0x88, 0xc1, 0x36,
	0xf1, 0x47, 0x4c, 0xd8, 0xa8, 0xd1, 0xc5, 0x68, 0x48, 0x49, 0x34, 0x71, 0x6b, 0xaf, 0x67, 0x69,
	0xc2, 0x71, 0xf5, 0x7a, 0x8d, 0xe9, 0x7f, 0xe2, 0xe8, 0x9f, 0xe2, 0xc8, 0xe8, 0xec, 0x40, 0x0c,
	0x6c, 0xfe, 0x0d, 0x0e, 0x86, 0xe3, 0xa8, 0x8d, 0x83, 0x5b, 0x5f, 0xdf, 0x7b, 0x9f, 0xf7, 0xf2,
	0x0e, 0x36, 0x7d, 0x4c, 0xed, 0x84, 0x33, 0xc1, 0x50, 0xcd, 0xc7, 0xb4, 0x7b, 0x18, 0xc5, 0x62,
	0x9a, 0x05, 0x36, 0x66, 0xd4, 0x89, 0x58, 0xc4, 0x1c, 0xc9, 0x05, 0xd9, 0x83, 0x44, 0x12, 0xc8,
	0xaf, 0xad, 0xa7, 0xdb, 0x49, 0x08, 0xa7, 0x71, 0x9a, 0xc6, 0x6c, 0xae, 0xfe, 0xb4, 0x31, 0xcf,
	0x13, 0xa1, 0xf8, 0xfd, 0xaf, 0x2a, 0xd4, 0x47, 0x18, 0xb3, 0x6c, 0x2e, 0xd0, 0x0d, 0xd4, 0x47,
	0x61, 0xc8, 0x49, 0x9a, 0x1a, 0xa0, 0x07, 0xfa, 0x6d, 0xf7, 0x64, 0xb1, 0xb4, 0x2a, 0xef, 0x4b,
	0x6b, 0x50, 0xca, 0x9c, 0xe6, 0x09, 0xe1, 0x33, 0x12, 0x46, 0x84, 0x3b, 0x41, 0xc6, 0x39, 0x7b,
	0x72, 0xd4, 0x41, 0xe5, 0xf5, 0x76, 0x47, 0xd0, 0x29, 0x6c, 0x8e, 0xb3, 0x60, 0x16, 0xe3, 0x2b,
	0x92, 0x1b, 0xd5, 0x1e, 0xe8, 0xb7, 0x8e, 0xfe, 0xdb, 0x4a, 0x5c, 0x10, 0xae, 0xb6, 0x09, 0xf1,
	0x7e, 0x94, 0xa8, 0x0b, 0x1b, 0x13, 0xf2, 0x98, 0x91, 0x39, 0x26, 0x46, 0xad, 0x07, 0xfa, 0x9a,
	0x57, 0x60, 0x64, 0x40, 0xdd, 0xf5, 0x67, 0xfe, 0x86, 0xd2, 0x24, 0xb5, 0x83, 0xe8, 0x00, 0xea,
	0x17, 0xb7, 0xd7, 0xe7, 0x2c, 0x24, 0xc6, 0x3f, 0x59, 0xbe, 0xa3, 0xca, 0x37, 0xdc, 0x5c, 0x10,
	0xcc, 0x42, 0xe2, 0xed, 0x04, 0xe8, 0x12, 0xb6, 0xc6, 0xc5, 0x2c, 0xa9, 0x51, 0x97, 0xd5, 0x4c,
	0xbb, 0x34, 0x95, 0x9a, 0xa4, 0xa4, 0x52, 0x3d, 0xcb, 0x46, 0x34, 0x84, 0x8d, 0xbb, 0xd1, 0x64,
	0x1b, 0xaa, 0xcb, 0x50, 0xf3, 0x77, 0xe8, 0xe7, 0xd2, 0x82, 0x03, 0x46, 0x63, 0x41, 0x68, 0x22,
	0x72, 0xaf, 0xd0, 0x0f, 0xb5, 0xe7, 0x17, 0xab, 0xe2, 0x9e, 0x2d, 0x56, 0x26, 0x78, 0x5b, 0x99,
	0xe0, 0x63, 0x65, 0x82, 0xd7, 0xb5, 0x09, 0x16, 0x6b, 0x13, 0xdc, 0xef, 0xfd, 0xbd, 0xb7, 0x8f,
	0x69, 0x50, 0x97, 0xcf, 0x77, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x65, 0x02, 0xa3, 0x8e, 0x1f,
	0x02, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Address.Size()))
	n1, err := m.Address.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.PublicKey.Size()))
	n2, err := m.PublicKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Sequence != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Sequence))
	}
	if m.Balance != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Balance))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.EVMCode.Size()))
	n3, err := m.EVMCode.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Permissions.Size()))
	n4, err := m.Permissions.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.WASMCode.Size()))
	n5, err := m.WASMCode.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAcm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.PublicKey.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.Sequence != 0 {
		n += 1 + sovAcm(uint64(m.Sequence))
	}
	if m.Balance != 0 {
		n += 1 + sovAcm(uint64(m.Balance))
	}
	l = m.EVMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.Permissions.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.WASMCode.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAcm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAcm(x uint64) (n int) {
	return sovAcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EVMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASMCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WASMCode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAcm
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
func skipAcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAcm
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
				return 0, ErrInvalidLengthAcm
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAcm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAcm
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAcm(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAcm
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAcm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcm   = fmt.Errorf("proto: integer overflow")
)
