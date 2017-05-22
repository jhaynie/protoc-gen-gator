package golang

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhaynie/go-gator/orm"
	"github.com/jhaynie/protoc-gen-gator/generator"
	"github.com/jhaynie/protoc-gen-gator/generators/sql"
	eproto "github.com/jhaynie/protoc-gen-gator/proto"
	"github.com/jhaynie/protoc-gen-gator/types"
	"github.com/serenize/snaker"
)

type gogenerator struct {
}

func init() {
	generator.Register2("golang", &gogenerator{})
}

var re = regexp.MustCompile("\\w+\\s*\\((\\d+)\\)")

func toLength(s string) int {
	if re.MatchString(s) {
		m := re.FindStringSubmatch(s)
		i, _ := strconv.Atoi(m[1])
		return i
	}
	return 999999
}

func findEnumValue(entity *types.Entity, e *descriptor.EnumValueDescriptorProto) string {
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

func findEnum(p *types.Property) (string, *descriptor.EnumDescriptorProto) {
	tok := strings.Split(p.Field.Descriptor.GetTypeName(), ".")
	n := tok[3]
	for _, e := range p.Entity.Message.Descriptor.EnumType {
		if e.GetName() == n {
			return n, e
		}
	}
	return "", nil
}

func toTestData(p *types.Property, forUpdate bool) string {
	if isTimestamp(p) {
		return "ToTimestampNow()"
	}
	if p.Nullable {
		return "nil"
	}
	if p.IsEnumeration() {
		n, en := findEnum(p)
		if en != nil {
			return n + "_" + en.Value[0].GetName()
		}
	}
	t := p.Field.Type
	switch t {
	case "string":
		{
			if p.SQLType() == "JSON" {
				return `"{}"`
			}
			// trim to max length of string if provided
			l := toLength(p.SQLType())
			// create a stable value so that subsequent generations will diff the same
			id := orm.HashStrings(p.Entity.Name, p.Name, fmt.Sprintf("%v", forUpdate))
			if len(id) > l {
				id = id[0:l]
			}
			return `"` + id + `"`
		}
	case "int32":
		{
			if forUpdate {
				return "int32(320)"
			}
			return "int32(32)"
		}
	case "int64":
		{
			if forUpdate {
				return "int64(640)"
			}
			return "int64(64)"
		}
	case "bool":
		{
			if forUpdate {
				return "false"
			}
			return "true"
		}
	case "float32":
		{
			if forUpdate {
				return "float32(32.1)"
			}
			return "float32(3.2)"
		}
	case "float64":
		{
			if forUpdate {
				return "float64(64.1)"
			}
			return "float64(6.4)"
		}
	}
	return t
}

func toGoEnumDefinitions(entity types.Entity) string {
	var buf bytes.Buffer
	for _, e := range entity.Message.Descriptor.EnumType {
		names := make([]string, 0)
		n := entity.Name + "_" + strings.Replace(e.GetName(), entity.Name, "", 1)
		for _, v := range e.Value {
			ev := findEnumValue(&entity, v)
			names = append(names, e.GetName()+"_"+v.GetName()+" "+n+" = \""+ev+"\"")
		}
		buf.WriteString("type " + n + " string\n\n")
		buf.WriteString("const (\n\t" + strings.Join(names, "\n\t"))
		buf.WriteString("\n)\n\n")
		buf.WriteString("func (x " + n + ") String() string {\n")
		buf.WriteString("\treturn string(x)\n")
		buf.WriteString("}\n\n")
		buf.WriteString("func enum" + e.GetName() + "ToString(v *" + n + ") string {\n")
		buf.WriteString("\tif v == nil {\n")
		buf.WriteString("\t\treturn \"\"\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn v.String()\n")
		buf.WriteString("}\n\n")
		buf.WriteString("func to" + n + "(v string) *" + n + " {\n")
		ename := strings.Replace(n, "_", "", 1)
		buf.WriteString("var ev *" + n + "\n")
		buf.WriteString("\tswitch v {\n")
		for _, v := range e.Value {
			buf.WriteString("\t\tcase ")
			ev := findEnumValue(&entity, v)
			eval := ename + "_" + v.GetName()
			if ev != v.GetName() && strings.ToLower(ev) != strings.ToLower(v.GetName()) {
				buf.WriteString("\"" + ev + "\", \"" + v.GetName() + "\", \"" + strings.ToLower(v.GetName()) + "\"")
			} else {
				if v.GetName() == strings.ToLower(v.GetName()) {
					buf.WriteString("\"" + v.GetName() + "\", \"" + strings.ToUpper(v.GetName()) + "\"")
				} else {
					buf.WriteString("\"" + v.GetName() + "\", \"" + strings.ToLower(v.GetName()) + "\"")
				}
			}
			buf.WriteString(": {\n")
			buf.WriteString("\t\t\tv := " + eval + "\n")
			buf.WriteString("\t\t\tev = &v\n")
			buf.WriteString("\t\t}\n")
		}
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn ev\n")
		buf.WriteString("}\n")
	}
	return buf.String()
}

func isTimestamp(property *types.Property) bool {
	tn := property.Field.Descriptor.GetTypeName()
	switch tn {
	case ".proto.DateTime", ".google.protobuf.Timestamp":
		{
			return true
		}
	}
	return false
}

func toGoType(property *types.Property) string {
	if isTimestamp(property) {
		return "*timestamp.Timestamp"
	}
	t := property.Field.Type
	if property.IsEnumeration() {
		t = property.Entity.Name + "_" + snaker.SnakeToCamel(property.Field.Descriptor.GetName())
	}
	if property.Nullable {
		return "*" + t
	}
	return t
}

func toGoTypeWithoutPointer(property *types.Property) string {
	if isTimestamp(property) {
		return "*timestamp.Timestamp"
	}
	if property.IsEnumeration() {
		_, en := findEnum(property)
		if en != nil {
			n := strings.Replace(en.GetName(), property.Entity.Name, "", 1)
			return property.Entity.Name + "_" + n
		}
	}
	return property.Field.Type
}

func toGoTags(property *types.Property) string {
	var buf bytes.Buffer
	n := snaker.CamelToSnake(property.Name)
	buf.WriteString("`json:\"")
	buf.WriteString(n)
	if property.Nullable {
		buf.WriteString(",omitempty")
	}
	buf.WriteString("\"")
	buf.WriteString("`")
	return buf.String()
}

func toGetterValue(property *types.Property, varname string, stringer bool) string {
	var buf bytes.Buffer
	if property.Nullable {
		deref := true
		buf.WriteString("if " + varname + "." + property.Field.Name + " == nil {\n")
		switch toGoType(property) {
		case "int32", "*int32":
			{
				buf.WriteString("\t\treturn int32(0)")
			}
		case "int64", "*int64":
			{
				buf.WriteString("\t\treturn int64(0)")
			}
		case "uint32", "*uint32":
			{
				buf.WriteString("\t\treturn uint32(0)")
			}
		case "uint64", "*uint64":
			{
				buf.WriteString("\t\treturn uint64(0)")
			}
		case "string", "*string":
			{
				buf.WriteString("\t\treturn \"\"")
			}
		case "bool", "*bool":
			{
				buf.WriteString("\t\treturn false")
			}
		case "float32", "*float32":
			{
				buf.WriteString("\t\treturn float32(0.0)")
			}
		case "float64", "*float64":
			{
				buf.WriteString("\t\treturn float64(0.0)")
			}
		case "[]byte":
			{
				buf.WriteString("\t\treturn nil")
			}
		case "*timestamp.Timestamp":
			{
				deref = false
				buf.WriteString("\t\treturn nil")
			}
		default:
			{
				if property.IsEnumeration() {
					if stringer {
						deref = false
						buf.WriteString("\t\treturn \"\"")
					} else {
						n, ev := findEnum(property)
						buf.WriteString("\t\treturn " + n + "_" + ev.Value[0].GetName())
					}
				}
			}
		}
		buf.WriteString("\n")
		buf.WriteString("\t}\n\t")
		if property.IsEnumeration() && stringer {
			buf.WriteString("return " + varname + "." + property.Field.Name + ".String()")
		} else {
			if deref {
				buf.WriteString("return *" + varname + "." + property.Field.Name)
			} else {
				buf.WriteString("return " + varname + "." + property.Field.Name)
			}
		}
	} else {
		if property.IsEnumeration() && stringer {
			buf.WriteString("return " + varname + "." + property.Field.Name + ".String()")
		} else {
			buf.WriteString("return " + varname + "." + property.Field.Name)
		}
	}
	return buf.String()
}

func toSetterValue(property *types.Property, clname string, varname string, stringer bool) string {
	if property.IsEnumeration() {
		if stringer {
			n := property.Entity.Name + "_" + property.Name
			if property.Nullable {
				return clname + "." + property.Field.Name + " = to" + n + "(" + varname + ")"
			}
			var buf bytes.Buffer
			buf.WriteString("var _" + property.Field.Name + " = to" + n + "(" + varname + ")\n")
			buf.WriteString("\tif _" + property.Field.Name + " != nil {\n")
			buf.WriteString("\t\t" + clname + "." + property.Field.Name + " = *_" + property.Field.Name + "\n")
			buf.WriteString("\t}")
			return buf.String()
		}
	}
	if property.Nullable && !isTimestamp(property) {
		return clname + "." + property.Field.Name + " = &" + varname
	}
	return clname + "." + property.Field.Name + " = " + varname
}

func toSQLConversion(property *types.Property) string {
	switch toGoType(property) {
	case "int32", "int64", "uint32", "uint64",
		"*int32", "*int64", "*uint32", "*uint64":
		{
			return "sql.NullInt64"
		}
	case "string", "*string":
		{
			return "sql.NullString"
		}
	case "bool", "*bool":
		{
			return "sql.NullBool"
		}
	case "float32", "float64", "*float32", "*float64":
		{
			return "sql.NullFloat64"
		}
	case "[]byte":
		{
			return "sql.NullString"
		}
	case "*timestamp.Timestamp":
		{
			return "NullTime"
		}
	default:
		{
			if property.IsEnumeration() {
				return "sql.NullString"
			}
		}
	}
	return property.Field.Type
}

func fromSQLConversion(property *types.Property) string {
	switch toGoType(property) {
	case "int32", "int64", "uint32", "uint64",
		"*int32", "*int64", "*uint32", "*uint64":
		{
			return "orm.ToSQLInt64"
		}
	case "string", "*string":
		{
			return "orm.ToSQLString"
		}
	case "bool", "*bool":
		{
			return "orm.ToSQLBool"
		}
	case "float32", "float64", "*float32", "*float64":
		{
			return "orm.ToSQLFloat64"
		}
	case "[]byte":
		{
			return "orm.ToSQLBlob"
		}
	case "*timestamp.Timestamp":
		{
			return "orm.ToSQLDate"
		}
	default:
		{
			if property.IsEnumeration() {
				return "orm.ToSQLString"
			}
		}
	}
	return property.Field.Type
}

func toPropertySetter(property *types.Property) string {
	switch toGoType(property) {
	case "int32", "*int32", "uint32", "*uint32":
		{
			return "int32(_" + property.Field.Name + ".Int64)"
		}
	case "int64", "*int64", "uint64", "*uint64":
		{
			return "_" + property.Field.Name + ".Int64"
		}
	case "string", "*string":
		{
			return "_" + property.Field.Name + ".String"
		}
	case "bool", "*bool":
		{
			return "_" + property.Field.Name + ".Bool"
		}
	case "float32", "*float32":
		{
			return "float32(_" + property.Field.Name + ".Float64)"
		}
	case "float64", "*float64":
		{
			return "_" + property.Field.Name + ".Float64"
		}
	case "[]byte":
		{
			return "[]byte(_" + property.Field.Name + ".String)"
		}
	case "*timestamp.Timestamp":
		{
			return "t.toTimestamp(_" + property.Field.Name + ".Time)"
		}
	default:
		{
			if property.IsEnumeration() {
				return "_" + property.Field.Name + ".String"
			}
		}
	}
	return ""
}

func toChecksumField(t types.Entity, value string) string {
	for _, column := range t.Properties {
		if column.IsChecksum() {
			if column.Nullable {
				return column.Field.Name + " = &" + value
			}
			return column.Field.Name + " = " + value
		}
	}
	return ""
}

func toSQL(table types.Entity) string {
	buf, _ := sql.GenerateSQL(&table)
	s := string(buf)
	var out bytes.Buffer
	out.WriteString(`"`)
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "-- ") {
			continue
		}
		// clean up the line
		out.WriteString(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.TrimSpace(line), "\t", " ", -1), "   ", "", -1), "  ", " ", -1), `"`, `\"`, -1))
	}
	out.WriteString(`"`)
	return out.String()
}

