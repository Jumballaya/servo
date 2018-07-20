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
