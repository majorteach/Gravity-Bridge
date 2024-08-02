// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/v1/auction.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// AuctionPeriod represents a period of auctions.
// An AuctionPeriod applies to as many auctionable tokens exist in the auction pool
// Note: see params for a list of non-auctionable tokens
type AuctionPeriod struct {
	StartBlockHeight uint64 `protobuf:"varint,1,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	EndBlockHeight   uint64 `protobuf:"varint,2,opt,name=end_block_height,json=endBlockHeight,proto3" json:"end_block_height,omitempty"`
}

func (m *AuctionPeriod) Reset()         { *m = AuctionPeriod{} }
func (m *AuctionPeriod) String() string { return proto.CompactTextString(m) }
func (*AuctionPeriod) ProtoMessage()    {}
func (*AuctionPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{0}
}
func (m *AuctionPeriod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionPeriod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionPeriod.Merge(m, src)
}
func (m *AuctionPeriod) XXX_Size() int {
	return m.Size()
}
func (m *AuctionPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionPeriod proto.InternalMessageInfo

func (m *AuctionPeriod) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *AuctionPeriod) GetEndBlockHeight() uint64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

// Auction represents a single auction.
// An Auction has a unique identifier relative to its Auction Period Id , an amount being auctioned, a status, and a highest bid.
type Auction struct {
	Id         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount     types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	HighestBid *Bid       `protobuf:"bytes,3,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{1}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Auction) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Auction) GetHighestBid() *Bid {
	if m != nil {
		return m.HighestBid
	}
	return nil
}

// Bid represents a bid on an Auction.
// A Bid includes the identifier of the Auction, the amount of the bid, and the address of the bidder.
type Bid struct {
	BidAmount     uint64 `protobuf:"varint,1,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
	BidderAddress string `protobuf:"bytes,2,opt,name=bidder_address,json=bidderAddress,proto3" json:"bidder_address,omitempty"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{2}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetBidAmount() uint64 {
	if m != nil {
		return m.BidAmount
	}
	return 0
}

func (m *Bid) GetBidderAddress() string {
	if m != nil {
		return m.BidderAddress
	}
	return ""
}

// AuctionId enables easy storage and retrieval of the most recently used AuctionId
type AuctionId struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *AuctionId) Reset()         { *m = AuctionId{} }
func (m *AuctionId) String() string { return proto.CompactTextString(m) }
func (*AuctionId) ProtoMessage()    {}
func (*AuctionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{3}
}
func (m *AuctionId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionId.Merge(m, src)
}
func (m *AuctionId) XXX_Size() int {
	return m.Size()
}
func (m *AuctionId) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionId.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionId proto.InternalMessageInfo

func (m *AuctionId) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*AuctionPeriod)(nil), "auction.v1.AuctionPeriod")
	proto.RegisterType((*Auction)(nil), "auction.v1.Auction")
	proto.RegisterType((*Bid)(nil), "auction.v1.Bid")
	proto.RegisterType((*AuctionId)(nil), "auction.v1.AuctionId")
}

func init() { proto.RegisterFile("auction/v1/auction.proto", fileDescriptor_efe336ece9e41ddd) }

