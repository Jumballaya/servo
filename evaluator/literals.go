package evaluator

import (
	"strings"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

// Eval Identifier
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if env.Silent && node.Value == "log" {
		return NULL
	}

	if builtin, ok := getBuiltin(node.Value, env); ok {
		return builtin
	}

	return newError("identifier not found: %s", node.Value)
}

// Eval Integer Literal
func evalIntegerLiteral(node *ast.IntegerLiteral, env *object.Environment) object.Object {
	return &object.Integer{Value: node.Value}
}

// Eval Integer Literal
func evalFloatLiteral(node *ast.FloatLiteral, env *object.Environment) object.Object {
	return &object.Float{Value: node.Value}
}

// Native Boolean To Boolean Object returns the singletons TRUE and FALSE of type object.Boolean
// from the native go true and false
func nativeBooleanToBooleanObject(val bool) *object.Boolean {
	if val {
		return TRUE
	}
	return FALSE
}

// Eval String Literal
func evalStringLiteral(node *ast.StringLiteral, env *object.Environment) object.Object {
	val := node.Value
	val = strings.Replace(val, "\\n", "\n", -1)
	val = strings.Replace(val, "\\t", "\t", -1)
	val = strings.Replace(val, "\\r", "\r", -1)
	return &object.String{Value: val}
}
