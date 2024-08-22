package parser

import (
	"fmt"
	"testing"

	"github.com/sushil-cmd-r/glox/token"
)

func TestParseExpr(t *testing.T) {
	input := `( -x + 3.14) / "hello"`

	expect := "(/ (+ (- x) 3.14) hello)"

	p := New(input)
	expr := p.parseExpr(token.PrecLowest)
	if expr == nil {
		t.Fatal("TestParseExpr: expected expression")
	}

	if p.errors.Len() > 0 {
		t.Fatal(p.errors)
	}

	exprStr := fmt.Sprintf("%s", expr)

	if expect != exprStr {
		t.Fatalf("TestParseExpr: expected expr %s, got %s", expect, exprStr)
	}
}
