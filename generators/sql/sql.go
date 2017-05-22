package sql

import (
	"bytes"

	"github.com/jhaynie/protoc-gen-gator/generator"
	"github.com/jhaynie/protoc-gen-gator/types"
)

type sqlgenerator struct {
}

func init() {
	generator.Register2("sql", &sqlgenerator{})
}

// GenerateSQL returns a SQL statement for the entity
func GenerateSQL(entity *types.Entity) ([]byte, error) {
	return entity.GenerateCode(sqltemplate, nil, nil)
}

func (g *sqlgenerator) Generate(scheme string, file *types.File, entities []types.Entity) ([]*types.Generation, error) {
	var out bytes.Buffer
	for _, entity := range entities {
		// don't generate this table since Goose will automatically create
		if entity.SQLTableName() == "GooseDbVersion" {
			continue
		}
		buf, err := GenerateSQL(&entity)
		if err != nil {
			return nil, err
		}
		out.Write(buf)
	}
	return []*types.Generation{
		&types.Generation{
			Filename: file.Package + "/db.sql",
			Output:   out.String(),
		},
	}, nil
}

const sqltemplate = `{{- with .Entity -}}
{{- if .Comment -}}
-- {{ .Comment -}}
{{- end -}}
CREATE TABLE {{tick .SQLTableName}} (
	{{- $cw := .ColumnWidth }}
	{{- $indexes := .SQLIndexes }}
	{{- $il := len $indexes }}
	{{- $cc := .ColumnCount }}
	{{- $l := add $cc $il }}
	{{- range $i, $col := .Properties }}
	{{ pad $col.SQLColumnNameWithTick $cw }} {{ $col.SQLColumnTypeWithAttributes }}{{- condctx "c" $l "," }}
	{{- addctx "c" 1 -}}
	{{- end }}
	{{ range $i, $col := $indexes -}}
	{{ .Type }} {{ .Name }} ({{.Fields}}){{- condctx "c" $l "," }}
	{{ addctx "c" 1 }}
	{{- end -}}
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
{{- end }}

`
