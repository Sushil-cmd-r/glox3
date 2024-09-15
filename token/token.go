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

	keywordStart
	TRUE  // true
	FALSE // false
	NIL   // nil
	keywordEnd
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

	TRUE:  "true",
	FALSE: "false",
	NIL:   "nil",
}

func (tok Token) String() string {
	return tokens[tok]
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keywordEnd-keywordStart+1)
	for i := keywordStart + 1; i < keywordEnd; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) (Token, string) {
	if keyword, ok := keywords[ident]; ok {
		return keyword, ident
	}
	return IDENTIFIER, ident
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
