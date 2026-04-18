package main

import (
	"fmt"
	"strings"
)

type Interpreter struct {
	line         string
	fields       []string
	recordNumber int
	fieldSeparator string
}

func NewInterpreter(line string, fields []string, recordNumber int, fieldSeparator string) *Interpreter {
	return &Interpreter{
		line:         line,
		fields:       fields,
		recordNumber: recordNumber,
		fieldSeparator: fieldSeparator,
	}
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

func parseFieldIndex(field string) int {
	var index int
	fmt.Sscanf(field, "%d", &index)
	return index
}