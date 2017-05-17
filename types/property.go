package types

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
	"github.com/serenize/snaker"
)

// IsBoolExtension returns true if the extension provided is defined on the property
func (p Property) IsBoolExtension(e *proto.ExtensionDesc, key string) bool {
	if p.Field.Descriptor.Options != nil {
		e, _ := proto.GetExtension(p.Field.Descriptor.GetOptions(), e)
		if s, ok := e.(*eproto.SQLFieldOptions); ok {
			rv := reflect.ValueOf(*s)
			if rv.IsValid() {
				f := rv.FieldByName(key)
				if f.IsValid() {
					return f.Bool()
				}
			}
		}
	}
	return false
}

// IsChecksum returns true if the field type is .proto.Checksum
func (p Property) IsChecksum() bool {
	return p.Field.Descriptor.GetTypeName() == ".proto.Checksum"
}

// IsEnumeration returns true if the field type is an enumeration
func (p Property) IsEnumeration() bool {
	return p.Field.Descriptor.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM
}

// NewProperty will convert a Field into a Property
func NewProperty(e *Entity, f *Field) *Property {
	p := Property{}
	p.Entity = *e
	p.Field = *f
	f.Name = snaker.SnakeToCamel(f.Name)
	p.Field.Name = f.Name
	p.Name = f.Name
	p.Comment = f.Comment
	p.PrimaryKey = p.IsBoolExtension(eproto.E_Column, "Primarykey")
	p.Nullable = !p.IsBoolExtension(eproto.E_Column, "Notnull")
	p.Index = p.IsBoolExtension(eproto.E_Column, "Index")
	p.Unique = p.IsBoolExtension(eproto.E_Column, "Unique")
	if p.IsSQLIDColumn() || p.PrimaryKey {
		p.Nullable = false
		p.PrimaryKey = true
	}
	return &p
}
