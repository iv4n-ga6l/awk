package main

const (
	TOKEN_IN     TokenType = "IN"
	TOKEN_DELETE TokenType = "DELETE"
)

func (l *Lexer) lexKeywordOrIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) && isAlphaNumeric(l.input[l.pos]) {
		l.pos++
	}
	value := l.input[start:l.pos]
	switch value {
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
	case '[':
		l.pos++
		return Token{Type: TOKEN_LBRACKET, Value: "["}
	case ']':
		l.pos++
		return Token{Type: TOKEN_RBRACKET, Value: "]"}
	default:
		if isAlpha(ch) {
			return l.lexKeywordOrIdentifier()
		}
		return l.NextToken()
	}
}