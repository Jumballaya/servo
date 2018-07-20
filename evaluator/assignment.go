package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

// Eval Let Statement
func evalLetStatement(node *ast.LetStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	return env.Set(node.Name.Value, val)
}

// Eval Assignment Statement
func evalAssignStatement(node *ast.AssignStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	exp, _ := node.Left.(ast.Expression)
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		return newError("left value is not an identifier")
	}

	return env.Set(ident.Value, val)
}

// Eval Return Statement
func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(node.ReturnValue, env)
	if isError(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}

// Unwrap Return Value gets the return expression from an object and evaluates its value
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

// Eval Import Statement
func evalImportStatement(importExp *ast.ImportExpression, env *object.Environment) object.Object {
	mod := importExp.Path.String()
	obj := importExp.Name.String()

	// Open the module's file
	// Evaluate it
	// Rip out the identifier that is being imported

	// Comes from a file
	if strings.HasPrefix(mod, "./") || strings.HasPrefix(mod, "../") || strings.HasPrefix(mod, "/") {
		var currentFile string
		if len(os.Args) < 2 {
			currentFile = os.Args[0]
		} else {
			currentFile = os.Args[1]
		}
		currentDir := "./" + strings.Join(strings.Split(currentFile, "/")[:1], "/")

		dir, err := filepath.Abs(currentDir + "/" + mod)
		if err != nil {
			fmt.Println(err.Error())
			return newError(err.Error())
		}
		pulled := GetObjectFromFile(dir, obj)
		env.Set(obj, pulled)
		return NULL
		// Comes from the standard lib
	} else {
		fn, ok := env.Get(mod)
		if !ok {
			return newError("improper import, function key not found")
		}

		hash, ok := fn.(*object.Hash)
		if !ok {
			return newError("Must import off of an exported hash")
		}

		key := (&object.String{Value: obj}).HashKey()

		found := hash.Pairs[key].Value
		_, ok = found.(*object.Null)
		if ok {
			env.Set(obj, NULL)
			return NULL
		}

		env.Set(obj, found)
		return NULL
	}

	return NULL
}
