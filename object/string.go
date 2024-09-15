package object

import "fmt"

type String struct {
	value string
}

func (s *String) Type() ObjType {
	return StringObj
}

func NewString(val string) *String {
	return &String{value: val}
}

func AsString(obj Obj) (string, error) {
	if obj.Type() != StringObj {
		return "", fmt.Errorf("invalid cast: casting %s to string", obj.Type())
	}
	return obj.(*String).value, nil
}
