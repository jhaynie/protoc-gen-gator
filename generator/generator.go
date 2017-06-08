package generator

import (
	"fmt"
	"sort"

	"github.com/jhaynie/protoc-gen-gator/types"
)

var generators = make(map[string]types.Generator)
var generators2 = make(map[string]types.Generator2)

// Register a generator
func Register(name string, g types.Generator) {
	generators[name] = g
}

// Register2 a generator
func Register2(name string, g types.Generator2) {
	generators2[name] = g
}

// GetAllTypes returns an array of all the register types
func GetAllTypes() []string {
	types := make([]string, 0)
	for k := range generators {
		types = append(types, k)
	}
	return types
}

// Generate will generate output for a specific type
func Generate(gentype []string, file *types.File) ([]*types.Generation, error) {
	entities := make([]types.Entity, 0)
	for _, message := range file.Messages {
		e := types.NewEntity(file.Package, file, message)
		entities = append(entities, e)
	}
	sort.Slice(entities, func(i, j int) bool { return entities[i].Name < entities[j].Name })
	results := make([]*types.Generation, 0)
	for _, t := range gentype {
		generator := generators2[t]
		if generator == nil {
			return nil, fmt.Errorf("no generator found for '%s'", t)
		}
		r, err := generator.Generate(t, file, entities)
		if err != nil {
			return nil, err
		}
		if r != nil {
			for _, result := range r {
				results = append(results, result)
			}
		}
	}
	return results, nil
}
