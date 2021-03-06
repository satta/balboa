// Code generated by protoc-gen-go. DO NOT EDIT.
// source: newdomain.proto

package format

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NewDomain struct {
	Domain               []byte         `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	TimeSeen             *uint32        `protobuf:"varint,2,opt,name=time_seen,json=timeSeen" json:"time_seen,omitempty"`
	Type                 *DnsDedupeType `protobuf:"varint,13,opt,name=type,enum=format.DnsDedupeType" json:"type,omitempty"`
	Count                *uint32        `protobuf:"varint,10,opt,name=count" json:"count,omitempty"`
	TimeFirst            *uint32        `protobuf:"varint,11,opt,name=time_first,json=timeFirst" json:"time_first,omitempty"`
	TimeLast             *uint32        `protobuf:"varint,12,opt,name=time_last,json=timeLast" json:"time_last,omitempty"`
	ZoneTimeFirst        *uint32        `protobuf:"varint,17,opt,name=zone_time_first,json=zoneTimeFirst" json:"zone_time_first,omitempty"`
	ZoneTimeLast         *uint32        `protobuf:"varint,18,opt,name=zone_time_last,json=zoneTimeLast" json:"zone_time_last,omitempty"`
	ResponseIp           []byte         `protobuf:"bytes,14,opt,name=response_ip,json=responseIp" json:"response_ip,omitempty"`
	Rrname               []byte         `protobuf:"bytes,3,opt,name=rrname" json:"rrname,omitempty"`
	Rrtype               *uint32        `protobuf:"varint,4,opt,name=rrtype" json:"rrtype,omitempty"`
	Rrclass              *uint32        `protobuf:"varint,5,opt,name=rrclass" json:"rrclass,omitempty"`
	Rrttl                *uint32        `protobuf:"varint,6,opt,name=rrttl" json:"rrttl,omitempty"`
	Rdata                [][]byte       `protobuf:"bytes,7,rep,name=rdata" json:"rdata,omitempty"`
	Response             []byte         `protobuf:"bytes,15,opt,name=response" json:"response,omitempty"`
	Bailiwick            []byte         `protobuf:"bytes,16,opt,name=bailiwick" json:"bailiwick,omitempty"`
	Keys                 [][]byte       `protobuf:"bytes,9,rep,name=keys" json:"keys,omitempty"`
	NewDomain            *bool          `protobuf:"varint,19,opt,name=new_domain,json=newDomain" json:"new_domain,omitempty"`
	NewRrname            *bool          `protobuf:"varint,20,opt,name=new_rrname,json=newRrname" json:"new_rrname,omitempty"`
	NewRrtype            *bool          `protobuf:"varint,21,opt,name=new_rrtype,json=newRrtype" json:"new_rrtype,omitempty"`
	NewRr                []bool         `protobuf:"varint,22,rep,name=new_rr,json=newRr" json:"new_rr,omitempty"`
	NewRrset             *bool          `protobuf:"varint,23,opt,name=new_rrset,json=newRrset" json:"new_rrset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NewDomain) Reset()         { *m = NewDomain{} }
func (m *NewDomain) String() string { return proto.CompactTextString(m) }
func (*NewDomain) ProtoMessage()    {}
func (*NewDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c8736039976de9f, []int{0}
}

func (m *NewDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewDomain.Unmarshal(m, b)
}
func (m *NewDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewDomain.Marshal(b, m, deterministic)
}
func (m *NewDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewDomain.Merge(m, src)
}
func (m *NewDomain) XXX_Size() int {
	return xxx_messageInfo_NewDomain.Size(m)
}
func (m *NewDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_NewDomain.DiscardUnknown(m)
}

var xxx_messageInfo_NewDomain proto.InternalMessageInfo

func (m *NewDomain) GetDomain() []byte {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *NewDomain) GetTimeSeen() uint32 {
	if m != nil && m.TimeSeen != nil {
		return *m.TimeSeen
	}
	return 0
}

func (m *NewDomain) GetType() DnsDedupeType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return DnsDedupeType_INSERTION
}

