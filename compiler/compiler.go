package compiler

import (
	"github.com/sushil-cmd-r/glox/ast"
	"github.com/sushil-cmd-r/glox/code"
	"github.com/sushil-cmd-r/glox/object"
)

type Compiler struct {
	code      []code.Inst
	constants []object.Obj
	constCnt  byte
}

func Init() *Compiler {
	return &Compiler{
		code:      make([]code.Inst, 0),
		constants: make([]object.Obj, code.MaxByte),
		constCnt:  0,
	}
}

func (c *Compiler) Compile(prog []ast.Stmt) *code.ByteCode {
	c.compileProgram(prog)
	return &code.ByteCode{Code: c.code, Constants: c.constants}
}

func (c *Compiler) compileProgram(prog []ast.Stmt) {
	for _, stmt := range prog {
		c.compileStmt(stmt)
	}
}

func (c *Compiler) emitSimpleInst(b byte) {
	inst := code.Inst{b, 0}
	c.code = append(c.code, inst)
}

func (c *Compiler) emitOperandInst(b1, b2 byte) {
	inst := code.Inst{b1, b2}
	c.code = append(c.code, inst)
}

func (c *Compiler) addConstant(obj object.Obj) byte {
	c.constants[c.constCnt] = obj
	c.constCnt += 1
	return c.constCnt - 1
}
