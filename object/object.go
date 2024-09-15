package object

type Obj interface {
	Type() ObjType
}

type ObjType int

const (
	NumberObj ObjType = iota
	StringObj
	BoolObj
	NilObj
)

var objTypes = [...]string{
	NumberObj: "number",
	StringObj: "string",
	BoolObj:   "bool",
	NilObj:    "<nil>",
}

func (ot ObjType) String() string {
	return objTypes[ot]
}
