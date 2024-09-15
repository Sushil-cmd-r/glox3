package parser

import (
	"fmt"
	"strings"

	"github.com/sushil-cmd-r/glox/token"
)

type Error struct {
	Msg string
	Loc token.Location
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Loc, e.Msg)
}

type ErrorList []*Error

func (e *ErrorList) Len() int {
	return len(*e)
}

func (e *ErrorList) Add(msg string, loc token.Location) {
	*e = append(*e, &Error{Msg: msg, Loc: loc})
}

func (e *ErrorList) Error() string {
	var sb strings.Builder

	for _, err := range *e {
		sb.WriteString(fmt.Sprintf("%s\n", err))
	}
	return sb.String()
}
