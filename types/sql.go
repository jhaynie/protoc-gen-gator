package types

import (
	"os"
	"reflect"
	"strings"

	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
	"github.com/serenize/snaker"
)

// SQLIndex describes a SQL index
type SQLIndex struct {
	Type   string
	Name   string
	Fields string
}

// SQLAssociationType is a constant defining the type of association
type SQLAssociationType int

const (
	SQLAssocationBelongsTo SQLAssociationType = iota
	SQLAssocationHasOne
	SQLAssocationHasMany
)

// SQLAssociation describes details about a specific association
type SQLAssociation struct {
	Type       SQLAssociationType
	Table      string
	PrimaryKey string
	ForeignKey string
	Name       string
}

// IsMultiKey will return true if the primary key is a multi-key reference
func (a SQLAssociation) IsMultiKey() bool {
	return strings.Count(a.PrimaryKey, ",") > 0
}

// IsArrayType returns true if the association returns an array
func (a SQLAssociation) IsArrayType() bool {
	return a.Type == SQLAssocationHasMany
}

// GenerateSQL returns true if the SQL DDL should be generated for this entity
func (e Entity) GenerateSQL() bool {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), eproto.E_Table)
		if s, ok := ex.(*eproto.SQLMessageOptions); ok {
			return !s.GetNogenerate()
		}
	}
	return true
}

// HasSQLAssociations returns true if there are one or more associations
func (e Entity) HasSQLAssociations() bool {
	a := e.SQLAssociations()
	return len(a) > 0
}

// SQLAssociationsUnique returns the name of each association
func (e Entity) SQLAssociationsUnique() []string {
	a := e.SQLAssociations()
	if len(a) > 0 {
		h := make(map[string]bool)
		s := make([]string, 0)
		for _, ua := range a {
			h[snaker.SnakeToCamel(ua.Table)] = true
		}
		for k := range h {
			s = append(s, k)
		}
		return s
	}
	return []string{}
}

