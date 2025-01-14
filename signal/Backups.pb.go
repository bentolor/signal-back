// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signal/Backups.proto

package signal

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

type SqlStatement struct {
	Statement            *string                      `protobuf:"bytes,1,opt,name=statement" json:"statement,omitempty"`
	Parameters           []*SqlStatement_SqlParameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SqlStatement) Reset()         { *m = SqlStatement{} }
func (m *SqlStatement) String() string { return proto.CompactTextString(m) }
func (*SqlStatement) ProtoMessage()    {}
func (*SqlStatement) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{0}
}

func (m *SqlStatement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlStatement.Unmarshal(m, b)
}
func (m *SqlStatement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlStatement.Marshal(b, m, deterministic)
}
func (m *SqlStatement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlStatement.Merge(m, src)
}
func (m *SqlStatement) XXX_Size() int {
	return xxx_messageInfo_SqlStatement.Size(m)
}
func (m *SqlStatement) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlStatement.DiscardUnknown(m)
}

var xxx_messageInfo_SqlStatement proto.InternalMessageInfo

func (m *SqlStatement) GetStatement() string {
	if m != nil && m.Statement != nil {
		return *m.Statement
	}
	return ""
}

func (m *SqlStatement) GetParameters() []*SqlStatement_SqlParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type SqlStatement_SqlParameter struct {
	StringParameter      *string  `protobuf:"bytes,1,opt,name=stringParameter" json:"stringParameter,omitempty"`
	IntegerParameter     *uint64  `protobuf:"varint,2,opt,name=integerParameter" json:"integerParameter,omitempty"`
	DoubleParameter      *float64 `protobuf:"fixed64,3,opt,name=doubleParameter" json:"doubleParameter,omitempty"`
	BlobParameter        []byte   `protobuf:"bytes,4,opt,name=blobParameter" json:"blobParameter,omitempty"`
	NullParameter        *bool    `protobuf:"varint,5,opt,name=nullParameter" json:"nullParameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqlStatement_SqlParameter) Reset()         { *m = SqlStatement_SqlParameter{} }
func (m *SqlStatement_SqlParameter) String() string { return proto.CompactTextString(m) }
func (*SqlStatement_SqlParameter) ProtoMessage()    {}
func (*SqlStatement_SqlParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{0, 0}
}

func (m *SqlStatement_SqlParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlStatement_SqlParameter.Unmarshal(m, b)
}
func (m *SqlStatement_SqlParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlStatement_SqlParameter.Marshal(b, m, deterministic)
}
func (m *SqlStatement_SqlParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlStatement_SqlParameter.Merge(m, src)
}
func (m *SqlStatement_SqlParameter) XXX_Size() int {
	return xxx_messageInfo_SqlStatement_SqlParameter.Size(m)
}
func (m *SqlStatement_SqlParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlStatement_SqlParameter.DiscardUnknown(m)
}

var xxx_messageInfo_SqlStatement_SqlParameter proto.InternalMessageInfo

func (m *SqlStatement_SqlParameter) GetStringParameter() string {
	if m != nil && m.StringParameter != nil {
		return *m.StringParameter
	}
	return ""
}

func (m *SqlStatement_SqlParameter) GetIntegerParameter() uint64 {
	if m != nil && m.IntegerParameter != nil {
		return *m.IntegerParameter
	}
	return 0
}

func (m *SqlStatement_SqlParameter) GetDoubleParameter() float64 {
	if m != nil && m.DoubleParameter != nil {
		return *m.DoubleParameter
	}
	return 0
}

func (m *SqlStatement_SqlParameter) GetBlobParameter() []byte {
	if m != nil {
		return m.BlobParameter
	}
	return nil
}

func (m *SqlStatement_SqlParameter) GetNullParameter() bool {
	if m != nil && m.NullParameter != nil {
		return *m.NullParameter
	}
	return false
}

type SharedPreference struct {
	File                 *string  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Key                  *string  `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	BooleanValue         *bool    `protobuf:"varint,4,opt,name=booleanValue" json:"booleanValue,omitempty"`
	StringSetValue       []string `protobuf:"bytes,5,rep,name=stringSetValue" json:"stringSetValue,omitempty"`
	IsStringSetValue     *bool    `protobuf:"varint,6,opt,name=isStringSetValue" json:"isStringSetValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedPreference) Reset()         { *m = SharedPreference{} }
func (m *SharedPreference) String() string { return proto.CompactTextString(m) }
func (*SharedPreference) ProtoMessage()    {}
func (*SharedPreference) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{1}
}

func (m *SharedPreference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedPreference.Unmarshal(m, b)
}
func (m *SharedPreference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedPreference.Marshal(b, m, deterministic)
}
func (m *SharedPreference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedPreference.Merge(m, src)
}
func (m *SharedPreference) XXX_Size() int {
	return xxx_messageInfo_SharedPreference.Size(m)
}
func (m *SharedPreference) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedPreference.DiscardUnknown(m)
}

var xxx_messageInfo_SharedPreference proto.InternalMessageInfo

func (m *SharedPreference) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

func (m *SharedPreference) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *SharedPreference) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *SharedPreference) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *SharedPreference) GetStringSetValue() []string {
	if m != nil {
		return m.StringSetValue
	}
	return nil
}

func (m *SharedPreference) GetIsStringSetValue() bool {
	if m != nil && m.IsStringSetValue != nil {
		return *m.IsStringSetValue
	}
	return false
}

type Attachment struct {
	RowId                *uint64  `protobuf:"varint,1,opt,name=rowId" json:"rowId,omitempty"`
	AttachmentId         *uint64  `protobuf:"varint,2,opt,name=attachmentId" json:"attachmentId,omitempty"`
	Length               *uint32  `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{2}
}

