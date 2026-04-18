package main

// ControlFlow constructs

type ForInLoop struct {
	Variable  string
	ArrayName string
	Body      []*Action
}

type DeleteStatement struct {
	ArrayName string
	Key       Expression
}

func (f *ForInLoop) Execute(interpreter *Interpreter) {
	array := interpreter.GetArray(f.ArrayName)
	if array != nil {
		keys := array.Keys()
		for _, key := range keys {
			interpreter.SetVariable(f.Variable, key)
			for _, act := range f.Body {
				interpreter.Execute(act)
			}
		}
	}
}

func (d *DeleteStatement) Execute(interpreter *Interpreter) {
	array := interpreter.GetArray(d.ArrayName)
	if array != nil {
		key := d.Key.Evaluate(interpreter).(string)
		array.Delete(key)
	}
}