// SQLAssociations returns the associations (if any) for the table or nil if none are defined
func (e Entity) SQLAssociations() []SQLAssociation {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), eproto.E_Table)
		if s, ok := ex.(*eproto.SQLMessageOptions); ok {
			a := make([]SQLAssociation, 0)
			if s.BelongsTo != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo.Table,
					Name:       s.BelongsTo.Name,
					PrimaryKey: s.BelongsTo.Pk,
					ForeignKey: s.BelongsTo.Fk,
				})
			}
			if s.BelongsTo2 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo2.Table,
					Name:       s.BelongsTo2.Name,
					PrimaryKey: s.BelongsTo2.Pk,
					ForeignKey: s.BelongsTo2.Fk,
				})
			}
			if s.BelongsTo3 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo3.Table,
					Name:       s.BelongsTo3.Name,
					PrimaryKey: s.BelongsTo3.Pk,
					ForeignKey: s.BelongsTo3.Fk,
				})
			}
			if s.BelongsTo4 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo4.Table,
					Name:       s.BelongsTo4.Name,
					PrimaryKey: s.BelongsTo4.Pk,
					ForeignKey: s.BelongsTo4.Fk,
				})
			}
			if s.BelongsTo5 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo5.Table,
					Name:       s.BelongsTo5.Name,
					PrimaryKey: s.BelongsTo5.Pk,
					ForeignKey: s.BelongsTo5.Fk,
				})
			}
			if s.BelongsTo6 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo6.Table,
					Name:       s.BelongsTo6.Name,
					PrimaryKey: s.BelongsTo6.Pk,
					ForeignKey: s.BelongsTo6.Fk,
				})
			}
			if s.BelongsTo7 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo7.Table,
					Name:       s.BelongsTo7.Name,
					PrimaryKey: s.BelongsTo7.Pk,
					ForeignKey: s.BelongsTo7.Fk,
				})
			}
			if s.BelongsTo8 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo8.Table,
					Name:       s.BelongsTo8.Name,
					PrimaryKey: s.BelongsTo8.Pk,
					ForeignKey: s.BelongsTo8.Fk,
				})
			}
			if s.BelongsTo5 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo5.Table,
					Name:       s.BelongsTo5.Name,
					PrimaryKey: s.BelongsTo5.Pk,
					ForeignKey: s.BelongsTo5.Fk,
				})
			}
			if s.BelongsTo9 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo9.Table,
					Name:       s.BelongsTo9.Name,
					PrimaryKey: s.BelongsTo9.Pk,
					ForeignKey: s.BelongsTo9.Fk,
				})
			}
			if s.BelongsTo10 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationBelongsTo,
					Table:      s.BelongsTo10.Table,
					Name:       s.BelongsTo10.Name,
					PrimaryKey: s.BelongsTo10.Pk,
					ForeignKey: s.BelongsTo10.Fk,
				})
			}
			if s.HasMany != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany.Table,
					Name:       s.HasMany.Name,
					PrimaryKey: s.HasMany.Pk,
					ForeignKey: s.HasMany.Fk,
				})
			}
			if s.HasMany2 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany2.Table,
					Name:       s.HasMany2.Name,
					PrimaryKey: s.HasMany2.Pk,
					ForeignKey: s.HasMany2.Fk,
				})
			}
			if s.HasMany3 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany3.Table,
					Name:       s.HasMany3.Name,
					PrimaryKey: s.HasMany3.Pk,
					ForeignKey: s.HasMany3.Fk,
				})
			}
			if s.HasMany4 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany4.Table,
					Name:       s.HasMany4.Name,
					PrimaryKey: s.HasMany4.Pk,
					ForeignKey: s.HasMany4.Fk,
				})
			}
			if s.HasMany5 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany5.Table,
					Name:       s.HasMany5.Name,
					PrimaryKey: s.HasMany5.Pk,
					ForeignKey: s.HasMany5.Fk,
				})
			}
			if s.HasMany6 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany6.Table,
					Name:       s.HasMany6.Name,
					PrimaryKey: s.HasMany6.Pk,
					ForeignKey: s.HasMany6.Fk,
				})
			}
			if s.HasMany7 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany7.Table,
					Name:       s.HasMany7.Name,
					PrimaryKey: s.HasMany7.Pk,
					ForeignKey: s.HasMany7.Fk,
				})
			}
			if s.HasMany8 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany8.Table,
					Name:       s.HasMany8.Name,
					PrimaryKey: s.HasMany8.Pk,
					ForeignKey: s.HasMany8.Fk,
				})
			}
			if s.HasMany9 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany9.Table,
					Name:       s.HasMany9.Name,
					PrimaryKey: s.HasMany9.Pk,
					ForeignKey: s.HasMany9.Fk,
				})
			}
			if s.HasMany10 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasMany,
					Table:      s.HasMany10.Table,
					Name:       s.HasMany10.Name,
					PrimaryKey: s.HasMany10.Pk,
					ForeignKey: s.HasMany10.Fk,
				})
			}
			if s.HasOne != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne.Table,
					Name:       s.HasOne.Name,
					PrimaryKey: s.HasOne.Pk,
					ForeignKey: s.HasOne.Fk,
				})
			}
			if s.HasOne2 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne2.Table,
					Name:       s.HasOne2.Name,
					PrimaryKey: s.HasOne2.Pk,
					ForeignKey: s.HasOne2.Fk,
				})
			}
			if s.HasOne3 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne3.Table,
					Name:       s.HasOne3.Name,
					PrimaryKey: s.HasOne3.Pk,
					ForeignKey: s.HasOne3.Fk,
				})
			}
			if s.HasOne4 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne4.Table,
					Name:       s.HasOne4.Name,
					PrimaryKey: s.HasOne4.Pk,
					ForeignKey: s.HasOne4.Fk,
				})
			}
			if s.HasOne5 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne5.Table,
					Name:       s.HasOne5.Name,
					PrimaryKey: s.HasOne5.Pk,
					ForeignKey: s.HasOne5.Fk,
				})
			}
			if s.HasOne6 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne6.Table,
					Name:       s.HasOne6.Name,
					PrimaryKey: s.HasOne6.Pk,
					ForeignKey: s.HasOne6.Fk,
				})
			}
			if s.HasOne7 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne7.Table,
					Name:       s.HasOne7.Name,
					PrimaryKey: s.HasOne7.Pk,
					ForeignKey: s.HasOne7.Fk,
				})
			}
			if s.HasOne8 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne8.Table,
					Name:       s.HasOne8.Name,
					PrimaryKey: s.HasOne8.Pk,
					ForeignKey: s.HasOne8.Fk,
				})
			}
			if s.HasOne9 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne9.Table,
					Name:       s.HasOne9.Name,
					PrimaryKey: s.HasOne9.Pk,
					ForeignKey: s.HasOne9.Fk,
				})
			}
			if s.HasOne10 != nil {
				a = append(a, SQLAssociation{
					Type:       SQLAssocationHasOne,
					Table:      s.HasOne10.Table,
					Name:       s.HasOne10.Name,
					PrimaryKey: s.HasOne10.Pk,
					ForeignKey: s.HasOne10.Fk,
				})
			}
			return a
		}
	}
	return nil
}

