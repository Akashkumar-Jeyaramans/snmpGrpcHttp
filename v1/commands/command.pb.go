//
// SNMP Protobuf.
//
// A protobuf interface to SNMP functions.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: v1/commands/command.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
// Asn1Ber Enum Type definitions
type Asn1BER int32

const (
	Asn1BER_EndOfContents     Asn1BER = 0
	Asn1BER_UnknownType       Asn1BER = 0
	Asn1BER_Boolean           Asn1BER = 1
	Asn1BER_Integer           Asn1BER = 2
	Asn1BER_BitString         Asn1BER = 3
	Asn1BER_OctetString       Asn1BER = 4
	Asn1BER_Null              Asn1BER = 5
	Asn1BER_ObjectIdentifier  Asn1BER = 6
	Asn1BER_ObjectDescription Asn1BER = 7
	Asn1BER_IPAddress         Asn1BER = 64
	Asn1BER_Counter32         Asn1BER = 65
	Asn1BER_Gauge32           Asn1BER = 66
	Asn1BER_TimeTicks         Asn1BER = 67
	Asn1BER_Opaque            Asn1BER = 68
	Asn1BER_NsapAddress       Asn1BER = 69
	Asn1BER_Counter64         Asn1BER = 70
	Asn1BER_Uinteger32        Asn1BER = 71
	Asn1BER_OpaqueFloat       Asn1BER = 120
	Asn1BER_OpaqueDouble      Asn1BER = 121
	Asn1BER_NoSuchObject      Asn1BER = 128
	Asn1BER_NoSuchInstance    Asn1BER = 129
	Asn1BER_EndOfMibView      Asn1BER = 130
)

// Enum value maps for Asn1BER.
var (
	Asn1BER_name = map[int32]string{
		0: "EndOfContents",
		// Duplicate value: 0: "UnknownType",
		1:   "Boolean",
		2:   "Integer",
		3:   "BitString",
		4:   "OctetString",
		5:   "Null",
		6:   "ObjectIdentifier",
		7:   "ObjectDescription",
		64:  "IPAddress",
		65:  "Counter32",
		66:  "Gauge32",
		67:  "TimeTicks",
		68:  "Opaque",
		69:  "NsapAddress",
		70:  "Counter64",
		71:  "Uinteger32",
		120: "OpaqueFloat",
		121: "OpaqueDouble",
		128: "NoSuchObject",
		129: "NoSuchInstance",
		130: "EndOfMibView",
	}
	Asn1BER_value = map[string]int32{
		"EndOfContents":     0,
		"UnknownType":       0,
		"Boolean":           1,
		"Integer":           2,
		"BitString":         3,
		"OctetString":       4,
		"Null":              5,
		"ObjectIdentifier":  6,
		"ObjectDescription": 7,
		"IPAddress":         64,
		"Counter32":         65,
		"Gauge32":           66,
		"TimeTicks":         67,
		"Opaque":            68,
		"NsapAddress":       69,
		"Counter64":         70,
		"Uinteger32":        71,
		"OpaqueFloat":       120,
		"OpaqueDouble":      121,
		"NoSuchObject":      128,
		"NoSuchInstance":    129,
		"EndOfMibView":      130,
	}
)

func (x Asn1BER) Enum() *Asn1BER {
	p := new(Asn1BER)
	*p = x
	return p
}

func (x Asn1BER) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Asn1BER) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_commands_command_proto_enumTypes[0].Descriptor()
}

func (Asn1BER) Type() protoreflect.EnumType {
	return &file_v1_commands_command_proto_enumTypes[0]
}

func (x Asn1BER) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Asn1BER.Descriptor instead.
func (Asn1BER) EnumDescriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{0}
}

//
// Represents SNMP OID.
type Oid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"` // SNMP OID
}

func (x *Oid) Reset() {
	*x = Oid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_commands_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oid) ProtoMessage() {}

func (x *Oid) ProtoReflect() protoreflect.Message {
	mi := &file_v1_commands_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oid.ProtoReflect.Descriptor instead.
func (*Oid) Descriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{0}
}

func (x *Oid) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

//
// Represents list of SNMP Oids
type OidList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oids []string `protobuf:"bytes,1,rep,name=oids,proto3" json:"oids,omitempty"` // List of SNMP OIDs
}

func (x *OidList) Reset() {
	*x = OidList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_commands_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OidList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OidList) ProtoMessage() {}

func (x *OidList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_commands_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OidList.ProtoReflect.Descriptor instead.
func (*OidList) Descriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{1}
}

func (x *OidList) GetOids() []string {
	if x != nil {
		return x.Oids
	}
	return nil
}

