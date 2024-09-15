package code

import "fmt"

type Inst [2]byte

func (i Inst) Opcode() byte {
	return i[0]
}

func (i Inst) Operand() byte {
	return i[1]
}

func (i Inst) String() string {
	op := i.Opcode()
	switch op {
	case OpReturn:
		return simpleInst(op)
	case OpConstant:
		operand := i.Operand()
		return operandInst(op, operand)
	default:
		panic(fmt.Sprintf("unreachable: unknown opcode %d", op))
	}
}

func simpleInst(op byte) string {
	return opcodes[op]
}

func operandInst(op byte, operand byte) string {
	return fmt.Sprintf("%s %d", opcodes[op], operand)
}