// IsBoolExtension returns true if the extension provided is defined on the file or message
func (e Entity) IsBoolExtension(d *proto.ExtensionDesc, key string) bool {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), d)
		if s, ok := ex.(*eproto.SQLFieldOptions); ok {
			rv := reflect.ValueOf(*s)
			if rv.IsValid() {
				f := rv.FieldByName(key)
				if f.IsValid() {
					return f.Bool()
				}
			}
		}
	}
	if e.File.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.File.Descriptor.GetOptions(), d)
		if s, ok := ex.(*eproto.SQLFileOptions); ok {
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

func generateIndex(message Message, e *proto.ExtensionDesc, key string) *SQLIndex {
	ext, _ := proto.GetExtension(message.Descriptor.GetOptions(), e)
	if ev, ok := ext.(*eproto.SQLMessageOptions); ok {
		rv := reflect.ValueOf(*ev)
		if rv.IsValid() {
			f := rv.FieldByName(key)
			if f.IsValid() {
				intf := f.Interface()
				if i, ok := intf.(*eproto.SQLIndex); ok {
					if i == nil {
						return nil
					}
					if i.Name == "" {
						// create the name of the index if not provided
						i.Name = strings.ToLower(snaker.CamelToSnake(message.Name) + "_" + strings.Replace(strings.Replace(i.Fields, ",", "_", -1), " ", "", -1) + "_index")
					}
					if i.Type == "" {
						i.Type = "INDEX"
					}
					return &SQLIndex{i.Type, i.Name, i.Fields}
				}
			}
		}
	}
	return nil
}

// SQLColumnPlaceholders returns a string of placeholder values for the SQL query
func (e Entity) SQLColumnPlaceholders() string {
	return "?" + strings.Repeat(",?", len(e.Properties)-1)
}

// SQLColumnSetterList returns a string of placeholders for setter
func (e Entity) SQLColumnSetterList() string {
	l := make([]string, 0)
	for _, c := range e.Properties {
		if !c.PrimaryKey {
			l = append(l, Backtick(snaker.CamelToSnake(c.Name))+"=?")
		}
	}
	return strings.Join(l, ",")
}

// SQLColumnList returns a comma-separated list of column names
func (e Entity) SQLColumnList() string {
	s := make([]string, 0)
	for _, p := range e.Properties {
		s = append(s, Backtick(e.SQLTableName())+"."+Backtick(snaker.CamelToSnake(p.Name)))
	}
	return strings.Join(s, ",")
}

// SQLColumnUpsertList returns a comma-separated list of upsert column setter
func (e Entity) SQLColumnUpsertList() string {
	l := make([]string, 0)
	for _, c := range e.Properties {
		if !c.PrimaryKey {
			n := Backtick(snaker.CamelToSnake(c.Name))
			l = append(l, n+"=VALUES("+n+")")
		}
	}
	return strings.Join(l, ",")
}

func (e Entity) validateIndex(i *SQLIndex) error {
	for _, tok := range strings.Split(i.Fields, ",") {
		tok = strings.TrimSpace(tok)
		var found bool
		for _, p := range e.Properties {
			if p.SQLColumnName() == tok {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("couldn't find column '%s' for index '%s' on table '%s'", tok, i.Fields, e.SQLTableName())
		}
	}
	return nil
}

// SQLIndexes returns information about any indexes on the entity
func (e Entity) SQLIndexes() []SQLIndex {
	indicies := make([]SQLIndex, 0)
	iprefix := strings.ToLower(e.SQLTableName()) + "_"
	for _, p := range e.Properties {
		if p.Index {
			cn := p.SQLColumnName()
			t := getExtensionString(p.Field.Descriptor.GetOptions(), eproto.E_Column, "Indextype")
			if t == "" {
				t = "INDEX"
			}
			indicies = append(indicies, SQLIndex{Type: t, Name: iprefix + strings.ToLower(cn) + "_index", Fields: cn})
		}
	}
	if e.Message.Descriptor.Options != nil {
		i := generateIndex(e.Message, eproto.E_Table, "Index")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index2")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index3")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index4")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index5")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index6")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index7")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index8")
		if i != nil {
			indicies = append(indicies, *i)
		}
		i = generateIndex(e.Message, eproto.E_Table, "Index9")
		if i != nil {
			indicies = append(indicies, *i)
		}
	}
	for _, index := range indicies {
		if err := e.validateIndex(&index); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return indicies
}