//
// Represents a single SNMP PDU
// consisting of oid, type and a value.
//
// The value is in any one of the following fields,
// determined by the type of value it stores.
type SnmpPDU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`                              // OID
	Type Asn1BER `protobuf:"varint,2,opt,name=Type,proto3,enum=thebinary.snmp.Asn1BER" json:"Type,omitempty"` // PDU Type (Asn1BER encoding type)
	// Types that are assignable to Value:
	//	*SnmpPDU_I32
	//	*SnmpPDU_I64
	//	*SnmpPDU_UI32
	//	*SnmpPDU_UI64
	//	*SnmpPDU_Str
	Value isSnmpPDU_Value `protobuf_oneof:"Value"`
}

func (x *SnmpPDU) Reset() {
	*x = SnmpPDU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_commands_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpPDU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpPDU) ProtoMessage() {}

func (x *SnmpPDU) ProtoReflect() protoreflect.Message {
	mi := &file_v1_commands_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpPDU.ProtoReflect.Descriptor instead.
func (*SnmpPDU) Descriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{2}
}

func (x *SnmpPDU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SnmpPDU) GetType() Asn1BER {
	if x != nil {
		return x.Type
	}
	return Asn1BER_EndOfContents
}

func (m *SnmpPDU) GetValue() isSnmpPDU_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SnmpPDU) GetI32() int32 {
	if x, ok := x.GetValue().(*SnmpPDU_I32); ok {
		return x.I32
	}
	return 0
}

func (x *SnmpPDU) GetI64() int64 {
	if x, ok := x.GetValue().(*SnmpPDU_I64); ok {
		return x.I64
	}
	return 0
}

func (x *SnmpPDU) GetUI32() uint32 {
	if x, ok := x.GetValue().(*SnmpPDU_UI32); ok {
		return x.UI32
	}
	return 0
}

func (x *SnmpPDU) GetUI64() uint64 {
	if x, ok := x.GetValue().(*SnmpPDU_UI64); ok {
		return x.UI64
	}
	return 0
}

func (x *SnmpPDU) GetStr() string {
	if x, ok := x.GetValue().(*SnmpPDU_Str); ok {
		return x.Str
	}
	return ""
}

type isSnmpPDU_Value interface {
	isSnmpPDU_Value()
}

type SnmpPDU_I32 struct {
	I32 int32 `protobuf:"varint,3,opt,name=I32,proto3,oneof"` // Stores 32-bit integer
}

type SnmpPDU_I64 struct {
	I64 int64 `protobuf:"varint,4,opt,name=I64,proto3,oneof"` // Stores 64-bit signed integer
}

type SnmpPDU_UI32 struct {
	UI32 uint32 `protobuf:"varint,5,opt,name=UI32,proto3,oneof"` // Stores 32-bit unsigned integer
}

type SnmpPDU_UI64 struct {
	UI64 uint64 `protobuf:"varint,6,opt,name=UI64,proto3,oneof"` // Stores 64-bit unsigned integer
}

type SnmpPDU_Str struct {
	Str string `protobuf:"bytes,7,opt,name=Str,proto3,oneof"` // Stores string
}

func (*SnmpPDU_I32) isSnmpPDU_Value() {}

func (*SnmpPDU_I64) isSnmpPDU_Value() {}

func (*SnmpPDU_UI32) isSnmpPDU_Value() {}

func (*SnmpPDU_UI64) isSnmpPDU_Value() {}

func (*SnmpPDU_Str) isSnmpPDU_Value() {}

//
// Represents multiple SNMP PDU
type SnmpPDUs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdus []*SnmpPDU `protobuf:"bytes,1,rep,name=pdus,proto3" json:"pdus,omitempty"`
}

func (x *SnmpPDUs) Reset() {
	*x = SnmpPDUs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_commands_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpPDUs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpPDUs) ProtoMessage() {}

func (x *SnmpPDUs) ProtoReflect() protoreflect.Message {
	mi := &file_v1_commands_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpPDUs.ProtoReflect.Descriptor instead.
func (*SnmpPDUs) Descriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{3}
}

func (x *SnmpPDUs) GetPdus() []*SnmpPDU {
	if x != nil {
		return x.Pdus
	}
	return nil
}

type SnmpPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    uint32     `protobuf:"varint,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Variable []*SnmpPDU `protobuf:"bytes,2,rep,name=Variable,proto3" json:"Variable,omitempty"`
}

func (x *SnmpPacket) Reset() {
	*x = SnmpPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_commands_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpPacket) ProtoMessage() {}

func (x *SnmpPacket) ProtoReflect() protoreflect.Message {
	mi := &file_v1_commands_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpPacket.ProtoReflect.Descriptor instead.
func (*SnmpPacket) Descriptor() ([]byte, []int) {
	return file_v1_commands_command_proto_rawDescGZIP(), []int{4}
}

func (x *SnmpPacket) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SnmpPacket) GetVariable() []*SnmpPDU {
	if x != nil {
		return x.Variable
	}
	return nil
}

var File_v1_commands_command_proto protoreflect.FileDescriptor