var fileDescriptor_efe336ece9e41ddd = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xbf, 0xee, 0xd3, 0x30,
	0x10, 0xc7, 0x93, 0xb6, 0x2a, 0xaa, 0xab, 0x96, 0xca, 0x62, 0x28, 0x45, 0x04, 0x14, 0x09, 0xa9,
	0x03, 0xd8, 0x04, 0x06, 0x46, 0xd4, 0x30, 0x00, 0x62, 0x41, 0xd9, 0x60, 0x89, 0xe2, 0x9c, 0x95,
	0x9c, 0x68, 0xe2, 0x2a, 0x76, 0x22, 0xba, 0xf3, 0x00, 0x3c, 0x56, 0xc7, 0x8e, 0x4c, 0x08, 0xb5,
	0x2f, 0x82, 0x1a, 0xbb, 0x82, 0x5f, 0xb7, 0xf3, 0xf7, 0x3e, 0xf7, 0xe7, 0xeb, 0x23, 0xcb, 0xac,
	0xcd, 0x0d, 0xaa, 0x9a, 0x77, 0x11, 0x77, 0x21, 0xdb, 0x35, 0xca, 0x28, 0x4a, 0xae, 0xcf, 0x2e,
	0x5a, 0x05, 0xb9, 0xd2, 0x95, 0xd2, 0x5c, 0x64, 0x5a, 0xf2, 0x2e, 0x12, 0xd2, 0x64, 0x11, 0xcf,
	0x15, 0x3a, 0x76, 0xf5, 0xa0, 0x50, 0x85, 0xea, 0x43, 0x7e, 0x89, 0xac, 0x1a, 0x16, 0x64, 0xb6,
	0xb1, 0x3d, 0x3e, 0xcb, 0x06, 0x15, 0xd0, 0xe7, 0x84, 0x6a, 0x93, 0x35, 0x26, 0x15, 0x5b, 0x95,
	0x7f, 0x4b, 0x4b, 0x89, 0x45, 0x69, 0x96, 0xfe, 0x53, 0x7f, 0x3d, 0x4a, 0x16, 0x7d, 0x26, 0xbe,
	0x24, 0x3e, 0xf4, 0x3a, 0x5d, 0x93, 0x85, 0xac, 0xe1, 0x2e, 0x3b, 0xe8, 0xd9, 0xb9, 0xac, 0xe1,
	0x3f, 0x32, 0xfc, 0xe1, 0x93, 0x7b, 0x6e, 0x12, 0x9d, 0x93, 0x01, 0x82, 0xeb, 0x39, 0x40, 0xa0,
	0x6f, 0xc8, 0x38, 0xab, 0x54, 0x5b, 0xdb, 0xda, 0xe9, 0xab, 0x87, 0xcc, 0x7a, 0x61, 0x17, 0x2f,
	0xcc, 0x79, 0x61, 0xef, 0x14, 0xd6, 0xf1, 0xe8, 0xf0, 0xfb, 0x89, 0x97, 0x38, 0x9c, 0xbe, 0x24,
	0xd3, 0x12, 0x8b, 0x52, 0x6a, 0x93, 0x0a, 0x84, 0xe5, 0xb0, 0xaf, 0xbe, 0xcf, 0xfe, 0xfd, 0x0a,
	0x8b, 0x11, 0x12, 0xe2, 0x98, 0x18, 0x21, 0xfc, 0x44, 0x86, 0x31, 0x02, 0x7d, 0x4c, 0x88, 0x40,
	0x48, 0xdd, 0x54, 0xbb, 0xc9, 0x44, 0x20, 0x6c, 0x6c, 0xdf, 0x67, 0x64, 0x2e, 0x10, 0x40, 0x36,
	0x69, 0x06, 0xd0, 0x48, 0xad, 0xfb, 0xc5, 0x26, 0xc9, 0xcc, 0xaa, 0x1b, 0x2b, 0x86, 0x8f, 0xc8,
	0xc4, 0x59, 0xfa, 0x08, 0xb7, 0xa6, 0xe2, 0x2f, 0x87, 0x53, 0xe0, 0x1f, 0x4f, 0x81, 0xff, 0xe7,
	0x14, 0xf8, 0x3f, 0xcf, 0x81, 0x77, 0x3c, 0x07, 0xde, 0xaf, 0x73, 0xe0, 0x7d, 0x7d, 0x5b, 0xa0,
	0x29, 0x5b, 0xc1, 0x72, 0x55, 0xf1, 0xf7, 0x4d, 0xd6, 0xa1, 0xd9, 0xbf, 0x88, 0x1b, 0x84, 0x42,
	0xde, 0x3e, 0x2b, 0x05, 0xed, 0x56, 0xf2, 0xef, 0xd7, 0xb3, 0x73, 0xb3, 0xdf, 0x49, 0x2d, 0xc6,
	0xfd, 0xed, 0x5e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x39, 0xdf, 0xe1, 0xbf, 0x19, 0x02, 0x00,
	0x00,
}

func (m *AuctionPeriod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionPeriod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionPeriod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlockHeight != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.EndBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.StartBlockHeight != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.StartBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HighestBid != nil {
		{
			size, err := m.HighestBid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidderAddress) > 0 {
		i -= len(m.BidderAddress)
		copy(dAtA[i:], m.BidderAddress)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.BidderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BidAmount != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.BidAmount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AuctionId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionPeriod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartBlockHeight != 0 {
		n += 1 + sovAuction(uint64(m.StartBlockHeight))
	}
	if m.EndBlockHeight != 0 {
		n += 1 + sovAuction(uint64(m.EndBlockHeight))
	}
	return n
}

func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.HighestBid != nil {
		l = m.HighestBid.Size()
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BidAmount != 0 {
		n += 1 + sovAuction(uint64(m.BidAmount))
	}
	l = len(m.BidderAddress)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func (m *AuctionId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuctionPeriod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionPeriod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionPeriod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlockHeight", wireType)
			}
			m.StartBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockHeight", wireType)
			}
			m.EndBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HighestBid == nil {
				m.HighestBid = &Bid{}
			}
			if err := m.HighestBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidAmount", wireType)
			}
			m.BidAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BidAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *AuctionId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)