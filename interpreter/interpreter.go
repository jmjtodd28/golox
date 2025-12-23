package interpreter

import (
	"fmt"
	"strconv"

	"github.com/jmjtodd28/golox/ast"
	"github.com/jmjtodd28/golox/token"
)

type Interpreter struct {
}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i *Interpreter) Evaluate(expr ast.Expr) any {

	switch expr := expr.(type) {
	case *ast.Literal:
		return i.evalLiteral(expr)
	case *ast.BinaryExpr:
		return i.evalBinaryExpr(expr)
	}

	return nil
}

func (i *Interpreter) evalLiteral(expr *ast.Literal) any {
	tok := expr.Token
	switch expr.Token.TokenType {
	case token.NUMBER:
		val, err := strconv.ParseFloat(tok.Lexeme, 64)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error evaluating number token: %s", err))
		}
		return val
	}

	return nil
}

func (i *Interpreter) evalBinaryExpr(expr *ast.BinaryExpr) any {

	left := i.Evaluate(expr.Left)
	right := i.Evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case token.PLUS:
		return i.applyBinaryOperator(token.PLUS, left, right)
	}
	return nil
}

func (i *Interpreter) applyBinaryOperator(operator token.Type, left any, right any) any {
	switch operator {
	case token.PLUS:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum + rightNum
			}
		}

		panic("Cannot add two non-numbers")
	}

	return nil
}