func fixcode(buf []byte) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", "goimp")
	if err != nil {
		return nil, err
	}
	tmpfile.Write(buf)
	tmpfile.Close()
	var b bytes.Buffer
	c := exec.Command("goimports", tmpfile.Name())
	c.Stdout = &b
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return nil, err
	}
	tmpfile, err = ioutil.TempFile("", "goimp")
	if err != nil {
		return nil, err
	}
	tmpfile.Write(b.Bytes())
	tmpfile.Close()
	b.Reset()
	c = exec.Command("gofmt", "-s", tmpfile.Name())
	c.Stdout = &b
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func generate(tmpl string, name string, file *types.File, entity *types.Entity, fn map[string]interface{}, results []*types.Generation) (*types.Generation, error) {
	buf, err := entity.GenerateCode(tmpl, nil, fn)
	if err != nil {
		return nil, err
	}
	// fmt.Fprintln(os.Stderr, string(buf))
	buf, err = fixcode(buf)
	if err != nil {
		return nil, err
	}
	return &types.Generation{
		Filename: file.Package + "/golang/" + name + ".go",
		Output:   string(buf),
	}, nil
}

func toGoEnumPointer(property *types.Property) string {
	if property.Nullable {
		return ""
	}
	return "&"
}

func toSetterSuffix(property *types.Property) string {
	if property.IsEnumeration() {
		return "String"
	}
	return ""
}

