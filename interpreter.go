package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Interpreter struct {
	line           string
	fields         []string
	recordNumber   int
	FieldSeparator string
}

func NewInterpreter(fieldSeparator string) *Interpreter {
	return &Interpreter{
		FieldSeparator: fieldSeparator,
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
				output = append(output, arg)
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

func parseFieldIndex(field string) int {
	index, _ := strconv.Atoi(field)
	return index
}