package interpreter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmjtodd28/golox/ast"
	"github.com/jmjtodd28/golox/token"
)

type Interpreter struct {
}

func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (i *Interpreter) EvaluateExpr(expr ast.Expr) any {

	switch expr := expr.(type) {
	case *ast.Literal:
		return i.evalLiteralExpr(expr)
	case *ast.BinaryExpr:
		return i.evalBinaryExpr(expr)
	case *ast.Grouping:
		return i.evalGroupingExpr(expr)
	case *ast.Unary:
		return i.evalUnaryExpr(expr)
	}

	return nil
}

func (i *Interpreter) evalLiteralExpr(expr *ast.Literal) any {
	tok := expr.Token
	switch expr.Token.TokenType {
	case token.NUMBER:
		val, err := strconv.ParseFloat(tok.Lexeme, 64)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error evaluating number token: %s", err))
		}
		return val
	case token.STRING:
		return expr.Token.Lexeme
	case token.TRUE:
		return true
	case token.FALSE:
		return false
	case token.NIL:
		return nil
	default:
		panic("unexpected literal type")
	}

}

func (i *Interpreter) evalBinaryExpr(expr *ast.BinaryExpr) any {

	left := i.EvaluateExpr(expr.Left)
	right := i.EvaluateExpr(expr.Right)

	switch expr.Operator.TokenType {
	case token.PLUS:
		return i.applyBinaryOperator(token.PLUS, left, right)
	case token.MINUS:
		return i.applyBinaryOperator(token.MINUS, left, right)
	case token.STAR:
		return i.applyBinaryOperator(token.STAR, left, right)
	case token.SLASH:
		return i.applyBinaryOperator(token.SLASH, left, right)
	case token.BANG_EQUAL:
		return i.applyBinaryOperator(token.BANG_EQUAL, left, right)
	case token.EQUAL_EQUAL:
		return i.applyBinaryOperator(token.EQUAL_EQUAL, left, right)
	case token.GREATER:
		return i.applyBinaryOperator(token.GREATER, left, right)
	case token.GREATER_EQUAL:
		return i.applyBinaryOperator(token.GREATER_EQUAL, left, right)
	case token.LESS:
		return i.applyBinaryOperator(token.LESS, left, right)
	case token.LESS_EQUAL:
		return i.applyBinaryOperator(token.LESS_EQUAL, left, right)

	}
	return nil
}

func (i *Interpreter) evalGroupingExpr(expr *ast.Grouping) any {
	return i.EvaluateExpr(expr.Expression)
}

func (i *Interpreter) evalUnaryExpr(expr *ast.Unary) any {
	right := i.EvaluateExpr(expr.Expression)

	switch expr.Operator.TokenType {
	case token.MINUS:
		if num, ok := right.(float64); ok {
			return -num
		}
		panic("Cannot subtract a non-number")
	case token.BANG:
		return !i.isTruthy(right)
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
		if leftString, ok := left.(string); ok {
			if rightString, ok := right.(string); ok {
				return leftString + rightString
			}
		}
		panic("Cannot add two non-numbers/strings")
	case token.MINUS:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum - rightNum
			}
		}
		panic("Cannot subtract two non-numbers")
	case token.STAR:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum * rightNum
			}
		}
		panic("Cannot multiply two non-numbers")
	case token.SLASH:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				if rightNum == 0 {
					panic("Cannot divide by 0")
				}
				return leftNum / rightNum
			}
		}
		panic("Cannot divide two non-numbers")
	case token.BANG_EQUAL:
		return left != right
	case token.EQUAL_EQUAL:
		return left == right
	case token.GREATER:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum > rightNum
			}
		}
		panic("Cannot > two non-numbers")
	case token.GREATER_EQUAL:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum >= rightNum
			}
		}
		panic("Cannot >= two non-numbers")
	case token.LESS:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum < rightNum
			}
		}
		panic("Cannot < two non-numbers")
	case token.LESS_EQUAL:
		if leftNum, ok := left.(float64); ok {
			if rightNum, ok := right.(float64); ok {
				return leftNum <= rightNum
			}
		}
		panic("Cannot <= two non-numbers")
	default:
		panic("unknown binary operator")
	}
}

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}

	if b, ok := value.(bool); ok {
		return b
	}

	return true
}
