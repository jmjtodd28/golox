package scanner

import (
	"fmt"

	"github.com/jmjtodd28/golox/token"
)

type Scanner struct {
	source  string
	Tokens  []token.Token
	start   int
	current int
	line    int
	errs    []error
}

func NewScanner(source string) Scanner {
	var tokens []token.Token
	return Scanner{
		source:  source,
		Tokens:  tokens,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() {

	for !s.isAtEnd() {
		s.start = s.current
		if nextToken, ok := s.scanToken(); ok {
			s.Tokens = append(s.Tokens, nextToken)
		}
	}

	s.Tokens = append(s.Tokens, token.NewToken(token.EOF, "", nil, s.line))
}

// scanToken returns a token, if there is no token to return
// then it returns false
func (s *Scanner) scanToken() (token.Token, bool) {
	c := s.advance()

	var tokenType token.Type
	switch c {
	case '(':
		tokenType = token.LEFT_PAREN
	case ')':
		tokenType = token.RIGHT_PAREN
	case '{':
		tokenType = token.LEFT_BRACE
	case '}':
		tokenType = token.RIGHT_BRACE
	case ',':
		tokenType = token.COMMA
	case '.':
		tokenType = token.DOT
	case '-':
		tokenType = token.MINUS
	case '+':
		tokenType = token.PLUS
	case ';':
		tokenType = token.SEMICOLON
	case '*':
		tokenType = token.STAR
	case '!':
		if s.match('=') {
			tokenType = token.BANG_EQUAL
		} else {
			tokenType = token.BANG
		}
	case '=':
		if s.match('=') {
			tokenType = token.EQUAL_EQUAL
		} else {
			tokenType = token.EQUAL
		}
	case '<':
		if s.match('=') {
			tokenType = token.LESS_EQUAL
		} else {
			tokenType = token.LESS
		}
	case '>':
		if s.match('=') {
			tokenType = token.GREATER_EQUAL
		} else {
			tokenType = token.EQUAL
		}
	case '/':
		if s.match('/') {
			for s.source[s.current] != '\n' && !s.isAtEnd() {
				s.advance()
			}
			return token.Token{}, false
		} else {
			tokenType = token.SLASH
		}
	case ' ':
		return token.Token{}, false
	case '\r':
		return token.Token{}, false
	case '\t':
		return token.Token{}, false
	case '\n':
		s.line++
		return token.Token{}, false
	case '"':
		tokenType = token.STRING
		s.string()
		return token.NewToken(tokenType, s.source[s.start+1:s.current-1], nil, s.line), true
	default:
		if isDigit(c) {
			tokenType = token.NUMBER
			s.number()
		} else if isAlpha(c) {
			tokenType = s.identifier()

		} else {
			tokenType = token.ILLEGAL
			s.errs = append(s.errs, fmt.Errorf("Line %d: Invalid character.", s.line))
		}
	}

	return token.NewToken(tokenType, s.source[s.start:s.current], nil, s.line), true
}

func (s *Scanner) identifier() token.Type {
	for !s.isAtEnd() && isAlphaNumeric(s.source[s.current]) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	if wordType, ok := token.StringToKeyWordType[text]; ok {
		return wordType
	}

	return token.IDENTIFIER
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isDigit(c) || isAlpha(c)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) string() {
	for !s.isAtEnd() && s.source[s.current] != '"' {
		if s.source[s.current] == '\n' {
			s.line++
		}
		s.advance()
	}

	// todo properly handle unterminated strings
	if s.isAtEnd() {
		s.errs = append(s.errs, fmt.Errorf("Line %d: Unterminated string", s.line))
		return
	}

	s.advance()
}

func (s *Scanner) number() {
	for isDigit(s.source[s.current]) {
		s.advance()
	}

	// look for fractional
	next, ok := s.peekNext()
	if s.source[s.current] == '.' && ok && isDigit(next) {
		// consume the "."
		s.advance()
		for isDigit(s.source[s.current]) {
			s.advance()
		}
	}
}

// returns the next character unless its the end of the
// file, then it will return false
func (s *Scanner) peekNext() (byte, bool) {
	if s.current+1 >= len(s.source) {
		return 0, false
	}

	return s.source[s.current+1], true
}

// isAtEnd determines whether we are at the end of the source code
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	char := s.source[s.current]
	s.current++
	return char
}
