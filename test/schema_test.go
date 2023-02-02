package test

import (
	"gophoenix/edge"
	"gophoenix/field"
	"gophoenix/schema"
)

type Person struct {
	*schema.Schema
}

func (p *Person) Field() []field.Field {
	return []field.Field{

	}
}

func (p *Person) Edge() []edge.Edge{
	return []edge.Edge{

	}
}



