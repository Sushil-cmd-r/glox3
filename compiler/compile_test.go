package compiler

import (
	"fmt"
	"testing"

	"github.com/sushil-cmd-r/glox/ast"
	"github.com/sushil-cmd-r/glox/code"
	"github.com/sushil-cmd-r/glox/object"
	"github.com/sushil-cmd-r/glox/parser"
	"github.com/sushil-cmd-r/glox/token"
)

func parseProgram(input string) []ast.Stmt {
	f := token.NewFile("test.glox")
	p := parser.New(f, input)

	return p.Parse()
}

func TestExprStmt(t *testing.T) {
	input := `1; 3.14;
            "hello"
          true; false
          nil`

	expect := []code.Inst{
		{code.OpConstant, 0},
		{code.OpConstant, 1},
		{code.OpConstant, 2},
		{code.OpConstant, 3},
		{code.OpConstant, 4},
		{code.OpConstant, 5},
	}

	prog := parseProgram(input)
	c := Init()
	bc := c.Compile(prog)

	for i, tc := range expect {
		inst := bc.Code[i]
		if inst.Opcode() != tc.Opcode() && inst.Operand() != tc.Operand() {
			t.Fatalf("TestExprStmt [%d]: expected inst %s, got %s", i+1, tc, inst)
		}
	}
}

func TestConstansts(t *testing.T) {
	input := `1; 3.14;
            "hello"
          true; false
          nil`

	expect := []object.ObjType{
		object.NumberObj,
		object.NumberObj,
		object.StringObj,
		object.BoolObj,
		object.BoolObj,
		object.NilObj,
	}

	prog := parseProgram(input)
	c := Init()
	bc := c.Compile(prog)

	for i, tc := range expect {
		obj := bc.Constants[i]
		fmt.Println(obj)
		if obj.Type() != tc {
			t.Fatalf("TestConstansts [%d]: expected objType %s, got %s", i+1, tc, obj.Type())
		}
	}
}