func (g *gogenerator) Generate(scheme string, file *types.File, entities []types.Entity) ([]*types.Generation, error) {
	results := make([]*types.Generation, 0)
	fn := make(map[string]interface{})
	fn["GoType"] = toGoType
	fn["GoTypeWithoutPointer"] = toGoTypeWithoutPointer
	fn["GoTags"] = toGoTags
	fn["ConvertToSQL"] = toSQLConversion
	fn["ConvertFromSQL"] = fromSQLConversion
	fn["GoGetterValue"] = toGetterValue
	fn["GoSetterValue"] = toSetterValue
	fn["GoSetterSuffix"] = toSetterSuffix
	fn["GoPropertySetter"] = toPropertySetter
	fn["GoChecksum"] = toChecksumField
	fn["GoEnumDefinitions"] = toGoEnumDefinitions
	fn["GoEnumPointer"] = toGoEnumPointer
	fn["GoTestData"] = toTestData
	fn["SQL"] = toSQL
	for _, entity := range entities {
		fmt.Fprintln(os.Stderr, "generating "+entity.Name)
		result, err := generate(goTemplate, strings.ToLower(entity.Name), file, &entity, fn, results)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
		result, err = generate(goTestTemplate, strings.ToLower(entity.Name)+"_test", file, &entity, fn, results)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	buf, err := types.GenerateCode(goTestMainTemplate, map[string]interface{}{"PkgName": file.Package}, nil)
	if err != nil {
		return nil, err
	}
	buf, err = fixcode(buf)
	if err != nil {
		return nil, err
	}
	results = append(results, &types.Generation{
		Filename: file.Package + "/golang/testmain_test.go",
		Output:   string(buf),
	})
	buf, err = types.GenerateCode(goUtilTemplate, map[string]interface{}{"PkgName": file.Package}, nil)
	if err != nil {
		return nil, err
	}
	buf, err = fixcode(buf)
	if err != nil {
		return nil, err
	}
	results = append(results, &types.Generation{
		Filename: file.Package + "/golang/" + strings.ToLower(file.Package) + "_util.go",
		Output:   string(buf),
	})
	return results, nil
}

const goTemplate = `
{{- with .Entity -}}
{{- $m := .Name }}
{{- $w := .ColumnWidth }}
{{- $cl := .SQLColumnList }}
{{- $tn := .SQLTableName }}
{{- $pkc := .PrimaryKey }}
{{- $tnp := .TableNamePlural }}
{{- $tns := .TableNameSingular }}
{{- $columns := .Properties }}
{{- $hpk := .HasPrimaryKey }}
{{- $pkp := .PrimaryKeyProperty }}
{{- $tnt := tick $tn }}
package {{.Package}}

import (
	"github.com/jhaynie/go-gator/orm"
)

{{ GoEnumDefinitions . }}

// {{$m}} table
type {{$m}} struct {
	{{- range $i, $col := .Properties }}
	{{- $gt := GoType $col }}
	{{- $tags := GoTags $col }}
	{{ pad $col.Field.Name $w }}  {{ pad $gt 27 }} {{ $tags }}
	{{- end }}
}

{{- range $i, $col := .Properties }}

// Get{{ $col.Field.Name }} will return the {{ $m }} {{ $col.Field.Name }} value
func (t *{{$m}}) Get{{ $col.Field.Name }}() {{ GoTypeWithoutPointer $col }} {
	{{ GoGetterValue $col "t" false}}
}

// Set{{ $col.Field.Name }} will set the {{ $m }} {{ $col.Field.Name }} value
func (t *{{$m}}) Set{{ $col.Field.Name }}(v {{ GoTypeWithoutPointer $col }}) {
	{{ GoSetterValue $col "t" "v" false }}
}

{{- if $col.IsEnumeration }}
// Get{{ $col.Field.Name }}String will return the {{ $m }} {{ $col.Field.Name }} value as a string
func (t *{{$m}}) Get{{ $col.Field.Name }}String() string {
	{{ GoGetterValue $col "t" true}}
}

// Set{{ $col.Field.Name }}String will set the {{ $m }} {{ $col.Field.Name }} value from a string
func (t *{{$m}}) Set{{ $col.Field.Name }}String(v string) {
	{{ GoSetterValue $col "t" "v" true }}
}
{{- end }}

{{- if not $col.PrimaryKey }}
{{- if $col.Index }}

// Find{{$tnp}}By{{ $col.Field.Name }} will find all {{$m}}s by the {{ $col.Field.Name }} value
func Find{{$tnp}}By{{ $col.Field.Name }}(ctx context.Context, db *sql.DB, value {{ GoTypeWithoutPointer $col }}) ([]*{{$m}}, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$col.SQLColumnNameWithTick}} = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, {{ ConvertFromSQL $pkp }}(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*{{$m}}, 0)
	for rows.Next() {
		{{- range $ii, $c := $columns }}
		var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
		{{- end }}
		err := rows.Scan(
			{{- range $f, $c := $columns }}
			&_{{ $c.Field.Name }},
			{{- end }}
		)
		if err != nil {
			return nil, err
		}
		t := &{{$m}}{}
		{{- range $ii, $c := $columns }}
		if _{{ $c.Field.Name }}.Valid {
			t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
		}
		{{- end }}
		results = append(results, t)
	}
	return results, nil
}

// Find{{$tnp}}By{{ $col.Field.Name }}Tx will find all {{$m}}s by the {{ $col.Field.Name }} value using the provided transaction
func Find{{$tnp}}By{{ $col.Field.Name }}Tx(ctx context.Context, tx *sql.Tx, value {{ GoTypeWithoutPointer $col }}) ([]*{{$m}}, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$col.SQLColumnNameWithTick}} = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, {{ ConvertFromSQL $pkp }}(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*{{$m}}, 0)
	for rows.Next() {
		{{- range $ii, $c := $columns }}
		var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
		{{- end }}
		err := rows.Scan(
			{{- range $f, $c := $columns }}
			&_{{ $c.Field.Name }},
			{{- end }}
		)
		if err != nil {
			return nil, err
		}
		t := &{{$m}}{}
		{{- range $ii, $c := $columns }}
		if _{{ $c.Field.Name }}.Valid {
			t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
		}
		{{- end }}
		results = append(results, t)
	}
	return results, nil
}
{{- end }}
{{- else }}

{{- if $hpk }}

// Find{{$tns}}By{{ $col.Field.Name }} will find a {{$m}} by {{ $col.Field.Name }}
func Find{{$tns}}By{{ $col.Field.Name }}(ctx context.Context, db *sql.DB, value {{ GoTypeWithoutPointer $col }}) (*{{$m}}, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$pkc}} = ?"
	{{- range $ii, $c := $columns }}
	var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
	{{- end }}
	err := db.QueryRowContext(ctx, q, value).Scan(
		{{- range $f, $c := $columns }}
		&_{{ $c.Field.Name }},
		{{- end }}
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &{{$m}}{}
	{{- range $ii, $c := $columns }}
	if _{{ $c.Field.Name }}.Valid {
		t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
	}
	{{- end }}
	return t, nil
}

// Find{{$tns}}By{{ $col.Field.Name }}Tx will find a {{$m}} by {{ $col.Field.Name }} using the provided transaction
func Find{{$tns}}By{{ $col.Field.Name }}Tx(ctx context.Context, tx *sql.Tx, value {{ GoTypeWithoutPointer $col }}) (*{{$m}}, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$pkc}} = ?"
	{{- range $ii, $c := $columns }}
	var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
	{{- end }}
	err := tx.QueryRowContext(ctx, q, value).Scan(
		{{- range $f, $c := $columns }}
		&_{{ $c.Field.Name }},
		{{- end }}
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &{{$m}}{}
	{{- range $ii, $c := $columns }}
	if _{{ $c.Field.Name }}.Valid {
		t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
	}
	{{- end }}
	return t, nil
}
{{- end }}

{{- end }}

{{- end }}

func (t *{{$m}}) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreate{{$m}}Table will create the {{$m}} table
func DBCreate{{$m}}Table(ctx context.Context, db *sql.DB) error {
	q := {{ SQL . }}
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreate{{$m}}TableTx will create the {{$m}} table using the provided transction
func DBCreate{{$m}}TableTx(ctx context.Context, tx *sql.Tx) error {
	q := {{ SQL . }}
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDrop{{$m}}Table will drop the {{$m}} table
func DBDrop{{$m}}Table(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS {{$tnt}}"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDrop{{$m}}TableTx will drop the {{$m}} table using the provided transaction
func DBDrop{{$m}}TableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS {{$tnt}}"
	_, err := tx.ExecContext(ctx, q)
	return err
}

{{- if .HasChecksum }}
// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *{{$m}}) CalculateChecksum() string {
	return orm.HashStrings(
		{{- range $i, $col := $columns }}
		{{- if not $col.IsChecksum }}
		{{- if $col.IsEnumeration }}
		enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }}),
		{{- else }}
		orm.ToString(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
		{{- end }}
	)
}
{{- end }}

// DBCreate will create a new {{$m}} record in the database
func (t *{{$m}}) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}})"
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{ GoChecksum . "checksum" }}
	{{- end }}
	return db.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
}

// DBCreateTx will create a new {{$m}} record in the database using the provided transaction
func (t *{{$m}}) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}})"
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{ GoChecksum . "checksum" }}
	{{- end }}
	return tx.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
}

{{- if .HasPrimaryKey }}

// DBCreateIgnoreDuplicate will upsert the {{$m}} record in the database
func (t *{{$m}}) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}}) ON DUPLICATE KEY UPDATE {{$pkc}} = {{$pkc}}"
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{ GoChecksum . "checksum" }}
	{{- end }}
	return db.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
}

// DBCreateIgnoreDuplicateTx will upsert the {{$m}} record in the database using the provided transaction
func (t *{{$m}}) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}}) ON DUPLICATE KEY UPDATE {{$pkc}} = {{$pkc}}"
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{ GoChecksum . "checksum" }}
	{{- end }}
	return tx.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
}

{{- end }}

// DeleteAll{{$tnp}} deletes all {{$m}} records in the database with optional filters
func DeleteAll{{$tnp}}(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	_, err := db.ExecContext(ctx, "DELETE "+ q, p...)
	return err
}

// DeleteAll{{$tnp}}Tx deletes all {{$m}} records in the database with optional filters using the provided transaction
func DeleteAll{{$tnp}}Tx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	_, err := tx.ExecContext(ctx, "DELETE "+ q, p...)
	return err
}

{{- if .HasPrimaryKey }}

// DBDelete will delete this {{$m}} record in the database
func (t *{{$m}}) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM {{$tnt}} WHERE {{$pkc}} = ?"
	r, err := db.ExecContext(ctx, q, {{ ConvertFromSQL $pkp }}(t.{{$pkp.Field.Name}}))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBDeleteTx will delete this {{$m}} record in the database using the provided transaction
func (t *{{$m}}) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM {{$tnt}} WHERE {{$pkc}} = ?"
	r, err := tx.ExecContext(ctx, q, {{ ConvertFromSQL $pkp }}(t.{{$pkp.Field.Name}}))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBUpdate will update the {{$m}} record in the database
func (t *{{$m}}) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{GoChecksum . "checksum"}}
	{{- end }}
	q := "UPDATE {{$tnt}} SET {{.SQLColumnSetterList}} WHERE {{$pkc}}=?"
	return db.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if not $col.PrimaryKey }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
		{{- end }}
		{{ ConvertFromSQL $pkp }}(t.{{$pkp.Field.Name}}),
	)
}

// DBUpdateTx will update the {{$m}} record in the database using the provided transaction
func (t *{{$m}}) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return nil, nil
	}
	t.{{GoChecksum . "checksum"}}
	{{- end }}
	q := "UPDATE {{$tnt}} SET {{.SQLColumnSetterList}} WHERE {{$pkc}}=?"
	return tx.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if not $col.PrimaryKey }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
		{{- end }}
		{{ ConvertFromSQL $pkp }}(t.{{$pkp.Field.Name}}),
	)
}

{{- end }}

// DBUpsert will upsert the {{$m}} record in the database
func (t *{{$m}}) DBUpsert(ctx context.Context, db *sql.DB) (bool, bool, error) {
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return false, false, nil
	}
	t.{{GoChecksum . "checksum"}}
	{{- end }}
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}}) ON DUPLICATE KEY UPDATE {{.SQLColumnUpsertList}}"
	r, err := db.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the {{$m}} record in the database using the provided transaction
func (t *{{$m}}) DBUpsertTx(ctx context.Context, tx *sql.Tx) (bool, bool, error) {
	{{- if .HasChecksum }}
	checksum := t.CalculateChecksum()
	if t.Get{{.Checksum}}() == checksum {
		return false, false, nil
	}
	t.{{GoChecksum . "checksum"}}
	{{- end }}
	q := "INSERT INTO {{$tnt}} ({{$cl}}) VALUES ({{.SQLColumnPlaceholders}}) ON DUPLICATE KEY UPDATE {{.SQLColumnUpsertList}}"
	r, err := tx.ExecContext(ctx, q,
		{{- range $i, $col := $columns }}
		{{- if $col.IsEnumeration }}
		{{ ConvertFromSQL $col }}(enum{{$m}}{{$col.Field.Name}}ToString({{ GoEnumPointer $col -}}t.{{ $col.Field.Name }})),
		{{- else }}
		{{ ConvertFromSQL $col }}(t.{{ $col.Field.Name }}),
		{{- end }}
		{{- end }}
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

{{- if .HasPrimaryKey }}

// DBFindOne will find a {{$m}} record in the database with the primary key
func (t *{{$m}}) DBFindOne(ctx context.Context, db *sql.DB, value {{GoTypeWithoutPointer $pkp}}) (bool, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$pkc}} = ? LIMIT 1"	
	row := db.QueryRowContext(ctx, q, {{ ConvertFromSQL $pkp }}(value))
	{{- range $i, $col := $columns }}
	var _{{ $col.Field.Name }} {{ ConvertToSQL $col }}
	{{- end }}
	err := row.Scan(
		{{- range $i, $col := $columns }}
		&_{{ $col.Field.Name }},
		{{- end }}
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _{{$pkp.Field.Name}}.Valid == false {
		return false, nil
	}
	{{- range $i, $col := $columns }}
	if _{{ $col.Field.Name }}.Valid {
		t.Set{{ $col.Field.Name }}{{ GoSetterSuffix $col -}}({{ GoPropertySetter $col }})
	}
	{{- end }}
	return true, nil
}

// DBFindOneTx will find a {{$m}} record in the database with the primary key using the provided transaction
func (t *{{$m}}) DBFindOneTx(ctx context.Context, tx *sql.Tx, value {{GoTypeWithoutPointer $pkp}}) (bool, error) {
	q := "SELECT {{$cl}} FROM {{$tnt}} WHERE {{$pkc}} = ? LIMIT 1"	
	row := tx.QueryRowContext(ctx, q, {{ ConvertFromSQL $pkp }}(value))
	{{- range $i, $col := $columns }}
	var _{{ $col.Field.Name }} {{ ConvertToSQL $col }}
	{{- end }}
	err := row.Scan(
		{{- range $i, $col := $columns }}
		&_{{ $col.Field.Name }},
		{{- end }}
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _{{$pkp.Field.Name}}.Valid == false {
		return false, nil
	}
	{{- range $i, $col := $columns }}
	if _{{ $col.Field.Name }}.Valid {
		t.Set{{ $col.Field.Name }}{{ GoSetterSuffix $col -}}({{ GoPropertySetter $col }})
	}
	{{- end }}
	return true, nil
}

{{- end }}

// Find{{$tnp}} will find a {{$m}} record in the database with the provided parameters
func Find{{$tnp}}(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*{{$m}}, error) {
	params := []interface{}{
	{{- range $i, $col := $columns }}
		orm.Column("{{$col.SQLColumnName}}"),
	{{- end }}
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	rows, err := db.QueryContext(ctx, q, p...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*{{$m}}, 0)
	for rows.Next() {
		{{- range $ii, $c := $columns }}
		var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
		{{- end }}
		err := rows.Scan(
			{{- range $f, $c := $columns }}
			&_{{ $c.Field.Name }},
			{{- end }}
		)
		if err != nil {
			return nil, err
		}
		t := &{{$m}}{}
		{{- range $ii, $c := $columns }}
		if _{{ $c.Field.Name }}.Valid {
			t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
		}
		{{- end }}
		results = append(results, t)
	}
	return results, nil
}

// Find{{$tnp}}Tx will find a {{$m}} record in the database with the provided parameters using the provided transaction
func Find{{$tnp}}Tx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*{{$m}}, error) {
	params := []interface{}{
	{{- range $i, $col := $columns }}
		orm.Column("{{$col.SQLColumnName}}"),
	{{- end }}
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	rows, err := tx.QueryContext(ctx, q, p...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*{{$m}}, 0)
	for rows.Next() {
		{{- range $ii, $c := $columns }}
		var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
		{{- end }}
		err := rows.Scan(
			{{- range $f, $c := $columns }}
			&_{{ $c.Field.Name }},
			{{- end }}
		)
		if err != nil {
			return nil, err
		}
		t := &{{$m}}{}
		{{- range $ii, $c := $columns }}
		if _{{ $c.Field.Name }}.Valid {
			t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
		}
		{{- end }}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a {{$m}} record in the database with the provided parameters
func (t *{{$m}}) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
	{{- range $i, $col := $columns }}
		orm.Column("{{$col.SQLColumnName}}"),
	{{- end }}
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	{{- range $i, $c := $columns }}
	var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
	{{- end }}
	err := row.Scan(
		{{- range $i, $c := $columns }}
		&_{{ $c.Field.Name }},
		{{- end }}
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	{{- range $i, $c := $columns }}
	if _{{ $c.Field.Name }}.Valid {
		t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
	}
	{{- end }}
	return true, nil
}

// DBFindTx will find a {{$m}} record in the database with the provided parameters using the provided transaction
func (t *{{$m}}) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
	{{- range $i, $col := $columns }}
		orm.Column("{{$col.SQLColumnName}}"),
	{{- end }}
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	{{- range $i, $c := $columns }}
	var _{{ $c.Field.Name }} {{ ConvertToSQL $c }}
	{{- end }}
	err := row.Scan(
		{{- range $i, $c := $columns }}
		&_{{ $c.Field.Name }},
		{{- end }}
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	{{- range $i, $c := $columns }}
	if _{{ $c.Field.Name }}.Valid {
		t.Set{{ $c.Field.Name }}{{ GoSetterSuffix $c -}}({{ GoPropertySetter $c }})
	}
	{{- end }}
	return true, nil
}

// Count{{$tnp}} will find the count of {{$m}} records in the database
func Count{{$tnp}}(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := db.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// Count{{$tnp}}Tx will find the count of {{$m}} records in the database using the provided transaction
func Count{{$tnp}}Tx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := tx.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// DBCount will find the count of {{$m}} records in the database
func (t *{{$m}}) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := db.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// DBCountTx will find the count of {{$m}} records in the database using the provided transaction
func (t *{{$m}}) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table("{{$tn}}"),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := tx.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

{{- if .HasPrimaryKey }}

// DBExists will return true if the {{$m}} record exists in the database
func (t *{{$m}}) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT {{$pkc}} FROM {{$tnt}} WHERE {{$pkc}} = ? LIMIT 1"	
	var _{{$pkp.Name}} sql.NullString
	err := db.QueryRowContext(ctx, q, {{ ConvertFromSQL $pkp }}(t.{{$pkp.Name}})).Scan(&_{{$pkp.Name}})
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _{{$pkp.Name}}.Valid, nil
}

// DBExistsTx will return true if the {{$m}} record exists in the database using the provided transaction
func (t *{{$m}}) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT {{$pkc}} FROM {{$tnt}} WHERE {{$pkc}} = ? LIMIT 1"	
	var _{{$pkp.Name}} sql.NullString
	err := tx.QueryRowContext(ctx, q, {{ ConvertFromSQL $pkp }}(t.{{$pkp.Name}})).Scan(&_{{$pkp.Name}})
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _{{$pkp.Name}}.Valid, nil
}

{{- end -}}

{{- end -}}
`

const goTestTemplate = `
{{- with .Entity -}}
{{- $m := .Name }}
{{- $w := .ColumnWidth }}
{{- $cl := .SQLColumnList }}
{{- $tn := .SQLTableName }}
{{- $pkc := .PrimaryKey }}
{{- $tnp := .TableNamePlural }}
{{- $tns := .TableNameSingular }}
{{- $columns := .Properties }}
{{- $hpk := .HasPrimaryKey }}
{{- $pkp := .PrimaryKeyProperty }}
{{- $tnt := tick $tn }}
package {{.Package}}

import (
	"github.com/jhaynie/go-gator/orm"
)

func TestCreate{{$m}}Table(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreate{{$m}}TableTx(context.Background(), tx)
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
}

func TestCreate{{$m}}Delete(t *testing.T) {
	r := &{{$m}}{
		{{- range $i, $col := $columns }}
		{{ $col.Field.Name }}: {{ GoTestData $col false }},
		{{- end }}
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAll{{$tnp}}(ctx, db)
	result, err := r.DBCreate(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected result to be non-nil")
	}
	count, err := r.DBCount(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("count should have been 1 but was %d", count)
	}
	exists, err := r.DBExists(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("exists should have been true but was false")
	}
	found, err := Find{{$tns}}By{{$pkp.Name}}(ctx, db, r.{{$pkp.Name}})
	if err != nil {
		t.Fatal(err)
	}
	if found == nil {
		t.Fatal("expected found to be a value but was nil")
	}
	if found.{{$pkp.Name}} != r.{{$pkp.Name}} {
		t.Fatalf("expected found primary key to be %v but was %v", r.{{$pkp.Name}}, found.{{$pkp.Name}})
	}
	if orm.Stringify(r) != orm.Stringify(found) {
		t.Fatalf("expected r to be found but was different")
	}
	results, err := Find{{.TableNamePlural}}(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if results == nil {
		t.Fatal("expected results to be a value but was nil")
	}
	if len(results) != 1 {
		t.Log(orm.Stringify(results))
		t.Fatalf("expected results length to be 1 but was %d", len(results))
	}
	f, err := r.DBFindOne(ctx, db, r.Get{{$pkp.Name}}())
	if err != nil {
		t.Fatal(err)
	}
	if f == false {
		t.Fatal("expected found to be a true but was false")
	}
	{{- if .HasChecksum }}
	a, b, err := r.DBUpsert(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if a {
		t.Fatal("expected a to be false but was true")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	{{- range $i, $col := $columns }}
	{{- if not $col.PrimaryKey }}
	{{- if not $col.Nullable }}
	r.Set{{ $col.Name }}({{ GoTestData $col true }})
	{{ addctx "nullable" 1}}
	{{- end }}
	{{- end }}
	{{- end }}
	{{- if hasctx "nullable" }}
	a, b, err = r.DBUpsert(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if !a {
		t.Fatal("expected a to be true but was false")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	{{- end }}
	{{- end }}
	_, err = r.DBDelete(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	count, err = r.DBCount(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("count should have been 0 but was %d", count)
	}
	exists, err = r.DBExists(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("exists should have been false but was true")
	}
}

func TestCreate{{$m}}DeleteTx(t *testing.T) {
	r := &{{$m}}{
		{{- range $i, $col := $columns }}
		{{ $col.Field.Name }}: {{ GoTestData $col false }},
		{{- end }}
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAll{{$tnp}}(ctx, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	result, err := r.DBCreateTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected result to be non-nil")
	}
	count, err := r.DBCountTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("count should have been 1 but was %d", count)
	}
	exists, err := r.DBExistsTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("exists should have been true but was false")
	}
	found, err := Find{{$tns}}By{{$pkp.Name}}Tx(ctx, tx, r.{{$pkp.Name}})
	if err != nil {
		t.Fatal(err)
	}
	if found == nil {
		t.Fatal("expected found to be a value but was nil")
	}
	if found.{{$pkp.Name}} != r.{{$pkp.Name}} {
		t.Fatalf("expected found primary key to be %v but was %v", r.{{$pkp.Name}}, found.{{$pkp.Name}})
	}
	if orm.Stringify(r) != orm.Stringify(found) {
		t.Fatalf("expected r to be found but was different")
	}
	results, err := Find{{.TableNamePlural}}Tx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if results == nil {
		t.Fatal("expected results to be a value but was nil")
	}
	if len(results) != 1 {
		t.Log(orm.Stringify(results))
		t.Fatalf("expected results length to be 1 but was %d", len(results))
	}
	f, err := r.DBFindOneTx(ctx, tx, r.Get{{$pkp.Name}}())
	if err != nil {
		t.Fatal(err)
	}
	if f == false {
		t.Fatal("expected found to be a true but was false")
	}
	{{- if .HasChecksum }}
	a, b, err := r.DBUpsertTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if a {
		t.Fatal("expected a to be false but was true")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	{{- rmctx "nullable" }}
	{{- range $i, $col := $columns }}
	{{- if not $col.PrimaryKey }}
	{{- if not $col.Nullable }}
	r.Set{{ $col.Name }}({{ GoTestData $col true }})
	{{ addctx "nullable" 1}}
	{{- end }}
	{{- end }}
	{{- end }}
	{{- if hasctx "nullable" }}
	a, b, err = r.DBUpsertTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if !a {
		t.Fatal("expected a to be true but was false")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	{{- end }}
	{{- end }}
	_, err = r.DBDeleteTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	count, err = r.DBCountTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("count should have been 0 but was %d", count)
	}
	exists, err = r.DBExistsTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("exists should have been false but was true")
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}

{{- end }}
`

const goTestMainTemplate = `package {{ .PkgName }}

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jhaynie/go-gator/orm"
)

var (
	database string
	username string
	password string
	hostname string
	port     int
	db       *sql.DB
	createdb = true
)

func init() {
	var defuser = "root"
	var defdb = fmt.Sprintf("test_%s", orm.UUID()[0:9])
	flag.StringVar(&username, "username", defuser, "database username")
	flag.StringVar(&password, "password", "", "database password")
	flag.StringVar(&hostname, "hostname", "localhost", "database hostname")
	flag.IntVar(&port, "port", 3306, "database port")
	database = defdb
}

func GetDatabase() *sql.DB {
	return db
}

func GetDSN(name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, name)
}

func openDB(name string) *sql.DB {
	dsn := GetDSN(name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}

func dropDB() {
	if createdb {
		_, err := db.Exec(fmt.Sprintf("drop database %s", database))
		if err != nil {
			fmt.Printf("error dropping database named %s\n", database)
		}
	}
}

func ToTimestampNow() *timestamp.Timestamp {
	t := time.Now()
	// truncate to 24 hours
	t = t.Truncate(time.Hour * 24)
	ts, _ := ptypes.TimestampProto(t)
	// since mysql is only second precision we truncate
	ts.Nanos = 0
	return ts
}

func TestMain(m *testing.M) {
	flag.Parse()
	if createdb {
		// open without a database so we can create a temp one
		d := openDB("")
		_, err := d.Exec(fmt.Sprintf("create database %s", database))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		d.Close()
	}
	// reopen now with the temp database
	db = openDB(database)
	x := m.Run()
	dropDB()
	db.Close()
	os.Exit(x)
}
`

const goUtilTemplate = `package {{ .PkgName }}
// NullTime taken from
// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// NullTime represents a time.Time that may be NULL.
// NullTime implements the Scanner interface so
// it can be used as a scan destination:
//
//  var nt NullTime
//  err := db.QueryRow("SELECT time FROM foo WHERE id=?", id).Scan(&nt)
//  ...
//  if nt.Valid {
//     // use nt.Time
//  } else {
//     // NULL value
//  }
//

const timeFormat = "2006-01-02 15:04:05.999999"

// This NullTime implementation is not driver-specific
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
// The value type must be time.Time or string / []byte (formatted time-string),
// otherwise Scan fails.
func (nt *NullTime) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return
	case []byte:
		nt.Time, err = parseDateTime(string(v), time.UTC)
		nt.Valid = (err == nil)
		return
	case string:
		nt.Time, err = parseDateTime(v, time.UTC)
		nt.Valid = (err == nil)
		return
	}

	nt.Valid = false
	return fmt.Errorf("Can't convert %T to time.Time", value)
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func parseDateTime(str string, loc *time.Location) (t time.Time, err error) {
	base := "0000-00-00 00:00:00.0000000"
	switch len(str) {
	case 10, 19, 21, 22, 23, 24, 25, 26: // up to "YYYY-MM-DD HH:MM:SS.MMMMMM"
		if str == base[:len(str)] {
			return
		}
		t, err = time.Parse(timeFormat[:len(str)], str)
	default:
		err = fmt.Errorf("invalid time string: %s", str)
		return
	}

	// Adjust location
	if err == nil && loc != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc), nil
	}

	return
}

func parseBinaryDateTime(num uint64, data []byte, loc *time.Location) (driver.Value, error) {
	switch num {
	case 0:
		return time.Time{}, nil
	case 4:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			0, 0, 0, 0,
			loc,
		), nil
	case 7:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			int(data[4]),                              // hour
			int(data[5]),                              // minutes
			int(data[6]),                              // seconds
			0,
			loc,
		), nil
	case 11:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			int(data[4]),                              // hour
			int(data[5]),                              // minutes
			int(data[6]),                              // seconds
			int(binary.LittleEndian.Uint32(data[7:11]))*1000, // nanoseconds
			loc,
		), nil
	}
	return nil, fmt.Errorf("invalid DATETIME packet length %d", num)
}
`
