package scanner

import (
	"testing"

	"github.com/sushil-cmd-r/glox/token"
)

func TestScan(t *testing.T) {
	input := `12 3.14 "hello""a" "" wor1d 123abc  , ;&  x
  + =/   -* ! != >< >= <= "unterminated `

	expect := []struct {
		tok token.Token
		lit string
	}{
		{token.NUMBER, "12"},
		{token.NUMBER, "3.14"},
		{token.STRING, "hello"},
		{token.STRING, "a"},
		{token.STRING, ""},
		{token.IDENTIFIER, "wor1d"},
		{token.ILLEGAL, "invalid number"},
		{token.COMMA, ","},
		{token.SEMI, ";"},
		{token.ILLEGAL, "&"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.SLASH, "/"},
		{token.MINUS, "-"},
		{token.STAR, "*"},
		{token.NOT, "!"},
		{token.NEQ, "!="},
		{token.GTR, ">"},
		{token.LSS, "<"},
		{token.GEQ, ">="},
		{token.LEQ, "<="},
		{token.ILLEGAL, "unterminated string"},
		{token.EOF, "eof"},
	}

	s := Init([]byte(input))
	for i, tc := range expect {
		tok, lit := s.Scan()
		if tok != tc.tok {
			t.Fatalf("TestScan [%d]: expected tok %s, got %s", i+1, tc.tok, tok)
		}
		if lit != tc.lit {
			t.Fatalf("TestScan [%d]: expected lit %s, got %s", i+1, tc.lit, lit)
		}
	}
}
