package graphql

import (
	"bytes"
	"strings"

	"github.com/jhaynie/protoc-gen-gator/generator"
	"github.com/jhaynie/protoc-gen-gator/types"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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
import { Filter, Query } from 'gator-js';
{{- range $i, $a := $e.SQLAssociationsUnique }}
import {{ $a }} from './{{ $a }}';
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
	static hasAssociation(fieldNodes, name) {
		for (let f = 0; f < fieldNodes.length; f++) {
			const fieldNode = fieldNodes[f];
			if (fieldNode.selectionSet && fieldNode.selectionSet.selections && fieldNode.selectionSet.selections.length) {
				for (let s = 0; s < fieldNode.selectionSet.selections.length; s++) {
					const selection = fieldNode.selectionSet.selections[s];
					if (selection.kind == 'Field' && selection.name.value === name) {
						return selection;
					}
				}
			}
		}
		return false;
	}
	static annotateEntities(db, fieldNodes, results, filter) {
		return new Promise (
			async (resolve, reject) => {
				try {
					{{- if $e.HasSQLAssociations }}
					if (results && !Array.isArray(results)) {
						results = [results];
					}
					if (results && results.length) {
						{{- range $i, $a := $e.SQLAssociations }}
						const has{{ $a.Name }} = this.hasAssociation(fieldNodes, '{{ $a.Name }}'), {{ $a.Name }}indices = [];
						{{- end }}
						if ({{- range $i, $a := $e.SQLAssociations -}} has{{ $a.Name }} || {{ end -}} false) {
							const promises = [];
							results.forEach((r, i) => {
								{{- range $i, $a := $e.SQLAssociations }}
								if (has{{ $a.Name }}) {
									{{if not $a.IsMultiKey -}}
									const v = r.{{ $a.PrimaryKey }};
									if (v !== null && v !== undefined) {
									{{ end -}}
										promises.push(new Promise(
											async (resolve, reject) => {
												try {
													{{if $a.IsMultiKey -}}
													const filter = Filter.toJoinWithParams(null, '{{ $a.PrimaryKey }}', '{{ $a.ForeignKey }}', '{{ $a.Table }}', '{{ $tn }}', r);
													filter.tables = ['{{ $a.Table }}', '{{ $tn }}'];
													const pr = await {{ SnakeToCamel $a.Table }}.find(db, filter);
													{{ else -}}
													const pr = await {{ SnakeToCamel $a.Table }}.findBy{{ SnakeToCamel $a.ForeignKey }}(db, v);
													{{ end -}}
													await {{ SnakeToCamel $a.Table }}.annotateEntities(db, [has{{ $a.Name }}], pr);
													{{- if $a.IsArrayType -}}
													if (Array.isArray(pr)) {
														resolve(pr);
													} else {
														resolve(pr ? [pr] : []);	
													}
													{{ else }}
													if (Array.isArray(pr)) {
														resolve(pr && pr.length && pr[0]);
													} else {
														resolve(pr);
													}
													{{- end }}
												} catch (ex) {
													reject(ex);
												}
											}
										));
									{{if not $a.IsMultiKey -}}
									} else {
										promises.push(Promise.resolve());
									}
									{{ end -}}
									{{ $a.Name }}indices.push({r:i,p:promises.length-1});
								}
								{{- end }}
							});
							if (promises.length) {
								const presults = await Promise.all(promises);
								{{- range $i, $a := $e.SQLAssociations }}
								if (has{{ $a.Name }}) {
									for (let p = 0; p < {{ $a.Name }}indices.length; p++) {
										const e = {{ $a.Name }}indices[p];
										results[e.r].{{ $a.Name }} = presults[e.p];
									}
								}
								{{- end }}
							}
						}
					}
					{{- else }}
					// there are no associations
					{{- end }}
					if (filter && results && results.length) {
						results = filter(results);
					}
					resolve(results);
				} catch (ex) {
					reject(ex);
				}
			}
		);
	}
	static columns() {
		return columnNames;
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
		{{- $l := len $e.Properties }}
		{{- range $i, $col := $e.Properties }}
		{{- if $col.PrimaryKey }}
		resolvers.Query.{{ lowerfc $e.TableNameSingular }}By{{$col.Name}} = async (root, { {{$col.SQLColumnName}} }, context, info) => {
			const _r = await this.findBy{{$col.Name}}(db, {{$col.SQLColumnName}});
			return this.annotateEntities(context.db || db, info.fieldNodes, [_r], res => res[0]);
		};
		{{- else }}
		{{- if $col.Index}}
		resolvers.Query.{{ lowerfc $e.TableNamePlural }}By{{$col.Name}} = async (root, { {{$col.SQLColumnName}}, filter }, context, info) => {
			return this.annotateEntities(context.db || db, info.fieldNodes, await this.findBy{{$col.Name}}(db, {{$col.SQLColumnName}}, filter));
		};
		{{- end }}
		{{- end }}
		{{- end }}
		resolvers.Query.{{ lowerfc $e.TableNamePlural }} = async(root, { filter }, context, info) => {
			return this.annotateEntities(context.db || db, info.fieldNodes, await this.find(db, filter));
		};
		const isValidField = a => !/^__/.test(a);
		const isAggFn = a => /^(distinct|sum|avg|min|max|count)$/.test(a);
		resolvers.Query.{{ lowerfc $e.TableNameSingular }} = async(root, { filter }, context, info) => {
			const detail = {
				columns: {},
				groupby: [],
				joins: [],
				tables: ['{{$tnt}}'],
				operations: [],
				sql: [],
				nonagg: false
			};
			info.fieldNodes.map(n => n.selectionSet).filter(n => n).forEach(node => {
				node.selections.forEach(s => {
					if (!isValidField(s.name.value)) {
						return;
					}
					if (isAggFn(s.name.value)) {
						let fn;
						detail.op = s.name.value;
						if (detail.op == 'count') {
							detail.sql = ['COUNT(*) as count'];
							detail.operations.push((o, r) => {
								o.count = r.count;
								return Promise.resolve();
							});
						} else {
							const fields = s.selectionSet.selections.map(ss => ss.name.value);
							if (s.name.value === 'distinct') {
								detail.nonagg = true;
								const args = s.arguments.map(a => a.value.value);
								fn = args.length ? args[0] : 'id';
								detail.sql = ['DISTINCT({{$tnt}}.' + Query.escapeId(fn) + ') as ' + Query.escapeId(fn)];
								fields.filter(f => isValidField(f) && f !== fn).forEach(f => {
									detail.sql.push('{{$tnt}}.' + Query.escapeId(f));
									detail.columns[f] = 1;
								});
								detail.operations.push((o, r) => {
									o.distinct = o.distinct || [];
									o.distinct.push(r);
									return Promise.resolve();
								});
								detail.columns[fn] = 1;
							} else {
								s.selectionSet.selections.forEach(sel => {
									if (!isValidField(sel.name.value)) {
										return;
									}
									const fn = sel.name.value;
									detail.sql.push(s.name.value.toUpperCase() + '({{$tnt}}.' + Query.escapeId(fn) + ') as ' + Query.escapeId(fn));
									detail.operations.push((o, r) => {
										o[s.name.value] = o[s.name.value] || {};
										o[s.name.value][fn] = r[fn];
										return Promise.resolve();
									});
								});
							}
						}
					} else {
						const a = this.getAssociation(s.name.value);
						detail.nonagg = true;
						if (a) {
							detail.tables.push(Query.escapeId(a.table));
							detail.joins.push('{{$tnt}}.' + Query.escapeId(a.primarykey) + '=' + Query.escapeId(a.table) + '.' + Query.escapeId(a.foreignkey));
							detail.operations.push((o, r) => {
								return new Promise (
									async (resolve, reject) => {
										try {
											const res = await a.finder(context.db || db, r[a.primarykey]);
											const pv = Array.isArray(res) ? res : [res];
											const an = await a.ref.annotateEntities(context.db || db, [s], pv, res => res[0]);
											o[s.name.value] = res;
											resolve();
										} catch (ex) {
											reject(ex);
										}
									}
								);
							});
							detail.sql.push(Query.escapeId(a.table) + '.' + Query.escapeId(a.foreignkey) + ' as ' + Query.escapeId(a.primarykey));
							detail.groupby.push(Query.escapeId(a.table) + '.' + Query.escapeId(a.foreignkey));
						} else {
							detail.sql.push('{{$tnt}}.' + Query.escapeId(s.name.value) + ' as ' + Query.escapeId(s.name.value));
							detail.groupby.push('{{$tnt}}.' + Query.escapeId(s.name.value));
							detail.operations.push((o, r) => {
								o[s.name.value] = r[s.name.value];
								return Promise.resolve();
							});
						}
					}
				});
			});
			if (filter && filter.order && filter.order.length && detail.nonagg) {
				filter.order.forEach(o => {
					if (!(o.field in detail.columns)) {
						detail.sql.push('{{$tnt}}.' + Query.escapeId(o.field));
					}
					detail.groupby.push(Query.escapeId(o.field));
				});
			}
			return this.aggregation(db, detail, filter);
		};
	}
	static aggregation(db, detail, filter) {
		return new Promise (
			async (resolve, reject) => {
				try {
					filter = filter || {};
					filter.table = '{{$tn}}';
					if (detail.op !== 'distinct' && detail.groupby.length === 0) {
						filter.limit = 1;
					}
					const cond = Filter.toWhereConditions(filter, detail.joins, detail.groupby.join(', '));
					const sql = 'SELECT ' + detail.sql.join(', ') + ' FROM ' + detail.tables.join(', ') + ' ' + cond.query;
					db.query(sql, cond.params, async (err, results) => {
						if (err) {
							return reject(err);
						}
						if (results && results.length) {
							const res = [];
							const p = [];
							results.forEach(r => {
								const o = {};
								res.push(o);
								detail.operations.forEach(op => p.push(op(o, r)));
							});
							try {
								await Promise.all(p);
								return resolve(res);
							} catch (ex) {
								reject(ex);
							}
						}
						resolve();
					});
				} catch (ex) {
					reject(ex);
				}
			}
		);
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
	static findByPrimaryKey(db, _{{$col.SQLColumnName}}) {
		return new Promise (
			async (resolve, reject) => {
				try {
					const q = queryPrefix + 'WHERE {{$col.SQLColumnNameWithTick}} = ? LIMIT 1';
					const r = await Query.exec(db, q, [_{{$col.SQLColumnName}}], {{$e.TableNameSingular}}, columnNames);
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
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}) {
		return {{ $e.TableNameSingular }}.findByPrimaryKey(db, _{{$col.SQLColumnName}});
	}
	{{- else }}
	{{- if $col.Index}}
	static findBy{{ $col.Name }}(db, _{{$col.SQLColumnName}}, filter) {
		const cond = Filter.toWherePrepend(filter, '{{$col.SQLColumnName}}', _{{$col.SQLColumnName}});
		const q = queryPrefix + cond.query;
		return Query.exec(db, q, cond.params, {{$e.TableNameSingular}}, columnNames);
	}
	{{- end }}
	{{- end }}
	{{- end }}
	static find(db, filter) {
		let cond;
		let sql = queryPrefix;
		if (filter && filter.query && filter.params) {
			cond = filter;
			if (filter.tables) {
				// add additional tables
				const tl = filter.tables.filter(t => t !== '{{$tn}}').map(t => Query.escapeId(t))
				if (tl.length) {
					sql += ',' + tl.join(', ');
				}
			}
		} else {
			cond = Filter.toWhere(filter);
		}
		sql += cond.query;
		return Query.exec(db, sql, cond.params, {{$e.TableNameSingular}}, columnNames);
	}
}
{{- end }}
`

const graphqlSchemaTemplate = `
`