func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (m *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(m, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetRowId() uint64 {
	if m != nil && m.RowId != nil {
		return *m.RowId
	}
	return 0
}

func (m *Attachment) GetAttachmentId() uint64 {
	if m != nil && m.AttachmentId != nil {
		return *m.AttachmentId
	}
	return 0
}

func (m *Attachment) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type Sticker struct {
	RowId                *uint64  `protobuf:"varint,1,opt,name=rowId" json:"rowId,omitempty"`
	Length               *uint32  `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sticker) Reset()         { *m = Sticker{} }
func (m *Sticker) String() string { return proto.CompactTextString(m) }
func (*Sticker) ProtoMessage()    {}
func (*Sticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{3}
}

func (m *Sticker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sticker.Unmarshal(m, b)
}
func (m *Sticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sticker.Marshal(b, m, deterministic)
}
func (m *Sticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sticker.Merge(m, src)
}
func (m *Sticker) XXX_Size() int {
	return xxx_messageInfo_Sticker.Size(m)
}
func (m *Sticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Sticker.DiscardUnknown(m)
}

var xxx_messageInfo_Sticker proto.InternalMessageInfo

func (m *Sticker) GetRowId() uint64 {
	if m != nil && m.RowId != nil {
		return *m.RowId
	}
	return 0
}

func (m *Sticker) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type Avatar struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RecipientId          *string  `protobuf:"bytes,3,opt,name=recipientId" json:"recipientId,omitempty"`
	Length               *uint32  `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Avatar) Reset()         { *m = Avatar{} }
func (m *Avatar) String() string { return proto.CompactTextString(m) }
func (*Avatar) ProtoMessage()    {}
func (*Avatar) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{4}
}

func (m *Avatar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Avatar.Unmarshal(m, b)
}
func (m *Avatar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Avatar.Marshal(b, m, deterministic)
}
func (m *Avatar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Avatar.Merge(m, src)
}
func (m *Avatar) XXX_Size() int {
	return xxx_messageInfo_Avatar.Size(m)
}
func (m *Avatar) XXX_DiscardUnknown() {
	xxx_messageInfo_Avatar.DiscardUnknown(m)
}

var xxx_messageInfo_Avatar proto.InternalMessageInfo

func (m *Avatar) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Avatar) GetRecipientId() string {
	if m != nil && m.RecipientId != nil {
		return *m.RecipientId
	}
	return ""
}

