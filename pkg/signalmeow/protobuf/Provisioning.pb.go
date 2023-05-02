//*
// Copyright (C) 2014-2016 Open Whisper Systems
//
// Licensed according to the LICENSE file in this repository.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: Provisioning.proto

package signalpb

import (
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

type ProvisioningVersion int32

const (
	ProvisioningVersion_INITIAL        ProvisioningVersion = 0
	ProvisioningVersion_TABLET_SUPPORT ProvisioningVersion = 1
	ProvisioningVersion_CURRENT        ProvisioningVersion = 1
)

// Enum value maps for ProvisioningVersion.
var (
	ProvisioningVersion_name = map[int32]string{
		0: "INITIAL",
		1: "TABLET_SUPPORT",
		// Duplicate value: 1: "CURRENT",
	}
	ProvisioningVersion_value = map[string]int32{
		"INITIAL":        0,
		"TABLET_SUPPORT": 1,
		"CURRENT":        1,
	}
)

func (x ProvisioningVersion) Enum() *ProvisioningVersion {
	p := new(ProvisioningVersion)
	*p = x
	return p
}

func (x ProvisioningVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProvisioningVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_Provisioning_proto_enumTypes[0].Descriptor()
}

func (ProvisioningVersion) Type() protoreflect.EnumType {
	return &file_Provisioning_proto_enumTypes[0]
}

func (x ProvisioningVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProvisioningVersion) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProvisioningVersion(num)
	return nil
}

// Deprecated: Use ProvisioningVersion.Descriptor instead.
func (ProvisioningVersion) EnumDescriptor() ([]byte, []int) {
	return file_Provisioning_proto_rawDescGZIP(), []int{0}
}

type ProvisioningUuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (x *ProvisioningUuid) Reset() {
	*x = ProvisioningUuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provisioning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningUuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningUuid) ProtoMessage() {}

func (x *ProvisioningUuid) ProtoReflect() protoreflect.Message {
	mi := &file_Provisioning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningUuid.ProtoReflect.Descriptor instead.
func (*ProvisioningUuid) Descriptor() ([]byte, []int) {
	return file_Provisioning_proto_rawDescGZIP(), []int{0}
}

func (x *ProvisioningUuid) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

type ProvisionEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	Body      []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"` // Encrypted ProvisionMessage
}

func (x *ProvisionEnvelope) Reset() {
	*x = ProvisionEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provisioning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionEnvelope) ProtoMessage() {}

func (x *ProvisionEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_Provisioning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionEnvelope.ProtoReflect.Descriptor instead.
func (*ProvisionEnvelope) Descriptor() ([]byte, []int) {
	return file_Provisioning_proto_rawDescGZIP(), []int{1}
}

func (x *ProvisionEnvelope) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *ProvisionEnvelope) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ProvisionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AciIdentityKeyPublic  []byte  `protobuf:"bytes,1,opt,name=aciIdentityKeyPublic" json:"aciIdentityKeyPublic,omitempty"`
	AciIdentityKeyPrivate []byte  `protobuf:"bytes,2,opt,name=aciIdentityKeyPrivate" json:"aciIdentityKeyPrivate,omitempty"`
	PniIdentityKeyPublic  []byte  `protobuf:"bytes,11,opt,name=pniIdentityKeyPublic" json:"pniIdentityKeyPublic,omitempty"`
	PniIdentityKeyPrivate []byte  `protobuf:"bytes,12,opt,name=pniIdentityKeyPrivate" json:"pniIdentityKeyPrivate,omitempty"`
	Aci                   *string `protobuf:"bytes,8,opt,name=aci" json:"aci,omitempty"`
	Pni                   *string `protobuf:"bytes,10,opt,name=pni" json:"pni,omitempty"`
	Number                *string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
	ProvisioningCode      *string `protobuf:"bytes,4,opt,name=provisioningCode" json:"provisioningCode,omitempty"`
	UserAgent             *string `protobuf:"bytes,5,opt,name=userAgent" json:"userAgent,omitempty"`
	ProfileKey            []byte  `protobuf:"bytes,6,opt,name=profileKey" json:"profileKey,omitempty"`
	ReadReceipts          *bool   `protobuf:"varint,7,opt,name=readReceipts" json:"readReceipts,omitempty"`
	ProvisioningVersion   *uint32 `protobuf:"varint,9,opt,name=provisioningVersion" json:"provisioningVersion,omitempty"` // NEXT ID: 13
}

