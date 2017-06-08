// Code generated by protoc-gen-go.
// source: annotations.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	annotations.proto
	types.proto

It has these top-level messages:
	SQLIndex
	SQLAssociation
	SQLFieldOptions
	SQLMessageOptions
	SQLFileOptions
	SQLEnumValueOptions
	ID
	DateTime
	Checksum
	UID
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SQLIndex struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Fields string `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *SQLIndex) Reset()                    { *m = SQLIndex{} }
func (m *SQLIndex) String() string            { return proto1.CompactTextString(m) }
func (*SQLIndex) ProtoMessage()               {}
func (*SQLIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SQLIndex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SQLIndex) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *SQLIndex) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SQLAssociation struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Pk    string `protobuf:"bytes,2,opt,name=pk" json:"pk,omitempty"`
	Fk    string `protobuf:"bytes,3,opt,name=fk" json:"fk,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *SQLAssociation) Reset()                    { *m = SQLAssociation{} }
func (m *SQLAssociation) String() string            { return proto1.CompactTextString(m) }
func (*SQLAssociation) ProtoMessage()               {}
func (*SQLAssociation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SQLAssociation) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *SQLAssociation) GetPk() string {
	if m != nil {
		return m.Pk
	}
	return ""
}

func (m *SQLAssociation) GetFk() string {
	if m != nil {
		return m.Fk
	}
	return ""
}