func (m *NewDomain) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *NewDomain) GetTimeFirst() uint32 {
	if m != nil && m.TimeFirst != nil {
		return *m.TimeFirst
	}
	return 0
}

func (m *NewDomain) GetTimeLast() uint32 {
	if m != nil && m.TimeLast != nil {
		return *m.TimeLast
	}
	return 0
}

func (m *NewDomain) GetZoneTimeFirst() uint32 {
	if m != nil && m.ZoneTimeFirst != nil {
		return *m.ZoneTimeFirst
	}
	return 0
}

func (m *NewDomain) GetZoneTimeLast() uint32 {
	if m != nil && m.ZoneTimeLast != nil {
		return *m.ZoneTimeLast
	}
	return 0
}

func (m *NewDomain) GetResponseIp() []byte {
	if m != nil {
		return m.ResponseIp
	}
	return nil
}

func (m *NewDomain) GetRrname() []byte {
	if m != nil {
		return m.Rrname
	}
	return nil
}

func (m *NewDomain) GetRrtype() uint32 {
	if m != nil && m.Rrtype != nil {
		return *m.Rrtype
	}
	return 0
}

func (m *NewDomain) GetRrclass() uint32 {
	if m != nil && m.Rrclass != nil {
		return *m.Rrclass
	}
	return 0
}

func (m *NewDomain) GetRrttl() uint32 {
	if m != nil && m.Rrttl != nil {
		return *m.Rrttl
	}
	return 0
}

func (m *NewDomain) GetRdata() [][]byte {
	if m != nil {
		return m.Rdata
	}
	return nil
}

func (m *NewDomain) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *NewDomain) GetBailiwick() []byte {
	if m != nil {
		return m.Bailiwick
	}
	return nil
}

func (m *NewDomain) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *NewDomain) GetNewDomain() bool {
	if m != nil && m.NewDomain != nil {
		return *m.NewDomain
	}
	return false
}

func (m *NewDomain) GetNewRrname() bool {
	if m != nil && m.NewRrname != nil {
		return *m.NewRrname
	}
	return false
}

func (m *NewDomain) GetNewRrtype() bool {
	if m != nil && m.NewRrtype != nil {
		return *m.NewRrtype
	}
	return false
}

func (m *NewDomain) GetNewRr() []bool {
	if m != nil {
		return m.NewRr
	}
	return nil
}

func (m *NewDomain) GetNewRrset() bool {
	if m != nil && m.NewRrset != nil {
		return *m.NewRrset
	}
	return false
}

func init() {
	proto.RegisterType((*NewDomain)(nil), "format.NewDomain")
}

func init() { proto.RegisterFile("newdomain.proto", fileDescriptor_6c8736039976de9f) }

