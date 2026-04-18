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
	var recordSeparator string
	var programFile string
	var variableAssignments stringArray

	flag.StringVar(&fieldSeparator, "F", "\t", "Field separator")
	flag.StringVar(&recordSeparator, "RS", "\n", "Record separator")
	flag.StringVar(&programFile, "f", "", "Program file")
	flag.Var(&variableAssignments, "v", "Variable assignments")
	flag.Parse()

	args := flag.Args()
	if programFile == "" && len(args) < 1 {
		fmt.Println("Usage: ccawk [-F separator] [-RS separator] [-f programfile | 'program'] [-v var=value] [file ...]")
		os.Exit(1)
	}

	var program string
	if programFile != "" {
		content, err := os.ReadFile(programFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading program file: %v\n", err)
			os.Exit(1)
		}
		program = string(content)
	} else {
		program = args[0]
	}

	files := args[1:]
	if len(files) == 0 {
		files = append(files, "-") // Default to stdin
	}

	lexer := NewLexer(program)
	parser := NewParser(lexer)
	programAST, parseErr := parser.Parse()
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Error parsing program: %v\n", parseErr)
		os.Exit(1)
	}

	interpreter := NewInterpreter(fieldSeparator, recordSeparator)

	// Set ARGV and ARGC
	interpreter.SetARGV(append([]string{os.Args[0]}, args...))
	interpreter.SetARGC(len(os.Args))

	// Apply variable assignments
	for _, assignment := range variableAssignments {
		parts := strings.SplitN(assignment, "=", 2)
		if len(parts) == 2 {
			interpreter.SetVariable(parts[0], parts[1])
		}
	}

	// Execute BEGIN actions
	interpreter.ExecuteBegin(programAST.BeginActions)

	for _, file := range files {
		var input *os.File
		var err error
		if file == "-" {
			input = os.Stdin
			interpreter.SetFilename("stdin")
		} else {
			input, err = os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", file, err)
				os.Exit(1)
			}
			defer input.Close()
			interpreter.SetFilename(file)
		}

		scanner := bufio.NewScanner(input)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if idx := strings.Index(string(data), recordSeparator); idx >= 0 {
				return idx + len(recordSeparator), data[:idx], nil
			}
			if atEOF {
				return len(data), data, nil
			}
			return 0, nil, nil
		})

		recordNumber := 0
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
	}

	// Execute END actions
	interpreter.ExecuteEnd(programAST.EndActions)
}

type stringArray []string

func (s *stringArray) String() string {
	return strings.Join(*s, ",")
}

func (s *stringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}