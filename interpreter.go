package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type FunctionDefinition struct {
	Name       string
	Parameters []string
	Body       []*Action
}

type ReturnStatement struct {
	Expression Expression
}

func (f *FunctionDefinition) Execute(interpreter *Interpreter) {
	interpreter.DefineFunction(f.Name, f)
}

func (r *ReturnStatement) Execute(interpreter *Interpreter) {
	interpreter.SetReturnValue(r.Expression.Evaluate(interpreter))
}

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
	case "function":
		if action.FunctionDefinition != nil {
			action.FunctionDefinition.Execute(i)
		}
	case "return":
		if action.ReturnStatement != nil {
			action.ReturnStatement.Execute(i)
		}
	case "int":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(int(value))
	case "sqrt":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(math.Sqrt(value))
	case "sin":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(math.Sin(value))
	case "cos":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(math.Cos(value))
	case "atan2":
		y := action.YExpression.Evaluate(i).(float64)
		x := action.XExpression.Evaluate(i).(float64)
		fmt.Println(math.Atan2(y, x))
	case "exp":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(math.Exp(value))
	case "log":
		value := action.NumericExpression.Evaluate(i).(float64)
		fmt.Println(math.Log(value))
	case "rand":
		fmt.Println(rand.Float64())
	case "srand":
		seed := action.NumericExpression.Evaluate(i).(float64)
		rand.Seed(int64(seed))
	// Other cases remain unchanged
	}
}

func (i *Interpreter) DefineFunction(name string, function *FunctionDefinition) {
	i.functions[name] = function
}

func (i *Interpreter) CallFunction(name string, args []Expression) interface{} {
	function, exists := i.functions[name]
	if !exists {
		panic(fmt.Sprintf("Undefined function: %s", name))
	}
	localScope := make(map[string]interface{})
	for idx, param := range function.Parameters {
		localScope[param] = args[idx].Evaluate(i)
	}
	prevScope := i.scope
	i.scope = localScope
	for _, action := range function.Body {
		i.Execute(action)
		if i.returnValue != nil {
			result := i.returnValue
			i.returnValue = nil
			i.scope = prevScope
			return result
		}
	}
	i.scope = prevScope
	return nil
}

func (i *Interpreter) SetReturnValue(value interface{}) {
	i.returnValue = value
}