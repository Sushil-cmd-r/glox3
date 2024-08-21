package ast

import (
	"fmt"
	"strconv"

	"github.com/sushil-cmd-r/glox/token"
)

type Expr interface {
	exprNode()
}

type (
	BadExpr struct{}

	BinaryExpr struct {
		Op    token.Token
		Left  Expr
		Right Expr
	}

	UnaryExpr struct {
		Op   token.Token
		Left Expr
	}

	GroupExpr struct {
		Expression Expr
	}

	NumberLit struct {
		Value float64
	}

	StringLit struct {
		Value string
	}

	IdentExpr struct {
		Name string
	}

	NilExpr struct{}

	BoolExpr struct {
		Value bool
	}
)

func (*BadExpr) exprNode()    {}
func (*BinaryExpr) exprNode() {}
func (*UnaryExpr) exprNode()  {}
func (*GroupExpr) exprNode()  {}
func (*NumberLit) exprNode()  {}
func (*StringLit) exprNode()  {}
func (*IdentExpr) exprNode()  {}
func (*NilExpr) exprNode()    {}
func (*BoolExpr) exprNode()   {}

func (*BadExpr) String() string {
	return fmt.Sprint("BadExpr:{}")
}

func (b *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", b.Op, b.Left, b.Right)
}

func (u *UnaryExpr) String() string {
	return fmt.Sprintf("(%s %s)", u.Op, u.Left)
}

func (g *GroupExpr) String() string {
	return fmt.Sprintf("%s", g.Expression)
}

func (n *NumberLit) String() string {
	return strconv.FormatFloat(n.Value, 'g', -1, 64)
}

func (s *StringLit) String() string {
	return s.Value
}

func (i *IdentExpr) String() string {
	return i.Name
}

func (*NilExpr) String() string {
	return "<nil>"
}

func (b *BoolExpr) String() string {
	return fmt.Sprintf("%v", b.Value)
}
