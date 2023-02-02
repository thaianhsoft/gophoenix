package field

type GeOp uint8
const (
	AI GeOp = iota
	NOTNULL
	NULL
	UNIQUE
	_UNSIGNED // only private for internal numeric fields
)

type IntOp uint8

type Field interface {
	Name(v string)
	SetOp(ops ...GeOp) // set generic option of 4 ops

}