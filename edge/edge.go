package edge

type Edge interface{
	Ref(on Edge) Edge
	Unique() Edge
}

type edge struct {
	unique bool
	name string
}

func (e edge) Ref(on Edge) Edge {
	panic("implement me")
}

func (e edge) Unique() Edge {
	panic("implement me")
}


