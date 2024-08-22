package parser

import (
	"fmt"
	"strings"
)

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}

type ErrorList []*Error

func (e *ErrorList) Len() int {
	return len(*e)
}

func (e *ErrorList) Add(msg string) {
	*e = append(*e, &Error{Msg: msg})
}

func (e *ErrorList) Error() string {
	var sb strings.Builder

	for _, err := range *e {
		sb.WriteString(fmt.Sprintf("%s\n", err))
	}
	return sb.String()
}
