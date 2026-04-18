package main

import (
	"regexp"
	"strings"
	"unicode"
)

type TokenType string

const (
	TOKEN_PRINT   TokenType = "PRINT"
	TOKEN_DOLLAR  TokenType = "DOLLAR"
	TOKEN_NUMBER  TokenType = "NUMBER"
	TOKEN_LBRACE  TokenType = "LBRACE"
	TOKEN_RBRACE  TokenType = "RBRACE"
	TOKEN_STRING  TokenType = "STRING"
	TOKEN_REGEX   TokenType = "REGEX"
	TOKEN_BEGIN   TokenType = "BEGIN"
	TOKEN_END     TokenType = "END"
	TOKEN_OPERATOR TokenType = "OPERATOR"
	TOKEN_IDENTIFIER TokenType = "IDENTIFIER"
	TOKEN_EOF     TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF}
	}

	ch := l.input[l.pos]

	if unicode.IsSpace(rune(ch)) {
		l.pos++
		return l.NextToken()
	}

	switch ch {
	case '{':
		l.pos++
		return Token{Type: TOKEN_LBRACE, Value: "{"}
	case '}':
		l.pos++
		return Token{Type: TOKEN_RBRACE, Value: "}"}
	case '$':
		l.pos++
		return Token{Type: TOKEN_DOLLAR, Value: "$"}
	case '/':
		return l.lexRegex()
	case '=', '!', '<', '>':
		return l.lexOperator()
	}

	if isAlpha(ch) {
		return l.lexIdentifier()
	}

	if isDigit(ch) {
		return l.lexNumber()
	}

	l.pos++
	return l.NextToken()
}

func (l *Lexer) lexRegex() Token {
	start := l.pos
	l.pos++ // Skip initial '/'
	for l.pos < len(l.input) && l.input[l.pos] != '/' {
		l.pos++
	}
	if l.pos < len(l.input) {
		l.pos++ // Skip closing '/'
		return Token{Type: TOKEN_REGEX, Value: l.input[start:l.pos]}
	}
	return Token{Type: TOKEN_EOF}
}

func (l *Lexer) lexOperator() Token {
	start := l.pos
	l.pos++
	if l.pos < len(l.input) && (l.input[l.pos] == '=' || l.input[l.pos] == '~') {
		l.pos++
	}
	return Token{Type: TOKEN_OPERATOR, Value: l.input[start:l.pos]}
}

func (l *Lexer) lexIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) && (isAlphaNumeric(l.input[l.pos]) || l.input[l.pos] == '_') {
		l.pos++
	}
	value := l.input[start:l.pos]
	if value == "print" {
		return Token{Type: TOKEN_PRINT, Value: value}
	} else if value == "BEGIN" {
		return Token{Type: TOKEN_BEGIN, Value: value}
	} else if value == "END" {
		return Token{Type: TOKEN_END, Value: value}
	}
	return Token{Type: TOKEN_IDENTIFIER, Value: value}
}

func (l *Lexer) lexNumber() Token {
	start := l.pos
	for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
		l.pos++
	}
	return Token{Type: TOKEN_NUMBER, Value: l.input[start:l.pos]}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}