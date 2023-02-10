package engine

import (
	"gophoenix/edge"
	"gophoenix/field"
)

type Migrator interface{
	Migrate()
}

type SchemaMigrate interface{
	migrate()
	Fields() []field.Field
	Edges() []edge.Edge
}

type Engine struct {

}

func (e *Engine) createTable(sc SchemaMigrate, driverDB string) {

}

func NewEngine() *Engine {

}