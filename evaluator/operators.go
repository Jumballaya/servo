package evaluator

import (
	"math"

	"github.com/jumballaya/servo/object"
)

// Eval Prefix Expression
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

// Eval Infix Expression
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBooleanToBooleanObject(left == right)
	case operator == "!=":
		return nativeBooleanToBooleanObject(left != right)
	case operator == "&&":
		return evalBooleanInfixExpression(operator, left, right)
	case operator == "||":
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Eval Boolean Infix Expression
func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	l, ok := left.(*object.Boolean)
	if !ok {
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}
	r, ok := right.(*object.Boolean)
	if !ok {
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}

	switch operator {
	case "&&":
		return nativeBooleanToBooleanObject(l.Value && r.Value)
	case "||":
		return nativeBooleanToBooleanObject(l.Value || r.Value)
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Eval Bang Operator Expression
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// Eval Minus Prefix Operator Expression
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// Eval Integer Infix Expression
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBooleanToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBooleanToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBooleanToBooleanObject(leftVal >= rightVal)
	case "<=":
		return nativeBooleanToBooleanObject(leftVal <= rightVal)
	case "==":
		return nativeBooleanToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBooleanToBooleanObject(leftVal != rightVal)
	case ">>":
		return &object.Integer{Value: int64(uint(leftVal) >> uint(rightVal))}
	case "<<":
		return &object.Integer{Value: int64(uint(leftVal) << uint(rightVal))}
	case "^":
		return &object.Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "|":
		return &object.Integer{Value: leftVal | rightVal}
	case "&":
		return &object.Integer{Value: leftVal & rightVal}
	case "&^":
		return &object.Integer{Value: leftVal &^ rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Eval String Infix Expression
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return &object.Boolean{Value: leftVal == rightVal}
	case "!=":
		return &object.Boolean{Value: leftVal != rightVal}
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
