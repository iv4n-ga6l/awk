package main

import "fmt"

func (p *Parser) parseForInLoop() (*ForInLoop, error) {
	forInLoop := &ForInLoop{}
	p.nextToken()
	if p.token.Type != TOKEN_IDENTIFIER {
		return nil, fmt.Errorf("expected identifier for loop variable")
	}
	forInLoop.Variable = p.token.Value
	p.nextToken()
	if p.token.Type != TOKEN_IN {
		return nil, fmt.Errorf("expected 'in' keyword after loop variable")
	}
	p.nextToken()
	if p.token.Type != TOKEN_IDENTIFIER {
		return nil, fmt.Errorf("expected array identifier after 'in'")
	}
	forInLoop.ArrayName = p.token.Value
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after 'for (key in array)'")
	}
	p.nextToken()
	body, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	forInLoop.Body = body
	p.nextToken()
	return forInLoop, nil
}

func (p *Parser) parseDeleteStatement() (*DeleteStatement, error) {
	deleteStmt := &DeleteStatement{}
	p.nextToken()
	if p.token.Type != TOKEN_IDENTIFIER {
		return nil, fmt.Errorf("expected array identifier after 'delete'")
	}
	deleteStmt.ArrayName = p.token.Value
	p.nextToken()
	if p.token.Type != TOKEN_LBRACKET {
		return nil, fmt.Errorf("expected '[' after array identifier")
	}
	p.nextToken()
	key, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	deleteStmt.Key = key
	p.nextToken()
	if p.token.Type != TOKEN_RBRACKET {
		return nil, fmt.Errorf("expected ']' after array key")
	}
	p.nextToken()
	return deleteStmt, nil
}