var fileDescriptor_6c8736039976de9f = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x0d, 0xb6, 0xf0, 0xe0, 0x3e, 0x5a, 0x74, 0x7c, 0x7d, 0xde, 0x54, 0x8d, 0xc4, 0x18, 0x83,
	0x9b, 0x2e, 0xfc, 0x86, 0xc6, 0xc4, 0xc4, 0xb8, 0xc0, 0xee, 0xc9, 0x58, 0x6e, 0x13, 0x52, 0x18,
	0xc8, 0xcc, 0x34, 0xa4, 0xee, 0xfc, 0xf3, 0x17, 0xee, 0x00, 0xed, 0x8e, 0x73, 0xce, 0x9d, 0x73,
	0xb8, 0x67, 0x06, 0x12, 0x45, 0x7d, 0xd9, 0x36, 0xb2, 0x52, 0xbb, 0x4e, 0xb7, 0xb6, 0x15, 0xc1,
	0xa9, 0xd5, 0x8d, 0xb4, 0xdb, 0xa4, 0x54, 0xa6, 0xa4, 0xf2, 0xd2, 0x91, 0x13, 0x3e, 0xff, 0xf7,
	0x21, 0xfa, 0x4d, 0xfd, 0x9e, 0x87, 0xc5, 0x33, 0x04, 0xee, 0x18, 0x7a, 0xa9, 0x97, 0xc5, 0xf9,
	0x88, 0xc4, 0x7b, 0x88, 0x6c, 0xd5, 0x50, 0x61, 0x88, 0x14, 0xbe, 0x4a, 0xbd, 0x6c, 0x95, 0x87,
	0x03, 0xf1, 0x87, 0x48, 0x89, 0x6f, 0xb0, 0xb4, 0xd7, 0x8e, 0x70, 0x95, 0x7a, 0xd9, 0xfa, 0xfb,
	0x66, 0xe7, 0xa2, 0x76, 0x7b, 0x65, 0xf6, 0x9c, 0x74, 0xb8, 0x76, 0x94, 0xf3, 0x88, 0x78, 0x02,
	0xff, 0xd8, 0x5e, 0x94, 0x45, 0x60, 0x0f, 0x07, 0xc4, 0x47, 0x00, 0x76, 0x3f, 0x55, 0xda, 0x58,
	0x7c, 0x64, 0x89, 0xf3, 0x7e, 0x0c, 0xc4, 0x1c, 0x5e, 0x4b, 0x63, 0x31, 0xbe, 0x85, 0xff, 0x92,
	0xc6, 0x8a, 0xaf, 0x90, 0xfc, 0x6b, 0x15, 0x15, 0x77, 0x06, 0x6f, 0x78, 0x64, 0x35, 0xd0, 0x87,
	0xd9, 0xe4, 0x0b, 0xac, 0x6f, 0x73, 0xec, 0x24, 0x78, 0x2c, 0x9e, 0xc6, 0xd8, 0xed, 0x13, 0x3c,
	0x6a, 0x32, 0x5d, 0xab, 0x0c, 0x15, 0x55, 0x87, 0x6b, 0x2e, 0x01, 0x26, 0xea, 0x67, 0x37, 0x14,
	0xa4, 0xb5, 0x92, 0x0d, 0xe1, 0xc2, 0x15, 0xe4, 0x90, 0xe3, 0xb9, 0x85, 0x25, 0xdb, 0x8e, 0x48,
	0x20, 0x3c, 0x68, 0x7d, 0xac, 0xa5, 0x31, 0xe8, 0xb3, 0x30, 0xc1, 0xa1, 0x0a, 0xad, 0xad, 0xad,
	0x31, 0x70, 0x55, 0x30, 0x60, 0xb6, 0x94, 0x56, 0xe2, 0x43, 0xba, 0xc8, 0xe2, 0xdc, 0x01, 0xb1,
	0x85, 0x70, 0xfa, 0x07, 0x4c, 0x38, 0x77, 0xc6, 0xe2, 0x03, 0x44, 0x7f, 0x65, 0x55, 0x57, 0x7d,
	0x75, 0x3c, 0xe3, 0x6b, 0x16, 0x6f, 0x84, 0x10, 0xb0, 0x3c, 0xd3, 0xd5, 0x60, 0xc4, 0x76, 0xfc,
	0x3d, 0xd4, 0xad, 0xa8, 0x2f, 0xc6, 0x8b, 0x7e, 0x9b, 0x7a, 0x59, 0x98, 0x47, 0x6a, 0x7e, 0x03,
	0xa3, 0x3c, 0xae, 0xf9, 0x34, 0xcb, 0xb9, 0xdb, 0x74, 0x96, 0x79, 0xdb, 0xcd, 0x9d, 0xcc, 0x0b,
	0x6f, 0x20, 0x70, 0x32, 0x3e, 0xa7, 0x8b, 0x2c, 0xcc, 0x7d, 0x96, 0x86, 0x3b, 0x74, 0xb4, 0x21,
	0x8b, 0xef, 0xf8, 0x50, 0xc8, 0x8a, 0x21, 0xfb, 0x12, 0x00, 0x00, 0xff, 0xff, 0x46, 0x1e, 0x03,
	0xca, 0xae, 0x02, 0x00, 0x00,
}