var file_v1_commands_command_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x68, 0x65,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x03, 0x4f, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x69, 0x64, 0x22, 0x1d, 0x0a, 0x07, 0x4f, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x69, 0x64,
	0x73, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x53, 0x6e, 0x6d, 0x70, 0x50, 0x44, 0x55, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x74, 0x68, 0x65, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70,
	0x2e, 0x41, 0x73, 0x6e, 0x31, 0x42, 0x45, 0x52, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x03, 0x49, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x49,
	0x33, 0x32, 0x12, 0x12, 0x0a, 0x03, 0x49, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x03, 0x49, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04, 0x55, 0x49, 0x33, 0x32, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x55, 0x49, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x04,
	0x55, 0x49, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x04, 0x55, 0x49,
	0x36, 0x34, 0x12, 0x12, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x53, 0x74, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x37, 0x0a, 0x08, 0x53, 0x6e, 0x6d, 0x70, 0x50, 0x44, 0x55, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x70,
	0x64, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x68, 0x65, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x53, 0x6e, 0x6d, 0x70, 0x50,
	0x44, 0x55, 0x52, 0x04, 0x70, 0x64, 0x75, 0x73, 0x22, 0x57, 0x0a, 0x0a, 0x53, 0x6e, 0x6d, 0x70,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x08,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x68, 0x65, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e,
	0x53, 0x6e, 0x6d, 0x70, 0x50, 0x44, 0x55, 0x52, 0x08, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x2a, 0xf6, 0x02, 0x0a, 0x07, 0x41, 0x73, 0x6e, 0x31, 0x42, 0x45, 0x52, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42,
	0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x63,
	0x74, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x75, 0x6c, 0x6c, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10,
	0x40, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x33, 0x32, 0x10, 0x41,
	0x12, 0x0b, 0x0a, 0x07, 0x47, 0x61, 0x75, 0x67, 0x65, 0x33, 0x32, 0x10, 0x42, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x10, 0x43, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x10, 0x44, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x73, 0x61, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x45, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x36, 0x34, 0x10, 0x46, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x33, 0x32, 0x10, 0x47, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x78, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x79, 0x12, 0x11, 0x0a, 0x0c, 0x4e,
	0x6f, 0x53, 0x75, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x80, 0x01, 0x12, 0x13,
	0x0a, 0x0e, 0x4e, 0x6f, 0x53, 0x75, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x10, 0x81, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x4f, 0x66, 0x4d, 0x69, 0x62, 0x56,
	0x69, 0x65, 0x77, 0x10, 0x82, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x32, 0x57, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74,
	0x68, 0x65, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x4f, 0x69,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x68, 0x65, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x53, 0x6e, 0x6d, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x65, 0x63, 0x68, 0x6f,
	0x3a, 0x01, 0x2a, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x6b, 0x61, 0x73, 0x68, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x2d, 0x4a, 0x65, 0x79,
	0x61, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x73, 0x2f, 0x73, 0x6e, 0x6d, 0x70, 0x47, 0x72, 0x70, 0x63,
	0x48, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_commands_command_proto_rawDescOnce sync.Once
	file_v1_commands_command_proto_rawDescData = file_v1_commands_command_proto_rawDesc
)

func file_v1_commands_command_proto_rawDescGZIP() []byte {
	file_v1_commands_command_proto_rawDescOnce.Do(func() {
		file_v1_commands_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_commands_command_proto_rawDescData)
	})
	return file_v1_commands_command_proto_rawDescData
}

var file_v1_commands_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_commands_command_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_commands_command_proto_goTypes = []interface{}{
	(Asn1BER)(0),       // 0: thebinary.snmp.Asn1BER
	(*Oid)(nil),        // 1: thebinary.snmp.Oid
	(*OidList)(nil),    // 2: thebinary.snmp.OidList
	(*SnmpPDU)(nil),    // 3: thebinary.snmp.SnmpPDU
	(*SnmpPDUs)(nil),   // 4: thebinary.snmp.SnmpPDUs
	(*SnmpPacket)(nil), // 5: thebinary.snmp.SnmpPacket
}
var file_v1_commands_command_proto_depIdxs = []int32{
	0, // 0: thebinary.snmp.SnmpPDU.Type:type_name -> thebinary.snmp.Asn1BER
	3, // 1: thebinary.snmp.SnmpPDUs.pdus:type_name -> thebinary.snmp.SnmpPDU
	3, // 2: thebinary.snmp.SnmpPacket.Variable:type_name -> thebinary.snmp.SnmpPDU
	2, // 3: thebinary.snmp.Command.Get:input_type -> thebinary.snmp.OidList
	5, // 4: thebinary.snmp.Command.Get:output_type -> thebinary.snmp.SnmpPacket
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_commands_command_proto_init() }
func file_v1_commands_command_proto_init() {
	if File_v1_commands_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_commands_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_commands_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OidList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_commands_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpPDU); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_commands_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpPDUs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_commands_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_v1_commands_command_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SnmpPDU_I32)(nil),
		(*SnmpPDU_I64)(nil),
		(*SnmpPDU_UI32)(nil),
		(*SnmpPDU_UI64)(nil),
		(*SnmpPDU_Str)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_commands_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_commands_command_proto_goTypes,
		DependencyIndexes: file_v1_commands_command_proto_depIdxs,
		EnumInfos:         file_v1_commands_command_proto_enumTypes,
		MessageInfos:      file_v1_commands_command_proto_msgTypes,
	}.Build()
	File_v1_commands_command_proto = out.File
	file_v1_commands_command_proto_rawDesc = nil
	file_v1_commands_command_proto_goTypes = nil
	file_v1_commands_command_proto_depIdxs = nil
}