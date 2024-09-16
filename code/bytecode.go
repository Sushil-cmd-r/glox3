package code

import "github.com/sushil-cmd-r/glox/object"

type ByteCode struct {
	Code      []Inst
	Constants []object.Obj
}
