package graphql

import (
	"bytes"
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
			t = "Int"
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

func toGraphQLAssocationTypeName(a *types.SQLAssociation) string {
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

func toGraphQLAssocationType(a *types.SQLAssociation) string {
	switch a.Type {
	case types.SQLAssocationHasMany:
		{
			return a.Name + ": [" + snaker.SnakeToCamel(a.Table) + "]"
		}
	case types.SQLAssocationBelongsTo, types.SQLAssocationHasOne:
		{
			return a.Name + ": " + snaker.SnakeToCamel(a.Table)
		}
	}
	return ""
}

func toGraphQLAssocationTypeOptional(a *types.SQLAssociation) string {
	switch a.Type {
	case types.SQLAssocationHasMany:
		{
			return a.Name + ": [" + snaker.SnakeToCamel(a.Table) + "Optionals]"
		}
	case types.SQLAssocationBelongsTo, types.SQLAssocationHasOne:
		{
			return a.Name + ": " + snaker.SnakeToCamel(a.Table) + "Optionals"
		}
	}
	return ""
}

func (g *gqlgenerator) Generate(scheme string, file *types.File, entities []types.Entity) ([]*types.Generation, error) {
	results := make([]*types.Generation, 0)
	fn := make(map[string]interface{})
	fn["GraphQLType"] = toGraphQLType
	fn["GraphQLTypeOptional"] = toGraphQLTypeOptional
	fn["GraphQLVariable"] = toGraphQLVariableName
	fn["GraphQLTypeEnumDefinition"] = toGraphQLTypeEnumDefinitions
	fn["DefaultJSValue"] = toDefaultJSValue
	fn["GraphQLAssociationType"] = toGraphQLAssocationType
	fn["GraphQLAssocationTypeName"] = toGraphQLAssocationTypeName
	fn["GraphQLAssociationTypeOptional"] = toGraphQLAssocationTypeOptional
	fn["GraphQLAggregationMathFields"] = toGraphQLAggregationMathFields
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

`

const graphqlRootQueryTemplate = `{{- $e := .Entity }}
	{{ lowerfc $e.TableNameSingular }}(filter: QueryFilter):[{{ $e.TableNameSingular -}} Aggregation]
	{{ lowerfc $e.TableNamePlural }}(filter: QueryFilter):[{{ $e.TableNameSingular }}]
	{{ lowerfc $e.TableNameSingular }}By {{- $e.PrimaryKeyProperty.Field.Name}}( {{- GraphQLVariable $e.PrimaryKeyProperty -}}: {{ GraphQLType $e.PrimaryKeyProperty -}}):{{ $e.TableNameSingular }}
	{{- range $i, $value := $e.Properties }}
	{{- if .Index }}
	{{ lowerfc $e.TableNamePlural }}By {{- $value.Field.Name}}(  {{- GraphQLVariable . -}}: {{GraphQLType . -}}, filter: QueryFilter):[{{ $e.TableNameSingular }}]
	{{- end }}
	{{- end }}
`

const graphqlTemplate = `{{- $e := .Entity -}}
{{ GraphQLTypeEnumDefinition $e }}
type {{ $e.TableNameSingular }} {
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

const columnNames = [
	{{- $l := len $e.Properties }}
	{{- range $i, $col := $e.Properties }}
	'{{$col.SQLColumnName}}'{{ cond $i $l "," }}
	{{- end }}	
];

const queryPrefix = 'SELECT {{$cl}} FROM {{$tnt}} ';

/**
 * {{ $e.TableNameSingular }}
 * @class
 */
export default class {{ $e.TableNameSingular }} {
	constructor(props = {}) {
		{{- range $i, $col := $e.Properties }}
		this.{{$col.SQLColumnName}} = {{DefaultJSValue $col}};
		{{- end }}
		Object.keys(props).filter(k => columnNames.indexOf(k) >= 0).forEach(k => this[k] = props[k]);
	}
	static columns() {
		return columnNames;
	}
	static table() {
		return '{{$tn}}';
	}
	static getAssociation(name) {
		{{- if $e.HasSQLAssociations }}
		switch (name) {
			{{- range $i, $a := $e.SQLAssociations }}
			case '{{ $a.Name }}': {
				return {
					ref: {{ SnakeToCamel $a.Table }},
					table: '{{ $a.Table }}',
					primarykey: '{{ $a.PrimaryKey }}',
					foreignkey: '{{ $a.ForeignKey }}',
					name: '{{ $a.Name }}',
					type: '{{ GraphQLAssocationTypeName $a }}',
					finder: {{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.ForeignKey }}
				};
			}
			{{- end }}
		}
		{{- end }}
	}
	static createQueryResolver(resolvers, db) {
		const cls = this;
		{{- $l := len $e.Properties }}
		{{- range $i, $col := $e.Properties }}
		{{- if $col.PrimaryKey }}
		resolvers.Query.{{ lowerfc $e.TableNameSingular }}By{{$col.Name}} = (root, { {{$col.SQLColumnName}} }, context, info) => {
			return cls.findBy{{$col.Name}}(db, {{$col.SQLColumnName}}, context);
		};
		{{- else }}
		{{- if $col.Index}}
		resolvers.Query.{{ lowerfc $e.TableNamePlural }}By{{$col.Name}} = (root, { {{$col.SQLColumnName}}, filter }, context, info) => {
			const cond = Filter.toWherePrepend(helper.augmentFilter(filter, context, cls), '{{$col.SQLColumnName}}', {{$col.SQLColumnName}});
			const q = queryPrefix + cond.query;
			return Query.exec(db, q, cond.params, {{$e.TableNameSingular}}, columnNames);
		};
		{{- end }}
		{{- end }}
		{{- end }}
		const associationResolvers = {
			{{- if $e.HasSQLAssociations }}
			{{- range $i, $a := $e.SQLAssociations }}
			{{ $a.Name }}: async function(obj, args, context, info) {
				return helper.returnAssociation(cls.getAssociation('{{$a.Name}}').finder, context.db || db, obj.{{$a.PrimaryKey}}, info, context);
			},
			{{- end }}
			{{- end }}
		};
		resolvers.{{ $e.TableNameSingular }} = Object.assign({}, associationResolvers);
		resolvers.{{ $e.TableNameSingular }}Optionals = Object.assign({}, associationResolvers);
		resolvers.{{ $e.TableNameSingular }}Aggregation = Object.assign({
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
		resolvers.Query.{{ lowerfc $e.TableNamePlural }} = (root, { filter }, context, info) => {
			const where = Filter.toWhere(helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'), context, cls));
			const sql = queryPrefix + where.query;
			return Query.exec(context.db || db, sql, where.params, {{$e.TableNameSingular}}, columnNames);
		};
		resolvers.Query.{{ lowerfc $e.TableNameSingular }} = async (root, { filter }, context, info) => {
			const aggQuery = helper.findAggregationQuery(info, cls);
			let sql, params, fn;
			if (aggQuery) {
				let a = helper.scopeFilter(filter, '{{$tn}}');
				aggQuery.agg.forEach(agg => {
					a = helper.buildAggregationFilter(a, agg.name, '{{ $tn }}', agg.fields, aggQuery.groups, aggQuery.fields, agg.args, columnNames, '{{ $pkc }}');
				});
				if (a.count) {
					fn = (mi, row) => {mi.count = row.count; mi};
				}
				const where = Filter.toWhere(helper.augmentFilter(a, context, cls));
				params = where.params;
				sql = 'SELECT ' + a.fields.join(', ') + ' FROM ' + a.tables.map(t => Query.escapeId(t)).join(', ') + ' ' + where.query;
			} else {
				const where = Filter.toWhere(helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'), context, cls));
				params = where.params;
				sql = queryPrefix + where.query;
			}
			return Query.exec(context.db || db, sql, params, {{$e.TableNameSingular}}, fn);
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
	static findByPrimaryKey(db, _{{$col.SQLColumnName}}, filter, context) {
		return new Promise (
			async (resolve, reject) => {
				try {
					filter = filter || {};
					filter.limit = 1;
					filter.condition = filter.condition || [];
					filter.condition.push({
						conditions: [{
							table: '{{$tn}}',
							field: '{{$col.SQLColumnName}}',
							operator: 'EQUAL',
							value: _{{$col.SQLColumnName}}
						}]
					});
					const where = Filter.toWhere(helper.augmentFilter(filter, context, {{ $e.TableNameSingular }}));
					const q = queryPrefix + where.query;
					const r = await Query.exec(db, q, where.params, {{$e.TableNameSingular}}, columnNames);
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
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}, filter, context) {
		return {{ $e.TableNameSingular }}.findByPrimaryKey(db, _{{$col.SQLColumnName}}, filter, context);
	}
	{{- else }}
	{{- if $col.Index}}
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}, filter, context) {
		const cond = Filter.toWherePrepend(helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'), context, {{ $e.TableNameSingular }}), '{{$col.SQLColumnName}}', _{{$col.SQLColumnName}});
		const q = queryPrefix + cond.query;
		return Query.exec(db, q, cond.params, {{$e.TableNameSingular}}, columnNames);
	}
	{{- end }}
	{{- end }}
	{{- end }}
	static find(db, filter, context) {
		let cond;
		let sql = queryPrefix;
		if (filter && filter.query && filter.params) {
			cond = helper.scopeFilter(filter, '{{$tn}}');
			if (filter.tables) {
				// add additional tables
				const tl = filter.tables.filter(t => t !== '{{$tn}}').map(t => Query.escapeId(t))
				if (tl.length) {
					sql += ',' + tl.join(', ');
				}
			}
		} else {
			cond = Filter.toWhere(helper.augmentFilter(helper.scopeFilter(filter, '{{$tn}}'), context, {{ $e.TableNameSingular }}));
		}
		sql += cond.query;
		return Query.exec(db, sql, cond.params, {{$e.TableNameSingular}}, columnNames);
	}
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
			const fn = args.length ? args[0] : primary_key;
			filter.fields = ['DISTINCT('+  fieldScope(table, fn) + ') as ' + Query.escapeId(fn)];
			filter.fields = filter.fields.concat(all_fields.filter(a => a !== fn).map(field => fieldScope(table, field) + ' as ' + Query.escape(field)));
			break;
		}
		default: {
			filter.fields = fields.map(field => agg + '('+ fieldScope(table, field) + ') as ' + Query.escapeId(field));
			filter.fields = filter.fields.concat(columns.map(c => fieldScope(c.table, c.field) + ' as ' + Query.escape(c.field)));
			break;
		}
	}
	if (grouping && grouping.length) {
		const cond = [];
		const groupby = [];
		grouping.forEach(group => {
			filter.fields.push(fieldScope(group.table, group.pk) + ' as ' + Query.escapeId(group.fk));
			cond.push({
				table: group.table,
				field: fieldScope(group.table, group.pk) + ' = ' + fieldScope(table, group.fk),
				operator: 'JOIN'
			});
			groupby.push(fieldScope(table, group.fk));
		});
		filter.condition.push({conditions:cond});
		filter.groupby = groupby.join(', ');
	}
	if (columns && columns.length) {
		filter.groupby = filter.groupby || '';
		filter.groupby += (filter.groupby ? ',' : '') + columns.map(c => fieldScope(c.table, c.field)).join(', ');
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
					table: cls.table(),
					name: s.name.value,
					fields: s.selectionSet && s.selectionSet.selections.filter(sel => isValidField(sel.name.value)).map(sel => sel.name.value),
					args: s.arguments && s.arguments.map(a => a.value.value)
				});
			} else {
				const assoc = cls.getAssociation(s.name.value);
				if (assoc) {
					groups.push({
						table: assoc.table,
						pk: assoc.foreignkey,
						fk: assoc.primarykey,
						fields: s.selectionSet && s.selectionSet.selections.filter(sel => isValidField(sel.name.value)).map(sel => sel.name.value)
					});
				} else if (isValidField(s.name.value)) {
					fields.push({
						table: cls.table(),
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

export async function returnAssociation(finder, db, pk, info, context) {
	return await finder(db, pk, null, context);
}

export function scopeFilter(filter, table) {
	if (filter && filter.condition && filter.condition.length) {
		filter.condition.forEach(cond => {
			cond.conditions.forEach(c => {
				c.table = c.table || table;
			});
		});
	}
	return filter;
}

`
