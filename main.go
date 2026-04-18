package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ccawk '{ print ... }' [file]")
		os.Exit(1)
	}

	program := os.Args[1]
	var input *os.File
	var err error

	if len(os.Args) > 2 {
		input, err = os.Open(os.Args[2])
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
	action, parseErr := parser.Parse()
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Error parsing program: %v\n", parseErr)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		interpreter := NewInterpreter(line, fields)
		interpreter.Execute(action)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}