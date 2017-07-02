package types

import (
	"github.com/golang/protobuf/proto"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
)

// GraphQLType is an additional type to add to the schema
type GraphQLType struct {
	Name string
	Type string
}

// AdditionalGraphQLTypes returns an array of GraphQLType if specified for a table or nil if not found
func (e Entity) AdditionalGraphQLTypes() []GraphQLType {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), eproto.E_Table)
		if s, ok := ex.(*eproto.SQLMessageOptions); ok {
			a := make([]GraphQLType, 0)
			if s.AddType != nil {
				a = append(a, GraphQLType{
					Name: s.AddType.Name,
					Type: s.AddType.Type,
				})
			}
			if s.AddType2 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType2.Name,
					Type: s.AddType2.Type,
				})
			}
			if s.AddType3 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType3.Name,
					Type: s.AddType3.Type,
				})
			}
			if s.AddType4 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType4.Name,
					Type: s.AddType4.Type,
				})
			}
			if s.AddType5 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType5.Name,
					Type: s.AddType5.Type,
				})
			}
			if s.AddType6 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType6.Name,
					Type: s.AddType6.Type,
				})
			}
			if s.AddType7 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType7.Name,
					Type: s.AddType7.Type,
				})
			}
			if s.AddType8 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType8.Name,
					Type: s.AddType8.Type,
				})
			}
			if s.AddType9 != nil {
				a = append(a, GraphQLType{
					Name: s.AddType9.Name,
					Type: s.AddType9.Type,
				})
			}
			return a
		}
	}
	return nil
}
