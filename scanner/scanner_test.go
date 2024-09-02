package scanner

import (
	"testing"

	"github.com/sushil-cmd-r/glox/token"
)

func makeScanner(input string) *Scanner {
	f := token.NewFile("test.glox")
	return Init(f, []byte(input))
}

func TestLoc(t *testing.T) {
	input := "hello 345\n\t\txyz\n   \t\t+ ="

	expect := []struct {
		loc      int
		location string
	}{
		{0, "test.glox:1:1"},
		{6, "test.glox:1:7"},
		{9, "test.glox:1:10"},
		{12, "test.glox:2:3"},
		{15, "test.glox:2:6"},
		{21, "test.glox:3:6"},
		{23, "test.glox:3:8"},
		{24, "test.glox:3:9"},
	}

	sc := makeScanner(input)
	for i, tc := range expect {
		_, tok, loc := sc.Scan()

		if tc.loc != int(loc) {
			t.Fatalf("TestLoc[%d]:loc failed for %s expected %d, got %d", i+1, tok, tc.loc, loc)
		}

		location := sc.file.LocationFor(loc).String()
		if tc.location != location {
			t.Fatalf("TestLoc[%d]:location failed for %s expected %s, got %s", i+1, tok, tc.location, location)
		}
	}
}

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
		{token.SEMI, ";"},
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
		{token.SEMI, ";"},
		{token.EOF, "eof"},
	}

	s := makeScanner(input)
	for i, tc := range expect {
		tok, lit, _ := s.Scan()
		if tok != tc.tok {
			t.Fatalf("TestScan [%d]: expected tok %s, got %s", i+1, tc.tok, tok)
		}
		if lit != tc.lit {
			t.Fatalf("TestScan [%d]: expected lit %s, got %s", i+1, tc.lit, lit)
		}
	}
}
