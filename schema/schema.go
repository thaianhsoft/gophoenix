package schema

import (
	"gophoenix/edge"
	"gophoenix/field"
)

type SchemaMigrate interface{
	migrate()
	Field() []field.Field
	Edge() []edge.Edge
}

type Schema struct {

}


func (s *Schema) migrate() {

}