package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fieldSeparator string
	flag.StringVar(&fieldSeparator, "F", "\t", "Field separator")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ccawk [-F separator] 'program' [file]")
		os.Exit(1)
	}

	program := args[0]
	var input *os.File
	var err error

	if len(args) > 1 {
		input, err = os.Open(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	lexer := NewLexer(program)
	parser := NewParser(lexer)
	programAST, parseErr := parser.Parse()
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Error parsing program: %v\n", parseErr)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)
	recordNumber := 0
	interpreter := NewInterpreter(fieldSeparator)

	// Execute BEGIN actions
	interpreter.ExecuteBegin(programAST.BeginActions)

	for scanner.Scan() {
		recordNumber++
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return strings.ContainsRune(interpreter.FieldSeparator, r)
		})
		interpreter.SetRecord(line, fields, recordNumber)

		for _, rule := range programAST.Rules {
			if interpreter.EvaluatePattern(rule.Pattern) {
				interpreter.Execute(rule.Action)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Execute END actions
	interpreter.ExecuteEnd(programAST.EndActions)
}