package main

import "fmt"

func (p *Parser) parseIfStatement() (*IfStatement, error) {
	ifStmt := &IfStatement{}
	p.nextToken()
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	ifStmt.Condition = condition
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after if condition")
	}
	p.nextToken()
	trueBranch, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	ifStmt.TrueBranch = trueBranch
	p.nextToken()
	if p.token.Type == TOKEN_ELSE {
		p.nextToken()
		if p.token.Type != TOKEN_LBRACE {
			return nil, fmt.Errorf("expected '{' after else")
		}
		p.nextToken()
		falseBranch, err := p.parseActions()
		if err != nil {
			return nil, err
		}
		ifStmt.FalseBranch = falseBranch
		p.nextToken()
	}
	return ifStmt, nil
}

func (p *Parser) parseWhileLoop() (*WhileLoop, error) {
	whileLoop := &WhileLoop{}
	p.nextToken()
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	whileLoop.Condition = condition
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after while condition")
	}
	p.nextToken()
	body, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	whileLoop.Body = body
	p.nextToken()
	return whileLoop, nil
}

func (p *Parser) parseForLoop() (*ForLoop, error) {
	forLoop := &ForLoop{}
	p.nextToken()
	init, err := p.parseAssignment()
	if err != nil {
		return nil, err
	}
	forLoop.Init = init
	p.nextToken()
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	forLoop.Condition = condition
	p.nextToken()
	post, err := p.parseAssignment()
	if err != nil {
		return nil, err
	}
	forLoop.Post = post
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after for loop header")
	}
	p.nextToken()
	body, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	forLoop.Body = body
	p.nextToken()
	return forLoop, nil
}

func (p *Parser) parseDoWhileLoop() (*DoWhileLoop, error) {
	doWhileLoop := &DoWhileLoop{}
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after do")
	}
	p.nextToken()
	body, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	doWhileLoop.Body = body
	p.nextToken()
	if p.token.Type != TOKEN_WHILE {
		return nil, fmt.Errorf("expected 'while' after do-while body")
	}
	p.nextToken()
	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	doWhileLoop.Condition = condition
	p.nextToken()
	return doWhileLoop, nil
}

func (p *Parser) parseCommand() (interface{}, error) {
	switch p.token.Type {
	case TOKEN_BREAK:
		p.nextToken()
		return &BreakCommand{}, nil
	case TOKEN_CONTINUE:
		p.nextToken()
		return &ContinueCommand{}, nil
	case TOKEN_NEXT:
		p.nextToken()
		return &NextCommand{}, nil
	case TOKEN_EXIT:
		p.nextToken()
		var exitCode Expression
		if p.token.Type != TOKEN_SEMICOLON {
			exitCode, _ = p.parseExpression()
		}
		return &ExitCommand{ExitCode: exitCode}, nil
	default:
		return nil, fmt.Errorf("unexpected command token: %s", p.token.Type)
	}
}

func (p *Parser) parseTernaryExpression(condition Expression) (*TernaryExpression, error) {
	p.nextToken()
	trueValue, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	p.nextToken()
	if p.token.Type != TOKEN_COLON {
		return nil, fmt.Errorf("expected ':' in ternary expression")
	}
	p.nextToken()
	falseValue, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	return &TernaryExpression{
		Condition:  condition,
		TrueValue:  trueValue,
		FalseValue: falseValue,
	}, nil
}
