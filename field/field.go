package field

type GeOp uint8
/* Compress option and field types into 32 bit = 4 byte
3 first bit denote represent for 3 option of all fields
1 bit denote represent for unsigned or signed for number fields
27 last bit denote represent for 27 field types to check
 */
const (
	AI GeOp = iota
	NULL
	UNIQUE
)
type FormatField uint32
type Field interface {
	Name(v string) Field
	SetOp(ops ...GeOp) Field // set generic option of 4 ops
}

