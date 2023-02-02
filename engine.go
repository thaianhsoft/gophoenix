package engine

import "gophoenix/field"

type migrator interface{
	migrate()
}

func (e *engine) createTable(field field.Field) {

}