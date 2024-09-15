package code

type Opcode = byte

const (
	OpReturn Opcode = iota
	OpConstant
	Count
)

var opcodes = [...]string{
	OpReturn:   "OpReturn",
	OpConstant: "OpConstant",
}

func init() {
	if len(opcodes) != int(Count) {
		panic("opcodes array not updated")
	}
}
