package main

import (
	"testing"
)

func TestPipeOperator(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	program := &Program{
		Rules: []*Rule{
			{
				Pattern: nil,
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&FieldExpression{Field: 1},
						},
						PipeCommand: "sort",
					},
				},
			},
		},
	}

	sampleInput := "John\nJane\nBob\nAlice\nCharlie"
	output := interpreter.Run(program, sampleInput)
	// Verify output manually for sorted names
}

func TestGetline(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	program := &Program{
		BeginActions: []*Action{
			{
				Command: "getline",
				GetlineTarget: &VariableExpression{Name: "line"},
				GetlineSource: "< test.txt",
			},
			{
				Command: "print",
				PrintArgs: []Expression{
					&VariableExpression{Name: "line"},
				},
			},
		},
	}

	output := interpreter.Run(program, "")
	// Verify output matches the content of test.txt
}

func TestCloseFunction(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	program := &Program{
		Rules: []*Rule{
			{
				Pattern: nil,
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&LiteralExpression{Value: "hello"},
						},
						PipeCommand: "cat",
					},
					{
						Command: "close",
						CloseTarget: "cat",
					},
				},
			},
		},
	}

	output := interpreter.Run(program, "")
	// Verify that the pipe was closed properly
}