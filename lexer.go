package main

const (
	TOKEN_FUNCTION TokenType = "FUNCTION"
	TOKEN_RETURN   TokenType = "RETURN"
)

func (l *Lexer) lexKeywordOrIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) && isAlphaNumeric(l.input[l.pos]) {
		l.pos++
	}
	value := l.input[start:l.pos]
	switch value {
	case "function":
		return Token{Type: TOKEN_FUNCTION, Value: value}
	case "return":
		return Token{Type: TOKEN_RETURN, Value: value}
	case "in":
		return Token{Type: TOKEN_IN, Value: value}
	case "delete":
		return Token{Type: TOKEN_DELETE, Value: value}
	default:
		return Token{Type: TOKEN_IDENTIFIER, Value: value}
	}
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
	case '(':
		l.pos++
		return Token{Type: TOKEN_LPAREN, Value: "("}
	case ')':
		l.pos++
		return Token{Type: TOKEN_RPAREN, Value: ")"}
	case ',':
		l.pos++
		return Token{Type: TOKEN_COMMA, Value: ","}
	default:
		if isAlpha(ch) {
			return l.lexKeywordOrIdentifier()
		}
		return l.NextToken()
	}
}