func (x *ProvisionMessage) Reset() {
	*x = ProvisionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Provisioning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionMessage) ProtoMessage() {}

func (x *ProvisionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Provisioning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionMessage.ProtoReflect.Descriptor instead.
func (*ProvisionMessage) Descriptor() ([]byte, []int) {
	return file_Provisioning_proto_rawDescGZIP(), []int{2}
}

func (x *ProvisionMessage) GetAciIdentityKeyPublic() []byte {
	if x != nil {
		return x.AciIdentityKeyPublic
	}
	return nil
}

func (x *ProvisionMessage) GetAciIdentityKeyPrivate() []byte {
	if x != nil {
		return x.AciIdentityKeyPrivate
	}
	return nil
}

func (x *ProvisionMessage) GetPniIdentityKeyPublic() []byte {
	if x != nil {
		return x.PniIdentityKeyPublic
	}
	return nil
}

func (x *ProvisionMessage) GetPniIdentityKeyPrivate() []byte {
	if x != nil {
		return x.PniIdentityKeyPrivate
	}
	return nil
}

func (x *ProvisionMessage) GetAci() string {
	if x != nil && x.Aci != nil {
		return *x.Aci
	}
	return ""
}

func (x *ProvisionMessage) GetPni() string {
	if x != nil && x.Pni != nil {
		return *x.Pni
	}
	return ""
}

func (x *ProvisionMessage) GetNumber() string {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return ""
}

func (x *ProvisionMessage) GetProvisioningCode() string {
	if x != nil && x.ProvisioningCode != nil {
		return *x.ProvisioningCode
	}
	return ""
}

func (x *ProvisionMessage) GetUserAgent() string {
	if x != nil && x.UserAgent != nil {
		return *x.UserAgent
	}
	return ""
}

func (x *ProvisionMessage) GetProfileKey() []byte {
	if x != nil {
		return x.ProfileKey
	}
	return nil
}

func (x *ProvisionMessage) GetReadReceipts() bool {
	if x != nil && x.ReadReceipts != nil {
		return *x.ReadReceipts
	}
	return false
}

func (x *ProvisionMessage) GetProvisioningVersion() uint32 {
	if x != nil && x.ProvisioningVersion != nil {
		return *x.ProvisioningVersion
	}
	return 0
}

var File_Provisioning_proto protoreflect.FileDescriptor

var file_Provisioning_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0xe2, 0x03, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x63, 0x69, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x61, 0x63, 0x69, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x15, 0x61,
	0x63, 0x69, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x61, 0x63, 0x69, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x6e, 0x69, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x4b, 0x65, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x14, 0x70, 0x6e, 0x69, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x15, 0x70, 0x6e, 0x69, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x70, 0x6e, 0x69, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x63, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x69, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x6e, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6e, 0x69, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x47, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x44, 0x0a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x68, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_Provisioning_proto_rawDescOnce sync.Once
	file_Provisioning_proto_rawDescData = file_Provisioning_proto_rawDesc
)

func file_Provisioning_proto_rawDescGZIP() []byte {
	file_Provisioning_proto_rawDescOnce.Do(func() {
		file_Provisioning_proto_rawDescData = protoimpl.X.CompressGZIP(file_Provisioning_proto_rawDescData)
	})
	return file_Provisioning_proto_rawDescData
}

var file_Provisioning_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Provisioning_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Provisioning_proto_goTypes = []interface{}{
	(ProvisioningVersion)(0),  // 0: signalservice.ProvisioningVersion
	(*ProvisioningUuid)(nil),  // 1: signalservice.ProvisioningUuid
	(*ProvisionEnvelope)(nil), // 2: signalservice.ProvisionEnvelope
	(*ProvisionMessage)(nil),  // 3: signalservice.ProvisionMessage
}
var file_Provisioning_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Provisioning_proto_init() }
func file_Provisioning_proto_init() {
	if File_Provisioning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Provisioning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningUuid); i {
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
		file_Provisioning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionEnvelope); i {
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
		file_Provisioning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Provisioning_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Provisioning_proto_goTypes,
		DependencyIndexes: file_Provisioning_proto_depIdxs,
		EnumInfos:         file_Provisioning_proto_enumTypes,
		MessageInfos:      file_Provisioning_proto_msgTypes,
	}.Build()
	File_Provisioning_proto = out.File
	file_Provisioning_proto_rawDesc = nil
	file_Provisioning_proto_goTypes = nil
	file_Provisioning_proto_depIdxs = nil
}