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
	GraphQLType
	GraphQLUnion
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

type GraphQLType struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Generate   bool   `protobuf:"varint,3,opt,name=generate" json:"generate,omitempty"`
	Definition string `protobuf:"bytes,4,opt,name=definition" json:"definition,omitempty"`
}

func (m *GraphQLType) Reset()                    { *m = GraphQLType{} }
func (m *GraphQLType) String() string            { return proto1.CompactTextString(m) }
func (*GraphQLType) ProtoMessage()               {}
func (*GraphQLType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GraphQLType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphQLType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GraphQLType) GetGenerate() bool {
	if m != nil {
		return m.Generate
	}
	return false
}

func (m *GraphQLType) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

type GraphQLUnion struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tables string `protobuf:"bytes,4,opt,name=tables" json:"tables,omitempty"`
}

func (m *GraphQLUnion) Reset()                    { *m = GraphQLUnion{} }
func (m *GraphQLUnion) String() string            { return proto1.CompactTextString(m) }
func (*GraphQLUnion) ProtoMessage()               {}
func (*GraphQLUnion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GraphQLUnion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphQLUnion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GraphQLUnion) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GraphQLUnion) GetTables() string {
	if m != nil {
		return m.Tables
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
	GqlAddType  *GraphQLType    `protobuf:"bytes,51,opt,name=gql_add_type,json=gqlAddType" json:"gql_add_type,omitempty"`
	GqlAddType1 *GraphQLType    `protobuf:"bytes,52,opt,name=gql_add_type1,json=gqlAddType1" json:"gql_add_type1,omitempty"`
	GqlAddType2 *GraphQLType    `protobuf:"bytes,53,opt,name=gql_add_type2,json=gqlAddType2" json:"gql_add_type2,omitempty"`
	GqlAddType3 *GraphQLType    `protobuf:"bytes,54,opt,name=gql_add_type3,json=gqlAddType3" json:"gql_add_type3,omitempty"`
	GqlAddType4 *GraphQLType    `protobuf:"bytes,55,opt,name=gql_add_type4,json=gqlAddType4" json:"gql_add_type4,omitempty"`
	GqlAddType5 *GraphQLType    `protobuf:"bytes,56,opt,name=gql_add_type5,json=gqlAddType5" json:"gql_add_type5,omitempty"`
	GqlAddType6 *GraphQLType    `protobuf:"bytes,57,opt,name=gql_add_type6,json=gqlAddType6" json:"gql_add_type6,omitempty"`
	GqlAddType7 *GraphQLType    `protobuf:"bytes,58,opt,name=gql_add_type7,json=gqlAddType7" json:"gql_add_type7,omitempty"`
	GqlAddType8 *GraphQLType    `protobuf:"bytes,59,opt,name=gql_add_type8,json=gqlAddType8" json:"gql_add_type8,omitempty"`
	GqlAddType9 *GraphQLType    `protobuf:"bytes,60,opt,name=gql_add_type9,json=gqlAddType9" json:"gql_add_type9,omitempty"`
	GqlUnion    *GraphQLUnion   `protobuf:"bytes,61,opt,name=gql_union,json=gqlUnion" json:"gql_union,omitempty"`
	GqlUnion2   *GraphQLUnion   `protobuf:"bytes,62,opt,name=gql_union2,json=gqlUnion2" json:"gql_union2,omitempty"`
	GqlUnion3   *GraphQLUnion   `protobuf:"bytes,63,opt,name=gql_union3,json=gqlUnion3" json:"gql_union3,omitempty"`
	GqlUnion4   *GraphQLUnion   `protobuf:"bytes,64,opt,name=gql_union4,json=gqlUnion4" json:"gql_union4,omitempty"`
	GqlUnion5   *GraphQLUnion   `protobuf:"bytes,65,opt,name=gql_union5,json=gqlUnion5" json:"gql_union5,omitempty"`
	GqlUnion6   *GraphQLUnion   `protobuf:"bytes,66,opt,name=gql_union6,json=gqlUnion6" json:"gql_union6,omitempty"`
	GqlUnion7   *GraphQLUnion   `protobuf:"bytes,67,opt,name=gql_union7,json=gqlUnion7" json:"gql_union7,omitempty"`
	GqlUnion8   *GraphQLUnion   `protobuf:"bytes,68,opt,name=gql_union8,json=gqlUnion8" json:"gql_union8,omitempty"`
	GqlUnion9   *GraphQLUnion   `protobuf:"bytes,69,opt,name=gql_union9,json=gqlUnion9" json:"gql_union9,omitempty"`
}

func (m *SQLMessageOptions) Reset()                    { *m = SQLMessageOptions{} }
func (m *SQLMessageOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLMessageOptions) ProtoMessage()               {}
func (*SQLMessageOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

func (m *SQLMessageOptions) GetGqlAddType() *GraphQLType {
	if m != nil {
		return m.GqlAddType
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType1() *GraphQLType {
	if m != nil {
		return m.GqlAddType1
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType2() *GraphQLType {
	if m != nil {
		return m.GqlAddType2
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType3() *GraphQLType {
	if m != nil {
		return m.GqlAddType3
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType4() *GraphQLType {
	if m != nil {
		return m.GqlAddType4
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType5() *GraphQLType {
	if m != nil {
		return m.GqlAddType5
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType6() *GraphQLType {
	if m != nil {
		return m.GqlAddType6
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType7() *GraphQLType {
	if m != nil {
		return m.GqlAddType7
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType8() *GraphQLType {
	if m != nil {
		return m.GqlAddType8
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlAddType9() *GraphQLType {
	if m != nil {
		return m.GqlAddType9
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion2() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion2
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion3() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion3
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion4() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion4
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion5() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion5
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion6() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion6
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion7() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion7
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion8() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion8
	}
	return nil
}

func (m *SQLMessageOptions) GetGqlUnion9() *GraphQLUnion {
	if m != nil {
		return m.GqlUnion9
	}
	return nil
}

type SQLFileOptions struct {
	LowercaseEnums bool `protobuf:"varint,1,opt,name=lowercaseEnums" json:"lowercaseEnums,omitempty"`
}

func (m *SQLFileOptions) Reset()                    { *m = SQLFileOptions{} }
func (m *SQLFileOptions) String() string            { return proto1.CompactTextString(m) }
func (*SQLFileOptions) ProtoMessage()               {}
func (*SQLFileOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*SQLEnumValueOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	proto1.RegisterType((*GraphQLType)(nil), "proto.GraphQLType")
	proto1.RegisterType((*GraphQLUnion)(nil), "proto.GraphQLUnion")
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
	// 1089 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd7, 0xeb, 0x72, 0xdb, 0x44,
	0x14, 0x00, 0xe0, 0xa9, 0x49, 0x7c, 0x39, 0x09, 0xe9, 0x74, 0x9b, 0xa6, 0x87, 0x90, 0xb6, 0xa9,
	0xa1, 0x17, 0x28, 0x38, 0xb6, 0xee, 0x0a, 0xd7, 0x00, 0x2d, 0x97, 0x71, 0x29, 0x76, 0x0a, 0xc3,
	0xf0, 0x83, 0x8c, 0x12, 0xc9, 0x8e, 0x26, 0xf2, 0xca, 0xb1, 0xac, 0x82, 0x5f, 0x81, 0x37, 0xe4,
	0x2f, 0x4f, 0xc2, 0xec, 0xea, 0xb6, 0x52, 0x62, 0x77, 0x7f, 0x45, 0xbb, 0x3e, 0xdf, 0x39, 0xbb,
	0x3e, 0x8a, 0x67, 0x17, 0x6e, 0x39, 0x94, 0x86, 0x73, 0x67, 0xee, 0x87, 0x34, 0xea, 0x4c, 0x67,
	0xe1, 0x3c, 0x24, 0xeb, 0xfc, 0xcf, 0xee, 0xfe, 0x38, 0x0c, 0xc7, 0x81, 0x77, 0xc0, 0x47, 0xa7,
	0xf1, 0xe8, 0xc0, 0xf5, 0xa2, 0xb3, 0x99, 0x3f, 0x9d, 0x87, 0xb3, 0x24, 0xb0, 0xfd, 0x13, 0x34,
	0x8f, 0x07, 0xfd, 0x1f, 0xa9, 0xeb, 0xfd, 0x4d, 0x08, 0xac, 0x51, 0x67, 0xe2, 0xe1, 0x8d, 0xfd,
	0x1b, 0x4f, 0x5b, 0x43, 0xfe, 0x4c, 0x76, 0xa0, 0x3e, 0xf2, 0xbd, 0xc0, 0x8d, 0xb0, 0xc6, 0x67,
	0xd3, 0x11, 0x8b, 0x9d, 0x2f, 0xa6, 0x1e, 0xbe, 0x93, 0xc4, 0xb2, 0xe7, 0xf6, 0x1f, 0xb0, 0x75,
	0x3c, 0xe8, 0x1f, 0x45, 0x51, 0x78, 0xe6, 0xf3, 0xd5, 0x90, 0x6d, 0x58, 0x9f, 0x3b, 0xa7, 0x41,
	0x96, 0x32, 0x19, 0x90, 0x2d, 0xa8, 0x4d, 0x2f, 0xd2, 0x7c, 0xb5, 0xe9, 0x05, 0x1b, 0x8f, 0x2e,
	0xd2, 0x4c, 0xb5, 0xd1, 0x45, 0xbe, 0x8e, 0xb5, 0x62, 0x1d, 0xed, 0x7f, 0x6f, 0xc0, 0xcd, 0xe3,
	0x41, 0xff, 0x05, 0xab, 0xfe, 0x6a, 0xca, 0xb7, 0xca, 0xd6, 0x16, 0x53, 0xff, 0x32, 0x4e, 0xd2,
	0x37, 0x87, 0xe9, 0x88, 0x55, 0xf5, 0xd9, 0x86, 0x78, 0x89, 0xe6, 0x30, 0x19, 0x90, 0x3d, 0x68,
	0xf1, 0x07, 0x61, 0xd9, 0xc5, 0x44, 0xbe, 0x9f, 0xb5, 0x62, 0x3f, 0xe4, 0x3e, 0xc0, 0x74, 0xe6,
	0x4f, 0x9c, 0xd9, 0xe2, 0xc2, 0x5b, 0xe0, 0x3a, 0x4f, 0x26, 0xcc, 0x10, 0x84, 0x06, 0x0d, 0xe7,
	0x34, 0x0e, 0x02, 0xac, 0xf3, 0x0f, 0xb3, 0x61, 0xbe, 0x83, 0x86, 0xf0, 0x4d, 0x22, 0x34, 0x5c,
	0x6f, 0xe4, 0xc4, 0xc1, 0x1c, 0x9b, 0x7c, 0x3a, 0x1b, 0xb6, 0x2f, 0x61, 0xe3, 0xfb, 0x99, 0x33,
	0x3d, 0x1f, 0xf4, 0x5f, 0xa7, 0x4b, 0xb9, 0xd2, 0x86, 0x6c, 0x79, 0x35, 0x61, 0x79, 0xbb, 0xd0,
	0x1c, 0x7b, 0xd4, 0x9b, 0x39, 0xf3, 0x64, 0x3f, 0xcd, 0x61, 0x3e, 0x66, 0x4b, 0x77, 0xbd, 0x91,
	0x4f, 0x7d, 0xf6, 0x4d, 0xa5, 0x9b, 0x12, 0x66, 0xda, 0x7f, 0xc2, 0x66, 0x5a, 0xf2, 0x57, 0xca,
	0x1a, 0x75, 0x5d, 0xcd, 0x2d, 0xa8, 0xf9, 0x6e, 0xd6, 0x26, 0xdf, 0xbd, 0xae, 0xe5, 0xac, 0x05,
	0xbc, 0xa7, 0x51, 0x5a, 0x23, 0x1d, 0xb5, 0xff, 0xbb, 0x0b, 0xb7, 0x8e, 0x07, 0xfd, 0x97, 0x5e,
	0x14, 0x39, 0x63, 0x2f, 0x6b, 0xd8, 0xa3, 0xac, 0x31, 0xac, 0xcc, 0x86, 0x72, 0x33, 0x79, 0x07,
	0x3b, 0xd9, 0x0b, 0x98, 0x75, 0xea, 0x09, 0xd4, 0xf9, 0x83, 0xc2, 0x8b, 0x5f, 0x13, 0x97, 0x7e,
	0x9c, 0x07, 0xaa, 0x7c, 0x4d, 0x4b, 0x03, 0xd5, 0x3c, 0x50, 0xe3, 0xcb, 0x5c, 0x1a, 0xa8, 0xe5,
	0x81, 0x3a, 0x6f, 0xf7, 0xd2, 0x40, 0x3d, 0x0f, 0x34, 0x78, 0xeb, 0x97, 0x06, 0x1a, 0x79, 0xa0,
	0xc9, 0x5f, 0x86, 0xa5, 0x81, 0x66, 0x1e, 0x68, 0xf1, 0xd7, 0x63, 0x69, 0xa0, 0x95, 0x07, 0xda,
	0xd8, 0x5a, 0x15, 0x68, 0xe7, 0x4d, 0x05, 0xa1, 0xa9, 0x1a, 0xc0, 0xa9, 0x17, 0x84, 0x74, 0x1c,
	0x9d, 0xcc, 0x43, 0xdc, 0xe6, 0x09, 0xee, 0x14, 0x09, 0x84, 0x7f, 0xde, 0x61, 0x2b, 0x0d, 0x7c,
	0x1d, 0x12, 0x03, 0x36, 0x0a, 0xa5, 0xe0, 0x9d, 0x55, 0x0c, 0x72, 0xa6, 0x94, 0x9d, 0x8a, 0x3b,
	0x72, 0x4e, 0x2d, 0x3b, 0x0d, 0xef, 0xca, 0x39, 0xad, 0xec, 0x74, 0x44, 0x39, 0xa7, 0x97, 0x9d,
	0x81, 0xef, 0xc9, 0x39, 0xa3, 0xec, 0x4c, 0xdc, 0x95, 0x73, 0x66, 0xd9, 0x59, 0xf8, 0xbe, 0x9c,
	0xb3, 0xca, 0xce, 0xc6, 0x3d, 0x39, 0x67, 0x13, 0x0b, 0x36, 0x0b, 0xd7, 0xeb, 0xe2, 0xbd, 0x55,
	0x70, 0x23, 0x87, 0xbd, 0x2e, 0xe9, 0x40, 0xe3, 0xdc, 0x89, 0x4e, 0x42, 0xea, 0xe1, 0xfd, 0x55,
	0xa8, 0x7e, 0xee, 0x44, 0xaf, 0xa8, 0x47, 0xba, 0xd0, 0x4c, 0xe3, 0x15, 0x7c, 0xb0, 0x0a, 0x34,
	0x12, 0xa0, 0x08, 0x42, 0xc5, 0x7d, 0x09, 0xa1, 0x0a, 0x42, 0xc3, 0x87, 0x12, 0x42, 0x13, 0x84,
	0x8e, 0x6d, 0x09, 0xa1, 0x0b, 0xc2, 0xc0, 0x0f, 0x24, 0x84, 0x21, 0x08, 0x13, 0x3f, 0x94, 0x10,
	0xa6, 0x20, 0x2c, 0x7c, 0x24, 0x21, 0x2c, 0x41, 0xd8, 0xf8, 0x58, 0x42, 0xd8, 0x44, 0x81, 0x56,
	0x2a, 0x7a, 0x5d, 0x7c, 0xb2, 0x8a, 0x34, 0x13, 0xd2, 0xeb, 0x66, 0x55, 0x26, 0x0e, 0x5d, 0xe0,
	0xd3, 0xb7, 0x55, 0x79, 0xe9, 0xd0, 0x45, 0x56, 0x85, 0x09, 0x05, 0x3f, 0x7a, 0x5b, 0x15, 0x46,
	0x14, 0xd1, 0xa8, 0xf8, 0xb1, 0x8c, 0x51, 0x45, 0xa3, 0xe1, 0x33, 0x19, 0xa3, 0x89, 0x46, 0xc7,
	0x4f, 0x64, 0x8c, 0x2e, 0x1a, 0x03, 0x3f, 0x95, 0x31, 0x86, 0x68, 0x4c, 0xec, 0xc8, 0x18, 0x53,
	0x34, 0x16, 0x1e, 0xc8, 0x18, 0x4b, 0x34, 0x36, 0x76, 0x65, 0x8c, 0xcd, 0x7e, 0xf5, 0x33, 0xd3,
	0xeb, 0x62, 0x6f, 0xe5, 0xaf, 0x7e, 0x8a, 0x7a, 0x5d, 0x76, 0x88, 0xa0, 0x61, 0x7e, 0xc4, 0x50,
	0x92, 0xf3, 0x4f, 0x31, 0x43, 0x34, 0xd8, 0x1c, 0x5f, 0x06, 0x27, 0x8e, 0xeb, 0x9e, 0xf0, 0x83,
	0x81, 0xca, 0xf3, 0x92, 0x34, 0xaf, 0x70, 0xa4, 0x19, 0xc2, 0xf8, 0x32, 0x38, 0x72, 0x5d, 0x7e,
	0xbc, 0x31, 0xe0, 0x5d, 0x51, 0xf5, 0x50, 0x5b, 0xca, 0x36, 0x0a, 0xd6, 0xab, 0x3a, 0x05, 0x75,
	0x19, 0xa7, 0x54, 0x9d, 0x8a, 0x86, 0x8c, 0x53, 0xab, 0x4e, 0x43, 0x53, 0xc6, 0x69, 0x55, 0xa7,
	0xa3, 0x25, 0xe3, 0xf4, 0xaa, 0x33, 0xd0, 0x96, 0x71, 0x46, 0xd5, 0x99, 0x78, 0x28, 0xe3, 0xcc,
	0xaa, 0xb3, 0xf0, 0x33, 0x19, 0x67, 0x55, 0x9d, 0x8d, 0x9f, 0xcb, 0x38, 0x9b, 0x74, 0xa1, 0xc5,
	0x5c, 0xcc, 0xce, 0x9b, 0xf8, 0x05, 0x37, 0xb7, 0xcb, 0x86, 0x1f, 0x45, 0x87, 0xcd, 0xf1, 0x65,
	0x90, 0x1c, 0x4a, 0x15, 0x80, 0x5c, 0x28, 0xf8, 0xe5, 0x72, 0xd2, 0xca, 0x88, 0x52, 0x32, 0x2a,
	0x7e, 0x25, 0x61, 0xd4, 0x92, 0xd1, 0xf0, 0x6b, 0x09, 0xa3, 0x95, 0x8c, 0x8e, 0x47, 0x12, 0x46,
	0x2f, 0x19, 0x03, 0xbf, 0x91, 0x30, 0x46, 0xc9, 0x98, 0xf8, 0xad, 0x84, 0x31, 0x4b, 0xc6, 0xc2,
	0xef, 0x24, 0x8c, 0x55, 0x32, 0x36, 0x3e, 0x97, 0x30, 0x76, 0xdb, 0xe2, 0xf7, 0xbd, 0x17, 0x7e,
	0x90, 0x1f, 0xf0, 0x1f, 0xc3, 0x56, 0x10, 0xfe, 0xe5, 0xcd, 0xce, 0x9c, 0xc8, 0x7b, 0x4e, 0xe3,
	0x49, 0x94, 0xde, 0xcc, 0x2a, 0xb3, 0xed, 0x67, 0x70, 0xfb, 0x78, 0xd0, 0x67, 0xcf, 0xbf, 0x39,
	0x41, 0x9c, 0xf3, 0x6d, 0x58, 0x7f, 0xc3, 0xc6, 0xd9, 0x75, 0x91, 0x0f, 0x0e, 0x7f, 0x80, 0xb5,
	0x91, 0x1f, 0x78, 0x64, 0xaf, 0x93, 0xdc, 0x66, 0x3b, 0xd9, 0x6d, 0xb6, 0x23, 0x94, 0xc6, 0x7f,
	0x9a, 0xd5, 0x5f, 0x35, 0xe1, 0xd3, 0x21, 0xcf, 0x70, 0xf8, 0x33, 0xd4, 0xcf, 0xc2, 0x20, 0x9e,
	0x50, 0x72, 0xef, 0x9a, 0x5c, 0xc5, 0xcd, 0x32, 0x4b, 0xb6, 0x23, 0x26, 0x2b, 0x3e, 0x1e, 0xa6,
	0x59, 0x0e, 0x7f, 0x49, 0xaf, 0xb7, 0xe4, 0xc1, 0x95, 0x74, 0xe5, 0x9b, 0x4f, 0x96, 0x10, 0x8b,
	0x84, 0xe5, 0x80, 0xf4, 0x6a, 0x7c, 0xf8, 0x3b, 0x34, 0x3c, 0x1a, 0x4f, 0xde, 0x38, 0x01, 0x79,
	0x78, 0x25, 0x67, 0xf5, 0xfb, 0xca, 0xb2, 0xee, 0x16, 0x59, 0xab, 0x21, 0xc3, 0x2c, 0xdd, 0x69,
	0x9d, 0xc7, 0xa8, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x93, 0x40, 0xcb, 0x42, 0x2d, 0x10, 0x00,
	0x00,
}