func (m *Avatar) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type DatabaseVersion struct {
	Version              *uint32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseVersion) Reset()         { *m = DatabaseVersion{} }
func (m *DatabaseVersion) String() string { return proto.CompactTextString(m) }
func (*DatabaseVersion) ProtoMessage()    {}
func (*DatabaseVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{5}
}

func (m *DatabaseVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseVersion.Unmarshal(m, b)
}
func (m *DatabaseVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseVersion.Marshal(b, m, deterministic)
}
func (m *DatabaseVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseVersion.Merge(m, src)
}
func (m *DatabaseVersion) XXX_Size() int {
	return xxx_messageInfo_DatabaseVersion.Size(m)
}
func (m *DatabaseVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseVersion.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseVersion proto.InternalMessageInfo

func (m *DatabaseVersion) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type Header struct {
	Iv                   []byte   `protobuf:"bytes,1,opt,name=iv" json:"iv,omitempty"`
	Salt                 []byte   `protobuf:"bytes,2,opt,name=salt" json:"salt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{6}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func (m *Header) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

type KeyValue struct {
	Key                  *string  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	BlobValue            []byte   `protobuf:"bytes,2,opt,name=blobValue" json:"blobValue,omitempty"`
	BooleanValue         *bool    `protobuf:"varint,3,opt,name=booleanValue" json:"booleanValue,omitempty"`
	FloatValue           *float32 `protobuf:"fixed32,4,opt,name=floatValue" json:"floatValue,omitempty"`
	IntegerValue         *int32   `protobuf:"varint,5,opt,name=integerValue" json:"integerValue,omitempty"`
	LongValue            *int64   `protobuf:"varint,6,opt,name=longValue" json:"longValue,omitempty"`
	StringValue          *string  `protobuf:"bytes,7,opt,name=stringValue" json:"stringValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{7}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetBlobValue() []byte {
	if m != nil {
		return m.BlobValue
	}
	return nil
}

func (m *KeyValue) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *KeyValue) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *KeyValue) GetIntegerValue() int32 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *KeyValue) GetLongValue() int64 {
	if m != nil && m.LongValue != nil {
		return *m.LongValue
	}
	return 0
}

func (m *KeyValue) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

type BackupFrame struct {
	Header               *Header           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Statement            *SqlStatement     `protobuf:"bytes,2,opt,name=statement" json:"statement,omitempty"`
	Preference           *SharedPreference `protobuf:"bytes,3,opt,name=preference" json:"preference,omitempty"`
	Attachment           *Attachment       `protobuf:"bytes,4,opt,name=attachment" json:"attachment,omitempty"`
	Version              *DatabaseVersion  `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	End                  *bool             `protobuf:"varint,6,opt,name=end" json:"end,omitempty"`
	Avatar               *Avatar           `protobuf:"bytes,7,opt,name=avatar" json:"avatar,omitempty"`
	Sticker              *Sticker          `protobuf:"bytes,8,opt,name=sticker" json:"sticker,omitempty"`
	KeyValue             *KeyValue         `protobuf:"bytes,9,opt,name=keyValue" json:"keyValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BackupFrame) Reset()         { *m = BackupFrame{} }
func (m *BackupFrame) String() string { return proto.CompactTextString(m) }
func (*BackupFrame) ProtoMessage()    {}
func (*BackupFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_70167e9ebde0f9d4, []int{8}
}

func (m *BackupFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupFrame.Unmarshal(m, b)
}
func (m *BackupFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupFrame.Marshal(b, m, deterministic)
}
func (m *BackupFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupFrame.Merge(m, src)
}
func (m *BackupFrame) XXX_Size() int {
	return xxx_messageInfo_BackupFrame.Size(m)
}
func (m *BackupFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupFrame.DiscardUnknown(m)
}

var xxx_messageInfo_BackupFrame proto.InternalMessageInfo

func (m *BackupFrame) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BackupFrame) GetStatement() *SqlStatement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *BackupFrame) GetPreference() *SharedPreference {
	if m != nil {
		return m.Preference
	}
	return nil
}

