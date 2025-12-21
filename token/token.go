package token

import "fmt"

// The type of the token
type Type int

const (
	ILLEGAL = iota

	// Single-character tokens.
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	TokenType Type
	Lexeme    string
	Literal   any
	Line      int
}

var StringToKeyWordType = map[string]Type{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

var TypeToStrings = map[Type]byte{
	LEFT_PAREN:  '(',
	RIGHT_PAREN: ')',
	LEFT_BRACE:  '{',
	RIGHT_BRACE: '}',
	COMMA:       ',',
	DOT:         '.',
	MINUS:       '-',
	PLUS:        '+',
	SEMICOLON:   ';',
	STAR:        '*',
}

func NewToken(tokenType Type, lexeme string, literal any, line int) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%d - %q - %v", t.TokenType, t.Lexeme, t.Literal)
}

// todo improve this error messaging, maybe store the column and line number
func ReportError(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}
