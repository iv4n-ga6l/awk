package main

func (i *Interpreter) Execute(action *Action) {
	switch action.Command {
	case "print":
		// Existing print logic
	case "if":
		ifStmt := action.IfStatement
		if i.ToNumber(ifStmt.Condition.Evaluate(i)) != 0 {
			for _, act := range ifStmt.TrueBranch {
				i.Execute(act)
			}
		} else {
			for _, act := range ifStmt.FalseBranch {
				i.Execute(act)
			}
		}
	case "while":
		whileLoop := action.WhileLoop
		for i.ToNumber(whileLoop.Condition.Evaluate(i)) != 0 {
			for _, act := range whileLoop.Body {
				i.Execute(act)
			}
		}
	case "for":
		forLoop := action.ForLoop
		forLoop.Init.Evaluate(i)
		for i.ToNumber(forLoop.Condition.Evaluate(i)) != 0 {
			for _, act := range forLoop.Body {
				i.Execute(act)
			}
			forLoop.Post.Evaluate(i)
		}
	case "do-while":
		doWhileLoop := action.DoWhileLoop
		for {
			for _, act := range doWhileLoop.Body {
				i.Execute(act)
			}
			if i.ToNumber(doWhileLoop.Condition.Evaluate(i)) == 0 {
				break
			}
		}
	case "break":
		// Break logic
	case "continue":
		// Continue logic
	case "next":
		// Skip to next record
	case "exit":
		exitCommand := action.ExitCommand
		if exitCommand.ExitCode != nil {
			exitCode := i.ToNumber(exitCommand.ExitCode.Evaluate(i))
			os.Exit(int(exitCode))
		} else {
			os.Exit(0)
		}
	}
}
