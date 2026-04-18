package main

type Action struct {
	Command string
	Args    []string
}

type Parser struct {
	lexer  *Lexer
	token  Token
	error  error
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) nextToken() {
	p.token = p.lexer.NextToken()
}

func (p *Parser) Parse() (*Action, error) {
	p.nextToken()

	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{', got %s", p.token.Value)
	}

	p.nextToken()
	if p.token.Type != TOKEN_PRINT {
		return nil, fmt.Errorf("expected 'print', got %s", p.token.Value)
	}

	action := &Action{Command: "print"}
	p.nextToken()

	for p.token.Type != TOKEN_RBRACE && p.token.Type != TOKEN_EOF {
		action.Args = append(action.Args, p.token.Value)
		p.nextToken()
	}

	if p.token.Type != TOKEN_RBRACE {
		return nil, fmt.Errorf("expected '}', got %s", p.token.Value)
	}

	return action, nil
}