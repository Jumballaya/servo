package evaluator

import (
	"fmt"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
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
	// 1. Get the ast.ClassLiteral out
	// 2. Get the arguments out and evaluate them
	// 3. Parse the class, including the constructor with the arguments
	// 4. Return the instance object
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

	newEnv := object.NewEnvironment()
	instance := &object.Instance{Class: classObj, Fields: newEnv}
	newEnv.Set("this", instance)
	// Parse the constructor function with the new environment
	// Add all of the class's fields to the environment as well
	for _, f := range classObj.Fields {
		name := f.Name.Value
		evaluated := Eval(f, newEnv)
		if name != "constructor" {
			newEnv.Set(name, evaluated)
		}
	}

	return instance
}
