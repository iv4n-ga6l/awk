package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Interpreter struct {
	line           string
	fields         []string
	recordNumber   int
	FieldSeparator string
	variables      map[string]interface{}
}

func NewInterpreter(fieldSeparator string) *Interpreter {
	return &Interpreter{
		FieldSeparator: fieldSeparator,
		variables:      make(map[string]interface{}),
	}
}

func (i *Interpreter) SetRecord(line string, fields []string, recordNumber int) {
	i.line = line
	i.fields = fields
	i.recordNumber = recordNumber
}

func (i *Interpreter) Execute(action *Action) {
	if action.Command == "print" {
		var output []string
		for _, arg := range action.Args {
			switch {
			case arg == "$0":
				output = append(output, i.line)
			case arg == "NR":
				output = append(output, fmt.Sprintf("%d", i.recordNumber))
			case arg == "NF":
				output = append(output, fmt.Sprintf("%d", len(i.fields)))
			case strings.HasPrefix(arg, "$"):
				index := parseFieldIndex(arg[1:])
				if index > 0 && index <= len(i.fields) {
					output = append(output, i.fields[index-1])
				}
			default:
				output = append(output, fmt.Sprintf("%v", i.GetVariable(arg)))
			}
		}
		fmt.Println(strings.Join(output, " "))
	}
}

func (i *Interpreter) EvaluatePattern(pattern Pattern) bool {
	// Placeholder: Implement pattern evaluation logic
	return true
}

func (i *Interpreter) ExecuteBegin(actions []*Action) {
	for _, action := range actions {
		i.Execute(action)
	}
}

func (i *Interpreter) ExecuteEnd(actions []*Action) {
	for _, action := range actions {
		i.Execute(action)
	}
}

func (i *Interpreter) GetVariable(name string) interface{} {
	if value, exists := i.variables[name]; exists {
		return value
	}
	return 0 // Default value for uninitialized variables
}

func (i *Interpreter) SetVariable(name string, value interface{}) {
	i.variables[name] = value
}

func (i *Interpreter) ToNumber(value interface{}) float64 {
	switch v := value.(type) {
	case string:
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			return num
		}
		return 0
	case float64:
		return v
	case int:
		return float64(v)
	default:
		return 0
	}
}

func parseFieldIndex(field string) int {
	index, _ := strconv.Atoi(field)
	return index
}