// SQLTableName return the tablename for this entity
func (e Entity) SQLTableName() string {
	if e.Message.Descriptor.Options != nil {
		n := getExtensionString(e.Message.Descriptor.GetOptions(), eproto.E_Table, "Name")
		if n != "" {
			return n
		}
	}
	return snaker.CamelToSnake(e.Name)
}

// IsLowercaseEnums returns true if enum values should be lower cased
func (e Entity) IsLowercaseEnums() bool {
	var lc bool
	if e.File.Descriptor.Options != nil {
		lc = e.IsBoolExtension(eproto.E_File, "LowercaseEnums")
	}
	return lc
}

func findEnumValue(entity *Entity, e *descriptor.EnumValueDescriptorProto) string {
	if e.Options != nil {
		opt, _ := proto.GetExtension(e.Options, eproto.E_Enumval)
		if o, ok := opt.(*eproto.SQLEnumValueOptions); ok {
			if o != nil {
				return o.GetValue()
			}
		}
	}
	ev := e.GetName()
	if entity.IsLowercaseEnums() {
		ev = strings.ToLower(ev)
	}
	return ev
}

// SQLEnum returns the SQL type for Enumeration
func (p Property) SQLEnum() string {
	c := strings.Count(p.Field.Descriptor.GetTypeName(), ".")
	switch c {
	case 3:
		{
			i := strings.LastIndex(p.Field.Descriptor.GetTypeName(), ".")
			n := p.Field.Descriptor.GetTypeName()[i+1:]
			for _, e := range p.Entity.Message.Descriptor.GetEnumType() {
				if e.GetName() == n {
					s := make([]string, 0)
					for _, v := range e.GetValue() {
						value := findEnumValue(&p.Entity, v)
						s = append(s, "'"+value+"'")
					}
					return "ENUM(" + strings.Join(s, ",") + ")"
				}
			}
		}
	}
	//??
	return "BINARY"
}

// IsSQLIDColumn returns true if the column is a .proto.ID field
func (p Property) IsSQLIDColumn() bool {
	return p.Field.Descriptor.GetTypeName() == ".proto.ID"
}

// SQLType returns the SQL type for the property
func (p Property) SQLType() string {
	if p.Field.Descriptor.Options != nil {
		t := getExtensionString(p.Field.Descriptor.GetOptions(), eproto.E_Column, "Type")
		if t != "" {
			return t
		}
	}
	switch p.Field.Descriptor.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		{
			return "DOUBLE"
		}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		{
			return "FLOAT"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		{
			return "BIGINT"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		{
			return "INT"
		}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		{
			return "BOOL"
		}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		{
			return "TEXT"
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		{
			return "BLOB"
		}
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		{
			switch p.Field.Descriptor.GetTypeName() {
			case ".proto.ID", ".proto.UID":
				{
					return "VARCHAR(64)"
				}
			case ".proto.Checksum":
				{
					return "CHAR(64)"
				}
			case ".proto.DateTime", ".google.protobuf.Timestamp":
				{
					return "DATETIME"
				}
			default:
				{
					return "JSON"
				}
			}
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		{
			return p.SQLEnum()
		}
	}
	return "BINARY"
}

// SQLColumnName return the column name for this property
func (p Property) SQLColumnName() string {
	if p.Field.Descriptor.Options != nil {
		t := getExtensionString(p.Field.Descriptor.GetOptions(), eproto.E_Column, "Name")
		if t != "" {
			return t
		}
	}
	return snaker.CamelToSnake(p.Name)
}

// SQLColumnNameWithTick returns the SQLColumnName with backticks surrounding it
func (p Property) SQLColumnNameWithTick() string {
	return Backtick(p.SQLColumnName())
}

// SQLColumnTypeWithAttributes returns a string with the column type + extra attributes (if any)
func (p Property) SQLColumnTypeWithAttributes() string {
	a := make([]string, 0)
	if !p.Nullable {
		a = append(a, "NOT NULL")
	}
	if p.PrimaryKey {
		a = append(a, "PRIMARY KEY")
	}
	if p.Unique {
		a = append(a, "UNIQUE")
	}
	if p.Field.Descriptor.Options != nil {
		def := getExtensionString(p.Field.Descriptor.GetOptions(), eproto.E_Column, "Default")
		if def != "" {
			c := p.SQLType()
			if strings.HasPrefix(c, "VARCHAR") || strings.HasSuffix(c, "TEXT") || strings.HasSuffix(c, "CHAR") {
				def = `"` + def + `"`
			}
			a = append(a, "DEFAULT "+def)
		}
	}
	if len(a) > 0 {
		return p.SQLType() + " " + strings.Join(a, " ")
	}
	return p.SQLType()
}
