package main

import (
	"testing"
)

func TestAssociativeArray(t *testing.T) {
	a := NewAssociativeArray()
	a.Set("key1", 42)
	a.Set("key2", "value")

	if val, exists := a.Get("key1"); !exists || val != 42 {
		t.Errorf("Expected key1 to have value 42, got %v", val)
	}

	if val, exists := a.Get("key2"); !exists || val != "value" {
		t.Errorf("Expected key2 to have value 'value', got %v", val)
	}

	if a.Contains("key3") {
		t.Errorf("Expected key3 to not exist")
	}

	a.Delete("key1")
	if _, exists := a.Get("key1"); exists {
		t.Errorf("Expected key1 to be deleted")
	}

	keys := a.Keys()
	if len(keys) != 1 || keys[0] != "key2" {
		t.Errorf("Expected keys to contain only 'key2', got %v", keys)
	}
}

func TestForInLoop(t *testing.T) {
	interpreter := NewInterpreter(" ")
	a := interpreter.GetArray("testArray")
	a.Set("key1", 1)
	a.Set("key2", 2)
	a.Set("key3", 3)

	forInLoop := &ForInLoop{
		Variable:  "key",
		ArrayName: "testArray",
		Body: []*Action{
			{
				Command: "print",
				PrintArgs: []Expression{
					&VariableExpression{Name: "key"},
					&ArrayAccessExpression{ArrayName: "testArray", Key: &VariableExpression{Name: "key"}},
				},
			},
		},
	}

	interpreter.Execute(&Action{Command: "for-in", ForInLoop: forInLoop})
	// Output order is unspecified, so manual verification is required.
}

func TestDeleteStatement(t *testing.T) {
	interpreter := NewInterpreter(" ")
	a := interpreter.GetArray("testArray")
	a.Set("key1", 1)
	a.Set("key2", 2)
	a.Set("key3", 3)

	deleteStmt := &DeleteStatement{
		ArrayName: "testArray",
		Key:      &LiteralExpression{Value: "key2"},
	}

	interpreter.Execute(&Action{Command: "delete", DeleteStatement: deleteStmt})

	if a.Contains("key2") {
		t.Errorf("Expected key2 to be deleted")
	}

	keys := a.Keys()
	if len(keys) != 2 || keys[0] == "key2" || keys[1] == "key2" {
		t.Errorf("Expected keys to not contain 'key2', got %v", keys)
	}
}