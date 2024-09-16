package compiler

import (
	"fmt"

	"github.com/sushil-cmd-r/glox/ast"
	"github.com/sushil-cmd-r/glox/code"
	"github.com/sushil-cmd-r/glox/object"
)

func (c *Compiler) compileExpr(expr ast.Expr) {
	switch expr := expr.(type) {
	case *ast.NumberLit:
		c.compileNumber(expr)
	case *ast.StringLit:
		c.compileString(expr)
	case *ast.BoolExpr:
		c.compileBool(expr)
	case *ast.NilExpr:
		c.compileNil(expr)
	default:

		panic(fmt.Sprintf("unreachable: unknown expression %v", expr))
	}
}

func (c *Compiler) compileNumber(expr *ast.NumberLit) {
	obj := object.NewNumber(expr.Value)
	c.emitOperandInst(code.OpConstant, c.addConstant(obj))
}

func (c *Compiler) compileString(expr *ast.StringLit) {
	obj := object.NewString(expr.Value)
	c.emitOperandInst(code.OpConstant, c.addConstant(obj))
}

func (c *Compiler) compileBool(expr *ast.BoolExpr) {
	obj := object.NewBool(expr.Value)
	c.emitOperandInst(code.OpConstant, c.addConstant(obj))
}

func (c *Compiler) compileNil(_ *ast.NilExpr) {
	obj := object.NewNil()
	c.emitOperandInst(code.OpConstant, c.addConstant(obj))
}
