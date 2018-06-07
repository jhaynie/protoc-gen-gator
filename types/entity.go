package types

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
	"github.com/serenize/snaker"
)

func getExtensionString(p proto.Message, e *proto.ExtensionDesc, key string) string {
	if p != nil {
		e, _ := proto.GetExtension(p, e)
		if ev, ok := e.(*eproto.SQLFieldOptions); ok {
			rv := reflect.ValueOf(*ev)
			if rv.IsValid() {
				f := rv.FieldByName(key)
				if f.IsValid() {
					return f.String()
				}
			}
		}
		if ev, ok := e.(*eproto.SQLMessageOptions); ok {
			rv := reflect.ValueOf(*ev)
			if rv.IsValid() {
				f := rv.FieldByName(key)
				if f.IsValid() {
					return f.String()
				}
			}
		}
	}
	return ""
}

// TableNameSingular returns the singular version of the name title cased
func (e Entity) TableNameSingular() string {
	return strings.Title(e.Name)
}

func plural(name string) string {
	if strings.HasSuffix(name, "s") {
		return name + "es"
	}
	if strings.HasSuffix(name, "y") && !strings.HasSuffix(name, "ey") {
		return name[0:len(name)-1] + "ies"
	}
	return name + "s"
}

// TableNamePlural returns an exteremly simplistic but good enough for now name
func (e Entity) TableNamePlural() string {
	return plural(strings.Title(e.Name))
}

// HasPrimaryKey returns true if the table has a primary key
func (e Entity) HasPrimaryKey() bool {
	return e.PrimaryKey() != ""
}

// PrimaryKey returns the primary key field name
func (e Entity) PrimaryKey() string {
	p := e.PrimaryKeyProperty()
	if p != nil {
		return Backtick(snaker.CamelToSnake(p.Name))
	}
	return ""
}

// PrimaryKeyProperty returns the primary key property or nil if not found
func (e Entity) PrimaryKeyProperty() *Property {
	for _, p := range e.Properties {
		if p.PrimaryKey {
			return &p
		}
	}
	return nil
}

// Cond is a simple function used in templates
func Cond(i, l int, v string) string {
	if i+1 < l {
		return v
	}
	return ""
}

// Backtick returns a string with backticks surrounding it
func Backtick(s string) string {
	return "`" + s + "`"
}

// BacktickArray returns a string with backticks surrounding it from an array of strings
func BacktickArray(s string) string {
	tok := strings.Split(s, ",")
	newarr := make([]string, 0)
	for _, s := range tok {
		newarr = append(newarr, Backtick(strings.TrimSpace(s)))
	}
	return strings.Join(newarr, ",")
}

// Add is a simple addition function for templates
func Add(a, b int) int {
	return a + b
}

// GenerateCode will generate code for this Entity using tmplcode
func GenerateCode(tmplcode string, state map[string]interface{}, funcs map[string]interface{}) ([]byte, error) {
	tpl := template.New("tmpl")
	ctx := make(map[string]interface{})
	fm := template.FuncMap{
		"add":       Add,
		"pad":       Pad,
		"cond":      Cond,
		"tick":      Backtick,
		"tickarray": BacktickArray,
		"addctx": func(key string, b int) string {
			v := ctx[key]
			if value, ok := v.(int); ok {
				ctx[key] = value + b
			} else {
				ctx[key] = b
			}
			return ""
		},
		"hasctx": func(key string) bool {
			v := ctx[key]
			if v == nil {
				return false
			}
			return true
		},
		"condctx": func(key string, l int, r string) string {
			v := ctx[key]
			if v == nil {
				return r
			}
			if value, ok := v.(int); ok {
				if value+1 < l {
					return r
				}
			}
			return ""
		},
		"rmctx": func(key string) string {
			delete(ctx, key)
			return ""
		},
		"lowerfc": func(key string) string {
			return strings.ToLower(key[0:1]) + key[1:]
		},
		"upcase": func(key string) string {
			return strings.ToUpper(key)
		},
		"title": func(key string) string {
			return strings.Title(key)
		},
		"snake": func(key string) string {
			return snaker.CamelToSnake(key)
		},
		"camel": func(key string) string {
			return snaker.SnakeToCamel(key)
		},
		"singular": func(key string) string {
			if strings.HasSuffix(key, "es") {
				return key[0 : len(key)-2]
			}
			if strings.HasSuffix(key, "s") {
				return key[0 : len(key)-1]
			}
			return key
		},
		"plural": func(key string) string {
			return plural(key)
		},
	}
	if funcs != nil {
		for k, v := range funcs {
			fm[k] = v
		}
	}
	tpl.Funcs(fm)
	tpl, err := tpl.Parse(tmplcode)
	if err != nil {
		return nil, err
	}
	if state == nil {
		state = make(map[string]interface{})
	}
	var w bytes.Buffer
	err = tpl.Execute(&w, state)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// GenerateCode will generate code for this Entity using tmplcode
func (e Entity) GenerateCode(tmplcode string, state map[string]interface{}, funcs map[string]interface{}) ([]byte, error) {
	if state == nil {
		state = make(map[string]interface{})
	}
	state["Entity"] = e
	return GenerateCode(tmplcode, state, funcs)
}

// ColumnWidth returns the maximum length of all the column names
func (e Entity) ColumnWidth() int {
	var max int
	for _, p := range e.Properties {
		l := len(p.Name)
		if l > max {
			max = l
		}
	}
	// space + backticks for column names
	return max + 4
}

// ColumnCount returns the number of columns in the entity
func (e Entity) ColumnCount() int {
	return len(e.Properties)
}

// Checksum returns the checksum property name
func (e Entity) Checksum() string {
	for _, p := range e.Properties {
		if p.IsChecksum() {
			return p.Name
		}
	}
	return ""
}

// HasChecksum returns true if any property is a Checksum type
func (e Entity) HasChecksum() bool {
	for _, p := range e.Properties {
		if p.IsChecksum() {
			return true
		}
	}
	return false
}

var entities = make([]*Entity, 0)

func findEntityByName(name string) (*Entity, error) {
	for _, e := range entities {
		if e.Name == name || e.SQLTableName() == name {
			return e, nil
		}
	}
	return nil, fmt.Errorf("couldn't find " + name)
}

// NewEntity converts a message into an Entity
func NewEntity(packageName string, file *File, message *Message) Entity {
	e := Entity{}
	e.File = *file
	e.Message = *message
	e.Package = packageName
	e.Name = message.Name
	e.Comment = message.Comment
	e.Properties = make([]Property, 0)
	for _, field := range message.Fields {
		p := NewProperty(&e, field)
		e.Properties = append(e.Properties, *p)
	}
	e.SortedProperties = make([]Property, len(e.Properties))
	copy(e.SortedProperties, e.Properties)
	sort.Slice(e.SortedProperties, func(i, j int) bool {
		// sort by json tag since an array of interfaces will then sort correctly even when untyped
		a := snaker.CamelToSnake(e.SortedProperties[i].Name)
		b := snaker.CamelToSnake(e.SortedProperties[j].Name)
		return a < b
	})
	entities = append(entities, &e)
	return e
}