func (m *SQLAssociation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SQLFieldOptions struct {
	Unique     bool   `protobuf:"varint,1,opt,name=unique" json:"unique,omitempty"`
	Index      bool   `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Indextype  string `protobuf:"bytes,3,opt,name=indextype" json:"indextype,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Primarykey bool   `protobuf:"varint,5,opt,name=primarykey" json:"primarykey,omitempty"`
	Notnull    bool   `protobuf:"varint,6,opt,name=notnull" json:"notnull,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	Default    string `protobuf:"bytes,8,opt,name=default" json:"default,omitempty"`
}

func (m *SQLFieldOptions) Reset()                    { *m = SQLFieldOptions{} }
func (m *SQLFieldOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLFieldOptions) ProtoMessage()               {}
func (*SQLFieldOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SQLFieldOptions) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

func (m *SQLFieldOptions) GetIndex() bool {
	if m != nil {
		return m.Index
	}
	return false
}

func (m *SQLFieldOptions) GetIndextype() string {
	if m != nil {
		return m.Indextype
	}
	return ""
}

func (m *SQLFieldOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SQLFieldOptions) GetPrimarykey() bool {
	if m != nil {
		return m.Primarykey
	}
	return false
}

func (m *SQLFieldOptions) GetNotnull() bool {
	if m != nil {
		return m.Notnull
	}
	return false
}

func (m *SQLFieldOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SQLFieldOptions) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

type SQLMessageOptions struct {
	Index       *SQLIndex       `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Index2      *SQLIndex       `protobuf:"bytes,2,opt,name=index2" json:"index2,omitempty"`
	Index3      *SQLIndex       `protobuf:"bytes,3,opt,name=index3" json:"index3,omitempty"`
	Index4      *SQLIndex       `protobuf:"bytes,4,opt,name=index4" json:"index4,omitempty"`
	Index5      *SQLIndex       `protobuf:"bytes,5,opt,name=index5" json:"index5,omitempty"`
	Index6      *SQLIndex       `protobuf:"bytes,6,opt,name=index6" json:"index6,omitempty"`
	Index7      *SQLIndex       `protobuf:"bytes,7,opt,name=index7" json:"index7,omitempty"`
	Index8      *SQLIndex       `protobuf:"bytes,8,opt,name=index8" json:"index8,omitempty"`
	Index9      *SQLIndex       `protobuf:"bytes,9,opt,name=index9" json:"index9,omitempty"`
	Name        string          `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
	BelongsTo   *SQLAssociation `protobuf:"bytes,20,opt,name=belongs_to,json=belongsTo" json:"belongs_to,omitempty"`
	BelongsTo2  *SQLAssociation `protobuf:"bytes,21,opt,name=belongs_to2,json=belongsTo2" json:"belongs_to2,omitempty"`
	BelongsTo3  *SQLAssociation `protobuf:"bytes,22,opt,name=belongs_to3,json=belongsTo3" json:"belongs_to3,omitempty"`
	BelongsTo4  *SQLAssociation `protobuf:"bytes,23,opt,name=belongs_to4,json=belongsTo4" json:"belongs_to4,omitempty"`
	BelongsTo5  *SQLAssociation `protobuf:"bytes,24,opt,name=belongs_to5,json=belongsTo5" json:"belongs_to5,omitempty"`
	BelongsTo6  *SQLAssociation `protobuf:"bytes,25,opt,name=belongs_to6,json=belongsTo6" json:"belongs_to6,omitempty"`
	BelongsTo7  *SQLAssociation `protobuf:"bytes,26,opt,name=belongs_to7,json=belongsTo7" json:"belongs_to7,omitempty"`
	BelongsTo8  *SQLAssociation `protobuf:"bytes,27,opt,name=belongs_to8,json=belongsTo8" json:"belongs_to8,omitempty"`
	BelongsTo9  *SQLAssociation `protobuf:"bytes,28,opt,name=belongs_to9,json=belongsTo9" json:"belongs_to9,omitempty"`
	BelongsTo10 *SQLAssociation `protobuf:"bytes,29,opt,name=belongs_to10,json=belongsTo10" json:"belongs_to10,omitempty"`
	HasOne      *SQLAssociation `protobuf:"bytes,30,opt,name=has_one,json=hasOne" json:"has_one,omitempty"`
	HasOne2     *SQLAssociation `protobuf:"bytes,31,opt,name=has_one2,json=hasOne2" json:"has_one2,omitempty"`
	HasOne3     *SQLAssociation `protobuf:"bytes,32,opt,name=has_one3,json=hasOne3" json:"has_one3,omitempty"`
	HasOne4     *SQLAssociation `protobuf:"bytes,33,opt,name=has_one4,json=hasOne4" json:"has_one4,omitempty"`
	HasOne5     *SQLAssociation `protobuf:"bytes,34,opt,name=has_one5,json=hasOne5" json:"has_one5,omitempty"`
	HasOne6     *SQLAssociation `protobuf:"bytes,35,opt,name=has_one6,json=hasOne6" json:"has_one6,omitempty"`
	HasOne7     *SQLAssociation `protobuf:"bytes,36,opt,name=has_one7,json=hasOne7" json:"has_one7,omitempty"`
	HasOne8     *SQLAssociation `protobuf:"bytes,37,opt,name=has_one8,json=hasOne8" json:"has_one8,omitempty"`
	HasOne9     *SQLAssociation `protobuf:"bytes,38,opt,name=has_one9,json=hasOne9" json:"has_one9,omitempty"`
	HasOne10    *SQLAssociation `protobuf:"bytes,39,opt,name=has_one10,json=hasOne10" json:"has_one10,omitempty"`
	HasMany     *SQLAssociation `protobuf:"bytes,40,opt,name=has_many,json=hasMany" json:"has_many,omitempty"`
	HasMany2    *SQLAssociation `protobuf:"bytes,41,opt,name=has_many2,json=hasMany2" json:"has_many2,omitempty"`
	HasMany3    *SQLAssociation `protobuf:"bytes,42,opt,name=has_many3,json=hasMany3" json:"has_many3,omitempty"`
	HasMany4    *SQLAssociation `protobuf:"bytes,43,opt,name=has_many4,json=hasMany4" json:"has_many4,omitempty"`
	HasMany5    *SQLAssociation `protobuf:"bytes,44,opt,name=has_many5,json=hasMany5" json:"has_many5,omitempty"`
	HasMany6    *SQLAssociation `protobuf:"bytes,45,opt,name=has_many6,json=hasMany6" json:"has_many6,omitempty"`
	HasMany7    *SQLAssociation `protobuf:"bytes,46,opt,name=has_many7,json=hasMany7" json:"has_many7,omitempty"`
	HasMany8    *SQLAssociation `protobuf:"bytes,47,opt,name=has_many8,json=hasMany8" json:"has_many8,omitempty"`
	HasMany9    *SQLAssociation `protobuf:"bytes,48,opt,name=has_many9,json=hasMany9" json:"has_many9,omitempty"`
	HasMany10   *SQLAssociation `protobuf:"bytes,49,opt,name=has_many10,json=hasMany10" json:"has_many10,omitempty"`
	Nogenerate  bool            `protobuf:"varint,50,opt,name=nogenerate" json:"nogenerate,omitempty"`
}

func (m *SQLMessageOptions) Reset()                    { *m = SQLMessageOptions{} }
func (m *SQLMessageOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLMessageOptions) ProtoMessage()               {}
func (*SQLMessageOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SQLMessageOptions) GetIndex() *SQLIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex2() *SQLIndex {
	if m != nil {
		return m.Index2
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex3() *SQLIndex {
	if m != nil {
		return m.Index3
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex4() *SQLIndex {
	if m != nil {
		return m.Index4
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex5() *SQLIndex {
	if m != nil {
		return m.Index5
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex6() *SQLIndex {
	if m != nil {
		return m.Index6
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex7() *SQLIndex {
	if m != nil {
		return m.Index7
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex8() *SQLIndex {
	if m != nil {
		return m.Index8
	}
	return nil
}

func (m *SQLMessageOptions) GetIndex9() *SQLIndex {
	if m != nil {
		return m.Index9
	}
	return nil
}

func (m *SQLMessageOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SQLMessageOptions) GetBelongsTo() *SQLAssociation {
	if m != nil {
		return m.BelongsTo
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo2() *SQLAssociation {
	if m != nil {
		return m.BelongsTo2
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo3() *SQLAssociation {
	if m != nil {
		return m.BelongsTo3
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo4() *SQLAssociation {
	if m != nil {
		return m.BelongsTo4
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo5() *SQLAssociation {
	if m != nil {
		return m.BelongsTo5
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo6() *SQLAssociation {
	if m != nil {
		return m.BelongsTo6
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo7() *SQLAssociation {
	if m != nil {
		return m.BelongsTo7
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo8() *SQLAssociation {
	if m != nil {
		return m.BelongsTo8
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo9() *SQLAssociation {
	if m != nil {
		return m.BelongsTo9
	}
	return nil
}

func (m *SQLMessageOptions) GetBelongsTo10() *SQLAssociation {
	if m != nil {
		return m.BelongsTo10
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne() *SQLAssociation {
	if m != nil {
		return m.HasOne
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne2() *SQLAssociation {
	if m != nil {
		return m.HasOne2
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne3() *SQLAssociation {
	if m != nil {
		return m.HasOne3
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne4() *SQLAssociation {
	if m != nil {
		return m.HasOne4
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne5() *SQLAssociation {
	if m != nil {
		return m.HasOne5
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne6() *SQLAssociation {
	if m != nil {
		return m.HasOne6
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne7() *SQLAssociation {
	if m != nil {
		return m.HasOne7
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne8() *SQLAssociation {
	if m != nil {
		return m.HasOne8
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne9() *SQLAssociation {
	if m != nil {
		return m.HasOne9
	}
	return nil
}

func (m *SQLMessageOptions) GetHasOne10() *SQLAssociation {
	if m != nil {
		return m.HasOne10
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany() *SQLAssociation {
	if m != nil {
		return m.HasMany
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany2() *SQLAssociation {
	if m != nil {
		return m.HasMany2
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany3() *SQLAssociation {
	if m != nil {
		return m.HasMany3
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany4() *SQLAssociation {
	if m != nil {
		return m.HasMany4
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany5() *SQLAssociation {
	if m != nil {
		return m.HasMany5
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany6() *SQLAssociation {
	if m != nil {
		return m.HasMany6
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany7() *SQLAssociation {
	if m != nil {
		return m.HasMany7
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany8() *SQLAssociation {
	if m != nil {
		return m.HasMany8
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany9() *SQLAssociation {
	if m != nil {
		return m.HasMany9
	}
	return nil
}

func (m *SQLMessageOptions) GetHasMany10() *SQLAssociation {
	if m != nil {
		return m.HasMany10
	}
	return nil
}

func (m *SQLMessageOptions) GetNogenerate() bool {
	if m != nil {
		return m.Nogenerate
	}
	return false
}

type SQLFileOptions struct {
	LowercaseEnums bool `protobuf:"varint,1,opt,name=lowercaseEnums" json:"lowercaseEnums,omitempty"`
}

func (m *SQLFileOptions) Reset()                    { *m = SQLFileOptions{} }
func (m *SQLFileOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLFileOptions) ProtoMessage()               {}
func (*SQLFileOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SQLFileOptions) GetLowercaseEnums() bool {
	if m != nil {
		return m.LowercaseEnums
	}
	return false
}

type SQLEnumValueOptions struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *SQLEnumValueOptions) Reset()                    { *m = SQLEnumValueOptions{} }
func (m *SQLEnumValueOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLEnumValueOptions) ProtoMessage()               {}
func (*SQLEnumValueOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SQLEnumValueOptions) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

var E_File = &proto1.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*SQLFileOptions)(nil),
	Field:         1034,
	Name:          "proto.file",
	Tag:           "bytes,1034,opt,name=file",
	Filename:      "annotations.proto",
}

var E_Column = &proto1.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*SQLFieldOptions)(nil),
	Field:         1034,
	Name:          "proto.column",
	Tag:           "bytes,1034,opt,name=column",
	Filename:      "annotations.proto",
}

var E_Table = &proto1.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*SQLMessageOptions)(nil),
	Field:         1034,
	Name:          "proto.table",
	Tag:           "bytes,1034,opt,name=table",
	Filename:      "annotations.proto",
}

var E_Enumval = &proto1.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*SQLEnumValueOptions)(nil),
	Field:         1034,
	Name:          "proto.enumval",
	Tag:           "bytes,1034,opt,name=enumval",
	Filename:      "annotations.proto",
}

func init() {
	proto1.RegisterType((*SQLIndex)(nil), "proto.SQLIndex")
	proto1.RegisterType((*SQLAssociation)(nil), "proto.SQLAssociation")
	proto1.RegisterType((*SQLFieldOptions)(nil), "proto.SQLFieldOptions")
	proto1.RegisterType((*SQLMessageOptions)(nil), "proto.SQLMessageOptions")
	proto1.RegisterType((*SQLFileOptions)(nil), "proto.SQLFileOptions")
	proto1.RegisterType((*SQLEnumValueOptions)(nil), "proto.SQLEnumValueOptions")
	proto1.RegisterExtension(E_File)
	proto1.RegisterExtension(E_Column)
	proto1.RegisterExtension(E_Table)
	proto1.RegisterExtension(E_Enumval)
}

func init() { proto1.RegisterFile("annotations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x92, 0xdb, 0x34,
	0x14, 0xc7, 0x67, 0x43, 0x3e, 0xb5, 0xcc, 0xee, 0x54, 0x6c, 0x97, 0xc3, 0xb2, 0x6d, 0xb7, 0x81,
	0x7e, 0x40, 0x21, 0xeb, 0xf8, 0xdb, 0xb9, 0xe3, 0x82, 0x0e, 0x30, 0x29, 0x25, 0x09, 0xc3, 0x30,
	0xdc, 0x74, 0x9c, 0x44, 0x49, 0x3d, 0x71, 0xa4, 0x10, 0xdb, 0x85, 0xbc, 0x02, 0x6f, 0xc8, 0x2d,
	0x4f, 0xc2, 0x58, 0xb6, 0x6c, 0x39, 0x6d, 0xb2, 0xba, 0x5a, 0x1d, 0xe9, 0xff, 0xfb, 0x1f, 0x9d,
	0xa3, 0x1d, 0xe7, 0xa0, 0x7b, 0x3e, 0xa5, 0x2c, 0xf6, 0xe3, 0x80, 0xd1, 0xa8, 0xb7, 0xd9, 0xb2,
	0x98, 0xe1, 0x06, 0xff, 0x73, 0x75, 0xb3, 0x64, 0x6c, 0x19, 0x92, 0x5b, 0x1e, 0x4d, 0x93, 0xc5,
	0xed, 0x9c, 0x44, 0xb3, 0x6d, 0xb0, 0x89, 0xd9, 0x36, 0x13, 0x76, 0x7f, 0x42, 0xed, 0xc9, 0x68,
	0xf8, 0x23, 0x9d, 0x93, 0xbf, 0x31, 0x46, 0x75, 0xea, 0xaf, 0x09, 0x9c, 0xdc, 0x9c, 0x3c, 0xef,
	0x8c, 0xf9, 0x1a, 0x5f, 0xa2, 0xe6, 0x22, 0x20, 0xe1, 0x3c, 0x82, 0x1a, 0xdf, 0xcd, 0xa3, 0x54,
	0x1b, 0xef, 0x36, 0x04, 0x3e, 0xca, 0xb4, 0xe9, 0xba, 0xfb, 0x07, 0x3a, 0x9b, 0x8c, 0x86, 0xdf,
	0x45, 0x11, 0x9b, 0x05, 0xfc, 0x36, 0xf8, 0x02, 0x35, 0x62, 0x7f, 0x1a, 0x0a, 0xcb, 0x2c, 0xc0,
	0x67, 0xa8, 0xb6, 0x59, 0xe5, 0x7e, 0xb5, 0xcd, 0x2a, 0x8d, 0x17, 0xab, 0xdc, 0xa9, 0xb6, 0x58,
	0x15, 0xf7, 0xa8, 0x97, 0xf7, 0xe8, 0xfe, 0x7b, 0x82, 0xce, 0x27, 0xa3, 0xe1, 0xcb, 0x34, 0xfb,
	0xeb, 0x0d, 0x2f, 0x35, 0xbd, 0x5b, 0x42, 0x83, 0x3f, 0x93, 0xcc, 0xbe, 0x3d, 0xce, 0xa3, 0x34,
	0x6b, 0x90, 0x16, 0xc4, 0x53, 0xb4, 0xc7, 0x59, 0x80, 0xaf, 0x51, 0x87, 0x2f, 0xa4, 0x6b, 0x97,
	0x1b, 0x45, 0x3d, 0xf5, 0xb2, 0x1e, 0xfc, 0x10, 0xa1, 0xcd, 0x36, 0x58, 0xfb, 0xdb, 0xdd, 0x8a,
	0xec, 0xa0, 0xc1, 0xcd, 0xa4, 0x1d, 0x0c, 0xa8, 0x45, 0x59, 0x4c, 0x93, 0x30, 0x84, 0x26, 0x3f,
	0x14, 0x61, 0x51, 0x41, 0x4b, 0xea, 0x24, 0xa0, 0xd6, 0x9c, 0x2c, 0xfc, 0x24, 0x8c, 0xa1, 0xcd,
	0xb7, 0x45, 0xd8, 0xfd, 0xef, 0x1c, 0xdd, 0x9b, 0x8c, 0x86, 0xaf, 0x48, 0x14, 0xf9, 0x4b, 0x22,
	0xaa, 0x7b, 0x22, 0xaa, 0x48, 0x8b, 0x3b, 0xd5, 0xcf, 0xb3, 0x07, 0xeb, 0x89, 0xd7, 0x12, 0x65,
	0x3d, 0x43, 0x4d, 0xbe, 0xd0, 0x79, 0xb5, 0x1f, 0xd0, 0xe5, 0xc7, 0x85, 0xd0, 0xe0, 0xc5, 0x1f,
	0x14, 0x1a, 0x85, 0xd0, 0xe4, 0xcd, 0x38, 0x28, 0x34, 0x0b, 0xa1, 0xc5, 0x7b, 0x73, 0x50, 0x68,
	0x15, 0x42, 0x9b, 0xf7, 0xe9, 0xa0, 0xd0, 0x2e, 0x84, 0x0e, 0xef, 0xdc, 0x41, 0xa1, 0x53, 0x08,
	0x5d, 0xde, 0xcb, 0x83, 0x42, 0xb7, 0x10, 0x7a, 0xd0, 0x39, 0x26, 0xf4, 0x8a, 0x27, 0x43, 0xd2,
	0x93, 0x99, 0x08, 0x4d, 0x49, 0xc8, 0xe8, 0x32, 0x7a, 0x13, 0x33, 0xb8, 0xe0, 0x06, 0xf7, 0x4b,
	0x03, 0xe9, 0x3f, 0x7d, 0xdc, 0xc9, 0x85, 0xbf, 0x32, 0x6c, 0xa3, 0xd3, 0x92, 0xd2, 0xe1, 0xfe,
	0x31, 0x0c, 0x15, 0x98, 0x5e, 0xe5, 0x0c, 0xb8, 0x54, 0xe3, 0x8c, 0x2a, 0x67, 0xc2, 0xa7, 0x6a,
	0x9c, 0x59, 0xe5, 0x2c, 0x00, 0x35, 0xce, 0xaa, 0x72, 0x36, 0x7c, 0xa6, 0xc6, 0xd9, 0x55, 0xce,
	0x81, 0x2b, 0x35, 0xce, 0xa9, 0x72, 0x2e, 0x7c, 0xae, 0xc6, 0xb9, 0x55, 0xce, 0x83, 0x6b, 0x35,
	0xce, 0xc3, 0x2e, 0xfa, 0xb8, 0xe4, 0xfa, 0x1a, 0x3c, 0x38, 0x06, 0x9e, 0x16, 0x60, 0x5f, 0xc3,
	0x3d, 0xd4, 0x7a, 0xeb, 0x47, 0x6f, 0x18, 0x25, 0xf0, 0xf0, 0x18, 0xd4, 0x7c, 0xeb, 0x47, 0xaf,
	0x29, 0xc1, 0x1a, 0x6a, 0xe7, 0x7a, 0x1d, 0x1e, 0x1d, 0x03, 0x5a, 0x19, 0xa0, 0x4b, 0x84, 0x01,
	0x37, 0x0a, 0x84, 0x21, 0x11, 0x26, 0x3c, 0x56, 0x20, 0x4c, 0x89, 0xb0, 0xa0, 0xab, 0x40, 0x58,
	0x12, 0x61, 0xc3, 0x17, 0x0a, 0x84, 0x2d, 0x11, 0x0e, 0x7c, 0xa9, 0x40, 0x38, 0x12, 0xe1, 0xc2,
	0x13, 0x05, 0xc2, 0x95, 0x08, 0x0f, 0x9e, 0x2a, 0x10, 0x1e, 0xd6, 0x51, 0x27, 0x27, 0xfa, 0x1a,
	0x3c, 0x3b, 0x86, 0xb4, 0x33, 0xa4, 0xaf, 0x89, 0x2c, 0x6b, 0x9f, 0xee, 0xe0, 0xf9, 0x5d, 0x59,
	0x5e, 0xf9, 0x74, 0x27, 0xb2, 0xa4, 0x84, 0x0e, 0x5f, 0xdd, 0x95, 0x25, 0x45, 0x74, 0x99, 0x31,
	0xe0, 0x6b, 0x15, 0xc6, 0x90, 0x19, 0x13, 0x5e, 0xa8, 0x30, 0xa6, 0xcc, 0x58, 0xf0, 0x8d, 0x0a,
	0x63, 0xc9, 0x8c, 0x0d, 0xdf, 0xaa, 0x30, 0xb6, 0xcc, 0x38, 0xd0, 0x53, 0x61, 0x1c, 0x99, 0x71,
	0xe1, 0x56, 0x85, 0x71, 0x65, 0xc6, 0x03, 0x4d, 0x85, 0xf1, 0xd2, 0xaf, 0xbe, 0x60, 0xfa, 0x1a,
	0xf4, 0x8f, 0x7e, 0xf5, 0x73, 0xa8, 0xaf, 0xa5, 0xc3, 0x02, 0x65, 0x4b, 0x42, 0xc9, 0xd6, 0x8f,
	0x09, 0xe8, 0xd9, 0xb0, 0x50, 0xee, 0x74, 0x5d, 0x3e, 0x1c, 0xbd, 0x0c, 0xc2, 0xe2, 0x07, 0xfe,
	0x29, 0x3a, 0x0b, 0xd9, 0x5f, 0x64, 0x3b, 0xf3, 0x23, 0xf2, 0x3d, 0x4d, 0xd6, 0x51, 0x3e, 0xc6,
	0xec, 0xed, 0x76, 0x5f, 0xa0, 0x4f, 0x26, 0xa3, 0x61, 0xba, 0xfe, 0xcd, 0x0f, 0x93, 0x02, 0xbf,
	0x40, 0x8d, 0x77, 0x69, 0x2c, 0x66, 0x2b, 0x1e, 0x0c, 0x7e, 0x40, 0xf5, 0x45, 0x10, 0x12, 0x7c,
	0xdd, 0xcb, 0x46, 0xbf, 0x9e, 0x18, 0xfd, 0x7a, 0x52, 0x6a, 0xf8, 0xa7, 0xbd, 0x5f, 0x95, 0x74,
	0x3a, 0xe6, 0x0e, 0x83, 0x9f, 0x51, 0x73, 0xc6, 0xc2, 0x64, 0x4d, 0xf1, 0x83, 0x0f, 0x78, 0x95,
	0x63, 0x98, 0x30, 0xbb, 0x94, 0xcd, 0xca, 0xe3, 0x71, 0xee, 0x32, 0xf8, 0x25, 0x9f, 0x05, 0xf1,
	0xa3, 0xf7, 0xec, 0xaa, 0x93, 0x8f, 0x30, 0x84, 0xd2, 0xb0, 0x2a, 0xc8, 0xe7, 0xc8, 0xc1, 0xef,
	0xa8, 0x45, 0x68, 0xb2, 0x7e, 0xe7, 0x87, 0xf8, 0xf1, 0x7b, 0x9e, 0xfb, 0xfd, 0x12, 0xae, 0x57,
	0xa5, 0xeb, 0xbe, 0x64, 0x2c, 0xec, 0xa6, 0x4d, 0xae, 0x31, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x17, 0x49, 0xc4, 0xa2, 0x5a, 0x0b, 0x00, 0x00,
}
