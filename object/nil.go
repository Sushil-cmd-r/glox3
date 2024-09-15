package object

type Nil struct{}

func (n *Nil) Type() ObjType {
	return NilObj
}

func NewNil() *Nil {
	return &Nil{}
}
