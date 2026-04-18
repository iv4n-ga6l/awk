package main

func (i *Interpreter) Execute(action *Action) {
	switch action.Command {
	case "print":
		// Existing print logic
	case "if":
		// Existing if logic
	case "while":
		// Existing while logic
	case "for":
		// Existing for logic
	case "do-while":
		// Existing do-while logic
	case "break":
		// Break logic
	case "continue":
		// Continue logic
	case "next":
		// Skip to next record
	case "exit":
		// Exit logic
	case "delete":
		deleteStmt := action.DeleteStatement
		array := i.GetArray(deleteStmt.ArrayName)
		if array != nil {
			key := deleteStmt.Key.Evaluate(i).(string)
			array.Delete(key)
		}
	case "for-in":
		forInLoop := action.ForInLoop
		array := i.GetArray(forInLoop.ArrayName)
		if array != nil {
			keys := array.Keys()
			for _, key := range keys {
				i.SetVariable(forInLoop.Variable, key)
				for _, act := range forInLoop.Body {
					i.Execute(act)
				}
			}
		}
	}
}

func (i *Interpreter) GetArray(name string) *AssociativeArray {
	value, exists := i.variables[name]
	if !exists {
		array := NewAssociativeArray()
		i.variables[name] = array
		return array
	}
	if arr, ok := value.(*AssociativeArray); ok {
		return arr
	}
	return nil
}