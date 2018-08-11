package evaluator

import (
	"fmt"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

// Eval is the evaluator function that recursively runs, evaluating the program
// and its statements.
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Main Program
	case *ast.Program:
		return evalProgram(node, env)

	// General Expression
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// Import Expression
	case *ast.ImportExpression:
		return evalImportStatement(node, env)

	// Integer
	case *ast.IntegerLiteral:
		return evalIntegerLiteral(node, env)

	// Float
	case *ast.FloatLiteral:
		return evalFloatLiteral(node, env)

	// Boolean
	case *ast.Boolean:
		return nativeBooleanToBooleanObject(node.Value)

	// String
	case *ast.StringLiteral:
		return evalStringLiteral(node, env)

	// Array
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)

	// Hash
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)

	// Class
	case *ast.ClassLiteral:
		return evalClassLiteral(node, env)

	// New Instance
	case *ast.NewInstance:
		return evalNewClassInstance(node, env)

	// Attribute Expression
	case *ast.AttributeExpression:
		return evalAttributeExpression(node, env)

	// Index Expression
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	// Prefix
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	// Infix
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	// Block
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	// If
	case *ast.IfExpression:
		return evalIfExpression(node, env)

	// Return
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)

	// Let
	case *ast.LetStatement:
		return evalLetStatement(node, env)

	// Assignment
	case *ast.AssignStatement:
		return evalAssignStatement(node, env)

	// Identifier
	case *ast.Identifier:
		return evalIdentifier(node, env)

	// Function
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)

	// Function Call
	case *ast.CallExpression:
		return evalCallFunction(node, env)

	// Default
	default:
		return NULL
	}
}

// New Error generates an error object
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// Eval Pogram
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

// Eval Statements
func evalStatements(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

// Eval Expressions
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

// Eval Block Statement
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range block.Statements {
		result = Eval(stmt, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

// Is Error checks if an object is an error or not
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
