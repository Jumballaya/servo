package evaluator

import (
	"fmt"
	"math"
	"strconv"

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
	case isNumber(left.Type()) && isNumber(right.Type()):
		return evalFloatInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case (left.Type() == object.STRING_OBJ || isNumber(left.Type())) && (right.Type() == object.STRING_OBJ || isNumber(right.Type())) && operator == "+":
		return evalMixStringIntegerInfixExpression(operator, left, right)
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

func isNumber(obj object.ObjectType) bool {
	return obj == object.INTEGER_OBJ || obj == object.FLOAT_OBJ
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
	if right.Type() == object.INTEGER_OBJ {
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	} else if right.Type() == object.FLOAT_OBJ {
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	}
	return newError("unknown operator: -%s", right.Type())
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

func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	var leftVal float64
	var rightVal float64
	l, ok := left.(*object.Float)
	if ok {
		leftVal = l.Value
	} else {
		leftVal = float64(left.(*object.Integer).Value)
	}

	r, ok := right.(*object.Float)
	if ok {
		rightVal = r.Value
	} else {
		rightVal = float64(right.(*object.Integer).Value)
	}

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
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
	case "^":
		return &object.Float{Value: math.Pow(leftVal, rightVal)}
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

func evalMixStringIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := ""
	rightValue := ""

	switch left.(type) {
	case *object.Integer:
		leftValue = strconv.FormatInt(left.(*object.Integer).Value, 10)
	case *object.String:
		leftValue = left.(*object.String).Value
	case *object.Float:
		leftValue = strconv.FormatFloat(left.(*object.Float).Value, 'f', -1, 64)
	}

	switch right.(type) {
	case *object.Integer:
		rightValue = strconv.FormatInt(right.(*object.Integer).Value, 10)
	case *object.String:
		rightValue = right.(*object.String).Value
	case *object.Float:
		rightValue = strconv.FormatFloat(right.(*object.Float).Value, 'f', -1, 64)
	}

	return &object.String{Value: fmt.Sprintf("%s%s", leftValue, rightValue)}
}
