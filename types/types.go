package types

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Field describes a specific message field
type Field struct {
	Descriptor *descriptor.FieldDescriptorProto
	Name       string
	Type       string
	Comment    string
}

// Message describes a specific message
type Message struct {
	Descriptor *descriptor.DescriptorProto
	Name       string
	Comment    string
	Fields     []*Field
}

// File describes a file (proto) to be generated
type File struct {
	Descriptor *descriptor.FileDescriptorProto
	Package    string
	Name       string
	Messages   []*Message
}

// Property is a property of an entity such as column in SQL or a property type in GraphQL
type Property struct {
	Field      Field
	Name       string
	Comment    string
	Entity     Entity
	Unique     bool
	Nullable   bool
	Index      bool
	PrimaryKey bool
}

// Entity is a conversion of a Message into a representation of a entity
type Entity struct {
	File       File
	Message    Message
	Package    string
	Name       string
	Comment    string
	Properties []Property
}

// Generation is what was generated from a Generator Generate
type Generation struct {
	Filename string
	Output   string
}

// Generator is an interface for generation of code
type Generator interface {
	// Generate output for a given file
	Generate(scheme string, file *File) ([]*Generation, error)
}

// Generator2 is an interface for generation of code
type Generator2 interface {
	// Generate output for a given file
	Generate(scheme string, file *File, entities []Entity) ([]*Generation, error)
}

// Pad will pad right a string up to ml size
func Pad(s string, ml int) string {
	for i := len(s); i < ml; i++ {
		s += " "
	}
	return s
}
