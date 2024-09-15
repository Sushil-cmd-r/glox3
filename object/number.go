package object

import "fmt"

type Number struct {
	value float64
}

func (n *Number) Type() ObjType {
	return NumberObj
}

func NewNumber(val float64) *Number {
	return &Number{value: val}
}

func AsNumber(obj Obj) (float64, error) {
	if obj.Type() != NumberObj {
		return -1, fmt.Errorf("invalid cast: casting %s to number", obj.Type())
	}
	return obj.(*Number).value, nil
}
