package main

import "fmt"

// ControlFlow constructs

type IfStatement struct {
	Condition Expression
	TrueBranch []*Action
	FalseBranch []*Action
}

type WhileLoop struct {
	Condition Expression
	Body []*Action
}

type ForLoop struct {
	Init      *AssignmentExpression
	Condition Expression
	Post      *AssignmentExpression
	Body      []*Action
}

type DoWhileLoop struct {
	Condition Expression
	Body      []*Action
}

// Commands

type BreakCommand struct{}

type ContinueCommand struct{}

type NextCommand struct{}

type ExitCommand struct{
	ExitCode Expression
}

// Ternary Conditional Operator

type TernaryExpression struct {
	Condition     Expression
	TrueValue     Expression
	FalseValue    Expression
}

func (t *TernaryExpression) Evaluate(interpreter *Interpreter) interface{} {
	if interpreter.ToNumber(t.Condition.Evaluate(interpreter)) != 0 {
		return t.TrueValue.Evaluate(interpreter)
	}
	return t.FalseValue.Evaluate(interpreter)
}
