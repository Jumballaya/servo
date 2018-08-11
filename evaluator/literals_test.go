package evaluator

import (
	"testing"

	"github.com/jumballaya/servo/object"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"100 % 15", 10},
		{"100 >> 2", 25},
		{"100 << 2", 400},
		{"100 &^ 67", 36},
		{"100 & 7", 4},
		{"112 | 54", 118},
		{"5 ^ 2", 25},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"3.14", 3.14},
		{"9.0", 9.0},
		{"-5.0", -5.0},
		{"-10.0", -10.0},
		{"5.5 + 5.5 + 5.5 + 5.5 - 10", 12.0},
		{"2.0 * 2.0 * 2.0 * 2.0 * 2", 32.0},
		{"-50.0 + 100.0 + -50", 0.0},
		{"5.0 * 2 + 10.0", 20.0},
		{"5.0 + 2 * 10.0", 25.0},
		{"20 + 2.0 * -10", 0.0},
		{"50.0 / 2 * 2 + 10", 60.0},
		{"2 * (5.0 + 10)", 30.0},
		{"3 * 3 * 3.0 + 10", 37.0},
		{"3 * (3 * 3.0) + 10", 37.0},
		{"(5 + 10 * 2 + 15.0 / 3) * 2 + -10", 50.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"10 <= 5", false},
		{"10 >= 100", false},
		{"1000 >= 1000", true},
		{"1000 <= 1000", true},
		{"!null", true},
		{"!!null", false},
		{"true && true", true},
		{"true && false", false},
		{"false && false", false},
		{"true || true", true},
		{"true || false", true},
		{"false || false", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Hello World!"`, "Hello World!"},
		{`'Hello World!'`, "Hello World!"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}

		if str.Value != tt.expected {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	}
}
