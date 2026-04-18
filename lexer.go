package main

const (
	TOKEN_IF       TokenType = "IF"
	TOKEN_ELSE     TokenType = "ELSE"
	TOKEN_WHILE    TokenType = "WHILE"
	TOKEN_FOR      TokenType = "FOR"
	TOKEN_DO       TokenType = "DO"
	TOKEN_BREAK    TokenType = "BREAK"
	TOKEN_CONTINUE TokenType = "CONTINUE"
	TOKEN_NEXT     TokenType = "NEXT"
	TOKEN_EXIT     TokenType = "EXIT"
	TOKEN_QUESTION TokenType = "QUESTION"
	TOKEN_COLON    TokenType = "COLON"
)

func (l *Lexer) lexControlFlow() Token {
	start := l.pos
	for l.pos < len(l.input) && isAlphaNumeric(l.input[l.pos]) {
		l.pos++
	}
	value := l.input[start:l.pos]
	switch value {
	case "if":
		return Token{Type: TOKEN_IF, Value: value}
	case "else":
		return Token{Type: TOKEN_ELSE, Value: value}
	case "while":
		return Token{Type: TOKEN_WHILE, Value: value}
	case "for":
		return Token{Type: TOKEN_FOR, Value: value}
	case "do":
		return Token{Type: TOKEN_DO, Value: value}
	case "break":
		return Token{Type: TOKEN_BREAK, Value: value}
	case "continue":
		return Token{Type: TOKEN_CONTINUE, Value: value}
	case "next":
		return Token{Type: TOKEN_NEXT, Value: value}
	case "exit":
		return Token{Type: TOKEN_EXIT, Value: value}
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
	case '?':
		l.pos++
		return Token{Type: TOKEN_QUESTION, Value: "?"}
	case ':':
		l.pos++
		return Token{Type: TOKEN_COLON, Value: ":"}
	default:
		if isAlpha(ch) {
			return l.lexControlFlow()
		}
		return l.NextToken()
	}
}
