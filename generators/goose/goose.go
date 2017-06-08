package goose

import (
	"bytes"

	"github.com/jhaynie/protoc-gen-gator/generator"
	"github.com/jhaynie/protoc-gen-gator/generators/sql"
	"github.com/jhaynie/protoc-gen-gator/types"
)

type goosegenerator struct {
}

func init() {
	generator.Register2("goose", &goosegenerator{})
}

func (g *goosegenerator) Generate(scheme string, file *types.File, entities []types.Entity) ([]*types.Generation, error) {
	var create bytes.Buffer
	var drop bytes.Buffer
	for _, entity := range entities {
		if entity.GenerateSQL() {
			buf, err := sql.GenerateSQL(&entity)
			if err != nil {
				return nil, err
			}
			create.Write(buf)
			drop.WriteString("DROP TABLE IF EXISTS `" + entity.SQLTableName() + "`;\n")
		}
	}
	var out bytes.Buffer
	out.WriteString(`-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

`)
	out.Write(create.Bytes())
	out.WriteString(`
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

`)
	out.Write(drop.Bytes())
	return []*types.Generation{
		&types.Generation{
			Filename: file.Package + "/goose_db.sql",
			Output:   out.String(),
		},
	}, nil
}
