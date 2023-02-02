package field

type intField struct {
	name string
	ops GeOp
	un bool // unsigned
}

func (i *intField) Name(v string) Field {
	i.name = v
	return i
}

func (i *intField) SetOp(ops ...GeOp) Field {
	for _, op := range ops {
		i.ops |= 1 << op
	}
	return i
}

