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

	for _, field := range node.Fields {
		encEnv := object.NewEnclosedEnvironment(env)
		ident := field.Name.Value
		this := &object.Instance{Class: class, Fields: encEnv}
		encEnv.Set("this", this)
		eval := Eval(field, encEnv)
		fn, ok := eval.(*object.Function)
		if ok {
			class.Methods[ident] = fn
		}
	}

	env.Set(node.Name, class)

	return class
}

// Eval New Class Instance
func evalNewClassInstance(node *ast.InstanceLiteral, env *object.Environment) object.Object {
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
	if instance.Class.Parent != nil {
		for _, stmt := range instance.Class.Parent.Fields {
			if stmt.Name.Value == "constructor" {
				superFn := evalLetStatement(stmt, newEnv)
				newEnv.Set("super", superFn)
			}
		}
	}
	newEnv.Set("this", instance)

	for _, f := range classObj.Fields {
		name := f.Name.Value
		switch f.Value.(type) {
		case *ast.FunctionLiteral:
			if name == "constructor" {
				if f.Value != nil {
					function, _ := f.Value.(*ast.FunctionLiteral)
					callExp := &ast.CallExpression{
						Token:     token.Token{Type: token.LPAREN, Literal: "("},
						Function:  function,
						Arguments: node.Arguments,
					}
					// Fix this
					evalCallFunction(callExp, newEnv)
				}
			} else {
				evaluated := Eval(f, newEnv)
				newEnv.Set(name, evaluated)
				fn, ok := evaluated.(*object.Function)
				if ok {
					instance.Class.Methods[name] = fn
				}
			}
		default:
			evaluated := Eval(f, newEnv)
			newEnv.Set(name, evaluated)
			fn, ok := evaluated.(*object.Function)
			if ok {
				instance.Class.Methods[name] = fn
			}
		}
	}

	return instance
}
