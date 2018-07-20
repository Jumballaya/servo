package evaluator

import (
	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

// Eval If Expression
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

// Is Truthy determines if a non-boolean value should act as a literal true or false
func isTruthy(obj object.Object) bool {
	// Check to see if a string is empty or not
	str, ok := obj.(*object.String)
	if ok {
		if str.Value != "" {
			return true
		}
		return false
	}

	// Check if a number is 0 or not
	num, ok := obj.(*object.Integer)
	if ok {
		if num.Value != 0 {
			return true
		}
		return false
	}

	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
