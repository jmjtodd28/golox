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

func (p *Parser) Parse() ast.Expr {
	ast := p.parseExpression()
	fmt.Println(ast.Print())
	return ast
}

func (p *Parser) parseExpression() ast.Expr {
	return p.parseEquality()
}

func (p *Parser) parseEquality() ast.Expr {
	expr := p.parseComparison()

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.parseComparison()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) parseComparison() ast.Expr {
	expr := p.parseTerm()

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.parseTerm()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) parseTerm() ast.Expr {
	expr := p.parseFactor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.parseFactor()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}

	return expr
}

func (p *Parser) parseFactor() ast.Expr {
	expr := p.parseUnary()

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.parseUnary()
		expr = ast.NewBinaryExpr(expr, right, operator)
	}
	return expr
}

func (p *Parser) parseUnary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.parseUnary()
		return ast.NewUnary(operator, right)
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() ast.Expr {

	tok := p.peek()

	switch tok.TokenType {
	case token.FALSE, token.TRUE, token.NIL, token.STRING, token.NUMBER:
		p.advance()
		return ast.NewLiteral(tok)
	case token.LEFT_PAREN:
		p.advance()
		expr := p.parseExpression()
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
