package main

import (
	"fmt"
	"strings"
)

func (i *Interpreter) Execute(action *Action) {
	switch action.Command {
	case "print":
		// Existing print logic
	case "printf":
		args := make([]interface{}, len(action.PrintfArgs))
		for idx, expr := range action.PrintfArgs {
			args[idx] = expr.Evaluate(i)
		}
		fmt.Printf(action.PrintfFormat, args...)
	case "length":
		value := action.StringExpression.Evaluate(i).(string)
		fmt.Println(length(value))
	case "substr":
		str := action.StringExpression.Evaluate(i).(string)
		start := int(action.StartExpression.Evaluate(i).(float64))
		length := int(action.LengthExpression.Evaluate(i).(float64))
		fmt.Println(substr(str, start, length))
	case "gsub":
		str := action.StringExpression.Evaluate(i).(string)
		regex := action.RegexExpression.Evaluate(i).(string)
		replacement := action.ReplacementExpression.Evaluate(i).(string)
		fmt.Println(gsub(regex, replacement, str))
	case "toupper":
		value := action.StringExpression.Evaluate(i).(string)
		fmt.Println(toupper(value))
	case "tolower":
		value := action.StringExpression.Evaluate(i).(string)
		fmt.Println(tolower(value))
	// Other cases remain unchanged
	}
}