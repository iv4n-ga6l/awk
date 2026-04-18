package main

import (
	"fmt"
)

type Interpreter struct {
	FieldSeparator string
	RecordSeparator string
	OutputFieldSeparator string
	OutputRecordSeparator string
	Filename string
	Argc int
	Argv []string
	functions map[string]*FunctionDefinition
	scope map[string]interface{}
	returnValue interface{}
}

func NewInterpreter(fieldSeparator, recordSeparator string) *Interpreter {
	return &Interpreter{
		FieldSeparator: fieldSeparator,
		RecordSeparator: recordSeparator,
		OutputFieldSeparator: " ",
		OutputRecordSeparator: "\n",
		functions: make(map[string]*FunctionDefinition),
		scope: make(map[string]interface{}),
	}
}

func (i *Interpreter) SetRecord(line string, fields []string, recordNumber int) {
	i.scope["$0"] = line
	for idx, field := range fields {
		i.scope[fmt.Sprintf("$%d", idx+1)] = field
	}
	i.scope["NR"] = recordNumber
}

func (i *Interpreter) SetFilename(filename string) {
	i.Filename = filename
	i.scope["FILENAME"] = filename
}

func (i *Interpreter) SetARGV(argv []string) {
	i.Argv = argv
	i.scope["ARGV"] = argv
}

func (i *Interpreter) SetARGC(argc int) {
	i.Argc = argc
	i.scope["ARGC"] = argc
}

func (i *Interpreter) SetVariable(name string, value interface{}) {
	i.scope[name] = value
}

func (i *Interpreter) Execute(action *Action) {
	switch action.Command {
	case "print":
		var output []string
		for _, expr := range action.PrintArgs {
			output = append(output, fmt.Sprintf("%v", expr.Evaluate(i)))
		}
		fmt.Print(strings.Join(output, i.OutputFieldSeparator) + i.OutputRecordSeparator)
	default:
		// Existing cases
	}
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