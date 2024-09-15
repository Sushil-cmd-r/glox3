package object

import "fmt"

type Bool struct {
	value bool
}

func (b *Bool) Type() ObjType {
	return BoolObj
}

func NewBool(val bool) *Bool {
	return &Bool{value: val}
}

func AsBool(obj Obj) (bool, error) {
	if obj.Type() != BoolObj {
		return false, fmt.Errorf("invalid cast: casting %s to bool", obj.Type())
	}
	return obj.(*Bool).value, nil
}
