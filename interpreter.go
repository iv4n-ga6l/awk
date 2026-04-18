package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Interpreter struct {
	FieldSeparator        string
	RecordSeparator       string
	OutputFieldSeparator  string
	OutputRecordSeparator string
	Filename              string
	Argc                  int
	Argv                  []string
	functions             map[string]*FunctionDefinition
	scope                 map[string]interface{}
	returnValue           interface{}
	openPipes             map[string]*exec.Cmd
	pipeWriters           map[string]*os.File
	pipeMutex             sync.Mutex
}

func NewInterpreter(fieldSeparator, recordSeparator string) *Interpreter {
	return &Interpreter{
		FieldSeparator:        fieldSeparator,
		RecordSeparator:       recordSeparator,
		OutputFieldSeparator:  " ",
		OutputRecordSeparator: "\n",
		functions:             make(map[string]*FunctionDefinition),
		scope:                 make(map[string]interface{}),
		openPipes:             make(map[string]*exec.Cmd),
		pipeWriters:           make(map[string]*os.File),
	}
}

func (i *Interpreter) Execute(action *Action) {
	switch action.Command {
	case "print":
		if action.PipeCommand != "" {
			i.pipeMutex.Lock()
			cmd := action.PipeCommand
			writer, exists := i.pipeWriters[cmd]
			if !exists {
				command := exec.Command("sh", "-c", cmd)
				stdin, err := command.StdinPipe()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error creating pipe: %v\n", err)
					i.pipeMutex.Unlock()
					return
				}
				writer = stdin
				i.pipeWriters[cmd] = writer
				i.openPipes[cmd] = command
				command.Stdout = os.Stdout
				command.Stderr = os.Stderr
				if err := command.Start(); err != nil {
					fmt.Fprintf(os.Stderr, "Error starting command: %v\n", err)
					i.pipeMutex.Unlock()
					return
				}
			}
			var output []string
			for _, expr := range action.PrintArgs {
				output = append(output, fmt.Sprintf("%v", expr.Evaluate(i)))
			}
			line := strings.Join(output, i.OutputFieldSeparator) + i.OutputRecordSeparator
			writer.Write([]byte(line))
			i.pipeMutex.Unlock()
		} else {
			var output []string
			for _, expr := range action.PrintArgs {
				output = append(output, fmt.Sprintf("%v", expr.Evaluate(i)))
			}
			fmt.Print(strings.Join(output, i.OutputFieldSeparator) + i.OutputRecordSeparator)
		}
	case "close":
		cmd := action.CloseTarget
		if pipe, exists := i.openPipes[cmd]; exists {
			i.pipeMutex.Lock()
			pipe.Process.Kill()
			pipe.Wait()
			delete(i.openPipes, cmd)
			if writer, exists := i.pipeWriters[cmd]; exists {
				writer.Close()
				delete(i.pipeWriters, cmd)
			}
			i.pipeMutex.Unlock()
		}
	case "getline":
		if action.GetlineTarget != nil {
			var line string
			if action.GetlineSource == "" {
				// Read from current input
				line = i.readNextLine()
			} else if strings.HasPrefix(action.GetlineSource, "|") {
				// Read from command
				cmd := strings.TrimPrefix(action.GetlineSource, "|")
				line = i.readFromCommand(cmd)
			} else {
				// Read from file
				line = i.readFromFile(action.GetlineSource)
			}
			i.SetVariable(action.GetlineTarget.Name, line)
		} else {
			// Read into $0
			line := i.readNextLine()
			fields := strings.FieldsFunc(line, func(r rune) bool {
				return strings.ContainsRune(i.FieldSeparator, r)
			})
			i.SetRecord(line, fields, i.scope["NR"].(int)+1)
		}
	default:
		// Existing cases
	}
}

func (i *Interpreter) readNextLine() string {
	// Implementation for reading the next line from the current input
	return ""
}

func (i *Interpreter) readFromCommand(cmd string) string {
	// Implementation for reading from a command
	return ""
}

func (i *Interpreter) readFromFile(filename string) string {
	// Implementation for reading from a file
	return ""
}
