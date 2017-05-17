package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhaynie/protoc-gen-gator/generator"
	_ "github.com/jhaynie/protoc-gen-gator/generators"
	"github.com/jhaynie/protoc-gen-gator/types"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	log.SetFlags(255)

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	req := new(plugin.CodeGeneratorRequest)
	err = proto.Unmarshal(b, req)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeFiles(os.Stdout, req)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeFiles(w io.Writer, req *plugin.CodeGeneratorRequest) error {
	files, err := getFiles(req)
	if err != nil {
		return err
	}

	var gentypes []string

	if req.GetParameter() == "" {
		gentypes = generator.GetAllTypes()

	} else {
		gentypes = strings.Split(req.GetParameter(), ",")
	}

	resp := &plugin.CodeGeneratorResponse{
		File: make([]*plugin.CodeGeneratorResponse_File, 0),
	}

	for _, file := range files {
		results, err := generator.Generate(gentypes, file)
		if err != nil {
			return err
		}
		for _, result := range results {
			resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
				Name:    proto.String(result.Filename),
				Content: proto.String(result.Output),
			})
		}
	}

	b, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func extractComments(file *descriptor.FileDescriptorProto) map[string]*descriptor.SourceCodeInfo_Location {
	comments := make(map[string]*descriptor.SourceCodeInfo_Location)
	for _, loc := range file.GetSourceCodeInfo().GetLocation() {
		if loc.LeadingComments == nil {
			continue
		}
		var p []string
		for _, n := range loc.Path {
			p = append(p, strconv.Itoa(int(n)))
		}
		comments[strings.Join(p, ",")] = loc
	}
	return comments
}

func getMessagePath(messageIndex int) string {
	return fmt.Sprintf("4,%d", messageIndex)
}

func getFieldPath(messageIndex int, fieldIndex int) string {
	return fmt.Sprintf("4,%d,2,%d", messageIndex, fieldIndex)
}

func getGoType(field *descriptor.FieldDescriptorProto) string {
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		{
			return "float64"
		}
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		{
			return "float64"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		{
			return "int64"
		}
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		{
			return "uint64"
		}
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		{
			return "int32"
		}
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		{
			return "uint64"
		}
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		{
			return "uint32"
		}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		{
			return "bool"
		}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		{
			return "string"
		}
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		{
			//?
		}
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		{
			//?
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		{
			return "[]byte"
		}
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		{
			return "uint32"
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		{
			//TODO
		}
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		{
			return "int32"
		}
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		{
			return "int64"
		}
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		{
			return "int32"
		}
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		{
			return "int64"
		}
	}
	return "string"
}

func getFiles(req *plugin.CodeGeneratorRequest) ([]*types.File, error) {
	filesToGenerate := make(map[string]bool)

	for _, fn := range req.GetFileToGenerate() {
		filesToGenerate[fn] = true
	}

	files := make([]*types.File, 0)

	for _, pf := range req.GetProtoFile() {
		if filesToGenerate[pf.GetName()] {

			file := &types.File{
				Descriptor: pf,
				Name:       strings.Replace(pf.GetName(), ".proto", "", 1),
				Package:    pkgName(pf),
				Messages:   make([]*types.Message, 0),
			}

			files = append(files, file)

			comments := extractComments(pf)

			// if pf.Options != nil {
			// 	a, _ := proto.GetExtension(pf.GetOptions(), eproto.E_ActionPrefix)
			// 	os.Stderr.WriteString(fmt.Sprintf("opt: %s", spew.Sdump(a)))
			// }

			for i, mt := range pf.GetMessageType() {
				message := &types.Message{
					Name:       mt.GetName(),
					Descriptor: mt,
					Comment:    comments[getMessagePath(i)].GetLeadingComments(),
					Fields:     make([]*types.Field, 0),
				}
				file.Messages = append(file.Messages, message)
				for j, ft := range mt.GetField() {
					field := &types.Field{
						Name:       ft.GetName(),
						Type:       getGoType(ft),
						Descriptor: ft,
						Comment:    comments[getFieldPath(i, j)].GetLeadingComments(),
					}
					message.Fields = append(message.Fields, field)
				}
			}
		}
	}

	return files, nil
}

// pkgName returns a suitable package name from file.
//
// Mostly borrowed from grpc-gateway.
func pkgName(file *descriptor.FileDescriptorProto) string {
	if file.Options != nil && file.Options.GoPackage != nil {
		gopkg := file.Options.GetGoPackage()
		i := strings.LastIndexByte(gopkg, '/')
		if i < 0 {
			return gopkg
		}
		return strings.Replace(gopkg[i+1:], ".", "_", -1)
	}
	if file.Package == nil {
		base := filepath.Base(file.GetName())
		ext := filepath.Ext(base)
		return strings.TrimSuffix(base, ext)
	}
	return strings.Replace(file.GetPackage(), ".", "_", -1)
}
