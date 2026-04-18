package main

import "fmt"

func (p *Parser) parseFunctionDefinition() (*FunctionDefinition, error) {
	function := &FunctionDefinition{}
	p.nextToken()
	if p.token.Type != TOKEN_IDENTIFIER {
		return nil, fmt.Errorf("expected function name after 'function'")
	}
	function.Name = p.token.Value
	p.nextToken()
	if p.token.Type != TOKEN_LPAREN {
		return nil, fmt.Errorf("expected '(' after function name")
	}
	p.nextToken()
	for p.token.Type == TOKEN_IDENTIFIER {
		function.Parameters = append(function.Parameters, p.token.Value)
		p.nextToken()
		if p.token.Type == TOKEN_COMMA {
			p.nextToken()
		}
	}
	if p.token.Type != TOKEN_RPAREN {
		return nil, fmt.Errorf("expected ')' after function parameters")
	}
	p.nextToken()
	if p.token.Type != TOKEN_LBRACE {
		return nil, fmt.Errorf("expected '{' after function signature")
	}
	p.nextToken()
	body, err := p.parseActions()
	if err != nil {
		return nil, err
	}
	function.Body = body
	p.nextToken()
	if p.token.Type != TOKEN_RBRACE {
		return nil, fmt.Errorf("expected '}' after function body")
	}
	return function, nil
}

func (p *Parser) parseReturnStatement() (*ReturnStatement, error) {
	returnStmt := &ReturnStatement{}
	p.nextToken()
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	returnStmt.Expression = expr
	return returnStmt, nil
}