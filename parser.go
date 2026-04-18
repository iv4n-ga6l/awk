package main

import (
	"fmt"
)

type Pattern interface {
	Evaluate(interpreter *Interpreter) bool
}

type Action struct {
	Command string
	Args    []string
}

type Rule struct {
	Pattern Pattern
	Action  *Action
}

type Program struct {
	BeginActions []*Action
	Rules        []Rule
	EndActions   []*Action
}

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) nextToken() {
	p.token = p.lexer.NextToken()
}

func (p *Parser) Parse() (*Program, error) {
	program := &Program{}
	p.nextToken()

	for p.token.Type != TOKEN_EOF {
		if p.token.Type == TOKEN_BEGIN {
			p.nextToken()
			if p.token.Type != TOKEN_LBRACE {
				return nil, fmt.Errorf("expected '{' after BEGIN, got %s", p.token.Value)
			}
			p.nextToken()
			action, err := p.parseAction()
			if err != nil {
				return nil, err
			}
			program.BeginActions = append(program.BeginActions, action)
			p.nextToken()
		} else if p.token.Type == TOKEN_END {
			p.nextToken()
			if p.token.Type != TOKEN_LBRACE {
				return nil, fmt.Errorf("expected '{' after END, got %s", p.token.Value)
			}
			p.nextToken()
			action, err := p.parseAction()
			if err != nil {
				return nil, err
			}
			program.EndActions = append(program.EndActions, action)
			p.nextToken()
		} else {
			pattern, err := p.parsePattern()
			if err != nil {
				return nil, err
			}
			var action *Action
			if p.token.Type == TOKEN_LBRACE {
				p.nextToken()
				action, err = p.parseAction()
				if err != nil {
					return nil, err
				}
				p.nextToken()
			}
			program.Rules = append(program.Rules, Rule{Pattern: pattern, Action: action})
		}
	}

	return program, nil
}

func (p *Parser) parsePattern() (Pattern, error) {
	// Placeholder: Implement pattern parsing (comparison, regex, etc.)
	return nil, nil
}

func (p *Parser) parseAction() (*Action, error) {
	action := &Action{Command: "print"}
	for p.token.Type != TOKEN_RBRACE && p.token.Type != TOKEN_EOF {
		action.Args = append(action.Args, p.token.Value)
		p.nextToken()
	}
	if p.token.Type != TOKEN_RBRACE {
		return nil, fmt.Errorf("expected '}', got %s", p.token.Value)
	}
	return action, nil
}