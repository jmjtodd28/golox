package parser

import (
	"fmt"
	"slices"

	"github.com/jmjtodd28/golox/ast"
	"github.com/jmjtodd28/golox/token"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func NewParser(tokens []token.Token) Parser {
	return Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) Parse() {
	ast := p.expression()
	fmt.Println(ast.Print())
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(token.BANG_EQUAL, token.BANG) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.term()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary()

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.NewUnary(operator, right)
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {

	tok := p.peek()

	switch tok.TokenType {
	case token.FALSE, token.TRUE, token.NIL, token.STRING, token.NUMBER:
		p.advance()
		return ast.NewLiteral(tok)
	case token.LEFT_PAREN:
		p.advance()
		expr := p.expression()
		if !p.match(token.RIGHT_PAREN) {
			// todo improve error handling
			panic("Expected ')' after expression")
		}

		return ast.NewGrouping(expr)
	default:
		panic(fmt.Sprintf("Unexpected token: %v", tok))
	}
}

func (p *Parser) match(tokenTypes ...token.Type) bool {
	if slices.ContainsFunc(tokenTypes, p.check) {
		p.advance()
		return true
	}

	return false
}

func (p *Parser) check(tokenType token.Type) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().TokenType == tokenType
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
