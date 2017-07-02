package graphql

import (
	"bytes"
	"sort"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhaynie/protoc-gen-gator/generator"
	"github.com/jhaynie/protoc-gen-gator/types"
	"github.com/serenize/snaker"
)

type gqlgenerator struct {
}

func init() {
	generator.Register2("graphql", &gqlgenerator{})
}

type graphql struct {
	entity        *types.Entity
	Types         []string
	TypesOptional []string
}

type graphqlagg struct {
	Name string
	Type string
}

func toGraphQLAggregationMathFields(entity types.Entity) []graphqlagg {
	fields := make([]graphqlagg, 0)
	for _, p := range entity.Properties {
		t := toGraphQLType(&p)
		if strings.HasSuffix(t, "!") {
			t = t[0 : len(t)-1]
		}
		if !p.PrimaryKey && (t == "Int" || t == "Float") {
			fields = append(fields, graphqlagg{Name: toGraphQLVariableName(&p), Type: t})
		}
	}
	return fields
}

func toGraphQLTypeEnumDefinitions(entity types.Entity) string {
	var buf bytes.Buffer
	for _, e := range entity.Message.Descriptor.EnumType {
		names := make([]string, 0)
		for _, v := range e.Value {
			names = append(names, strings.ToLower(v.GetName()))
		}
		buf.WriteString("enum " + e.GetName() + " {\n")
		buf.WriteString("\t" + strings.Join(names, "\n\t"))
		buf.WriteString("\n}\n")
	}
	return buf.String()
}

func toGraphQLTypeOptional(property *types.Property) string {
	t := toGraphQLType(property)
	if property.Nullable {
		return t
	}
	return t[0 : len(t)-1]
}

