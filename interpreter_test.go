package main

import (
	"testing"
)

func TestUserDefinedFunction(t *testing.T) {
	interpreter := NewInterpreter(" ")
	program := &Program{
		BeginActions: []*Action{
			{
				Command: "function",
				FunctionDefinition: &FunctionDefinition{
					Name: "max",
					Parameters: []string{"a", "b"},
					Body: []*Action{
						{
							Command: "return",
							ReturnStatement: &ReturnStatement{
								Expression: &ConditionalExpression{
									Condition: &BinaryExpression{
										Left:  &VariableExpression{Name: "a"},
										Right: &VariableExpression{Name: "b"},
										Operator: ">",
									},
									TrueExpr:  &VariableExpression{Name: "a"},
									FalseExpr: &VariableExpression{Name: "b"},
								},
							},
						},
				},
			},
		},
		Rules: []*Rule{
			{
				Pattern: nil,
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&FieldExpression{Field: 1},
							&FunctionCallExpression{
								Name: "max",
								Arguments: []Expression{
									&FieldExpression{Field: 2},
									&LiteralExpression{Value: 30},
								},
							},
						},
					},
				},
			},
		},
	}
	sampleInput := "John 25\nJane 30\nBob 22"
	output := interpreter.Run(program, sampleInput)
	expected := "John 30\nJane 30\nBob 30\n"
	if output != expected {
		t.Errorf("Expected output: %s, got: %s", expected, output)
	}
}

func TestBuiltInArithmeticFunctions(t *testing.T) {
	interpreter := NewInterpreter(" ")
	program := &Program{
		Rules: []*Rule{
			{
				Pattern: nil,
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&FieldExpression{Field: 1},
							&FunctionCallExpression{
								Name: "sqrt",
								Arguments: []Expression{
									&FieldExpression{Field: 2},
								},
							},
						},
					},
				},
			},
		},
	}
	sampleInput := "John 25\nJane 30\nBob 22"
	output := interpreter.Run(program, sampleInput)
	expected := "John 5\nJane 5\nBob 4\n"
	if output != expected {
		t.Errorf("Expected output: %s, got: %s", expected, output)
	}
}

func TestRandAndSrand(t *testing.T) {
	interpreter := NewInterpreter(" ")
	program := &Program{
		BeginActions: []*Action{
			{
				Command: "srand",
				NumericExpression: &LiteralExpression{Value: 42},
			},
			{
				Command: "print",
				PrintArgs: []Expression{
					&FunctionCallExpression{
						Name: "rand",
						Arguments: []Expression{},
					},
				},
			},
		},
	}
	output := interpreter.Run(program, "")
	if len(output) == 0 {
		t.Errorf("Expected random numbers, got empty output")
	}
}