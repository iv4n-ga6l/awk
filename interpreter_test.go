package main

import (
	"testing"
)

func TestOutputFieldSeparatorAndRecordSeparator(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	interpreter.SetVariable("OFS", "-")
	interpreter.SetVariable("ORS", "\n")

	program := &Program{
		BeginActions: []*Action{
			{
				Command: "print",
				PrintArgs: []Expression{
					&LiteralExpression{Value: "Hello"},
					&LiteralExpression{Value: "World"},
				},
			},
		},
	}

	output := interpreter.Run(program, "")
	expected := "Hello-World\n"
	if output != expected {
		t.Errorf("Expected output: %s, got: %s", expected, output)
	}
}

func TestFilenameAndMultipleFiles(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	program := &Program{
		Rules: []*Rule{
			{
				Pattern: nil,
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&VariableExpression{Name: "FILENAME"},
							&VariableExpression{Name: "$0"},
						},
					},
				},
			},
		},
	}

	sampleInput1 := "John 25\nJane 30"
	sampleInput2 := "Bob 22\nAlice 35"

	output := interpreter.Run(program, sampleInput1, sampleInput2)
	expected := "test1.txt John 25\ntest1.txt Jane 30\ntest2.txt Bob 22\ntest2.txt Alice 35\n"
	if output != expected {
		t.Errorf("Expected output: %s, got: %s", expected, output)
	}
}

func TestVariableAssignmentViaFlag(t *testing.T) {
	interpreter := NewInterpreter(" ", "\n")
	interpreter.SetVariable("threshold", 25)

	program := &Program{
		Rules: []*Rule{
			{
				Pattern: &BinaryExpression{
					Left:  &FieldExpression{Field: 2},
					Right: &VariableExpression{Name: "threshold"},
					Operator: ">",
				},
				Action: []*Action{
					{
						Command: "print",
						PrintArgs: []Expression{
							&FieldExpression{Field: 1},
						},
					},
				},
			},
		},
	}

	sampleInput := "John 25\nJane 30\nBob 22"
	output := interpreter.Run(program, sampleInput)
	expected := "Jane\n"
	if output != expected {
		t.Errorf("Expected output: %s, got: %s", expected, output)
	}
}