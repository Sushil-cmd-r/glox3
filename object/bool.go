package object

type Bool struct {
	value bool
}

func (b *Bool) Type() ObjType {
	return BoolObj
}

func NewBool(val bool) *Bool {
	return &Bool{value: val}
}
