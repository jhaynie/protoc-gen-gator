package types

import (
	"strings"

	"github.com/golang/protobuf/proto"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
	"github.com/serenize/snaker"
)

// GraphQLType is an additional type to add to the schema
type GraphQLType struct {
	Name       string
	Type       string
	Generate   bool
	Definition string
}

// GraphUnionType is a generated union type
type GraphUnionType struct {
	Name    string
	ID      string
	Type    string
	Union   string
	Mapping map[string]string
	Tables  []string
}

// Definition will return the GraphQL definition
func (u GraphUnionType) Definition() string {
	return "union " + u.Union + " = " + strings.Join(u.Tables, " | ")
}

// AdditionalGraphQLUnions returns an array of GraphUnionType
func (e Entity) AdditionalGraphQLUnions() []GraphUnionType {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), eproto.E_Table)
		if s, ok := ex.(*eproto.SQLMessageOptions); ok {
			a := make([]GraphUnionType, 0)
			if s.GqlUnion != nil {
				a = append(a, e.makeUnion(s.GqlUnion))
			}
			return a
		}
	}
	return nil
}

func (e Entity) makeUnion(u *eproto.GraphQLUnion) GraphUnionType {
	typeName := snaker.SnakeToCamel(e.SQLTableName()) + snaker.SnakeToCamel(u.Name) + "Table"
	union := GraphUnionType{
		Name:    u.Name,
		ID:      u.Id,
		Type:    u.Type,
		Union:   typeName,
		Tables:  make([]string, 0),
		Mapping: make(map[string]string),
	}
	tok := strings.Split(u.Tables, ",")
	for _, line := range tok {
		kv := strings.Split(strings.TrimSpace(line), ":")
		table := snaker.SnakeToCamel(kv[1])
		union.Tables = append(union.Tables, table)
		union.Mapping[kv[0]] = table
	}
	return union
}

// AdditionalGraphQLTypes returns an array of GraphQLType if specified for a table or nil if not found
func (e Entity) AdditionalGraphQLTypes() []GraphQLType {
	if e.Message.Descriptor.Options != nil {
		ex, _ := proto.GetExtension(e.Message.Descriptor.GetOptions(), eproto.E_Table)
		if s, ok := ex.(*eproto.SQLMessageOptions); ok {
			a := make([]GraphQLType, 0)
			if s.GqlUnion != nil {
				u := e.makeUnion(s.GqlUnion)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion2 != nil {
				u := e.makeUnion(s.GqlUnion2)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion3 != nil {
				u := e.makeUnion(s.GqlUnion3)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion4 != nil {
				u := e.makeUnion(s.GqlUnion4)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion5 != nil {
				u := e.makeUnion(s.GqlUnion5)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion6 != nil {
				u := e.makeUnion(s.GqlUnion6)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion7 != nil {
				u := e.makeUnion(s.GqlUnion7)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion8 != nil {
				u := e.makeUnion(s.GqlUnion8)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlUnion9 != nil {
				u := e.makeUnion(s.GqlUnion9)
				a = append(a, GraphQLType{
					Name:       u.Name,
					Type:       u.Union,
					Generate:   true,
					Definition: u.Definition(),
				})
			}
			if s.GqlAddType != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType.Name,
					Type:       s.GqlAddType.Type,
					Generate:   s.GqlAddType.Generate,
					Definition: s.GqlAddType.Definition,
				})
			}
			if s.GqlAddType2 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType2.Name,
					Type:       s.GqlAddType2.Type,
					Generate:   s.GqlAddType2.Generate,
					Definition: s.GqlAddType2.Definition,
				})
			}
			if s.GqlAddType3 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType3.Name,
					Type:       s.GqlAddType3.Type,
					Generate:   s.GqlAddType3.Generate,
					Definition: s.GqlAddType3.Definition,
				})
			}
			if s.GqlAddType4 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType4.Name,
					Type:       s.GqlAddType4.Type,
					Generate:   s.GqlAddType4.Generate,
					Definition: s.GqlAddType4.Definition,
				})
			}
			if s.GqlAddType5 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType5.Name,
					Type:       s.GqlAddType5.Type,
					Generate:   s.GqlAddType5.Generate,
					Definition: s.GqlAddType5.Definition,
				})
			}
			if s.GqlAddType6 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType6.Name,
					Type:       s.GqlAddType6.Type,
					Generate:   s.GqlAddType6.Generate,
					Definition: s.GqlAddType6.Definition,
				})
			}
			if s.GqlAddType7 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType7.Name,
					Type:       s.GqlAddType7.Type,
					Generate:   s.GqlAddType7.Generate,
					Definition: s.GqlAddType7.Definition,
				})
			}
			if s.GqlAddType8 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType8.Name,
					Type:       s.GqlAddType8.Type,
					Generate:   s.GqlAddType8.Generate,
					Definition: s.GqlAddType8.Definition,
				})
			}
			if s.GqlAddType9 != nil {
				a = append(a, GraphQLType{
					Name:       s.GqlAddType9.Name,
					Type:       s.GqlAddType9.Type,
					Generate:   s.GqlAddType9.Generate,
					Definition: s.GqlAddType9.Definition,
				})
			}
			return a
		}
	}
	return nil
}
