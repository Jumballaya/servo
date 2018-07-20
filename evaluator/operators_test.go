package evaluator

import (
	"testing"

	"github.com/jumballaya/servo/object"
)

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestStringInfixOperations(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`"Hello" + " " + "World!"`, "Hello World!"},
		{`"Hello World!" == "Hello World!"`, true},
		{`"Hello World!" != "Hello World!"`, false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch tt.expected.(type) {
		case string:
			str, ok := evaluated.(*object.String)
			if !ok {
				t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
			}
			if str.Value != tt.expected {
				t.Errorf("String has wrong value. got=%q", str.Value)
			}
		case bool:
			b, ok := evaluated.(*object.Boolean)
			if !ok {
				t.Fatalf("object is not Boolean. got=%T (%+v)", evaluated, evaluated)
			}
			if b.Value != tt.expected {
				t.Errorf("Boolean has wrong value. got=%t", b.Value)
			}
		}
	}
}

func TestReassignOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{`let x = 0; x += 5;`, 5},
		{`let x = 10; x -= 5;`, 5},
		{`let x = 5; x *= 5;`, 25},
		{`let x = 10; x /= 5;`, 2},
		{`let x = 0; x += 5; x;`, 5},
		{`let x = 10; x -= 5; x;`, 5},
		{`let x = 5; x *= 5; x;`, 25},
		{`let x = 10; x /= 5; x;`, 2},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		num, ok := evaluated.(*object.Integer)
		if !ok {
			t.Fatalf("object is not an Integer. got=%T (%+v)", evaluated, evaluated)
		}
		if num.Value != tt.expected {
			t.Errorf("Integer has wrong value. got=%q", num.Value)
		}
	}
}
