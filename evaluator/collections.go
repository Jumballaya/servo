package evaluator

import (
	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/object"
)

// Eval Array Literal
func evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}
	return &object.Array{Elements: elements}
}

// Eval Index Expression
func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

// Eval Array Index Expression
func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	id := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if id < 0 || id > max {
		return NULL
	}

	return arrayObject.Elements[id]
}

// Eval Hash Literal
func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as a hash key: %s", key.Type())
		}

		value := Eval(valNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

// Eval Hash Index Expression
func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as a hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}

// Eval Attribute Expression
func evalAttributeExpression(node *ast.AttributeExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	var Left object.Object

	str, ok := node.Left.(*ast.Identifier)
	if ok {
		Left, ok = env.Get(str.Value)
		if !ok {
			return newError("identifier not found '%s'", str.Value)
		}
	} else {
		Left = left
	}

	instance, ok := Left.(*object.Instance)
	if !ok {
		return newError("left hand side not an instance, type: %T", left)
	}

	method := instance.GetMethod(node.Index.Value)

	if method == nil {
		field, ok := instance.Fields.Get(node.Index.Value)
		if !ok {
			field, ok = env.Get(node.Index.Value)
			if !ok {
				return NULL
			}
			return wrapInstanceEnvironment(field, instance.Fields)
		}
		return wrapInstanceEnvironment(field, instance.Fields)
	}
	return wrapInstanceEnvironment(method, instance.Fields)
}

func wrapInstanceEnvironment(obj object.Object, env *object.Environment) object.Object {
	fn, ok := obj.(*object.Function)
	if ok {
		fn.Env = env
		return fn
	}
	return obj
}
