package main

import (
	"fmt"
	"strings"
)

type Interpreter struct {
	line   string
	fields []string
}

func NewInterpreter(line string, fields []string) *Interpreter {
	return &Interpreter{line: line, fields: fields}
}

func (i *Interpreter) Execute(action *Action) {
	if action.Command == "print" {
		var output []string
		for _, arg := range action.Args {
			if arg == "$0" {
				output = append(output, i.line)
			} else if strings.HasPrefix(arg, "$") {
				index := parseFieldIndex(arg[1:])
				if index > 0 && index <= len(i.fields) {
					output = append(output, i.fields[index-1])
				}
			} else {
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