package compiler

import (
	"github.com/sushil-cmd-r/glox/ast"
	"github.com/sushil-cmd-r/glox/code"
)

type Compiler struct {
	code []code.Inst
}

func (c *Compiler) Compile(prog []ast.Stmt) *code.ByteCode {
	c.compileProgram(prog)
	return &code.ByteCode{Code: c.code}
}

func (c *Compiler) compileProgram(prog []ast.Stmt) {
	for _, stmt := range prog {
		c.compileStmt(stmt)
	}
}