func toGraphQLType(property *types.Property) string {
	var t string
	switch property.Field.Descriptor.GetType() {
	case
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		{
			t = "Int"
		}
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		{
			t = "Float"
		}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		{
			t = "Boolean"
		}
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		{
			t = "String"
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		{
			t = "String"
		}
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		{
			switch property.Field.Descriptor.GetTypeName() {
			case ".proto.ID":
				{
					t = "ID"
				}
			case ".proto.Checksum":
				{
					t = "String"
				}
			case ".proto.DateTime", ".google.protobuf.Timestamp":
				{
					t = "DateTime"
				}
			default:
				{
					t = "String"
				}
			}
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		{
			tok := strings.Split(*property.Field.Descriptor.TypeName, ".")
			t = tok[3]
		}
	}
	if !property.Nullable {
		t = t + "!"
	}
	return t
}

func toGraphQLVariableName(p *types.Property) string {
	return snaker.CamelToSnake(p.Name)
}

func toDefaultJSValue(p *types.Property) string {
	switch p.Field.Descriptor.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		{
			return "0"
		}
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		{
			return "false"
		}
	case descriptor.FieldDescriptorProto_TYPE_STRING,
		descriptor.FieldDescriptorProto_TYPE_BYTES,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		{
			return "null"
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		{
			//TODO:
		}
	}
	return "null"
}

func toGraphQLAssociationTypeName(a *types.SQLAssociation) string {
	switch a.Type {
	case types.SQLAssocationHasMany:
		{
			return "has_many"
		}
	case types.SQLAssocationBelongsTo:
		{
			return "belongs_to"
		}
	case types.SQLAssocationHasOne:
		{
			return "has_one"
		}
	}
	return ""
}

func toGraphQLAssociationType(a *types.SQLAssociation) string {
	switch a.Type {
	case types.SQLAssocationHasMany:
		{
			args := toGraphQLModelFieldArgs(*a.Entity, false)
			return a.Name + "(" + args + "): [" + snaker.SnakeToCamel(a.Table) + "]"
		}
	case types.SQLAssocationBelongsTo, types.SQLAssocationHasOne:
		{
			return a.Name + ": " + snaker.SnakeToCamel(a.Table)
		}
	}
	return ""
}

func toGraphQLAssociationTypeOptional(a *types.SQLAssociation) string {
	switch a.Type {
	case types.SQLAssocationHasMany:
		{
			args := toGraphQLModelFieldArgs(*a.Entity, false)
			return a.Name + "(" + args + "): [" + snaker.SnakeToCamel(a.Table) + "Optionals]"
		}
	case types.SQLAssocationBelongsTo, types.SQLAssocationHasOne:
		{
			return a.Name + ": " + snaker.SnakeToCamel(a.Table) + "Optionals"
		}
	}
	return ""
}

func toGraphQLAssociationTypeIs(a *types.SQLAssociation, typename string) bool {
	return toGraphQLAssociationTypeName(a) == typename
}

func toJSSafeVariable(name string) string {
	switch name {
	//ES6 reserved words
	case "do",
		"if",
		"in",
		"for",
		"let",
		"new",
		"try",
		"var",
		"case",
		"else",
		"enum",
		"eval",
		"null",
		"this",
		"true",
		"void",
		"with",
		"await",
		"break",
		"catch",
		"class",
		"const",
		"false",
		"super",
		"throw",
		"while",
		"yield",
		"delete",
		"export",
		"import",
		"public",
		"return",
		"static",
		"switch",
		"typeof",
		"default",
		"extends",
		"finally",
		"package",
		"private",
		"continue",
		"debugger",
		"function",
		"arguments",
		"interface",
		"protected",
		"implements",
		"instanceof",
		"helper", // our helper util
		// these are used as sql filters
		"limit",
		"offset",
		"sort",
		"sortOrder":
		{
			return name + "_"
		}
	}
	return name
}

func toGraphQLFieldParameters(e types.Entity, idcolumn bool) string {
	args := make([]string, 0)
	args = append(args, "limit", "offset")
	for _, p := range e.Properties {
		if !p.IsSQLIDColumn() || idcolumn && p.IsSQLIDColumn() {
			args = append(args, toJSSafeVariable(p.SQLColumnName()))
		}
	}
	return strings.Join(args, ", ")
}

func toGraphQLFieldParameterArgs(e types.Entity, idcolumn bool) string {
	args := make([]string, 0)
	for _, p := range e.Properties {
		if !p.IsSQLIDColumn() || idcolumn && p.IsSQLIDColumn() {
			args = append(args, "{name:'"+p.SQLColumnName()+"', value:"+toJSSafeVariable(p.SQLColumnName())+"}")
		}
	}
	return "[" + strings.Join(args, ", ") + "]"
}

func toGraphQLModelFieldArgs(e types.Entity, idcolumn bool) string {
	args := make([]string, 0)
	args = append(args, "limit:Int", "offset:Int", "sort:"+e.TableNameSingular()+"Fields", "sortOrder:QueryDirection")
	for _, p := range e.Properties {
		if !p.IsSQLIDColumn() || idcolumn && p.IsSQLIDColumn() {
			args = append(args, toJSSafeVariable(p.SQLColumnName())+": "+toGraphQLTypeOptional(&p))
		}
	}
	return strings.Join(args, ", ")
}

func (g *gqlgenerator) Generate(scheme string, file *types.File, entities []types.Entity) ([]*types.Generation, error) {
	results := make([]*types.Generation, 0)
	fn := make(map[string]interface{})
	fn["GraphQLType"] = toGraphQLType
	fn["GraphQLTypeOptional"] = toGraphQLTypeOptional
	fn["GraphQLVariable"] = toGraphQLVariableName
	fn["GraphQLTypeEnumDefinition"] = toGraphQLTypeEnumDefinitions
	fn["DefaultJSValue"] = toDefaultJSValue
	fn["GraphQLAssociationType"] = toGraphQLAssociationType
	fn["GraphQLAssocationTypeName"] = toGraphQLAssociationTypeName
	fn["GraphQLAssociationTypeOptional"] = toGraphQLAssociationTypeOptional
	fn["GraphQLAssocationTypeIs"] = toGraphQLAssociationTypeIs
	fn["GraphQLAggregationMathFields"] = toGraphQLAggregationMathFields
	fn["GraphQLFieldParameters"] = toGraphQLFieldParameters
	fn["GraphQLFieldParameterArgs"] = toGraphQLFieldParameterArgs
	fn["GraphQLModelFieldArgs"] = toGraphQLModelFieldArgs
	fn["SnakeToCamel"] = snaker.SnakeToCamel
	fn["CamelToSnake"] = snaker.CamelToSnake
	tbls := make([]string, 0)
	var rootbuf bytes.Buffer
	var tablebuf bytes.Buffer
	rootbuf.WriteString("type Query {")
	for _, entity := range entities {
		kv := make(map[string]interface{})
		t := make([]string, 0)
		to := make([]string, 0)
		for _, p := range entity.Properties {
			t = append(t, snaker.CamelToSnake(p.Name)+": "+toGraphQLType(&p))
			to = append(to, snaker.CamelToSnake(p.Name)+": "+toGraphQLTypeOptional(&p))
		}
		tbls = append(tbls, entity.TableNameSingular())
		kv["g"] = graphql{&entity, t, to}
		buf, err := entity.GenerateCode(graphqlTemplate, kv, fn)
		if err != nil {
			return nil, err
		}
		tablebuf.Write(buf)
		tablebuf.WriteString("\n")
		buf2, err := entity.GenerateCode(graphqlRootQueryTemplate, kv, fn)
		if err != nil {
			return nil, err
		}
		rootbuf.Write(buf2)
		buf, err = entity.GenerateCode(graphqlResolver, kv, fn)
		if err != nil {
			return nil, err
		}
		results = append(results, &types.Generation{
			Filename: file.Package + "/graphql/" + entity.Name + ".js",
			Output:   string(buf),
		})
	}
	rootbuf.WriteString("}\n")
	results = append(results, &types.Generation{
		Filename: file.Package + "/graphql/model.graphql",
		Output:   graphqlQueryUtils + "\n" + tablebuf.String() + rootbuf.String() + graphqlSchemaTemplate,
	})
	sort.Strings(tbls)
	buf, err := types.GenerateCode(graphqlIndex, map[string]interface{}{"Tables": tbls}, fn)
	if err != nil {
		return nil, err
	}
	results = append(results, &types.Generation{
		Filename: file.Package + "/graphql/index.js",
		Output:   string(buf),
	})
	results = append(results, &types.Generation{
		Filename: file.Package + "/graphql/_helper.js",
		Output:   graphqlQueryHelper,
	})
	return results, nil
}

const graphqlQueryUtils = `
scalar JSON
scalar Date
scalar Time
scalar DateTime

enum QueryDirection {
	ASCENDING
	DESCENDING
}

input QueryOrder {
	field: String!
	direction: QueryDirection
}

input QueryRange {
	offset: Int!
	limit: Int!
}

enum QueryConditionOperator {
	EQUAL
	NOT_EQUAL
	NULL
	NOT_NULL
	GREATER
	GREATER_EQ
	LESS
	LESS_EQ
	IN
	NOT_IN
	BETWEEN
	NOT_BETWEEN
	LIKE
	NOT_LIKE
}

enum QueryConditionGroupOperator {
	AND
	OR
}

input QueryCondition {
	field: String!
	operator: QueryConditionOperator!
	value: JSON
}

input QueryConditionGroup {
	conditions: [QueryCondition!]!
	operator: QueryConditionGroupOperator
}

input QueryFilter {
	order: [QueryOrder!]
	range: QueryRange
	limit: Int
	condition: [QueryConditionGroup!]
}

interface Table {
	_tablename: String
}

`

const graphqlRootQueryTemplate = `{{- $e := .Entity }}
	{{ lowerfc $e.TableNameSingular }}(filter: QueryFilter, {{GraphQLModelFieldArgs $e true}}):[{{ $e.TableNameSingular -}} Aggregation]
	{{ lowerfc $e.TableNamePlural }}(filter: QueryFilter, offset:Int, limit:Int, sort:{{ $e.TableNameSingular }}Fields, sortOrder:QueryDirection, {{GraphQLModelFieldArgs $e false}}):[{{ $e.TableNameSingular }}]
	{{ lowerfc $e.TableNameSingular }}By {{- $e.PrimaryKeyProperty.Field.Name}}( {{- GraphQLVariable $e.PrimaryKeyProperty -}}: {{ GraphQLType $e.PrimaryKeyProperty -}}):{{ $e.TableNameSingular }}
	{{- range $i, $value := $e.Properties }}
	{{- if .Index }}
	{{ lowerfc $e.TableNamePlural }}By {{- $value.Field.Name}}(  {{- GraphQLVariable . -}}: {{GraphQLType . -}}, filter: QueryFilter, offset:Int, limit:Int, sort:{{ $e.TableNameSingular }}Fields, sortOrder:QueryDirection):[{{ $e.TableNameSingular }}]
	{{- end }}
	{{- end }}
`

const graphqlTemplate = `{{- $e := .Entity -}}
{{ GraphQLTypeEnumDefinition $e }}
type {{ $e.TableNameSingular }} implements Table {
	{{- with .g }}
	{{- range $i, $value := .Types }}
	{{ . }}
	{{- end }}
	{{- end }}
	{{- if len $e.SQLAssociations }}
	{{- range $i, $a := $e.SQLAssociations }}
	{{ GraphQLAssociationType $a }}
	{{- end }}
	{{- end }}
	{{- range $i, $p := $e.AdditionalGraphQLTypes }}
	{{ $p.Name }}: {{ $p.Type }}
	{{- end }}
	_tablename: String
}

type {{ $e.TableNameSingular }}Optionals {
	{{- with .g }}
	{{- range $i, $value := .TypesOptional }}
	{{ . }}
	{{- end }}
	{{- end }}
	{{- if len $e.SQLAssociations }}
	{{- range $i, $a := $e.SQLAssociations }}
	{{ GraphQLAssociationType $a }}
	{{- end }}
	{{- end }}
}

enum {{ $e.TableNameSingular }}Fields {
	{{- with .g }}
	{{- range $i, $value := $e.Properties }}
	{{ snake $value.Name }}
	{{- end }}
	{{- end }}
}

{{- $l := GraphQLAggregationMathFields $e }}

type {{ $e.TableNameSingular }}Aggregation {
	count: Int
	distinct(field: {{ $e.TableNameSingular }}Fields):[{{ $e.TableNameSingular }}]
{{- if len $l }}
	min: {{ $e.TableNameSingular }}Min
	max: {{ $e.TableNameSingular }}Max
	sum: {{ $e.TableNameSingular }}Sum
	avg: {{ $e.TableNameSingular }}Avg
{{- end }}
	{{- with .g }}
	{{- range $i, $value := .TypesOptional }}
	{{ . }}
	{{- end }}
	{{- end }}
	{{- if len $e.SQLAssociations }}
	{{- range $i, $a := $e.SQLAssociations }}
	{{ GraphQLAssociationTypeOptional $a }}
	{{- end }}
	{{- end }}
}

{{ if len $l -}}
type {{ $e.TableNameSingular }}Min {
{{- range $i, $p := $l }}
	{{ $p.Name }}: {{ $p.Type }}
{{- end }}
}

type {{ $e.TableNameSingular }}Max {
{{- range $i, $p := $l }}
	{{ $p.Name }}: {{ $p.Type }}
{{- end }}
}

type {{ $e.TableNameSingular }}Sum {
{{- range $i, $p := $l }}
	{{ $p.Name }}: {{ $p.Type }}
{{- end }}
}

type {{ $e.TableNameSingular }}Avg {
{{- range $i, $p := $l }}
	{{ $p.Name }}: Float
{{- end }}
}

{{- end }}
`

const graphqlIndex = `import { Filter, Query } from 'gator-js';
{{- range $i, $col := .Tables }}
import {{ . }} from './{{ . }}';
{{- end}}
export {
	{{- range $i, $col := .Tables }}
	{{ . }},
	{{- end}}
	Filter,
	Query
}
export default function resolve(resolvers, connection) {
	{{- range $i, $col := .Tables }}
	{{ . }}.createQueryResolver(resolvers, connection);
	{{- end}}
}
`

const graphqlResolver = `{{- $e := .Entity -}}
{{- with .Entity -}}
{{- $cl := .SQLColumnList }}
{{- $tn := .SQLTableName }}
{{- $pkc := .PrimaryKey }}
{{- $tnp := .TableNamePlural }}
{{- $tns := .TableNameSingular }}
{{- $columns := .Properties }}
{{- $hpk := .HasPrimaryKey }}
{{- $pkp := .PrimaryKeyProperty }}
{{- $tnt := tick $tn -}}
{{- $amf := GraphQLAggregationMathFields $e }}
import * as helper from './_helper';
import { Filter, Query } from 'gator-js';
{{- range $i, $a := $e.SQLAssociationsUnique }}
{{- if ne $a $e.TableNameSingular }}
import {{ $a }} from './{{ $a }}';
{{- end }}
{{- end}}

const COLUMN_NAMES = [
	{{- $l := len $e.Properties }}
	{{- range $i, $col := $e.Properties }}
	'{{$col.SQLColumnName}}'{{ cond $i $l "," }}
	{{- end }}	
];

const QUERY_PREFIX = 'SELECT {{$cl}} FROM {{$tnt}} ';
const associationBeforeHooks = {};
const associationAfterHooks = {};
const queryBeforeHooks = {};
const queryAfterHooks = {};

/**
 * {{ $e.TableNameSingular }}
 * @class
 */
export default class {{ $e.TableNameSingular }} {
	constructor(props = {}) {
		{{- range $i, $col := $e.Properties }}
		this.{{$col.SQLColumnName}} = {{DefaultJSValue $col}};
		{{- end }}
		Object.keys(props).filter(k => COLUMN_NAMES.indexOf(k) >= 0).forEach(k => this[k] = props[k]);
	}
	/**
	 * install a hook before an association is invoked
	 */
	static hookBeforeAssociation(name, fn) {
		{{- if $e.HasSQLAssociations }}
		switch (name) {
			{{- range $i, $a := $e.SQLAssociations }}
			case '{{ $a.Name }}': {
				const array = associationBeforeHooks[name] || [];
				array.push(fn);
				associationBeforeHooks[name] = array;
				return;
			}
			{{- end }}
		}
		{{- end }}
		throw new Error('no association named ' + name + ' for {{$e.TableNameSingular}}');
	}
	/**
	 * install a hook after an association is returned but before it is returned
	 */
	static hookAfterAssociation(name, fn) {
		{{- if $e.HasSQLAssociations }}
		switch (name) {
			{{- range $i, $a := $e.SQLAssociations }}
			case '{{ $a.Name }}': {
				const array = associationAfterHooks[name] || [];
				array.push(fn);
				associationAfterHooks[name] = array;
				return;
			}
			{{- end }}
		}
		{{- end }}
		throw new Error('no association named ' + name + ' for {{$e.TableNameSingular}}');
	}
	{{- if $e.HasSQLAssociations }}
	{{- range $i, $a := $e.SQLAssociations }}
	static get {{ $a.Name }}Association() {
		return '{{ $a.Name }}';
	}
	{{- end }}
	{{- end }}
	{{- if len $e.Properties }}
	{{- range $i, $col := $e.Properties }}
	{{- if not $col.PrimaryKey }}
	{{- if $col.Index }}
	static get {{ lowerfc $e.TableNamePlural }}By{{$col.Name}}Query() {
		return '{{ lowerfc $e.TableNamePlural }}By{{$col.Name}}';
	}
	{{- end }}
	{{- end }}
	{{- end }} 
	{{- end }}
	static get {{ lowerfc $e.TableNamePlural }}Query() {
		return '{{ lowerfc $e.TableNamePlural }}';
	}
	static get {{ lowerfc $e.TableNameSingular }}Query() {
		return '{{ lowerfc $e.TableNameSingular }}';
	}
	/**
	 * install a hook before a query is invoked
	 */
	static hookBeforeQuery(name, fn) {
		{{- if len $e.Properties }}
		switch (name) {
		{{- range $i, $col := $e.Properties }}
		{{- if not $col.PrimaryKey }}
		{{- if $col.Index }}
			case '{{ lowerfc $e.TableNamePlural }}By{{$col.Name}}':
		{{- end }}
		{{- end }}
		{{- end }} 
			case '{{ lowerfc $e.TableNamePlural }}':
			case '{{ lowerfc $e.TableNameSingular }}': {
				const array = queryBeforeHooks[name] || [];
				array.push(fn);
				queryBeforeHooks[name] = array;
				return;
			}
		}
		{{- end }}
		throw new Error('no query named ' + name + ' for {{$e.TableNameSingular}}');
	}
	/**
	 * install a hook after a query is invoked but before the result is returned
	 */
	static hookAfterQuery(name, fn) {
		switch (name) {
		{{- range $i, $col := $e.Properties }}
		{{- if $col.PrimaryKey }}
			case '{{ lowerfc $e.TableNameSingular }}By{{$col.Name}}':
		{{- else }}
		{{- if $col.Index}}
			case '{{ lowerfc $e.TableNamePlural }}By{{$col.Name}}':
		{{- end }}
		{{- end }}
		{{- end }} 
			case '{{ lowerfc $e.TableNamePlural }}':
			case '{{ lowerfc $e.TableNameSingular }}': {
				const array = queryAfterHooks[name] || [];
				array.push(fn);
				queryAfterHooks[name] = array;
				return;
			}
		}
		throw new Error('no query named ' + name + ' for {{$e.TableNameSingular}}');
	}
	static columns() {
		return COLUMN_NAMES;
	}
	{{- $l := len $e.Properties }}
	{{- range $i, $col := $e.Properties }}
	static get {{$col.SQLColumnName | upcase}}() {
		return '{{$col.SQLColumnName}}';
	}
	{{- end }}	
	static get table() {
		return '{{$tn}}';
	}
	table() {
		return '{{$tn}}';
	}
	static getAssociation(name) {
		{{- if $e.HasSQLAssociations }}
		switch (name) {
			{{- range $i, $a := $e.SQLAssociations }}
			case '{{ $a.Name }}': {
				return {
					primarykey: '{{ $a.PrimaryKey }}',
					foreignkey: '{{ $a.ForeignKey }}',
					name: '{{ $a.Name }}',
					type: '{{ GraphQLAssocationTypeName $a }}',
					{{- if GraphQLAssocationTypeIs $a "belongs_to" }}
					table: '{{ $tn }}',
					ref: {{ $e.TableNameSingular }},
					finder: {{ $e.TableNameSingular }}.findBy{{ SnakeToCamel $a.PrimaryKey }},
					finderString: '{{ $e.TableNameSingular }}.findBy{{ SnakeToCamel $a.PrimaryKey }}',
					finderKey: obj => obj.{{ $a.PrimaryKey }}
					{{- else }}
					{{- if GraphQLAssocationTypeIs $a "has_many" }}
					table: '{{ $a.Table }}',
					ref: {{ SnakeToCamel $a.Table }},
					finder: {{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.PrimaryKey }},
					finderString: '{{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.PrimaryKey }}',
					finderKey: obj => obj instanceof {{ $e.TableNameSingular }} ? obj.{{ $a.ForeignKey }} : obj.{{ $a.PrimaryKey }}
					{{- else }}
					table: '{{ $a.Table }}',
					ref: {{ SnakeToCamel $a.Table }},
					finder: {{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.ForeignKey }},
					finderString: '{{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.ForeignKey }}',
					finderKey: obj => obj instanceof {{ $e.TableNameSingular }} ? obj.{{ $a.PrimaryKey }} : obj.{{ $a.ForeignKey }}
					{{- end}}
					{{- end}}
				};
			}
			{{- end }}
		}
		{{- end }}
	}
	static createQueryResolver(_resolvers, _db) {
		const _cls = this;
		const _invokeBeforeFilters = async (_before, filter, _context, _info) => {
			if (_before && _before.length) {
				for (let _c = 0; _c < _before.length; _c++) {
					filter = await _before[_c](filter, _context, _info);
				}
			}
			return Promise.resolve(filter);
		};
		const _invokeAfterFilters = async (_after, promise, _context, _info) => {
			if (_after && _after.length) {
				let _result = await promise;
				for (let _c = 0; _c < _after.length; _c++) {
					_result = await _after[_c](_result, _context, _info);
				}
				return _result;
			}
			return promise;
		};
		{{- $l := len $e.Properties }}
		{{- range $i, $col := $e.Properties }}
		{{- if $col.PrimaryKey }}
		_resolvers.Query.{{ lowerfc $e.TableNameSingular }}By{{$col.Name}} = (root, { {{$col.SQLColumnName}} }, context, info) => {
			return _invokeAfterFilters(
				queryAfterHooks['{{ lowerfc $e.TableNameSingular }}By{{$col.Name}}'],
				_cls.findBy{{$col.Name}}(context.db || _db, {{$col.SQLColumnName}}, context),
				context,
				info
			);
		};
		{{- else }}
		{{- if $col.Index}}
		_resolvers.Query.{{ lowerfc $e.TableNamePlural }}By{{$col.Name}} = async (root, { {{$col.SQLColumnName}}, filter, offset, limit, sort, sortOrder }, context, info) => {
			filter = helper.filterWithLimit(filter, offset, limit);
			filter = helper.filterWithSort(filter, '{{$tn}}', sort, sortOrder);
			filter = helper.augmentFilter(filter, context, _cls);
			filter = await _invokeBeforeFilters(
				queryBeforeHooks['{{ lowerfc $e.TableNamePlural }}By{{$col.Name}}'],
				filter,
				_context,
				_info
			);
			const cond = Filter.toWherePrepend(filter, '{{$col.SQLColumnName}}', {{$col.SQLColumnName}});
			const q = QUERY_PREFIX + cond.query;
			return _invokeAfterFilters(
				queryAfterHooks['{{ lowerfc $e.TableNamePlural }}By{{$col.Name}}'],
				Query.exec(context.db || _db, q, cond.params, {{$e.TableNameSingular}}, COLUMN_NAMES),
				_context,
				_info
			);
		};
		{{- end }}
		{{- end }}
		{{- end }}
		const associationResolvers = {
			{{- if $e.HasSQLAssociations }}
			{{- range $i, $a := $e.SQLAssociations }}
			{{ $a.Name }}: async function(obj, args, context, info) {
				{{- if $e.HasSQLAssociations }}
				const before = associationBeforeHooks['{{ $a.Name }}'],
					after = associationAfterHooks['{{ $a.Name }}'];
				if ((before && before.length) || (after && after.length)) {
					if (before && before.length) {
						for (let c = 0; c < before.length; c++) {
							obj = await before[c](obj);
						}
					}
					let result = await helper.returnAssociation(_cls.getAssociation('{{$a.Name}}'), context.db || _db, obj, info, context, args);
					if (after && after.length) {
						for (let c = 0; c < after.length; c++) {
							result = await after[c](result);
						}
					}
					return result;
				}
				{{- end }}
				return helper.returnAssociation(_cls.getAssociation('{{$a.Name}}'), context.db || _db, obj, info, context, args);
			},
			{{- end }}
			{{- end }}
		};
		_resolvers.{{ $e.TableNameSingular }} = Object.assign({}, associationResolvers);
		_resolvers.{{ $e.TableNameSingular }}._tablename = o => o.table();
		_resolvers.{{ $e.TableNameSingular }}Optionals = Object.assign({}, associationResolvers);
		_resolvers.{{ $e.TableNameSingular }}Aggregation = Object.assign({
			distinct: function(obj, args, context, info) {
				const result = helper.aggregateResult(info, obj);
				return [result];
			},
 			{{- if len $amf }}
			sum: function(obj, args, context, info) {
				return helper.aggregateResult(info, obj);
			},
			avg: function(obj, args, context, info) {
				return helper.aggregateResult(info, obj);
			},
			min: function(obj, args, context, info) {
				return helper.aggregateResult(info, obj);
			},
			max: function(obj, args, context, info) {
				return helper.aggregateResult(info, obj);
			},
			{{ end }}
		}, associationResolvers);
		_resolvers.Query.{{ lowerfc $e.TableNamePlural }} = async (_, { filter, sort, sortOrder, {{ GraphQLFieldParameters $e true }} }, _context, _info) => {
			filter = helper.filterWithLimit(filter, offset, limit);
			filter = helper.filterWithSort(filter, '{{$tn}}', sort, sortOrder);
			filter = helper.buildFilter(filter, {{ GraphQLFieldParameterArgs $e true }}, '{{$tn}}');
			filter = helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'));
			filter = await _invokeBeforeFilters(
				queryBeforeHooks['{{ lowerfc $e.TableNamePlural }}'],
				filter,
				_context,
				_info
			);
			const _where = Filter.toWhere(filter);
			const _sql = QUERY_PREFIX + _where.query;
			return _invokeAfterFilters(
				queryAfterHooks['{{ lowerfc $e.TableNamePlural }}'],
				Query.exec(_context.db || _db, _sql, _where.params, {{$e.TableNameSingular}}, COLUMN_NAMES),
				_context,
				_info
			);
		};
		_resolvers.Query.{{ lowerfc $e.TableNameSingular }} = async (_, { filter, sort, sortOrder, {{ GraphQLFieldParameters $e false }} }, _context, _info) => {
			filter = helper.filterWithLimit(filter, offset, limit);
			filter = helper.filterWithSort(filter, '{{$tn}}', sort, sortOrder);
			filter = helper.buildFilter(filter, {{ GraphQLFieldParameterArgs $e false }}, '{{$tn}}');
			const _aggQuery = helper.findAggregationQuery(_info, _cls);
			let _sql, _params, _fn;
			if (_aggQuery) {
				let _a = helper.scopeFilter(filter, '{{$tn}}');
				_aggQuery.agg.forEach(_agg => {
					_a = helper.buildAggregationFilter(_a, _agg.name, '{{ $tn }}', _agg.fields, _aggQuery.groups, _aggQuery.fields, _agg.args, COLUMN_NAMES, '{{ $pkc }}');
				});
				if (_a.count) {
					_fn = (_mi, _row) => {_mi.count = _row.count; _mi};
				}
				filter = helper.augmentFilter(_a, _context, _cls)
				filter = await _invokeBeforeFilters(
					queryBeforeHooks['{{ lowerfc $e.TableNameSingular }}'],
					filter,
					_context,
					_info
				);
				const _where = Filter.toWhere(filter);
				_params = _where.params;
				_sql = 'SELECT ' + _a.fields.join(', ') + ' FROM ' + _a.tables.map(_t => Query.escapeId(_t)).join(', ') + ' ' + _where.query;
			} else {
				filter = helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'));
				filter = await _invokeBeforeFilters(
					queryBeforeHooks['{{ lowerfc $e.TableNameSingular }}'],
					filter,
					_context,
					_info
				);
				const _where = Filter.toWhere(filter, _context, _cls);
				_params = _where.params;
				_sql = QUERY_PREFIX + _where.query;
			}
			return _invokeAfterFilters(
				queryAfterHooks['{{ lowerfc $e.TableNameSingular }}'],
				Query.exec(_context.db || _db, _sql, _params, {{$e.TableNameSingular}}, _fn),
				_context,
				_info
			);
		};
	}
	{{- range $i, $col := $e.Properties }}
	get{{$col.Name}}() {
		return this.{{$col.SQLColumnName}};
	}
	set{{$col.Name}}(_{{$col.SQLColumnName}}) {
		this.{{$col.SQLColumnName}} = _{{$col.SQLColumnName}};
		return this;
	}
	{{- if $col.PrimaryKey }}
	static findByPrimaryKey(db, _{{$col.SQLColumnName}}, context) {
		return new Promise (
			async (resolve, reject) => {
				try {
					const filter = helper.buildConditionFilter({ limit : 1 }, '{{$tn}}', '{{$col.SQLColumnName}}', _{{$col.SQLColumnName}});
					const where = Filter.toWhere(helper.augmentFilter(filter, context, {{ $e.TableNameSingular }}));
					const q = QUERY_PREFIX + where.query;
					const r = await Query.exec(context.db || db, q, where.params, {{$e.TableNameSingular}}, COLUMN_NAMES);
					if (r && r.length) {
						resolve(r[0]);
					} else {
						resolve();
					}
				} catch (ex) {
					reject(ex);
				}
			}
		);
	}
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}, context) {
		return {{ $e.TableNameSingular }}.findByPrimaryKey(db, _{{$col.SQLColumnName}}, context);
	}
	{{- else }}
	{{- if $col.Index}}
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}, context, args) {
		const filter = helper.buildArgsFilter('{{$tn}}', args);
		const cond = Filter.toWherePrepend(helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'), context, {{ $e.TableNameSingular }}), '{{$col.SQLColumnName}}', _{{$col.SQLColumnName}});
		const q = QUERY_PREFIX + cond.query;
		return Query.exec(context.db || db, q, cond.params, {{$e.TableNameSingular}}, COLUMN_NAMES);
	}
	{{- end }}
	{{- end }}
	{{- end }}
	static find(db, filter, context, args) {
		let sql = QUERY_PREFIX;
		filter = helper.buildArgsFilter('{{$tn}}', args, filter);
		filter = helper.scopeFilter(filter, '{{$tn}}');
		if (filter.tables) {
			// add additional tables
			const tl = filter.tables.filter(t => t !== '{{$tn}}').map(t => Query.escapeId(t))
			if (tl.length) {
				sql += ',' + tl.join(', ');
			}
		}
		const where = Filter.toWhere(helper.augmentFilter(filter, context, {{ $e.TableNameSingular }}));
		sql += where.query;
		return Query.exec(context.db || db, sql, where.params, {{$e.TableNameSingular}}, COLUMN_NAMES);
	}
	{{- if $e.HasSQLAssociations }}
	{{- range $i, $a := $e.SQLAssociations }}
	{{- if GraphQLAssocationTypeIs $a "belongs_to" }}
	static findBy{{ title $a.Name }}{{ camel $a.ForeignKey }}(_db, _{{ $a.PrimaryKey }}, _context, _args) {
		return new Promise(
			async(_resolve, _reject) => {
				try {
					const _filter = helper.buildArgsFilter('{{$tn}}', _args);
					const _cond = Filter.toWhere(helper.augmentFilter(helper.buildConditionFilter(_filter, '{{$tn}}', '{{$a.PrimaryKey}}', _{{ $a.PrimaryKey }}), _context, {{ $e.TableNameSingular }}));
					const _q = QUERY_PREFIX + _cond.query;
					const _result = await Query.exec(_context.db || _db, _q, _cond.params, {{$e.TableNameSingular}}, COLUMN_NAMES);
					_resolve(_result);
				} catch (_ex) {
					_reject(_ex);
				}
			}
		);
	}
	{{- end }}
	{{- end }}
	{{- end }}
}
{{- end }}
`

const graphqlSchemaTemplate = `
`

const graphqlQueryHelper = `import {Query, Filter} from 'gator-js';

const isValidField = (a) => !/^__/.test(a);
const isAgg = (s) => /^(distinct|min|max|sum|avg|count)$/.test(s);

function fieldScope(table, field) {
	return Query.escapeId(table) + '.' + Query.escapeId(field);
}

function findFields(info) {
	return info.fieldNodes.map(fn => fn.selectionSet.selections.map(s => s.name.value).filter(isValidField))[0];
}

export function buildAggregationFilter(filter, agg, table, fields, grouping, columns, args, all_fields, primary_key) {
	if (!filter) {
		filter = {condition:[]};
	}
	if (!filter.condition) {
		filter.condition = [];
	}
	filter.tables = [table].concat((grouping || []).map(g => g.table));
	const orderbys = {};
	switch (agg) {
		case 'count': {
			filter.fields = ['COUNT(*) as count'];
			if (columns) {
				filter.fields = filter.fields.concat(columns.map(c => fieldScope(c.table, c.field) + ' as ' + Query.escapeId(c.field)));
			}
			filter.count = true;
			break;
		}
		case 'distinct': {
			const fn = args.length ? args[0].value : primary_key;
			filter.fields = ['DISTINCT('+  fieldScope(table, fn) + ') as ' + Query.escapeId(fn)];
			filter.fields = filter.fields.concat(all_fields.filter(a => a !== fn).map(field => fieldScope(table, field) + ' as ' + Query.escape(field)));
			if (filter.order || grouping && grouping.length) {
				grouping = grouping || [];
				grouping.unshift({
					table: table,
					pk: fn,
					fk: fn
				});
				orderbys[fn] = 1;
				all_fields.filter(a => a !== fn).forEach(field => {
					grouping.push({
						table: table,
						pk: field,
						fk: field
					});
					orderbys[field] = 1;
				});
			}
			break;
		}
		default: {
			filter.fields = fields.map(field => agg + '('+ fieldScope(table, field) + ') as ' + Query.escapeId(field));
			filter.fields = filter.fields.concat(columns.map(c => fieldScope(c.table, c.field) + ' as ' + Query.escape(c.field)));
			break;
		}
	}
	if (filter.order && filter.order.length) {
		filter.order.forEach(order => {
			grouping = grouping || [];
			grouping.push({
				table: order.table || table,
				pk: order.field,
				fk: order.field
			});
			orderbys[order.field] = 1;
		});
	}
	if (grouping && grouping.length) {
		const cond = [];
		const groupby = [];
		grouping.forEach(group => {
			if (!orderbys[group.pk]) {
				filter.fields.push(fieldScope(group.table, group.pk) + ' as ' + Query.escapeId(group.fk));
				cond.push({
					table: group.table,
					field: fieldScope(group.table, group.pk) + ' = ' + fieldScope(table, group.fk),
					operator: 'JOIN'
				});
				groupby.push(fieldScope(table, group.fk));
			} else {
				groupby.push(fieldScope(group.table, group.fk));
			}
		});
		if (cond.length) {
			filter.condition.push({conditions:cond});
		}
		if (groupby.length) {
			filter.groupby = groupby.join(', ');
		}
	}
	if (columns && columns.length) {
		filter.groupby = filter.groupby || '';
		filter.groupby += (filter.groupby ? ',' : '') + columns.map(c => fieldScope(c.table, c.field)).join(', ');
	}
	return filter;
}

export function buildConditionFilter(filter = {}, table, field, value, operator = 'EQUAL') {
	if (!filter.condition) {
		filter.condition = [];
	}
	filter.condition.push({
		conditions:[{
			table: table,
			field: field,
			operator: operator,
			value: value
		}]
	});
	return filter;
}

export function buildFilter(filter, args, table) {
	if (args && args.length) {
		args.filter(arg => arg.value !== undefined).forEach(arg => {
			filter = buildConditionFilter(filter, table, arg.name, arg.value);
		});
	}
	return filter;
}

export function findAggregationQuery(info, cls) {
	const ss = info.operation.selectionSet.selections[0].selectionSet;
	if (ss.selections && ss.selections.length) {
		const agg = [];
		const groups = [];
		const fields = [];
		ss.selections.forEach(s => {
			if (isAgg(s.name.value)) {
				agg.push({
					table: cls.table,
					name: s.name.value,
					fields: s.selectionSet && s.selectionSet.selections.filter(sel => isValidField(sel.name.value)).map(sel => sel.name.value),
					args: s.arguments && s.arguments.map(a => ({name:a.name.value, value:a.value.value}))
				});
			} else {
				const assoc = cls.getAssociation(s.name.value);
				if (assoc) {
					groups.push({
						table: assoc.table,
						pk: assoc.type === 'has_many' ? assoc.primarykey : assoc.foreignkey,
						fk: assoc.type === 'has_many' ? assoc.foreignkey : assoc.primarykey,
						fields: s.selectionSet && s.selectionSet.selections.filter(sel => isValidField(sel.name.value)).map(sel => sel.name.value)
					});
				} else if (isValidField(s.name.value)) {
					fields.push({
						table: cls.table,
						field: s.name.value
					});
				}
			}
		});
		if (agg.length) {
			return {agg:agg, groups:groups, fields:fields};
		}
	}
}

export function aggregateResult(info, obj) {
	const fields = findFields(info);
	const result = {};
	fields.forEach(k => result[k] = obj[k]);
	return result;
}

// allow a function to be defined to modify / augment the filter before calling the query
export function augmentFilter(filter, context, cls) {
	if (context && context.filterAugmentation) {
		return context.filterAugmentation(filter, context, cls);
	}
	return filter;
}

export function returnAssociation(assoc, db, obj, info, context, args) {
	if (!assoc.finder) {
		throw new Error('invalid code generation. ' + assoc.finderString + ' is not defined');
	}
	return new Promise(async(resolve, reject) => {
		try {
			const pkvalue = assoc.finderKey(obj);
			if (pkvalue === undefined) {
				console.error('couldn\'t find primary key value');
				return resolve();
			}
			const result = await assoc.finder(context.db || db, pkvalue, context, args);
			if (result && Array.isArray(result) && String(info.returnType).charAt(0) !== '[') {
				return resolve(result && result[0]);
			} else if (result && !Array.isArray(result) && String(info.returnType).charAt(0) === '[') {
				return resolve([result]);
			}
			return resolve(result);
		} catch (ex) {
			console.log(assoc);
			console.error('query failed', ex);
			reject(ex);
		}
	});
}

export function filterWithLimit(filter = {}, offset, limit) {
	if (offset !== undefined) {
		filter.range = {offset:offset, limit:filter.limit || limit || 1000};
	} else if (!filter.range) {
		filter.limit = filter.limit || limit || 1000;
	}
	return filter;
}

export function filterWithSort(filter = {}, table, sort, sortOrder = 'ASCENDING') {
	if (sort) {
		filter.order = filter.order || [];
		filter.order.push({
			table: table,
			field: sort,
			direction: sortOrder
		});
	}
	return filter;
}

const sqlkey = /^(offset|limit)$/;
const isNotSQLKey = key => !sqlkey.test(key);

export function buildArgsFilter(table, args, filter = {}) {
	if (args) {
		filter = filterWithLimit(filter, args.offset, args.limit);
		filter = filterWithSort(filter, table, args.sort, args.sortOrder);
		filter = buildFilter(filter, Object.keys(args).filter(isNotSQLKey).map(k => ({name:k,value:args[k]})), table);
	}
	return filter;
}

export function scopeFilter(filter, table) {
	if (filter && filter.condition && filter.condition.length) {
		filter.condition.forEach(cond => {
			cond.conditions.forEach(c => {
				c.table = c.table || table;
			});
		});
	}
	if (filter && filter.order && filter.order.length) {
		filter.order.forEach(order => {
			order.table = order.table || table;
		});
	}
	return filter;
}

`
