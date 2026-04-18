package main

import (
	"fmt"
)

type Expression interface {
	Evaluate(interpreter *Interpreter) interface{}
}

type VariableExpression struct {
	Name string
}

func (v *VariableExpression) Evaluate(interpreter *Interpreter) interface{} {
	return interpreter.GetVariable(v.Name)
}

type LiteralExpression struct {
	Value string
}

func (l *LiteralExpression) Evaluate(interpreter *Interpreter) interface{} {
	return l.Value
}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpression) Evaluate(interpreter *Interpreter) interface{} {
	left := interpreter.ToNumber(b.Left.Evaluate(interpreter))
	right := interpreter.ToNumber(b.Right.Evaluate(interpreter))

	switch b.Operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right == 0 {
			return 0 // Avoid division by zero
		}
		return left / right
	case "%":
		return int(left) % int(right)
	case "^":
		result := 1.0
		for i := 0; i < int(right); i++ {
			result *= left
		}
		return result
	default:
		return nil
	}
}

type AssignmentExpression struct {
	Variable string
	Operator string
	Value    Expression
}

func (a *AssignmentExpression) Evaluate(interpreter *Interpreter) interface{} {
	value := a.Value.Evaluate(interpreter)
	if a.Operator == "=" {
		interpreter.SetVariable(a.Variable, value)
	} else {
		current := interpreter.ToNumber(interpreter.GetVariable(a.Variable))
		newValue := interpreter.ToNumber(value)
		switch a.Operator {
		case "+=":
			interpreter.SetVariable(a.Variable, current+newValue)
		case "-=":
			interpreter.SetVariable(a.Variable, current-newValue)
		case "*=":
			interpreter.SetVariable(a.Variable, current*newValue)
		case "/=":
			if newValue != 0 {
				interpreter.SetVariable(a.Variable, current/newValue)
			}
		case "%=":
			interpreter.SetVariable(a.Variable, int(current)%int(newValue))
		}
	}
	return interpreter.GetVariable(a.Variable)
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