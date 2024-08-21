package token

type Token int

const (
	ILLEGAL Token = iota // illegal
	EOF                  // eof

	NUMBER     // number
	IDENTIFIER // identifier
	STRING     // string

	ASSIGN // =
	LPAREN // (
	RPAREN // )
	LCURLY // {
	RCURLY // }
	COMMA  // ,
	SEMI   // ;
	NOT    // !

	PLUS  // +
	MINUS // -
	STAR  // *
	SLASH // /
	EQL   // ==
	NEQ   // !=
	GTR   // >
	LSS   // <
	GEQ   // >=
	LEQ   // <=
)

var tokens = [...]string{
	ILLEGAL: "illegal",
	EOF:     "eof",

	NUMBER:     "number",
	IDENTIFIER: "identifier",
	STRING:     "string",

	ASSIGN: "=",
	LPAREN: "(",
	RPAREN: ")",
	LCURLY: "{",
	RCURLY: "}",
	COMMA:  ",",
	SEMI:   ";",
	NOT:    "!",

	PLUS:  "+",
	MINUS: "-",
	STAR:  "*",
	SLASH: "/",
	EQL:   "==",
	NEQ:   "!=",
	GTR:   ">",
	LSS:   "<",
	GEQ:   ">=",
	LEQ:   "<=",
}

func (tok Token) String() string {
	return tokens[tok]
}

const (
	PrecLowest = iota
	PrecEquality
	PrecComparison
	PrecTerm
	PrecFactor
	PrecUnary
)

func (tok Token) Precedence() int {
	switch tok {
	case EQL, NEQ:
		return PrecEquality
	case GTR, LSS, GEQ, LEQ:
		return PrecComparison
	case PLUS, MINUS:
		return PrecTerm
	case STAR, SLASH:
		return PrecFactor
	default:
		return PrecLowest
	}
}