func (m *BackupFrame) GetAttachment() *Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *BackupFrame) GetVersion() *DatabaseVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *BackupFrame) GetEnd() bool {
	if m != nil && m.End != nil {
		return *m.End
	}
	return false
}

func (m *BackupFrame) GetAvatar() *Avatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *BackupFrame) GetSticker() *Sticker {
	if m != nil {
		return m.Sticker
	}
	return nil
}

func (m *BackupFrame) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*SqlStatement)(nil), "signal.SqlStatement")
	proto.RegisterType((*SqlStatement_SqlParameter)(nil), "signal.SqlStatement.SqlParameter")
	proto.RegisterType((*SharedPreference)(nil), "signal.SharedPreference")
	proto.RegisterType((*Attachment)(nil), "signal.Attachment")
	proto.RegisterType((*Sticker)(nil), "signal.Sticker")
	proto.RegisterType((*Avatar)(nil), "signal.Avatar")
	proto.RegisterType((*DatabaseVersion)(nil), "signal.DatabaseVersion")
	proto.RegisterType((*Header)(nil), "signal.Header")
	proto.RegisterType((*KeyValue)(nil), "signal.KeyValue")
	proto.RegisterType((*BackupFrame)(nil), "signal.BackupFrame")
}

func init() { proto.RegisterFile("signal/Backups.proto", fileDescriptor_70167e9ebde0f9d4) }

