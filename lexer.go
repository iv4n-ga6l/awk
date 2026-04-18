package main

import (
	"strings"
)

type TokenType string

const (
	TOKEN_PRINT TokenType = "PRINT"
	TOKEN_DOLLAR TokenType = "DOLLAR"
	TOKEN_NUMBER TokenType = "NUMBER"
	TOKEN_LBRACE TokenType = "LBRACE"
	TOKEN_RBRACE TokenType = "RBRACE"
	TOKEN_EOF TokenType = "EOF"
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

	switch l.input[l.pos] {
	case '{':
		l.pos++
		return Token{Type: TOKEN_LBRACE, Value: "{"}
	case '}':
		l.pos++
		return Token{Type: TOKEN_RBRACE, Value: "}"}
	case '$':
		l.pos++
		return Token{Type: TOKEN_DOLLAR, Value: "$"}
	case 'p':
		if strings.HasPrefix(l.input[l.pos:], "print") {
			l.pos += len("print")
			return Token{Type: TOKEN_PRINT, Value: "print"}
		}
	}

	if isDigit(l.input[l.pos]) {
		start := l.pos
		for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
			l.pos++
		}
		return Token{Type: TOKEN_NUMBER, Value: l.input[start:l.pos]}
	}

	l.pos++
	return l.NextToken()
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}