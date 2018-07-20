package evaluator

import (
	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

// Eval Function Literal
func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	params := node.Parameters
	body := node.Body
	return &object.Function{Parameters: params, Env: env, Body: body}
}

// Eval Call Function
func evalCallFunction(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyFunction(function, args)
}

// Apply Function parses the function call expressions
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

// Extend Function Env binds the enclosed environment of the function with the parameters
// it was called with
// e.g. `let add = fn(x, y) { x + y };` then we call it can evaluate it
//
// `add(1, 2)` would call extendFunctionEnv with the object.Function of `add` and the args of
// `x = 1` and `y = 2` before you can evaluate `x + y` you must bind them to the environment
// of that block. This function is doing exactly that: binding x to 1 and y to 2 in the local
// environment.
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramId, param := range fn.Parameters {
		env.Set(param.Value, args[paramId])
	}

	return env
}
