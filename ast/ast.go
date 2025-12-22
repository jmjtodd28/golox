package ast

import "github.com/jmjtodd28/golox/token"

type Expr interface {
	Print() string
}

type BinaryExpr struct {
	left, right Expr
	operator    token.Token
}

func NewBinaryExpr(left Expr, right Expr, operator token.Token) Expr {
	return &BinaryExpr{
		left:     left,
		right:    right,
		operator: operator,
	}
}

func (b *BinaryExpr) Print() string {
	return "(" + b.operator.Lexeme + " " + b.left.Print() + " " + b.right.Print() + ")"
}

type Grouping struct {
	expression Expr
}

func NewGrouping(expression Expr) Expr {
	return &Grouping{
		expression: expression,
	}
}

func (g *Grouping) Print() string {
	return "(group " + g.expression.Print() + ")"
}

type Literal struct {
	token token.Token
}

func NewLiteral(token token.Token) Expr {
	return &Literal{
		token: token,
	}
}

func (l *Literal) Print() string {
	return l.token.Lexeme
}

type Unary struct {
	operator   token.Token
	expression Expr
}

func NewUnary(value token.Token, expression Expr) Expr {
	return &Unary{
		operator:   value,
		expression: expression,
	}
}

func (u *Unary) Print() string {
	return "(" + u.operator.Lexeme + " " + u.expression.Print() + ")"
}
