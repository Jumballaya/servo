package evaluator

import (
	"fmt"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/token"
)

// Eval Class Literal
func evalClassLiteral(node *ast.ClassLiteral, env *object.Environment) object.Object {
	// Get the parent class from the environment
	var parent *object.Class
	if node.Parent != "" {
		pObj, ok := env.Get(node.Parent)
		if !ok {
			return newError(fmt.Sprintf("could not find parent %s of declared class %s", node.Parent, node.Name))
		}

		parent, ok = pObj.(*object.Class)
		if !ok {
			return newError(fmt.Sprintf("parent of class %s is not an object.Class", node.Name))
		}
	}

	class := &object.Class{
		Name:    node.Name,
		Parent:  parent,
		Fields:  node.Fields,
		Methods: make(map[string]object.ClassMethod),
	}
	env.Set(node.Name, class)

	return class
}

// Eval New Class Instance
func evalNewClassInstance(node *ast.NewInstance, env *object.Environment) object.Object {
	ident, ok := node.Class.(*ast.Identifier)
	if !ok {
		return newError("cannot create an instance of a class that doesn't exist")
	}

	class, ok := env.Get(ident.Value)
	if !ok {
		return newError("cannot create an instance of a class that doesn't exist")
	}

	classObj, ok := class.(*object.Class)
	if !ok {
		return newError("cannot create an instance of a class that doesn't exist")
	}

	newEnv := object.NewEnclosedEnvironment(env)
	instance := &object.Instance{Class: classObj, Fields: newEnv}
	newEnv.Set("this", instance)

	for _, f := range classObj.Fields {
		name := f.Name.Value
		evaluated := Eval(f, newEnv)
		if name == "constructor" {
			callExp := &ast.CallExpression{
				Token:     token.Token{Type: token.LPAREN, Literal: "("},
				Function:  f.Value,
				Arguments: node.Arguments,
			}
			Eval(callExp, newEnv)
		} else {
			newEnv.Set(name, evaluated)
		}
	}

	return instance
}
