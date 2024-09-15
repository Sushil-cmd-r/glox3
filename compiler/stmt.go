package compiler

import "github.com/sushil-cmd-r/glox/ast"

func (c *Compiler) compileStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		c.compileExprStmt(stmt)
	default:
		panic("unknown stmt")
	}
}

func (c *Compiler) compileExprStmt(stmt *ast.ExprStmt) {
	c.compileExpr(stmt.Expression)
}