var fileDescriptor_70167e9ebde0f9d4 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xce, 0xff, 0x49, 0xda, 0x06, 0x11, 0x36, 0x33, 0xca, 0x70, 0xcd, 0x28, 0xde, 0x56,
	0x32, 0x9a, 0x9b, 0xed, 0xb6, 0x65, 0x8c, 0x95, 0xdd, 0x14, 0x05, 0x7a, 0x39, 0x50, 0x9c, 0x53,
	0xc7, 0xc4, 0xb1, 0x33, 0x49, 0xc9, 0xe8, 0xc3, 0xec, 0x75, 0x76, 0xb5, 0x27, 0xd8, 0xe5, 0x9e,
	0x64, 0x48, 0xb2, 0x2d, 0x3b, 0x6d, 0xef, 0xa4, 0xef, 0x7c, 0xe7, 0xf8, 0xe8, 0x3b, 0x9f, 0x64,
	0x98, 0x88, 0x24, 0xce, 0x58, 0xfa, 0xe1, 0x9a, 0x45, 0xeb, 0xdd, 0x56, 0x4c, 0xb7, 0x3c, 0x97,
	0x39, 0xe9, 0x1a, 0x34, 0xf8, 0xe3, 0xc2, 0x68, 0xfe, 0x23, 0x9d, 0x4b, 0x26, 0x71, 0x83, 0x99,
	0x24, 0xa7, 0x30, 0x10, 0xe5, 0xc6, 0x73, 0x7c, 0x27, 0x1c, 0x50, 0x0b, 0x90, 0x2b, 0x80, 0x2d,
	0xe3, 0x6c, 0x83, 0x12, 0xb9, 0xf0, 0x5c, 0xbf, 0x15, 0x0e, 0x67, 0x67, 0x53, 0x53, 0x6b, 0x5a,
	0xaf, 0xa3, 0x36, 0xb7, 0x25, 0x93, 0xd6, 0x92, 0x5e, 0xfd, 0x75, 0xf4, 0x17, 0xab, 0x20, 0x09,
	0xe1, 0x44, 0x48, 0x9e, 0x64, 0x71, 0x05, 0x15, 0xdf, 0x3d, 0x84, 0xc9, 0x3b, 0x18, 0x27, 0x99,
	0xc4, 0x18, 0xb9, 0xa5, 0xba, 0xbe, 0x13, 0xb6, 0xe9, 0x23, 0x5c, 0x55, 0x5d, 0xe6, 0xbb, 0x45,
	0x8a, 0x96, 0xda, 0xf2, 0x9d, 0xd0, 0xa1, 0x87, 0x30, 0x79, 0x03, 0x47, 0x8b, 0x34, 0x5f, 0x58,
	0x5e, 0xdb, 0x77, 0xc2, 0x11, 0x6d, 0x82, 0x8a, 0x95, 0xed, 0x52, 0xdb, 0xb6, 0xd7, 0xf1, 0x9d,
	0xb0, 0x4f, 0x9b, 0x60, 0xf0, 0xdb, 0x81, 0xf1, 0x7c, 0xc5, 0x38, 0x2e, 0x6f, 0x39, 0xde, 0x23,
	0xc7, 0x2c, 0x42, 0x42, 0xa0, 0x7d, 0x9f, 0xa4, 0x58, 0x9c, 0x4a, 0xaf, 0xc9, 0x18, 0x5a, 0x6b,
	0x7c, 0xd0, 0xdd, 0x0f, 0xa8, 0x5a, 0x92, 0x09, 0x74, 0xf6, 0x2c, 0xdd, 0xa1, 0x6e, 0x73, 0x40,
	0xcd, 0x86, 0x04, 0x30, 0x5a, 0xe4, 0x79, 0x8a, 0x2c, 0xbb, 0xd3, 0xc1, 0xb6, 0xfe, 0x6a, 0x03,
	0x23, 0xe7, 0x70, 0x6c, 0x94, 0x9a, 0xa3, 0x34, 0xac, 0x8e, 0xdf, 0x0a, 0x07, 0xf4, 0x00, 0xd5,
	0xf2, 0x89, 0x79, 0x93, 0xd9, 0xd5, 0xf5, 0x1e, 0xe1, 0xc1, 0x77, 0x80, 0x2b, 0x29, 0x59, 0xb4,
	0xd2, 0x63, 0x9f, 0x40, 0x87, 0xe7, 0x3f, 0x6f, 0x96, 0xfa, 0x08, 0x6d, 0x6a, 0x36, 0xaa, 0x37,
	0x56, 0x71, 0x6e, 0x96, 0xc5, 0x28, 0x1a, 0x18, 0x79, 0x01, 0xdd, 0x14, 0xb3, 0x58, 0xae, 0xf4,
	0xb1, 0x8e, 0x68, 0xb1, 0x0b, 0x3e, 0x42, 0x6f, 0x2e, 0x93, 0x68, 0x8d, 0xfc, 0x99, 0xe2, 0x36,
	0xd1, 0x6d, 0x24, 0xde, 0x41, 0xf7, 0x6a, 0xcf, 0x24, 0xe3, 0x4a, 0xd6, 0x8c, 0x6d, 0x2a, 0x59,
	0xd5, 0x9a, 0xf8, 0x30, 0xe4, 0x18, 0x25, 0xdb, 0xc4, 0x74, 0x64, 0xa4, 0xac, 0x43, 0xcf, 0xd6,
	0x7d, 0x0f, 0x27, 0x9f, 0x99, 0x64, 0x0b, 0x26, 0xf0, 0x0e, 0xb9, 0x48, 0xf2, 0x8c, 0x78, 0xd0,
	0xdb, 0x9b, 0xa5, 0xfe, 0xc6, 0x11, 0x2d, 0xb7, 0xc1, 0x05, 0x74, 0xbf, 0x22, 0x5b, 0x22, 0x27,
	0xc7, 0xe0, 0x26, 0x7b, 0x1d, 0x1e, 0x51, 0x37, 0xd9, 0xab, 0xa6, 0x04, 0x4b, 0xa5, 0x2e, 0x3e,
	0xa2, 0x7a, 0x1d, 0xfc, 0x73, 0xa0, 0xff, 0x0d, 0x1f, 0xcc, 0x10, 0x8a, 0xc1, 0x3b, 0x76, 0xf0,
	0xa7, 0x30, 0x50, 0x56, 0x33, 0xf3, 0x30, 0x79, 0x16, 0x78, 0x64, 0x80, 0xd6, 0x13, 0x06, 0x78,
	0x0d, 0x70, 0x9f, 0xe6, 0x4c, 0x5a, 0x8b, 0xb8, 0xb4, 0x86, 0xa8, 0x1a, 0xc5, 0xfd, 0x28, 0xed,
	0xe1, 0x84, 0x1d, 0xda, 0xc0, 0x54, 0x17, 0x69, 0x9e, 0xc5, 0xd6, 0x15, 0x2d, 0x6a, 0x01, 0xa5,
	0xab, 0x31, 0x93, 0x89, 0xf7, 0x8c, 0xae, 0x35, 0x28, 0xf8, 0xd5, 0x82, 0xa1, 0x79, 0x62, 0xbe,
	0xa8, 0xcb, 0x40, 0xce, 0xa1, 0xbb, 0xd2, 0x12, 0xe9, 0xa3, 0x0e, 0x67, 0xc7, 0xe5, 0x2b, 0x61,
	0x84, 0xa3, 0x45, 0x94, 0xcc, 0xea, 0xef, 0x8d, 0xab, 0xa9, 0x93, 0xa7, 0x1e, 0x94, 0xfa, 0x2b,
	0xf4, 0x09, 0x60, 0x5b, 0x5d, 0x2f, 0xad, 0xc8, 0x70, 0xe6, 0x55, 0x49, 0x07, 0xd7, 0x8f, 0xd6,
	0xb8, 0x64, 0x06, 0x60, 0xed, 0xa9, 0x95, 0x1a, 0xce, 0x48, 0x99, 0x69, 0x0d, 0x4f, 0x6b, 0x2c,
	0x72, 0x69, 0x6d, 0xd0, 0xd1, 0x09, 0x2f, 0xcb, 0x84, 0x03, 0xc3, 0x54, 0xfe, 0x50, 0x43, 0xc6,
	0x6c, 0x59, 0x5c, 0x2e, 0xb5, 0x54, 0x72, 0x30, 0x6d, 0x5b, 0xad, 0x5d, 0x4d, 0x0e, 0x63, 0x66,
	0x5a, 0x44, 0xc9, 0x5b, 0xe8, 0x09, 0x73, 0x2f, 0xbc, 0xbe, 0x26, 0x9e, 0x54, 0xe7, 0x32, 0x30,
	0x2d, 0xe3, 0xe4, 0x02, 0xfa, 0xeb, 0xc2, 0x55, 0xde, 0x40, 0x73, 0xc7, 0x25, 0xb7, 0x74, 0x1b,
	0xad, 0x18, 0xd7, 0x97, 0x70, 0x96, 0xf3, 0x78, 0x2a, 0x57, 0xf9, 0x2e, 0x5e, 0xc9, 0x88, 0x27,
	0x1b, 0x9c, 0x0a, 0x8c, 0x76, 0x1c, 0xc5, 0x46, 0x4c, 0x17, 0x7a, 0x72, 0xd7, 0x23, 0x33, 0xc1,
	0x5b, 0xf5, 0x8b, 0x10, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x87, 0x19, 0x11, 0x3a, 0x06,
	0x00, 0x00,
}
