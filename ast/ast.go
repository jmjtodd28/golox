package ast

import "github.com/jmjtodd28/golox/token"

type Expr interface {
	Print() string
}

type BinaryExpr struct {
	Left, Right Expr
	Operator    token.Token
}

func NewBinaryExpr(left Expr, right Expr, operator token.Token) Expr {
	return &BinaryExpr{
		Left:     left,
		Right:    right,
		Operator: operator,
	}
}

func (b *BinaryExpr) Print() string {
	return "(" + b.Operator.Lexeme + " " + b.Left.Print() + " " + b.Right.Print() + ")"
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) Expr {
	return &Grouping{
		Expression: expression,
	}
}

func (g *Grouping) Print() string {
	return "(group " + g.Expression.Print() + ")"
}

type Literal struct {
	Token token.Token
}

func NewLiteral(token token.Token) Expr {
	return &Literal{
		Token: token,
	}
}

func (l *Literal) Print() string {
	return l.Token.Lexeme
}

type Unary struct {
	Operator   token.Token
	Expression Expr
}

func NewUnary(value token.Token, expression Expr) Expr {
	return &Unary{
		Operator:   value,
		Expression: expression,
	}
}

func (u *Unary) Print() string {
	return "(" + u.Operator.Lexeme + " " + u.Expression.Print() + ")"
}
